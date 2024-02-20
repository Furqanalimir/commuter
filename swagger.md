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
