# microservice-login-api

Language : golang, angular\
Purpose : microservice example.

-- DB restore --\
mongorestore --db microservice ./microservice_DB_dump/microservice

-- token api server --\
$go mod tidy\
$./startapp.sh

-- login api server --\
change IP to your DB server : env/env.go\
change IP to your token api server : microservices/token/token.go (obj.API_SERVER)\
$go mod tidy \
$./startapp.sh 

-- web server --\
change IP to your login service server : src/app/env.ts\
$npm start\
web : http://IP:4200/login \
ID(email)/PW(plain) is stored in microservice DB.
