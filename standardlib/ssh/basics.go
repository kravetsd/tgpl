package ssh

import (
	"bytes"
	"fmt"
	"log"
	"net"

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

	hostKeyCallback, err := knownhosts.New("/home/dkravets/.ssh/known_hosts")
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
	var b bytes.Buffer
	session.Stdout = &b

	if err := session.Run(command); err != nil {
		log.Fatal("Failed to run: " + err.Error())
	}
	fmt.Println(b.String())
}
