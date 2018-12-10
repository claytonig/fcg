#Car management system

##TO RUN 
```
make dep 
go run main.fo
```

##RUN TESTS 
```
make test
```

##API's 
```
GET /fcg/cars HTTP/1.1
Host: localhost:8085

GET /fcg/cars/{uuid} HTTP/1.1
Host: localhost:8085

POST /fcg/cars HTTP/1.1
Host: localhost:8085
Content-Type: application/json

{
"make":"chevrolet",
"year":"2005",
"model":"U-VA"
}

DELETE /fcg/cars/{uuid} HTTP/1.1
Host: localhost:8085
```
