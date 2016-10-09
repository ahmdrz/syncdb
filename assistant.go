package main

import (
	"database/sql"
	"fmt"
	"log"
)

func Panic(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func MakeStmt(columns []string) string {
	query := "INSERT INTO " + selected_table + " ("
	for i := range columns {
		query += columns[i] + ","
	}
	query = query[:len(query)-1] + ") VALUES ("
	for i := 0; i < len(columns); i++ {
		query += "?,"
	}
	query = query[:len(query)-1] + ") "

	return query
}

type Row struct {
	Primary_Value interface{}
	Values        []interface{}
}

func MakeValues(rows *sql.Rows, columns []string) []Row {
	count := len(columns)
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)

	returnValues := make([]Row, 0)

	for rows.Next() {
		for i, _ := range columns {
			valuePtrs[i] = &values[i]
		}

		rows.Scan(valuePtrs...)

		tmp := make([]interface{}, count)
		row := Row{}
		row.Primary_Value = nil

		for i := range columns {
			val := values[i]

			if columns[i] == selected_column {
				row.Primary_Value = val
			}

			tmp[i] = val
		}

		if row.Primary_Value == nil {
			Panic(fmt.Errorf("Can't found primary_value , check selected_column"))
		}

		row.Values = tmp
		returnValues = append(returnValues, row)
	}
	return returnValues
}
