package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/codegangsta/cli"
)

func login(c *cli.Context) {
	email := c.String("email")
	password := c.String("password")
	//host := c.String("bazooka-uri")

	if len(email) == 0 {
		email = interactiveInput("Email")
	}

	if len(password) == 0 {
		password = interactiveInput("Password")
	}

	_, err := NewClient(c.String("bazooka-uri"))
	if err != nil {
		log.Fatal(err)
	}
}

func interactiveInput(name string) string {
	fmt.Printf("%s: ", name)
	bio := bufio.NewReader(os.Stdin)
	input, isPrefix, err := bio.ReadLine()

	if err != nil {
		log.Fatal("error: ", err)
	}

	if isPrefix {
		log.Fatalf("%s is too long", name)
	}

	return string(input[:])
}
