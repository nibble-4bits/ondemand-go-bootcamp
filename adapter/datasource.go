package adapter

// DataSource represents a data source that can read a collection of items
//
// Examples of "collections" could be: a database table, a CSV file, a JSON array
type DataSource interface {
	// ReadCollection reads the collection that belongs to the current DataSource and returns
	// the items as a slice of slices of strings
	ReadCollection() ([][]string, error)
}
