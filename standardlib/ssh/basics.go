package ssh

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/knownhosts"
)

func Sshbasic(user string, remoteAddress, remotePort string, command string, key []byte) {

	// FROM THE DOCUMETNATION PAGE:
	// An SSH client is represented with a ClientConn.
	//
	// To authenticate with the remote server you must pass at least one
	// implementation of AuthMethod via the Auth field in ClientConfig,
	// and provide a HostKeyCallback.

	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		log.Fatal("Error with parsing private key : ", err)
	}

	hostKeyCallback, err := knownhosts.New(os.Getenv("HOME") + "/.ssh/known_hosts")
	if err != nil {
		log.Fatalf("known hostst, hostkeycallback error : %s", err)
	}

	config := &ssh.ClientConfig{
		Config: ssh.Config{},
		User:   user,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: hostKeyCallback,
	}

	// Implementing a new ssh client connection:

	// setting up a connection to temote host
	c, err := net.Dial("tcp", fmt.Sprintf(remoteAddress+":"+remotePort))
	if err != nil {
		log.Fatalf("Can't establich tcp connection to %s : %s", remoteAddress, err)
	}

	// establishing connection:
	connection, newchan, requestchan, err := ssh.NewClientConn(c, fmt.Sprintf(remoteAddress+":"+remotePort), config)
	if err != nil {
		log.Fatalf("Error to establich ssh connection to %s : %s", remoteAddress, err)
	}
	client := ssh.NewClient(connection, newchan, requestchan)
	session, err := client.NewSession()
	if err != nil {
		log.Fatalf("Establishing SSH session error: %s", err)
	}
	defer session.Close()

	// once session is created we can execute command:

	// creating a buffer for session's output storage
	//var b bytes.Buffer
	//session.Stdout = &b
	//
	//if err := session.Run(command); err != nil {
	//	log.Fatal("Failed to run: " + err.Error())
	//}
	//fmt.Println(b.String())

	// Associating stdin,out and err with remote session process
	stdin, err := session.StdinPipe()
	if err != nil {
		log.Fatalf("StdinPipe error: %s", err)
	}
	stdout, err := session.StdoutPipe()
	if err != nil {
		log.Fatalf("StdoutPipe error: %s", err)
	}
	stderr, err := session.StderrPipe()
	if err != nil {
		log.Fatalf("StderrPipe error: %s", err)
	}

	if err := session.Shell(); err != nil {
		log.Fatal("Failed to run: " + err.Error())
	}

	// creating channel:
	//wr := make(chan []byte, 10)

	go func() {
		for {
			out := bufio.NewScanner(stdout)
			out.Scan()
			fmt.Println(out.Text())
		}
	}()

	go func() {
		for {
			out := bufio.NewScanner(stderr)
			out.Scan()
			fmt.Println(out.Text())
		}
	}()

	for {
		in := bufio.NewScanner(os.Stdin)
		in.Scan()
		stdin.Write([]byte(in.Text() + "\n"))
	}

	//if err := session.Run(command); err != nil {
	//
	//}
}
