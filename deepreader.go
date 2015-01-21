package pandocfilter

import (
	"errors"
	"strconv"
)

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

// GetString retrieves a sting from a deeply nested json
func GetString(json interface{}, indices ...string) (string, error) {
	// walk all bu

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
