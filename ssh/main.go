package main

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"golang.org/x/term"
	"io/ioutil"
	"log"
	"strings"
	"syscall"
)

func main() {
	opts, err := ParseFlags()
	if err != nil {
		log.Fatalln(err)
	}

	var authMethods []ssh.AuthMethod
	if opts.Key != nil {
		authMethod, err := readPublicKey(opts.Key)
		if err != nil {
			log.Fatalln(err)
		}
		authMethods = append(authMethods, authMethod)
	} else {
		authMethod, err := getUserPassword()
		if err != nil {
			log.Fatalln(err)
		}
		authMethods = append(authMethods, authMethod)
	}

	sshConfig := &ssh.ClientConfig{
		User:            opts.User,
		Auth:            authMethods,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	host := opts.Server
	if strings.Contains(host, ":") == false {
		host = fmt.Sprintf("%s:22", host)
	}
	client, err := ssh.Dial("tcp", host, sshConfig)
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()

	session, err := client.NewSession()
	if err != nil {
		log.Fatalln(err)
	}
	defer session.Close()

	output, err := session.CombinedOutput("ls -l")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(output))
}

func getUserPassword() (ssh.AuthMethod, error) {
	fmt.Print("Enter Password: ")
	pass, err := term.ReadPassword(syscall.Stdin)
	if err != nil {
		return nil, err
	}
	fmt.Println()
	return ssh.Password(string(pass)), nil
}

func readPublicKey(key *string) (ssh.AuthMethod, error) {
	file, err := ioutil.ReadFile(*key)
	if err != nil {
		return nil, err
	}
	privateKey, err := ssh.ParsePrivateKey(file)
	if err != nil {
		return nil, err
	}

	return ssh.PublicKeys(privateKey), nil
}
