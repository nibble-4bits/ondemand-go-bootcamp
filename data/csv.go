package data

import (
	"encoding/csv"
	"log"
	"os"
)

type csvDataSource struct {
	// Collection is a file path to a CSV file
	Collection string
}

func NewCSVDataSource(csvPath string) csvDataSource {
	return csvDataSource{Collection: csvPath}
}

func (ds csvDataSource) ReadCollection() [][]string {
	file, err := os.Open(ds.Collection)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatalln(err)
	}

	return records
}
