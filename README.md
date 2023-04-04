Running the api:

docker-compose up --build

The container application will expose port 9090
The container database will expose port 3306


Routes and curls:

Signin
I used jwt with cookie auth
curl http://localhost:9090/signin -d '{"username" : "admin", "password" : "123456"}' -c cookie.txt -X POST

Logout
curl http://localhost:9090/logout -c cookie.txt

Refresh Token
curl http://localhost:9090/refresh -b cookie.txt

GetUsers
curl http://localhost:9090/users -b cookie.txt

GetUserById
curl http://localhost:9090/users/getbyid?id=1 -b cookie.txt

CreateUser
curl -d '{"name":"Samir", "age" : "38", "address" : "Canoas", "password" : "123123", "email" : "samirortiz@gmail.com"}' -H "Content-Type: application/json" http://localhost:9090/users/create -X POST -b cookie.txt 

UpdateUser
curl -d '{"name":"Samir Ortiz", "age" : "40", "address" : "Porto Alegre", "password" : "999000", "email" : "samiroquai@gmail.com"}' -H "Content-Type: application/json" -X POST http://localhost:9090/users/update?id=1 -X PUT -b cookie.txt

DeleteUser
curl http://localhost:9090/users/delete?id=1 -X DELETE -b cookie.txt

Testing
CGO_ENABLED=1 go test

Documentation available at https://documenter.getpostman.com/view/3547764/2s93RTSYpn

Thank you guys.

Samir
