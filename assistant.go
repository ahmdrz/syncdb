package main

import (
	"database/sql"
	"log"
	"strconv"
	_ "time"
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

		for i := range columns {
			val := values[i]

			/*temp_val := ""

			switch val.(type) {
			case int64:
				temp_val = ToString64(val.(int64))
			case string:
				temp_val = val.(string)
			case time.Time:
				temp_val = val.(time.Time).Format("2006-01-02 15:04:05")
			case []uint8:
				temp_val = ToStringUInt8Array(val.([]uint8))
			case bool:
				if val.(bool) {
					temp_val = "1"
				} else {
					temp_val = "0"
				}
			case int32:
				temp_val = ToString32(val.(int32))
			case int:
				temp_val = ToString(val.(int))
			default:
				temp_val = "''"
			}*/

			if columns[i] == selected_column {
				row.Primary_Value = val
			}

			tmp[i] = val
		}

		row.Values = tmp
		returnValues = append(returnValues, row)
	}
	return returnValues
}

func ToString(a int) string {
	return strconv.Itoa(a)
}

func ToString32(a int32) string {
	return strconv.Itoa(int(a))
}

func ToString64(a int64) string {
	return strconv.FormatInt(a, 10)
}

func ToStringUInt8Array(a []uint8) string {
	b := make([]byte, len(a))
	for i, v := range a {
		b[i] = byte(v)
	}
	return string(b)
}
