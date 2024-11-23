package gradables

import (
	"encoding/json"
	"log"
	"reflect"
	"testing"
)

//
// Validating umarshaling of json file to string array.
//
func TestStringArray_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		Input    string
		Expected StringArray
	}{
		{
			`"some string"`,
			*StringArrayFromString("some string"),
		},
		{
			`["ok", "now", "this", "is", "an", "array"]`,
			*NewStringArray([]string{"ok", "now", "this", "is", "an", "array"}),
		},
		{
			`"[\"this\",\"is\",\"a\",\"string\",\"not\",\"array\"]"`,
			*StringArrayFromString(`["this","is","a","string","not","array"]`),
		},
	}

	for _, tc := range tests {
		var actual StringArray

		if err := json.Unmarshal([]byte(tc.Input), &actual); err != nil {
			log.Fatalf("Input %v: Unmarshalling error: %v\n", tc.Input, err)
		}

		if !reflect.DeepEqual(actual, tc.Expected) {
			log.Fatalf("Input %v: expected %v, actually %v", tc.Input, tc.Expected, actual)
		}
	}
}

func TestStringArray_UnmarshalStruct(t *testing.T) {

	type Wrapper struct {
		Array StringArray
	}

	tests := []struct {
		Input    string
		Expected Wrapper
	}{
		{
			`{"Array":"hi"}`,
			Wrapper{*StringArrayFromString("hi")},
		},
		{
			`{"Array":"[\"this\",\"is\",\"a\",\"string\",\"not\",\"array\"]"}`,
			Wrapper{*StringArrayFromString(`["this","is","a","string","not","array"]`)},
		},
	}

	for idx, tc := range tests {
		var actual Wrapper

		if err := json.Unmarshal([]byte(tc.Input), &actual); err != nil {
			log.Fatalf("Test %d failed. Input: %v. Unmarshalling error: %v\n", idx, tc.Input, err)
		}

		if !reflect.DeepEqual(actual, tc.Expected) {
			log.Fatalf("Test %d failed. Input: %v. Expected %v (%d). Actually %v (%d)\n",
				idx, tc.Input, tc.Expected, len(tc.Expected.Array.strings), actual, len(actual.Array.strings))
		}
	}
}
