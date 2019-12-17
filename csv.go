package csvmgr

import (
	"encoding/csv"
	"log"
	"os"
	"reflect"
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

	if c.Record[field] == "" {
		return 0
	}
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

//GetFloat 获取Float类型的数据
func (c *CsvRecord) GetFloat(field string) float64 {
	var r float64
	var err error

	if c.Record[field] == "" {
		return float64(0)
	}

	if r, err = strconv.ParseFloat(c.Record[field], 64); err != nil {
		log.Fatalln(err.Error())
	}
	return r
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

//SaveCsvCfg 导出CSV文件
func SaveCsvCfg(table []map[string]interface{}, filename string) {
	rowCount := len(table)

	// 循环遍历传入的数据数组,取出key作为表格头
	if rowCount > 0 {
		// 创建csv文件
		file, err := os.Create(filename)
		if err != nil {
			log.Fatalln(err.Error())
		}
		defer file.Close()

		// 创建csv writer
		cw := csv.NewWriter(file)

		var wKey []string

		for key := range table[0] {
			wKey = append(wKey, key)
		}

		// 写入表头
		cw.Write(wKey)
		cw.Flush()

		columnCount := len(wKey)

		for _, item := range table {
			var wData []string
			// 按表头分列数据
			for index := 0; index < columnCount; index++ {
				wTmpData := item[wKey[index]]
				switch reflect.TypeOf(wTmpData).String() {
				case "bool":
					boolData := strconv.FormatBool(wTmpData.(bool))
					wData = append(wData, boolData)
				case "int":
					intData := strconv.Itoa(wTmpData.(int))
					wData = append(wData, intData)
				case "float64":
					floatData := strconv.FormatFloat(wTmpData.(float64), 'f', -1, 64)
					wData = append(wData, floatData)
				default:
					wData = append(wData, wTmpData.(string))
				}
			}

			// 写入每行数据
			cw.Write(wData)
		}
		cw.Flush()
	}
}
