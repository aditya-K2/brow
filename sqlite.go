package main

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

var (
	db *sql.DB
)

func GetSQLQuery() string {
	s := fmt.Sprintf("SELECT %%s FROM %s", subsystem)
	s = fmt.Sprintf(s, strings.Repeat("%s ", len(query.Arr)-1)+"%s")
	s = fmt.Sprintf(s, query.Arr...)
	return s
}

func ConnectDatabase() {
	var err error
	if db, err = sql.Open("sqlite3", path); err != nil {
		PrintC("RED", "Error Opening database: ")
		PrintC("BLUE", path+"\n")
		panic(err)
	}
}

func GetSubsytemInfo() {
	s := fmt.Sprintf("SELECT name FROM pragma_table_info('%s');", subsystem)
	if rows, err := db.Query(s); err != nil {
		PrintC("RED", "Error Executing SQL Query: ")
		PrintC("BLUE", s+"\n\n")
		panic(err)
	} else {
		var r string
		for rows.Next() {
			rows.Scan(&r)
			fmt.Println(r)
		}
	}
}

// func GetData() []interface{} {
// 	s := GetSQLQuery()
// 	if rows, err := db.Query(s); err != nil {
// 		PrintC("RED", "Error Executing SQL Query: ")
// 		PrintC("BLUE", s+"")
// 		panic(err)
// 	} else {
// 	}
// }
