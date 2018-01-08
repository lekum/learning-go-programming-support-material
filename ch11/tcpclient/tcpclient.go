package main

import (
	"bufio"
	"fmt"
	"net"
	"time"
)

var host, port = "127.0.0.1", "4040"
var addr = net.JoinHostPort(host, port)

const prompt = "curr"
const buffLen = 1024

func main() {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	var cmd, param string
	// repl - interactive shell for client
	for {
		fmt.Print(prompt, "> ")
		_, err = fmt.Scanf("%s %s", &cmd, &param)
		if err != nil {
			fmt.Println("Usage GET <search string or *>")
			continue
		}
		// send command line
		cmdLine := fmt.Sprintf("%s %s", cmd, param)
		if n, err := conn.Write([]byte(cmdLine)); n == 0 || err != nil {
			fmt.Println(err)
			return
		}

		// stream and display response
		conn.SetReadDeadline(time.Now().Add(time.Second * 5))
		conbuf := bufio.NewReaderSize(conn, 1024)
		for {
			str, err := conbuf.ReadString('\n')
			if err != nil {
				break
			}
			fmt.Print(str)
			conn.SetReadDeadline(time.Now().Add(time.Millisecond * 700))
		}
	}

}
