package csvMgr

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
)

type CsvTable struct {
	FileName string
	Records  []CsvRecord
}

type CsvRecord struct {
	Record map[string]string
}

func (c *CsvRecord) GetInt(field string) int {
	var r int
	var err error

	if r, err = strconv.Atoi(c.Record[field]); err != nil {
		log.Fatalln(err.Error())
	}
	return r
}

func (c *CsvRecord) GetString(field string) string {
	data, ok := c.Record[field]
	if ok {
		return data
	} else {
		log.SetPrefix("Warning")
		log.Println("Get fileld failed! field:", field)
		return ""
	}
}

func (c *CsvRecord) GetBool(field string) bool {
	var ret bool
	var err error
	if ret, err = strconv.ParseBool(c.Record[field]); err != nil {
		log.Fatalln(err.Error())
	}
	return ret
}

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
