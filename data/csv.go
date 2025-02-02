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
// Each record is a slice of strings itself.
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

	// Remove header from slice of records
	records = records[1:]

	return records, nil
}

// csvDataStore represents a data store that saves to a CSV file.
type csvDataStore struct {
	// store is a file path to a CSV file.
	store string
}

// NewCSVDataStore receives a path to a CSV file and returns an instance of csvDataStore.
func NewCSVDataStore(csvPath string) csvDataStore {
	return csvDataStore{store: csvPath}
}

// SaveRecord saves a record to the CSV file associated with the csvDataStore instance.
//
// If the save is successful, a nil error is returned.
func (c csvDataStore) SaveRecord(record []string) error {
	file, err := os.OpenFile(c.store, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return err
	}
	defer file.Close()

	csvWriter := csv.NewWriter(file)
	err = csvWriter.Write(record)
	if err != nil {
		return err
	}

	csvWriter.Flush()

	err = csvWriter.Error()
	if err != nil {
		return err
	}

	return nil
}
