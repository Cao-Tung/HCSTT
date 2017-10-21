package main

import (
	"git.apache.org/thrift.git/lib/go/thrift"
	"net"
	"fmt"
	"os"
	"gen-go/hello"
	"time"
)

func main() {
	transportFactory := thrift.NewTBufferedTransportFactory(1024)
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	transport, err := thrift.NewTSocket(net.JoinHostPort("127.0.0.1", "1996"))
	if err != nil {
		fmt.Fprintln(os.Stderr, "error resolving address:", err)
		os.Exit(1)
	}
	useTransport,_ := transportFactory.GetTransport(transport)
	client := hello.NewHelloClientFactory(useTransport, protocolFactory)
	if err := transport.Open(); err != nil {
		fmt.Fprintln(os.Stderr, "localhost:1996", " ", err)
		os.Exit(1)
	}

	time.Sleep(3000 * time.Millisecond)
	r1, e1 := client.HelloString("Hello GoLang")
	fmt.Println("Call->", r1, e1)

}
