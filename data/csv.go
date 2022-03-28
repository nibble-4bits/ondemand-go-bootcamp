package data

import (
	"encoding/csv"
	"os"
)

// csvDataSource represents a data source that reads from a CSV file.
type csvDataSource struct {
	// collection is a file path to a CSV file.
	collection string
}

// NewCSVDataSource receives a path to a CSV file and returns an instance of csvDataSource.
func NewCSVDataSource(csvPath string) csvDataSource {
	return csvDataSource{collection: csvPath}
}

// ReadCollection reads the CSV file associated with the csvDataSource instance
// and returns an slice of records.
// Each record is an slice of strings itself.
func (c csvDataSource) ReadCollection() ([][]string, error) {
	file, err := os.Open(c.collection)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	records, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}

	return records, nil
}
