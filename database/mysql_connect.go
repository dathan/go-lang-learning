package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

//
//connect to a database
// grab a row
// convert that row to a map of strings where each column name == the string value of the column
//
func main() {
	// connect to the db
	db, err := sql.Open("mysql", "root:@tcp(192.168.56.101:3306)/UserShard1?charset=utf8mb4")
	if err != nil {
		log.Fatal(err)
	}

	// grab the data into a stmt handler
	stmt, err := db.Prepare("SELECT * FROM UserInfo WHERE uid = ?")
	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	//actually get the data from the DB (execute the prepare)
	rows, err := stmt.Query(2)
	if err != nil {
		log.Fatal(err)
	}

	//clean up
	defer rows.Close()
	defer db.Close()

	// get the col names
	cols, err := rows.Columns()
	if err != nil {
		log.Fatal(err)
	}

	// make a generic map to store the desired result
	row_map := make(map[string]interface{})

	//grab the raw bytes because we are doing "SELECT * FROM"
	row_values := make([]sql.RawBytes, len(cols))

	//rows.Scan wants []interface{} as an argument so we have to copy the references into such a slice
	//https://github.com/go-sql-driver/mysql/wiki/Examples
	//http://code.google.com/p/go-wiki/wiki/InterfaceSlice
	scanArgs := make([]interface{}, len(row_values))
	for i := range row_values {
		scanArgs[i] = &row_values[i]
	}

	// fetch the row
	rows.Next()
	// add the entire scanArgs to the Scan "method" via ... syntax
	err = rows.Scan(scanArgs...)
	if err != nil {
		log.Fatal(err)
	}

	// copy the data into our desired map
	for i, col := range cols {
		row_map[col] = string(row_values[i])
	}

	fmt.Println("Rows: ", row_map)

}
