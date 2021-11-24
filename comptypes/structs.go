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

type Platform struct {
	Name        string
	Region      string
	Environment string
}

func PlatformRepoName(p *Platform) {
	fmt.Printf("pltf-%s-%s-%s", p.Name, p.Region, p.Environment)
}

func GetPlatform(name, region, env string) *Platform {
	p := &Platform{Name: name, Region: region, Environment: env}
	return p
}

func InputPlatformData(pltf_name, region, env string) {

	Platformp := GetPlatform(pltf_name, region, env)

	PlatformRepoName(Platformp)

}
