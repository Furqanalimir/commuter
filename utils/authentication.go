package helper

import (
	"errors"
	"fmt"
	"microservices/micro-service/commuter/data"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

func CreateToken(userid uint64) (string, error) {
	var err error
	//Creating Access Token
	os.Setenv("ACCESS_SECRET", "jdnfksdmfksd") //this should be in an env file
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userid
	atClaims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return "", err
	}
	return token, nil
}

func TokenValid(r *http.Request) error {
	token, err := VerifyToken(r)
	if err != nil {
		return err
	}
	// _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid
	if !token.Valid {
		return err
	}
	return nil
}

func VerifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString := ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//does this token conform to "SigningMethodHMAC" ?
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("ACCESS_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

// get the token from the request body
func ExtractToken(r *http.Request) string {
	keys := r.URL.Query()
	token := keys.Get("token")
	if token != "" {
		return token
	}
	bearToken := r.Header.Get("Authorization")
	//normally Authorization the_token_xxx
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

func ExtractTokenAuth(r *http.Request) (uint64, error) {
	token, err := VerifyToken(r)
	if err != nil {
		return 0, err
	}
	claims, ok := token.Claims.(jwt.MapClaims) //the token claims should conform to MapClaims
	if ok && token.Valid {
		// authUuid, ok := claims["auth_uuid"].(string) //convert the interface to string
		// if !ok {
		// 	return nil, err
		// }
		userId, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)
		// claims["user_id"] //||
		// userId, err
		if err != nil {
			return 0, err
		}
		return userId, nil
	}
	return 0, err
}

func verifyRole(r *http.Request, role string) bool {
	token, err := VerifyToken(r)
	if err != nil {
		return false
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		userId, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)
		if err != nil {
			return false
		}
		user, err := data.GetUserById(int(userId))
		if err != nil {
			return false
		}
		return user.VerifyRole(role)
	}
	return false
}

func VerifyAdmin(r *http.Request) error {
	role := verifyRole(r, "admin")
	if role {
		return nil
	}
	return errors.New("you don't have admin access")
}

func GetTokenValidity(r *http.Request) (*jwt.NumericDate, error) {
	token, err := VerifyToken(r)
	if err != nil {
		return jwt.NewNumericDate(time.Now()), err
	}
	auth, err := token.Claims.GetExpirationTime()
	if err != nil {
		return auth, err
	}
	return auth, nil
}
