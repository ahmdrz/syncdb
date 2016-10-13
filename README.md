# syncdb
Copy database tables from server to another server

This repository is just for fun !

I love syncing , I love synchronizing.

This code can `SELECT` and `INSERT` :smile: in your tables.

Change `driver.go` for your database driver then let's compile...

Note: *default is mssql* `go get github.com/denisenkom/go-mssqldb`

### Connection Strings

For `mssql` driver : 

1. local : `server=127.0.0.1;encrypt=disable;database=*`
2. server : `server=*;user id=*;password=*;database=*;port=*;`
