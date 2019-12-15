package main

import (
	"fmt"

	csvmgr "github.com/sersoong/go-csv"
)

type TestTable struct {
	Id    int
	Type  int
	Value int
	Text  string
}

var gTestTable map[int]*TestTable

func main() {
	LoadTestTable()
	fmt.Print(gTestTable[1])
}

//LoadTestTable 载入测试数据表
func LoadTestTable() bool {
	var result = csvmgr.LoadCsvCfg("./data.csv", 1)
	if result == nil {
		return false
	}
	gTestTable = make(map[int]*TestTable)

	for _, record := range result.Records {
		id := record.GetInt("id")
		item := &TestTable{
			id,
			record.GetInt("type"),
			record.GetInt("value"),
			record.GetString("text"),
		}
		gTestTable[id] = item
	}
	return true
}
