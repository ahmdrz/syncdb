# syncdb
Copy database tables from server to another server

This repository is just for fun and made me happy.

I love syncing , I love synchronizing.

This code can `SELECT` and `INSERT` :smile: in your tables.

Change `driver.go` for your database driver then let's compile...

Note: *default is mssql* `go get github.com/denisenkom/go-mssqldb`

### Connection Strings

For `mssql` driver : 

1. local : `server=127.0.0.1;encrypt=disable;database=*`
2. server : `server=*;user id=*;password=*;database=*;port=*;`

### JSON Configuration File

```json
{
  "__connectionstring_comment": "these connection strings tested on mssql",
  
  "home": "server=*;user id=*;password=*;port=*;database=*",
  "away": "server=*;user id=*;password=*;port=*;database=*",
  
  "__direction_comment": "if true (server to local) else (local to server)",
  "direction": true,
  
  "tables": [
    {
      "name": "*",
      "column": "*"
    },
    {
      "name": "*",
      "column": "*"
    }
  ]
}
```

### Keywords 

Sync database , Sync tables , Golang , sync two database , copy tables , mssql , mysql , etc ...
