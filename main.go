// syncdb project main.go
package main

import (
	"fmt"
	"reflect"

	_ "github.com/denisenkom/go-mssqldb" //i used mssqldb library for connection between two SQL Server databases.
	// replace your library for database driver.
)

/*const (
	homeConnectionString = "<HomeServer>"
	awayConnectionString = "<AwayServer>"

	database_driver = "mssql"

	selected_table  = "<Selected_Table>"
	selected_column = "<Selected_Column>" // for preventing conflicts
)*/

func main() {
	//making connection for home...
	home := Connection(homeConnectionString)
	defer home.Close()
	//retrieving home datas using SELECT query
	rows, err := home.Query("SELECT * FROM " + selected_table + ";")
	Panic(err)
	//retrieving home columns
	homeColumns, err := rows.Columns()
	Panic(err)
	homeValues := MakeValues(rows, homeColumns)
	rows.Close()
	// now we have home datas with primary keys in homeValues

	//try again for away datas...
	away := Connection(awayConnectionString)
	defer away.Close()
	rows, err = away.Query("SELECT * FROM " + selected_table + ";")
	Panic(err)
	defer rows.Close()
	awayColumns, err := rows.Columns()
	Panic(err)
	awayValues := MakeValues(rows, awayColumns)

	if len(awayColumns) != len(homeColumns) {
		Panic(fmt.Errorf("Incorrect length of tables , They aren't same !"))
	}

	stmt, err := away.Prepare(MakeStmt(awayColumns))
	Panic(err)
	defer stmt.Close()

	for _, homeValue := range homeValues {
		exists := false
		for _, awayValue := range awayValues {
			if reflect.ValueOf(awayValue.Primary_Value).Type() != reflect.ValueOf(homeValue.Primary_Value).Type() {
				Panic(fmt.Errorf("Incorrect type of two primary_values"))
			}

			if awayValue.Primary_Value == homeValue.Primary_Value {
				exists = true
				fmt.Println("Conflict with", homeValue.Primary_Value)
				break
			}
		}

		if exists {
			continue
		}

		// new row detected ... let's start uploading
		fmt.Println("Uploading", homeValue.Primary_Value)
		//stmt.Exec(homeValue.Values...)
	}
}
