package listeners

import (
	"bufio"
	"io"
	"net"
	"strings"
	"fmt"
)

func handleConnection(con net.Conn) {
	defer con.Close()
 
	clientReader := bufio.NewReader(con)
 
	for {
		// Waiting for the client request
		clientRequest, err := clientReader.ReadString('\n')
 
		switch err {
		case nil:
			clientRequest := strings.TrimSpace(clientRequest)
			if clientRequest == ":QUIT" {
				fmt.Println("client requested server to close the connection so closing")
				return
			} else {
				fmt.Println(clientRequest)
			}
		case io.EOF:
			fmt.Println("client closed the connection by terminating the process")
			return
		default:
			fmt.Printf("error: %v\n", err)
			return
		}
 
		// Responding to the client request
		if _, err = con.Write([]byte("GOT IT!\n")); err != nil {
			fmt.Printf("failed to respond to client: %v\n", err)
		}
	}
}

// StartTCPServer will start the new TCP Server
func StartTCPServer() {
	fmt.Println("Starting TCP Server...")
	listener, err := net.Listen("tcp", "0.0.0.0:9999")
	if err != nil {
		fmt.Println(err)
	}
	defer listener.Close()
 
	for {
		con, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
 
		// If you want, you can increment a counter here and inject to handleClientRequest below as client identifier
		go handleConnection(con)
	}


}