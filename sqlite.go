package main

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

var (
	db *sql.DB
)

func GetSQLQuery() string {
	_s := ""
	var end string
	var nstring string = ""
	var orderstring string = "DESC"
	var orderColumnName string
	if subsystem == "urls" {
		orderColumnName = "last_visit_time"
	} else {
		orderColumnName = "last_modified"
	}
	for k := range query.Arr {
		if k != len(query.Arr)-1 {
			end = ","
		} else {
			end = ""
		}
		_s += (query.Arr[k] + end)
	}
	if numberOfResults != -1 {
		if numberOfResults <= 0 {
			PrintC("RED", "Number of Results Should be Greater than or equal to 1")
			os.Exit(-1)
		} else {
			nstring = fmt.Sprintf(" LIMIT %d", numberOfResults)
		}
	}
	if ascending {
		orderstring = "ASC"
	}
	s := fmt.Sprintf("SELECT %s FROM %s ORDER BY %s %s%s", _s, subsystem, orderColumnName, orderstring, nstring)
	return s
}

func GetElement(typ string) interface{} {
	var (
		s *string
		i *int
		b *[]byte
	)
	if strings.Contains(typ, "VARCHAR") {
		_s := string(make([]byte, 0))
		s = &_s
		return s
	} else if typ == "INTEGER" {
		_i := 10
		i = &_i
		return i
	} else if typ == "BLOB" {
		_b := make([]byte, 0)
		b = &_b
		return b
	} else {
		return nil
	}
}

func GetResultInterface() []interface{} {
	var i []interface{}
	m := GetSubsytemInfo()
	if query.Arr[0] == "*" {
		query.Arr = make([]string, 0)
		for k, v := range m {
			query.Arr = append(query.Arr, k)
			i = append(i, GetElement(v))
		}
	} else {
		for _, v := range query.Arr {
			if val, ok := m[v]; !ok {
				PrintC("RED", fmt.Sprintf("No Such Option Available in %s: ", subsystem))
				PrintC("BLUE", v+"\n")
			} else {
				i = append(i, GetElement(val))
			}
		}
	}
	return i
}

func ConnectDatabase() {
	var err error
	if db, err = sql.Open("sqlite3", path); err != nil {
		PrintC("RED", "Error Opening database: ")
		PrintC("BLUE", path+"\n")
		panic(err)
	}
}

func GetSubsytemInfo() map[string]string {
	s := fmt.Sprintf("SELECT name, type FROM pragma_table_info('%s');", subsystem)
	if rows, err := db.Query(s); err != nil {
		PrintC("RED", "Error Executing SQL Query: ")
		PrintC("BLUE", s+"\n\n")
		panic(err)
	} else {
		var (
			options   = map[string]string{}
			name, typ string
		)
		for rows.Next() {
			rows.Scan(&name, &typ)
			options[name] = typ
		}
		return options
	}
}

func GetData() []string {
	i := GetResultInterface()
	s := GetSQLQuery()
	if rows, err := db.Query(s); err != nil {
		PrintC("RED", "Error Executing SQL Query: ")
		PrintC("BLUE", s+"")
		panic(err)
	} else {
		results := []string{}
		for rows.Next() {
			rows.Scan(i...)
			var _result string
			for k := range i {
				var sep string = fieldSepartor
				if k == len(i)-1 {
					sep = ""
				}
				switch i[k].(type) {
				case *int:
					{
						_result += strconv.Itoa(*i[k].(*int)) + sep
					}
				case *string:
					{
						_result += *i[k].(*string) + sep
					}
				case *[]uint8:
					{
						_result += string(*i[k].(*[]uint8)) + sep
					}
				}
			}
			results = append(results, _result)
		}
		return results
	}
}
