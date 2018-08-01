package state

import "reflect"

// DataMapper converts data between different systems (vm storage <> golang types)
type DataMapper interface {
	Update(value interface{})
	Get(out interface{})
}

func Open(db Database) *DB {
	if db == nil {
		return nil, errors.New("invalid database")
	}

	return &orm{Database: Database}	
}

func (db *DB) Get(out interface{}) {

} 

func (db *DB) Update(attrs ...interface{}) {

}

type BeforeUpdate interface {
	BeforeUpdate(fieldName string)
}

func Map(obj stateObject, value interface{}) error {
	storageVal := reflect.ValueOf(storage)

	for i := 0; i < storageT.NumField(); i++ {
		// omit empty by default
		if fieldVal := storageVal.Field(i); fieldVal == reflect.Zero(fieldVal.Type()).Interface() {
			continue
		}

		// before update
	}
}
