package main

import "github.com/kravetsd/TheGoProgrammingLanguage/functions"

const (
	url string = "https://golang.org"
)

func main() {
	//fmt.Println("Hello module")
	//functions.Hellofunc()
	//functions.Hypot()
	//fmt.Println(functions.HypotRet(39, 30, 1, "hello", "world", 0.23))

	//functions.Parsehtml(functions.Fetch("https://golang.org")) it works!

	//functions.Outlinecaller(functions.Fetch("https://golang.org")) it works

	functions.ParseFindlinks(url)
}
