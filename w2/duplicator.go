package w2

// Duplicator duplicates the input json into output json
//
// Duplicator is the most basic filter possible.
// It makes sense to embed the Duplicator in more advanced filters.
type Duplicator struct {
}

func (d Duplicator) List(key string, value []interface{}) (bool, interface{}) {
	return true, nil
}

func (d Duplicator) Set(key string, value map[string]interface{}) (bool, interface{}) {
	return true, nil
}

func (d Duplicator) Value(key string, value interface{}) interface{} {
	return value
}
