package main

import (
	"fmt"

	csvmgr "github.com/sersoong/go-csv"
)

//TestTable 测试表结构
type TestTable struct {
	Name  string
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

	for index, record := range result.Records {
		item := &TestTable{
			record.GetString("name"),
			record.GetInt("type"),
			record.GetInt("value"),
			record.GetString("text"),
		}
		gTestTable[index] = item
	}
	return true
}
