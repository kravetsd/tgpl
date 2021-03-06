package functions

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

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
	z = strconv.Itoa(a + b + c)
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

func Outlinecaller(url string) {
	doc, err := html.Parse(strings.NewReader(url))
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline: %v\n", err)
		os.Exit(1)
	}
	outline(nil, doc)
}

func outline(stack []string, n *html.Node) {
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data) // push tag
		fmt.Println(stack)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline(stack, c)
	}
}

func ParseFindlinks(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "http.Get: %v\n", err)
		os.Exit(1)
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Fprintf(os.Stderr, "statuscode: %v =>  %v\n", url, resp.StatusCode)
		os.Exit(1)
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "response parse error: %v\n", err)
		os.Exit(1)
	}

	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}

	outline(nil, doc)
}

// Function and error hanling in function: two strategies :

//  "PROPAGATE" strategy:

func SimpleHttpGet(url string) (content []string, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	doc, err := html.Parse(resp.Body)
	resp.Body.Close()

	if err != nil {
		return nil, fmt.Errorf("parse %s: %v", url, err)
	}

	return visit(nil, doc), nil
}

// the next strategy for Error handling is a RETRY:

func Waitforserver(url string) error {
	const timeout = 1 * time.Minute
	deadline := time.Now().Add(timeout)

	for tries := 0; time.Now().Before(deadline); tries++ {
		_, err := http.Get(url)
		if err == nil {
			return nil
		}
		fmt.Printf("server not responding: %v; ...retrying:", err)
		time.Sleep(time.Second << uint(tries)) // exponential back-off
	}
	return fmt.Errorf("server %s failed to respond after %s", url, timeout)
}
