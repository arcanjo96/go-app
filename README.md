## install dependencies
```
go mod download
```

## generate sqlite3 database
```
run: sqlite3 db.sqlite3
sqlite> create table order(id varchar not null primary key, price not null float, tax not null float, final_price not null float);
```
## run application
### run application order
```
go run cmd/order/main.go
```
### run api
```
go run cmd/api/main.go
```

## run tests
```
go test ./...
```