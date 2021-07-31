//go:build linux
//+build linux

package main

import (
	"fmt"
	"os"

	"golang.org/x/sys/unix"
)

func main() {

	termios, err := unix.IoctlGetTermios(unix.Stdin, unix.TCGETS)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("%+v\n", *termios)

}
