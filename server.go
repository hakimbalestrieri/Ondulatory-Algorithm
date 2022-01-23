package PRR_Labo04

import (
	"PRR-Labo04/Network"
	"log"
	"net"
)

func main() {
	addr, _ := net.ResolveUDPAddr("udp", Network.srvAddr)
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	buf := make([]byte, 1024)
	for {
		n, cliAddr, err := conn.ReadFromUDP(buf)
		if err != nil {
			log.Fatal(err)
		}
		s := string(buf[0:n-1]) + " from " + cliAddr.String() + "\n"
		if _, err := conn.WriteTo([]byte(s), cliAddr); err != nil {
			log.Fatal(err)
		}
	}
}