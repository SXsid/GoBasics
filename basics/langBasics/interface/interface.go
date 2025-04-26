package main

//rules
/*

1) iterface can't ahve more than 3 method
2) try to make interface as modular as you can it shouldn't vary of struct (we can accomplish that using the type assertion)
3) your struct can be aware of the interface he need to implement ?? why and figure out
*/

type permanentEmp struct {
	name   string
	empId  int
	salary int
}

type contractEmp struct {
	name       string
	hourlyRate int
}

//whole emp interface

type Emp interface {
	getSalary() int
	getName() string
}

func (pEmp permanentEmp) getSalary() int {
	return pEmp.salary
}
func (pEmp permanentEmp) getName() string {
	return pEmp.name
}
func (cEmp contractEmp) getName() string {
	return cEmp.name
}
func (cEmp contractEmp) getSalary() int {
	return cEmp.hourlyRate * 8 * 30
}
