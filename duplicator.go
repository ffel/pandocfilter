package pandocfilter

// Duplicator duplicates the input json into output json
//
// Duplicator is the most basic filter possible.
// It makes sense to embed the Duplicator in more advanced filters.
type Duplicator struct {
}

func (d Duplicator) Value(key string, value interface{}) (bool, interface{}) {
	return true, value
}
