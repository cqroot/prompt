package main

import (
	"errors"
	"fmt"
	"net"
	"os"

	"github.com/cqroot/prompt"
	"github.com/cqroot/prompt/input"
)

var ErrInvalidIP = errors.New("invalid ip")

func validateIP(ip string) error {
	if net.ParseIP(ip) == nil {
		return fmt.Errorf("%s: %w", ip, ErrInvalidIP)
	} else {
		return nil
	}
}

func main() {
	val, err := prompt.New().Ask("Please enter the server IP:").
		Input("127.0.0.1", input.WithValidateFunc(validateIP))
	if err != nil {
		if errors.Is(err, prompt.ErrUserQuit) {
			fmt.Fprintln(os.Stderr, "Error:", err)
			os.Exit(1)
		} else if errors.Is(err, ErrInvalidIP) {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		} else {
			panic(err)
		}
	}
	fmt.Println(val)
}
