package functions

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func Hellofunc() {
	fmt.Println("Hello from functions")
}

func Hypot() {
	fmt.Println(`This is a simple hypot function.\n 
	"hypot". The function itself doesn't returns anything and is needed only for making some actions.`)
}

func HypotRet(a, b, c int, st, ring string, _ float64) (z string) {
	fmt.Printf("a=%v, b=%v, c=%v\n", a, b, c)
	z = string(a + b + c)
	return z
}

func Parsehtml(url string) {
	doc, err := html.Parse(strings.NewReader(url))
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks: %v \n", err)
		os.Exit(1)
	}

	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}

func Fetch(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		os.Exit(1)
	}
	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
		os.Exit(1)
	}
	return string(b)
}
