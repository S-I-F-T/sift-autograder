package gradables

import (
	"encoding/json"
	"testing"
)

func TestConfigJSON(t *testing.T) {

	inputs := []string{
		"{\"testcases\":[]}",
	}

	expected := []Config{
		{},
	}

	for i, _ := range inputs {
		var actual Config
		if json.Unmarshal([]byte(inputs[i]), &actual) != nil {
			t.Errorf("Unmarshalling error for JSON '%s'", inputs[i])
		}

		if !equal(actual, expected[i]) {
			t.Errorf("Expected %v, actually %v", expected[i], actual)
		}
	}

}

func equal(c1, c2 Config) bool {
	return false
}
