package data

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/go-playground/validator"
)

type Fruit struct {
	ID        int     `json:"id"`
	Name      string  `json:"name" validate:"required,min=4,max=15"`
	Price     float32 `json:"price" validate:"required,numeric,gt=2"`
	Currency  string  `json:"currency" validate:"required,oneof=usd kd"`
	Origin    string  `origin:"origin" validate:"required"`
	CreatedAt string  `json:"-"`
	UpdatedAt string  `json:"-"`
	DeletedAt string  `json:"-"`
}

type Fruits []*Fruit

func (f *Fruit) Validate() error {
	validate := validator.New()
	//custom validation for field currency
	err := validate.Struct(f)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			log.Println(err.Field(), err.Tag())
		}
	}
	return err
}

// func translateError(err error, trans ut.Translator) (errs []error) {
// 	if err == nil {
// 		return nil
// 	}
// 	validatorErrs := err.(validator.ValidationErrors)
// 	for _, e := range validatorErrs {
// 		translatedErr := fmt.Errorf(e.Translate(trans))
// 		errs = append(errs, translatedErr)
// 	}
// 	return errs
// }

// func validateFName(fl validator.FieldLevel) bool {
// 	log.Println("[values]", fl.Field().String())
// 	if len(fl.Field().String()) < 1 {
// 		log.Println("values less than one")
// 	}
// 	re := regexp.MustCompile(`[a-b]+-[A-B]+`)
// 	matches := re.FindAllString(fl.Field().String(), -1)
// 	log.Println("[matches] ", matches)
// 	return len(matches) == 1
// }

func getNextId() int {
	f := fruitList[len(fruitList)-1]
	return f.ID + 1
}
func (f *Fruit) ToJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(f)
}
func GetFruit(id int) (*Fruit, error) {
	for _, f := range fruitList {
		if f.ID == id {
			return f, nil
		}
	}
	return &Fruit{}, fmt.Errorf("item with id %d not found", id)
}
func GetAllFuits() Fruits {
	return fruitList
}
func AddFruit(f *Fruit) (Fruits, error) {
	f.ID = getNextId()
	fruitList = append(fruitList, f)
	return fruitList, nil
}

func RemoveFruit(id int) error {
	for i, f := range fruitList {
		if f.ID == id {
			fruitList = append(fruitList[:i], fruitList[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("item with id %d not found", id)
}

var fruitList = Fruits{
	&Fruit{
		ID:        1,
		Name:      "Water Mellon",
		Price:     1.1,
		Currency:  "usd",
		Origin:    "africa",
		CreatedAt: time.Now().UTC().String(),
		UpdatedAt: time.Now().UTC().String(),
	},
	&Fruit{
		ID:        2,
		Name:      "Orange",
		Price:     0.99,
		Currency:  "usd",
		Origin:    "florida",
		CreatedAt: time.Now().UTC().String(),
		UpdatedAt: time.Now().UTC().String(),
	},
	&Fruit{
		ID:        3,
		Name:      "Apple",
		Price:     1.9,
		Currency:  "usd",
		Origin:    "kashmir",
		CreatedAt: time.Now().UTC().String(),
		UpdatedAt: time.Now().UTC().String(),
	},
}
