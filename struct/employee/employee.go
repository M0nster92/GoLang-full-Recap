package employee

import "fmt"

type Employee struct {
	FirstName   string
	LastName    string
	TotalLeaves int
	LeavesTaken int
}

func (e Employee) LeavesRemainning() {
	fmt.Println("Leaves remainning :  ", e.FirstName, e.LastName, (e.TotalLeaves - e.LeavesTaken))
}
