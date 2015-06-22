package pandocfilter

import (
	"errors"
	"strconv"
)

// GetString retrieves a sting from a deeply nested object
func GetString(json interface{}, indices ...string) (string, error) {
	json, err := getObject(json, indices)

	if err != nil {
		return "", err
	}

	str, strok := json.(string)

	if !strok {
		return "", errors.New("GetString error - no string")
	}

	return str, nil
}

// SetString replaces an existing member with a string
func SetString(json interface{}, value string, indices ...string) error {
	return setObject(json, value, indices)
}

// GetNumber retrieves a number from a deeply nested object
func GetNumber(json interface{}, indices ...string) (float64, error) {
	json, err := getObject(json, indices)

	if err != nil {
		return 0, err
	}

	// know how to deal with an int
	if ival, ivalok := json.(int); ivalok == true {
		return float64(ival), nil
	}

	val, valok := json.(float64)

	if !valok {
		return 0, errors.New("GetNumber error - no number")
	}

	return val, nil
}

// SetNumber replaces an existing member with a number
func SetNumber(json interface{}, value float64, indices ...string) error {
	return setObject(json, value, indices)
}

// GetObject retrieves an object from a deeply nested object
func GetObject(json interface{}, indices ...string) (interface{}, error) {
	return getObject(json, indices)
}

// SetObject replaces an existing member with an object
func SetObject(json interface{}, value interface{}, indices ...string) error {
	return setObject(json, value, indices)
}

// GetSlice retrieves a slice from a deeply nested object
func GetSlice(json interface{}, indices ...string) ([]interface{}, error) {
	json, err := getObject(json, indices)

	if err != nil {
		return make([]interface{}, 0), err
	}

	val, valok := json.([]interface{})

	if !valok {
		return make([]interface{}, 0), errors.New("GetSlice error - no slice")
	}

	return val, nil
}

// SetSlice replaces an existing member with a slice
func SetSlice(json interface{}, value []interface{}, indices ...string) error {
	return setObject(json, value, indices)
}

// GetMap retrieves a map from a deeply nested object
func GetMap(json interface{}, indices ...string) (map[string]interface{}, error) {
	json, err := getObject(json, indices)

	if err != nil {
		return make(map[string]interface{}), err
	}

	val, valok := json.(map[string]interface{})

	if !valok {
		return make(map[string]interface{}, 0), errors.New("GetMap error - no map")
	}

	return val, nil
}

// SetMap replaces an existing member with a map
func SetMap(json interface{}, value map[string]interface{}, indices ...string) error {
	return setObject(json, value, indices)
}

// getObject tries to retrieve a deep object following indices
func getObject(json interface{}, indices []string) (interface{}, error) {
	for _, index := range indices /*[:len(slice)-1]*/ {
		i, err := strconv.Atoi(index)
		if err == nil {
			// json is a slice?
			s, sok := json.([]interface{})

			if !sok {
				return "", errors.New("GetObject error - no slice for index " + index)
			}

			if i < 0 || i >= len(s) {
				return "", errors.New("GetObject error - slice out-of-range index " + index)
			}

			json = s[i]

		} else {
			// json is a map?
			m, mok := json.(map[string]interface{})

			if !mok {
				return "", errors.New("GetObject error - no map for index " + index)
			}

			_, ok := m[index]

			if !ok {
				return "", errors.New("GetObject error reading map - no such key " + index)
			}

			json = m[index]
		}
	}

	return json, nil
}

// setObject replaces a value in-place, thus indices are expected to exist
func setObject(json interface{}, value interface{}, indices []string) error {
	for j, index := range indices {
		i, err := strconv.Atoi(index)
		if err == nil {
			// json is a slice?
			s, sok := json.([]interface{})

			if !sok {
				return errors.New("setObject error - no slice for index " + index)
			}

			if i < 0 || i >= len(s) {
				return errors.New("setObject error - slice out-of-range index " + index)
			}

			if j == len(indices)-1 {
				s[i] = value
				return nil
			}

			json = s[i]

		} else {
			// json is a map?
			m, mok := json.(map[string]interface{})

			if !mok {
				return errors.New("setObject error - no map for index " + index)
			}

			_, ok := m[index]

			if !ok {
				return errors.New("setObject error reading map - no such key " + index)
			}

			if j == len(indices)-1 {
				m[index] = value
				return nil
			}

			json = m[index]
		}
	}

	// you should not reach this ..
	return errors.New("setObject - did not reach level")

}
