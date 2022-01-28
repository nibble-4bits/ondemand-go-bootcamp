package adapter

// DataStore represents some sort of storage where data can be written to
// but not read.
//
// Examples of data stores could be: a database table, a file or even an API.
type DataStore interface {
	SaveRecord([]string) error
}
