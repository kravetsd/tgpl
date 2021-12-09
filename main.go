package main

import (
	"flag"
	"github.com/kravetsd/TheGoProgrammingLanguage/standardlib/ssh"
	"io/ioutil"
	"log"
)

const (
	url string = "https://golang.org"
)

func main() {

	user := flag.String("user", "", "specify a user")
	remoteAddress := flag.String("ip", "", "specify an ip address to ssh")
	remotePort := flag.String("port", "22", "specify a port for ssh")
	privateKeyPath := flag.String("path", "", "specify a port for ssh")

	flag.Parse()

	key, err := ioutil.ReadFile(*privateKeyPath)
	if err != nil {
		log.Fatalf("Open private key file : %s", err)
	}

	ssh.Sshbasic(*user, *remoteAddress, *remotePort, key)
}
