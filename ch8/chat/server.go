package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}

	// 练习 8.13： 使聊天服务器能够断开空闲的客户端连接，比如最近五分钟之后没有发送任何消息的那些客户端。提示：可以在其它goroutine中调用conn.Close()来解除Read调用，就像input.Scanner()所做的那样。

	// 练习 8.14： 修改聊天服务器的网络协议，这样每一个客户端就可以在entering时提供他们的名字。将消息前缀由之前的网络地址改为这个名字。

	// 练习 8.15： 如果一个客户端没有及时地读取数据可能会导致所有的客户端被阻塞。修改broadcaster来跳过一条消息，而不是等待这个客户端一直到其准备好读写。或者为每一个客户端的消息发送channel建立缓冲区，这样大部分的消息便不会被丢掉；broadcaster应该用一个非阻塞的send向这个channel中发消息。
}

type client chan<- string // an  outgoing message channel

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string)
	connMap  = make(map[net.Conn]time.Time)
)

func broadcaster() {
	clients := make(map[client]bool) // all connected clients
	for {
		select {
		case msg := <-messages:
			// Boradcast incoming messages to
			// all clients' outgoing message channels
			for cli := range clients {
				cli <- msg
			}
		case cli := <-entering:
			clients[cli] = true
		case cli := <-leaving:
			delete(clients, cli)
			close(cli)
		}
	}
}

func handleConn(conn net.Conn) {
	ch := make(chan string) // outgoing client messages
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	ch <- "You are " + who
	messages <- who + " has arrived"
	entering <- ch

	// 练习 8.13： 使聊天服务器能够断开空闲的客户端连接，比如最近五分钟之后没有发送任何消息的那些客户端。
	// 提示：可以在其它goroutine中调用conn.Close()来解除Read调用，就像input.Scanner()所做的那样。
	// 计算时间，如果超时，则断开连接
	msgTime := make(chan time.Time)
	go func(msgTime <-chan time.Time) {
		tick := time.Tick(time.Second)
		lastTime := time.Now()
		for t := range tick {
			if t.Sub(lastTime) > 5*time.Second {
				fmt.Printf("idle connection timeout: %v \n", conn.RemoteAddr())
				conn.Close()
				return
			} else {
				fmt.Printf("tick: %T \n", t)
			}
		}
	}(msgTime)

	// 等待网络，接受客户端的消息
	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- who + ":" + input.Text()
		msgTime <- time.Now()
	}

	leaving <- ch
	messages <- who + " has left"
	conn.Close()
	delete(connMap, conn)
}

func clientWriter(conn net.Conn, ch chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg) // NOTE: ignoring network erros
	}
}
