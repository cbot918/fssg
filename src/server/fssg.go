package server

import (
	"fmt"
	"net"
	"os"
)

type Fssg struct {
	Url string
}

func NewFssg(url string) *Fssg {
	return &Fssg{
		Url: url,
	}
}

func (f *Fssg) Run() {
	lis, err := net.Listen("tcp", f.Url)
	if err != nil {
		fmt.Println("listen failed: ", err)
		os.Exit(1)
	}
	fmt.Println("listening: ", f.Url)

	for {
		conn, err := lis.Accept()
		if err != nil {
			fmt.Println("lis.Accept failed:", err)
			os.Exit(1)
		}

		go HandleConnection(conn)

	}

}

func HandleConnection(conn net.Conn) {
	fmt.Printf("%s connected\n", conn.RemoteAddr())
	buf := make([]byte, 1024)
	for {
		_, err := conn.Read(buf)
		if err != nil {
			if err.Error() == "EOF" {
				fmt.Printf("%s is disconnected\n", conn.RemoteAddr())
				conn.Close()
				return
			}
			fmt.Println("conn.Read failed: ", err.Error())
			os.Exit(1)
		}
		fmt.Println("c: ", string(buf))

		_, err = conn.Write(buf)
		if err != nil {
			fmt.Println("conn.Write failed: ", err)
			os.Exit(1)
		}

	}

}
