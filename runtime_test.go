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
//===========================================
	type Salary struct{
		Basic, HRA, TA float64
	}

	type Employee struct{
		FirstName, LastName, Email string
		Age int
		MonthlySalary []Salary
	}

	e := Employee{
		FirstName: "Mark",
		LastName: "Jones",
		Email: "mark@gmail.com",
		Age: 25,
		MonthlySalary: []Salary{
			Salary{
				Basic:15000.00,
				HRA:5000.00,
				TA:2000.00,
			},
			Salary{
				Basic:16000.00,
				HRA:5000.00,
				TA:2100.00,
			},
			Salary{
				Basic:17000.00,
				HRA:5000.00,
				TA:2200.00,
			},
		},
	}

	PrettyJson(e,"Mark Jones")
}