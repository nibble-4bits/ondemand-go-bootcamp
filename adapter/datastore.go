package adapter

// DataStore represents some sort of storage where data can be written to
// but not read.
//
// Examples of data stores could be: a database table, a file or even an API.
type DataStore interface {
	// SaveRecord saves a single record, represented as a slice of strings,
	// to the underlying data storage.
	SaveRecord([]string) error
}
