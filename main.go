package main

import (
	"fmt"
	//"github.com/kravetsd/TheGoProgrammingLanguage/basictypes"
	//"github.com/kravetsd/TheGoProgrammingLanguage/functions"
	"github.com/kravetsd/TheGoProgrammingLanguage/comptypes"
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

	comptypes.InputPlatformData("ga", "na", "dev")
}
