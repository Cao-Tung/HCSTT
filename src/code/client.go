package main

import (
	"git.apache.org/thrift.git/lib/go/thrift"
	"net"
	"fmt"
	"os"
	"gen-go/hello"
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

	r1, e1 := client.HelloString("I want buy a jackets,it's color is black,my tall is 1m73, my weight is 53kg")
	fmt.Println("Call->", r1, e1)

}