package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	if len(os.Args) < 2 { // validate input arguments
		fmt.Println("Usage: go run main.go <port>")
		os.Exit(1)
	}

	port := fmt.Sprintf(":%s", os.Args[1]) // construct port string
	listener, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println("failed to create listener, err:", err)
		os.Exit(1)
	}
	defer listener.Close() // defer means this will be executed when the surrounding function returns
	fmt.Println("Server is listening on port", os.Args[1])

	for { // accept connections in a loop infinite loop
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("failed to accept connection, err:", err)
			continue
		}

		go handleConnection(conn) //Every client gets a separate goroutine so connections don’t block each other.
//This is why Go is famous — goroutine concurrency is cheap and simple.
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close() // ensure the connection is closed when the function exits

	reader := bufio.NewReader(conn) // reading data from the connection

	for {
		bytes, err := reader.ReadBytes('\n')
		if err != nil {
			if err != io.EOF {
				fmt.Println("failed to read data, err:", err)
			}
			return
		}
		fmt.Printf("request: %s", string(bytes))  // logging the received data

		line := fmt.Sprintf("Echo: %s", string(bytes)) // preparing response
		fmt.Printf("response: %s", line) // writing response back to the connection

		_, err = conn.Write([]byte(line)) // echoing back the received data
		if err != nil {
			fmt.Println("failed to write data, err:", err)
			return
		}
	}
}

//Sprintf formats according to a format specifier and returns the resulting string.
//It is analogous to Printf, but instead of printing to standard output, it returns the formatted string.

//defer statement defers the execution of a function until the surrounding function returns.

// goroutine is a lightweight thread managed by the Go runtime.
//When you prefix a function or method call with the go keyword, it runs in its own goroutine.

//bufio.NewReader returns a new Reader whose buffer has the default size.
//ReadBytes reads until the first occurrence of delim in the input, returning a slice containing the data up to and including the delimiter.
