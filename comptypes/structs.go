package comptypes

import (
	"fmt"
	"time"
)

func Structs() {
	fmt.Println("hello structs!")
	var dilbert Employee
	dilbert.Salary = 5000
	position := &dilbert.Position
	*position = "Senior " + *position

	var employeeOfTheMonth *Employee = &dilbert
	employeeOfTheMonth.Position += "  (proactive team player)"

}

type Employee struct {
	ID      int
	Name    string
	Address string
	DoB     time.Time

	Position  string
	Salary    int
	ManagerId int
}
