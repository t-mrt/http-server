package main

import (
	"fmt"
	"os"
	"syscall"
)

func main() {

	fmt.Println(os.Getpid())

	fd, _ := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, 0)
	sa := syscall.SockaddrInet4{Port: 8888}
	if err := syscall.Bind(fd, &sa); err != nil {
		fmt.Println(err)
	}
	if err := syscall.Listen(fd, 5); err != nil {
		fmt.Println(err)
	}

	for {

		nfd, _, err := syscall.Accept(fd)
		if err != nil {
			fmt.Println(err)
		}
		buffer := make([]byte, 1024)
		syscall.Read(nfd, buffer)
		fmt.Println(string(buffer))
		syscall.Close(nfd)
	}
}
