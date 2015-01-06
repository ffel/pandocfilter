package pandocfilter

import "testing"

func TestGetString(t *testing.T) {
	data := Jmap{
		"c": Jslice{
			Jmap{
				"c": Jslice{},
				"t": "deepT",
			},
			"lastC",
		},
		"t": "shallowT",
	}

	val, err := GetString(data, "t")

	if err != nil {
		t.Errorf("unexpected error %v\n", err)
	} else if val != "shallowT" {
		t.Errorf("unexpected value %q", val)
	}

	val, err = GetString(data, "c", "1")

	if err != nil {
		t.Errorf("unexpected error %v\n", err)
	} else if val != "lastC" {
		t.Errorf("unexpected value %q", val)
	}

	val, err = GetString(data, "c", "0", "t")

	if err != nil {
		t.Errorf("unexpected error %v\n", err)
	} else if val != "deepT" {
		t.Errorf("unexpected value %q", val)
	}

}

func TestGetString_errors(t *testing.T) {
	data := Jmap{
		"c": Jslice{
			Jmap{
				"c": Jslice{},
				"t": "deepT",
			},
			"lastC",
		},
		"t": true,
	}

	_, err := GetString(data, "0")

	if err.Error() != "GetString error - no slice for index 0" {
		t.Errorf("unexpected error %v", err)
	}

	_, err = GetString(data, "c", "-1")

	if err.Error() != "GetString error - slice out-of-range index -1" {
		t.Errorf("unexpected error %v", err)
	}

	_, err = GetString(data, "c", "2")

	if err.Error() != "GetString error - slice out-of-range index 2" {
		t.Errorf("unexpected error %v", err)
	}

	_, err = GetString(data, "c", "t")

	if err.Error() != "GetString error - no map for index t" {
		t.Errorf("unexpected error %v", err)
	}

	_, err = GetString(data, "tc")

	if err.Error() != "GetString error reading map - no such key tc" {
		t.Errorf("unexpected error %v", err)
	}

	_, err = GetString(data, "t")

	if err.Error() != "GetString error - no string" {
		t.Errorf("unexpected error %v", err)
	}

}
