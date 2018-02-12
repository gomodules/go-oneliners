package oneliners

import (
	"encoding/json"
	"testing"
)

func TestDumpJson(t *testing.T) {
	str := map[string]interface{}{
		"Name": "Wednesday",
		"Age":  6,
		"Parents": []interface{}{
			"Gomez",
			"Morticia",
		},
	}
	PrettyJson(str, "Asad")
	PrettyJson(str, "Asad", "jacob")

	// Pretty Print Marshalled Object
	if jsn, err := json.Marshal(str); err == nil {
		PrettyJson(jsn, "Marshalled JSON")
	}
}

func TestPrintStacktrace(t *testing.T) {
	PrintStacktrace()
}
