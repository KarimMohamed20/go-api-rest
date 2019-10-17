# GoApiRest
Simple example api rest with Gin framework and GORM


## Install packages and dependencies 

```
   $ make get
```

## Compile packages and dependencies

```
   $ make build
```

## Compile and run Go program

```
   $ make run
```

## Getting with Curl 

```
    $ curl -H 'content-type: application/json' -v -X GET http://127.0.0.1:8080/api/v1/books 
    $ curl -H 'content-type: application/json' -v -X GET http://127.0.0.1:8080/api/v1/books/:id
    $ curl -H 'content-type: application/json' -v -X POST -d '{"title":"Foo bar","description":"Lorem ipsum"}' http://127.0.0.1:8080/api/v1/books 
    $ curl -H 'content-type: application/json' -v -X PUT -d '{"title":"Foo bar","description":"Lorem ipsum"}' http://127.0.0.1:8080/api/v1/books/:id
    $ curl -H 'content-type: application/json' -v -X DELETE http://127.0.0.1:8080/api/v1/books/:id
```
