package main

import (
	"fmt"

	"github.com/sersoong/go-csv"
)

type TestTable struct {
	Id    int
	Type  int
	Value int
	Text  string
}

var g_allTestTable map[int]*TestTable

func main() {
	LoadTestTable()
	fmt.Print(g_allTestTable[1])
}

func LoadTestTable() bool {
	var result = csvMgr.LoadCsvCfg("./data.csv", 1)
	if result == nil {
		return false
	}
	g_allTestTable = make(map[int]*TestTable)

	for _, record := range result.Records {
		id := record.GetInt("id")
		item := &TestTable{
			id,
			record.GetInt("type"),
			record.GetInt("value"),
			record.GetString("text"),
		}
		g_allTestTable[id] = item
	}
	return true
}
