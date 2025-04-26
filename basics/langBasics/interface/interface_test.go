package main

import "testing"

func TestInterface(t *testing.T) {

	//a genric emp doest maater permane or contract
	assertErrorMsg := func(t testing.TB, emp Emp, want int) {

		t.Helper()
		got := emp.getSalary()

		if got != want {
			t.Errorf("got %d but wanted %d", got, want)
		}

	}

	t.Run("test permanent emp ", func(t *testing.T) {
		emp1 := permanentEmp{
			name:   "rajesh",
			salary: 10000,
			empId:  101,
		}

		assertErrorMsg(t, emp1, 10000)
	})
	t.Run("test contract  emp ", func(t *testing.T) {
		emp2 := contractEmp{
			name:       "rajesh",
			hourlyRate: 20,
		}

		assertErrorMsg(t, emp2, 10000)
	})

}
