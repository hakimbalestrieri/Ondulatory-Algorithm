package Network

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)


type Client struct {
	conn net.Conn
	Commands chan string
}

const srvAddr = "127.0.0.1:6000"

func main() {
	conn, err := net.Dial("udp", srvAddr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	go func() {
		mustCopy(os.Stdout, conn)
	}()
	mustCopy(conn, os.Stdin)
}
func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}

//read client input
func (c *Client) Read() {
	for {
		var input string
		_, err := fmt.Scanln(&input)
		if err != nil {
			log.Fatal(err)
		}
		c.Commands <- input
	}
}
