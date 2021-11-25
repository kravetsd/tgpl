package comptypes

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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

func InputPlatformData() {

	scanner := bufio.NewReader(os.Stdin)

	fmt.Println("The platform name is: ")
	pltf_name, _ := scanner.ReadString('\n')
	pltf_name = strings.TrimSuffix(pltf_name, "\r\n")

	fmt.Println("Region is: ")
	region, _ := scanner.ReadString('\n')
	region = strings.TrimSuffix(region, "\r\n")

	fmt.Println("environment: ")
	env, _ := scanner.ReadString('\n')
	env = strings.TrimSuffix(env, "\r\n")

	Platformp := GetPlatform(pltf_name, region, env)

	PlatformRepoName(Platformp)

}

var P = Platform{Name: "name1", Region: "ny", Environment: "dev"}
var Q = Platform{Name: "name3", Region: "ny", Environment: "dev"}

func StructCompare(p, q Platform) {
	fmt.Println("Comparing structs with '==' sign p == q : ", p == q)
	fmt.Println("Comparing structs field by field p.Name == q.Name : ", q.Name == p.Name && q.Environment == p.Environment && q.Region == p.Region)

	// struct as a key of a map:
	structmap := make(map[Platform]int)
	for i := 0; i < 10; i++ {
		structmap[Platform{"here", "I", "dev"}]++
	}
	fmt.Println(structmap)
}
