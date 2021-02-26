package listeners

import (
	"fmt"
	"net"


)

func handleConnection(conn net.Conn) {
  // try to read data from the connection
  data := make([]byte, 512)
  n, err := conn.Read(data)
  if err != nil { panic(err)  }
  s := string(data[:n])

  // print the request data
  fmt.Println(s)

}

// StartTCPServer will start the new TCP Server
func StartTCPServer() {
	fmt.Println("Starting TCP Server...")

	l, err := net.Listen("tcp", "127.0.0.1:8888")

	if err != nil {
		fmt.Println("Cannot create socket: ",err.Error())
	}

	for {
		conn,err := l.Accept()
		if err != nil {
			fmt.Println(err.Error())
		}
		go handleConnection(conn)

	}


}