package gradables

import (
	"encoding/json"
)

type StringArray struct {
	strings []string
}

func StringArrayFromString(s string) *StringArray {
	return &StringArray{[]string{s}}
}

func NewStringArray(s []string) *StringArray {
	return &StringArray{s}
}

func (sa *StringArray) Strings() []string {
	return sa.strings
}

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
