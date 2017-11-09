package oneliners

import (
	"testing"
)

func TestDumpJson(t *testing.T) {
	str:= map[string]interface{}{
		"Name": "Wednesday",
		"Age":  6,
		"Parents": []interface{}{
			"Gomez",
			"Morticia",
		},
	}
	PrettyJson(str,"Asad")
	PrettyJson(str,"Asad","jacob")
	PrettyJson(str)
}