// syncdb project main.go
package main

import (
	"fmt"
	"reflect"

	_ "github.com/denisenkom/go-mssqldb" //i used mssqldb library for connection between two SQL Server databases.
	// replace your library for database driver.
)

var config Config

const (
	database_driver = "mssql"
)

func main() {
	config = Read()

	// config.Direction -> if true : home to away , if false : away to home

	if config.Direction {
		config.Home, config.Away = config.Away, config.Home
		fmt.Println("Syncing home to away")
	} else {
		fmt.Println("Syncing away to home")
	}

	//making connection for home...
	home := Connection(config.Home)
	defer home.Close()
	away := Connection(config.Away)
	defer away.Close()

	for _, table := range config.Tables {
		fmt.Println("Checking for", table.Name, "in", table.Column)

		//retrieving home datas using SELECT query
		rows, err := home.Query("SELECT * FROM " + table.Name + ";")
		Panic(err)
		//retrieving home columns
		homeColumns, err := rows.Columns()
		Panic(err)
		homeValues := MakeValues(rows, homeColumns, table.Column)
		rows.Close()
		// now we have home datas with primary keys in homeValues

		//try again for away datas...
		rows, err = away.Query("SELECT * FROM " + table.Name + ";")
		Panic(err)
		defer rows.Close()
		awayColumns, err := rows.Columns()
		Panic(err)
		awayValues := MakeValues(rows, awayColumns, table.Column)

		if len(awayColumns) != len(homeColumns) {
			Panic(fmt.Errorf("Incorrect length of tables , They aren't same !"))
		}

		stmt, err := away.Prepare(MakeStmt(awayColumns, table.Name))
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

}
