package gradables

import "encoding/json"

type StringArray struct {
	strings []string
}

func StringArrayFromString(s string) *StringArray {
	return &StringArray{[]string{s}}
}

func (sa *StringArray) Strings() []string {
	return sa.strings
}

func (sa *StringArray) UnmarshalJSON(data []byte) error {
	if string(data) == "" || string(data) == "null" {
		return nil
	}

	var s string
	if err := json.Unmarshal(data, &s); err == nil {
		sa.strings = append(sa.strings, s)
		return nil
	}

	return json.Unmarshal(data, &sa.strings)
}
