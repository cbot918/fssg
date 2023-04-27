package fssc

import (
	"fmt"
	"net"
	"os"
)

type Fssc struct {
	Url  string
	Conn net.Conn
}

func NewFssc(url string) *Fssc {
	return &Fssc{
		Url: url,
	}
}

func (c *Fssc) Run() {
	var err error
	c.Conn, err = net.Dial("tcp", c.Url)
	if err != nil {
		fmt.Println("net.Dial failed: ", err)
		os.Exit(1)
	}

	c.mainLoop()

}

func (c *Fssc) mainLoop() {
	buf := make([]byte, 1024)
	var input string
	for {
		fmt.Printf("> ")
		_, err := fmt.Scanln(&input)
		if err != nil {
			fmt.Println("fmt.Scanln failed: ", err)
			os.Exit(1)
		}

		switch input {
		case "exit()":
			{
				fmt.Println("exit fssc")
				c.Conn.Close()
				return
			}
		default:
			{
				c.Conn.Write([]byte(input))
				c.Conn.Read(buf)
				fmt.Println("s: ", string(buf))
			}

		}

	}
}
