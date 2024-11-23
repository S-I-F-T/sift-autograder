package gradables

import (
	"encoding/json"
)

// Defining a StringArray struct to hold a slice of strings
type StringArray struct {
	strings []string
}

// Constructor function to create a new StringArray with a single string element.
func StringArrayFromString(s string) *StringArray {
	return &StringArray{[]string{s}}
}

// Constructor function that initializes and returns a pointer to a StringArray struct.
func NewStringArray(s []string) *StringArray {
	return &StringArray{s}
}

// Getter function for slice of strings
func (sa *StringArray) Strings() []string {
	return sa.strings
}

// JSON unmarshaling method to handle both single string and array of strings formats.
func (sa *StringArray) UnmarshalJSON(data []byte) error {
	if string(data) == "" || string(data) == "null" {
		return nil
	}

	var arr []string
	if err := json.Unmarshal(data, &arr); err == nil {
		sa.strings = append(sa.strings, arr...)
		return nil
	}

	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}

	sa.strings = append(sa.strings, str)
	return nil
}
