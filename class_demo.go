package main

import (
	"fmt"
)

// Employee is an employee class
type Employee struct {
	FirstName   string
	LastName    string
	TotalLeaves int
	LeavesTaken int
}

// LeavesRemaining gives no. of total leaves balance
func (e Employee) LeavesRemaining() {
	fmt.Printf("%s %s has %d leaves remaining\n",
		e.FirstName, e.LastName, (e.TotalLeaves - e.LeavesTaken))
}

func main() {
	e := Employee{
		FirstName:   "Foo",
		LastName:    "Bar",
		TotalLeaves: 30,
		LeavesTaken: 15,
	}
	e.LeavesRemaining()
}
