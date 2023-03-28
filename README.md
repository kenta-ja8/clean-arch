# Clean Architecture

## Setup Database
```
go run script/script.go setup
```

## Run
```
go run main.go
```

## Test
```
# Get User
curl localhost:8080/api/user | jq .
curl -H "Accept: application/xml" localhost:8080/api/user

# Add User
curl -X POST -d '{"name":"xxxx", "birthday":"2022-04-01T00:00:00Z"}' localhost:8080/api/user  | jq .
```

## Clean Database
```
go run script/script.go clean
```

