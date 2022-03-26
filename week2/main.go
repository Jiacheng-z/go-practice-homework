package main

import (
	"database/sql"
	"log"
)
/*
sql.ErrNoRows表示空结果集，可能是没有查找到结果，不需要返回给上层，业务认为正常且没有查到数据即可
*/

var db *sql.DB

func init() {
	d, err := sql.Open("mysql", "root:123@tcp(127.0.0.1:3305)/test")
	if err != nil {
		panic("main: mysql connect " + err.Error())
	}
	db = d
}

func main() {
	log.Println(GetData())
}


func GetData() (interface{}, error){
	var dst interface{}
	err := db.QueryRow("SELECT * FROM test WHERE id = ?", 1).Scan(dst)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return dst, nil
}
