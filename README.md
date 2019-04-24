# GoApiRest
Simple example api rest with Gin framework and GORM

## Compile packages and dependencies

```
   $ go build src/main.go
```

## Compile and run Go program

```
   $ go run src/main.go
```

## Getting with Curl 

```
    $ curl -H 'content-type: application/json' -v -X GET http://127.0.0.1:8080/api/books 
    $ curl -H 'content-type: application/json' -v -X GET http://127.0.0.1:8080/api/books/:id
    $ curl -H 'content-type: application/json' -v -X POST -d '{"title":"Foo bar","description":"Lorem ipsum"}' http://127.0.0.1:8080/api/books 
    $ curl -H 'content-type: application/json' -v -X PUT -d '{"title":"Foo bar","description":"Lorem ipsum"}' http://127.0.0.1:8080/api/books/:id
    $ curl -H 'content-type: application/json' -v -X DELETE http://127.0.0.1:8080/api/books/:id
```
