package w2

// pandoc types
const (
	Header = "Header" // section header
	Para   = "Para"   // paragraph
	Space  = "Space"  // white space
	Str    = "Str"    // word (possibly with some interpunction)
)

// CString checks whether value is a tc object with c as a string
func CString(value interface{}) (bool, string, string) {
	isTC, t, c := IsTypeContents(value)
	if !isTC {
		return false, "", ""
	}
	cstr, cok := c.(string)
	if !cok {
		return false, t, ""
	}
	return true, t, cstr
}

// WrapCString wraps a string value as c in a tc object
func WrapCString(t, c string) interface{} {
	return map[string]interface{}{"t": t, "c": c}
}

// isContentsType checks for typical pandoc ct object
// and returns type and contents in case it is
func IsTypeContents(value interface{}) (bool, string, interface{}) {
	set, isSet := value.(map[string]interface{})
	if !isSet {
		return false, "", nil
	}
	if len(set) != 2 {
		return false, "", nil
	}
	typeval, tok := set["t"]
	if !tok {
		return false, "", nil
	}
	str, isStr := typeval.(string)
	if !isStr {
		return false, "", nil
	}
	contents, cok := set["c"]
	if !cok {
		return false, "", nil
	}
	return true, str, contents
}
