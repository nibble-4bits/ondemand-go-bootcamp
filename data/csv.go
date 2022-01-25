package data

import (
	"encoding/csv"
	"log"
	"os"
)

type csvDataSource struct {
	// collection is a file path to a CSV file.
	collection string
}

func NewCSVDataSource(csvPath string) csvDataSource {
	return csvDataSource{collection: csvPath}
}

func (ds csvDataSource) ReadCollection() [][]string {
	file, err := os.Open(ds.collection)
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
