# [Step-2] Run Swagger on docker
#   [https://stackoverflow.com/questions/46235656/how-to-install-swagger-on-ubuntu]
    docker pull swaggerapi/swagger-editor
    docker run -p 80:8080 swaggerapi/swagger-editor

# [Step-2]
If you are using docker, simply pull and run the swaggerapi/swagger-editor

docker pull swaggerapi/swagger-editor
docker run -p 80:8080 swaggerapi/swagger-editor
Open your browser to http://localhost:80/

Note, make sure your browser does not silently redirect to https://localhost:80/

Like rugby2312 mentions, you can optionally pass in an existing swagger.json.

mkdir tmp
cp swagger.json /tmp/swagger.json
docker run -d -p 80:8080 -v $PWD/tmp:/tmp -e SWAGGER_FILE=/tmp/swagger.json swaggerapi/swagger-editor

# [Step-3]
As I have already lamp server running, I run container with: docker run -d -p 8080:8080 swaggerapi/swagger-editor to avoid error – 
bcag2
 Nov 4, 2019 at 14:28
1
If you want to pass an existing swagger.json to the container , remember to do the folder mapping and declare the file name as a variable docker run -d -p 8080:8080 -v /home/ubuntu/dswagger:/tmp -e SWAGGER_FILE=/tmp/swagger.json swaggerapi/swagger-editor – 
rugby2312
 Mar 8, 2022 at 11:16 


#   [Tutorial] [youtube]
# Install swagger cmd
    $ go get -u github.com/swaggo/swag/cmd/swag
                or
    $ go get -u github.com/swaggo/swag/cmd/swag@latest

# add libraries
    $   go get -u github.com/swaggo/gin-swagger
    $   go get -u github.com/swaggo/files

# [warning]
    if swag command not found, run:
        $ export PATH=$PATH:$GOPATH/bin
            of
        export PATH="/home/hussain/go/bin/:$PATH"
        (path to swag which inside [user/go/bin])

# Initialize swagger [if creating new]
    $ swag init
    now new folder is created as docs, which has docs.go, swagger.yaml and swagger.json

# add new router in main function
    import (
        ginSwagger   "github.com/swaggo/gin-swagger"
        swaggerFiles "github.com/swaggo/files"
    )
    func NewRouter(u *UserConfig){
    router := gin.Default()
#   //Add this line to router
    router.GET("/docs/*any", ginSwagger.WrapGandler(swaggerFiles.Handler))

    }

# Add below to main.go
    _ "github.com/furqanalimir/commuter/docs"
#   Add below to above main func:
        // @BasePath /api/v0.1
        // @title Go + Gin User API
        // @version 1.0
        // @description This is a sample server user server. You can visit the GitHub repository at https://github.com/Furqanalimir/commuter
        
        // @contact.name API Support
        // @contact.url https://furqanali.vercel.app/
        // @contact.email mirfurqan89@gmail.com
        
        // @license.name MIT
        // @license.url https://opensource.org/licenses/MIT
        
        // @host localhost:8080
        // @query.collection.format multi


#  Add comments on routes (handlers)
    // AddUser		godoc
    // @Summary		Create user
    // @Description	Save User data
    // @Param		user body data.User true "create user"      #expectes in body with type User (ex: {Email, Phone etc...})
    // @Param		email formData string true "email address"  #expectes in formData or x-www-form-urlencooded string (check how to differentiate)
    // @Param		id path int true "get fruit by id"          #expects id in url of type int
    // @produce		applicaton/json                             #response type produced
    // @Tags		user                                        #tag
    // @Success		200	{object} gin.H  "create response"       #status {objecct} type "heading/message"
    // @Success		400	{object} gin.H  "error response"        #diffstatus {objecct} type "heading/message"
    // @Router		/user [post]                                #route type
    // @Security ApiKeyAuth                                     #middleware validation (will send token in header)

# Run
    swag init
        or
    swag init --parseDependency
        or
    swag init -g path/to/main.go --parseDependency

# Run swagger command to generate [httpClient]
    can be used for testing
    $   swagger generate client -f ../path/to/doc/swagger.yamlm -A <main-filename>
