package main

import (
	"errors"
	"fmt"
	"net"
	"os"

	"github.com/cqroot/prompt"
)

func validateIP(ip string) error {
	if net.ParseIP(ip) == nil {
		return errors.New(ip + " is not a valid ip")
	} else {
		return nil
	}
}

func main() {
	val, err := prompt.New().Ask("Please enter the server IP:").
		Input("127.0.0.1", prompt.WithValidateFunc(validateIP))
	if err != nil {
		if errors.Is(err, prompt.ErrUserQuit) {
			fmt.Fprintln(os.Stderr, "Error:", err)
			os.Exit(1)
		} else {
			panic(err)
		}
	}
	fmt.Println(val)
}
