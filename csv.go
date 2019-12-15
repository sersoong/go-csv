package csvmgr

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
)

//CsvTable 定义Table结构
type CsvTable struct {
	FileName string
	Records  []CsvRecord
}

//CsvRecord 定义数据条目结构
type CsvRecord struct {
	Record map[string]string
}

//GetInt 获取Int类型的数据
func (c *CsvRecord) GetInt(field string) int {
	var r int
	var err error

	if r, err = strconv.Atoi(c.Record[field]); err != nil {
		log.Fatalln(err.Error())
	}
	return r
}

//GetString 获取string类型的数据
func (c *CsvRecord) GetString(field string) string {
	data, ok := c.Record[field]
	if ok {
		return data
	}
	log.SetPrefix("Warning")
	log.Println("Get fileld failed! field:", field)
	log.Println("Wrong ret Data is :", data)
	return ""
}

//GetBool 获取bool类型的数据
func (c *CsvRecord) GetBool(field string) bool {
	var ret bool
	var err error
	if ret, err = strconv.ParseBool(c.Record[field]); err != nil {
		log.Fatalln(err.Error())
	}
	return ret
}

//LoadCsvCfg 载入csv文件
func LoadCsvCfg(filename string, row int) *CsvTable {
	file, err := os.Open(filename)
	if err != nil {
		log.SetPrefix("Error")
		log.Println(err.Error())
	}
	defer file.Close()

	reader := csv.NewReader(file)

	if reader == nil {
		log.SetPrefix("Error")
		log.Println("NewReader return nil, file:", file)
		return nil
	}

	records, err := reader.ReadAll()
	if err != nil {
		log.SetPrefix("Error")
		log.Println(err.Error())
		return nil
	}

	if len(records) < row {
		log.SetPrefix("Warning")
		log.Println(filename, " is empty")
		return nil
	}

	colNum := len(records[0])
	recordNum := len(records)

	var allRecords []CsvRecord
	for i := row; i < recordNum; i++ {
		record := &CsvRecord{make(map[string]string)}
		for k := 0; k < colNum; k++ {
			record.Record[records[0][k]] = records[i][k]
		}
		allRecords = append(allRecords, *record)
	}
	var result = &CsvTable{
		filename,
		allRecords,
	}
	return result
}
