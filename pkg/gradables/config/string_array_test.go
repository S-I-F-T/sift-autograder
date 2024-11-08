package gradables

import (
	"encoding/json"
	"log"
	"reflect"
	"testing"
)

/* Validating umarshalling of json file to
string array. */
func TestStringArray_UnmarshalJSON(t *testing.T) {

	const tests = 2

	inputs := [tests]string{
		`"some string"`,
		`"[\"this\",\"is\",\"a\",\"string\",\"not\",\"array\"]"`,
	}

	expected := [tests]StringArray{
		*StringArrayFromString("some string"),
		*StringArrayFromString(`["this","is","a","string","not","array"]`),
	}

	for idx, input := range inputs {
		var actual StringArray

		if err := json.Unmarshal([]byte(input), &actual); err != nil {
			log.Fatalf("Unmarshalling error: %v\n", err)
		}

		if !reflect.DeepEqual(actual, expected[idx]) {
			log.Fatalf("Expected %v, actually %v", expected[idx], actual)
		}
	}
}
