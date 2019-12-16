package main

import (
	"encoding/json"
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

//RespData RespData
type RespData struct {
	Data []map[string]interface{} `json:"data"`
}

var jsondata = `{"data":[{"id":0,"name":"张三","value":"aaaaa"},{"id":1,"name":"李四","value":"bbbbb"}]}`
var gTestTable map[int]*TestTable

func main() {
	LoadTestTable()

	var ret RespData
	json.Unmarshal([]byte(jsondata), &ret)

	csvmgr.SaveCsvCfg(ret.Data, "./out.csv")
	for _, item := range gTestTable {
		fmt.Println(item)
	}
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
