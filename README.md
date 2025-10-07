# structure

test different source code structures in golang


## flat 

```
main.go         // entry point
message.go      // model
handler.go      // http handlers
service.go      // business logic
repository.go   // data access layer
```

## how to run

```bash
export FLAT_PORT=5000
export FLAT_DB_DRIVER_NAME=mysql
export FLAT_DB_DATA_SOURCE_NAME="[user]:[password]@tcp([hostname]:[port])/[dbname]" 
```

build

```bash
go build main.go repository.go service.go handler.go message.go 
    ./main
```

or run directly 

```bash
go run .
```

curl ping

    curl -s 'http://localhost:5000/ping' | jq .
