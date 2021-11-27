package main

import (
	"fmt"

	"github.com/kravetsd/TheGoProgrammingLanguage/standardlib"
)

const (
	url string = "https://golang.org"
)

func main() {
	fmt.Println("Hello module")
	//functions.Hellofunc()
	//functions.Hypot()
	//fmt.Println(functions.HypotRet(39, 30, 1, "hello", "world", 0.23))

	//functions.Parsehtml(functions.Fetch("https://golang.org")) it works!

	//functions.Outlinecaller(functions.Fetch("https://golang.org")) it works

	//functions.ParseFindlinks(url) It works

	//fmt.Println(functions.SimpleHttpGet("https://golang.org")) it works

	//errmes := functions.Waitforserver("https://adsafasdfasfg.org")
	//fmt.Println(errmes.Error())
	//basictypes.Strs()

	//comptypes.StructCompare(comptypes.P, comptypes.Q)

	//comptypes.Embstructs()

	//comptypes.JsonBasics()

	// issuetrack testing:
	// result, err := github.SearchIssues(os.Args[1:])

	// if err != nil {
	// 	log.Fatalf("Search Issue main failed: %s ", err)
	// 	os.Exit(1)
	// }

	// fmt.Printf("%d issues:\n", result.TotalCount)
	// for _, item := range result.Items {
	// 	fmt.Printf("#%-5d %9.9s %.55s\n",
	// 		item.Number, item.User.Login, item.Title)
	// }

	//interfaces.Shopping()
	envvars := standardlib.EnvVarToMap()
	fmt.Println(envvars["Path"])
}
