package main

import (
	"git.apache.org/thrift.git/lib/go/thrift"
	"gen-go/hello"
	"fmt"
	"time"
	"strings"
)

var job int
var jobs chan string
var dict []string
var data []string

type HelloHandler struct {
	str string
}

func NewHelloHandler() *HelloHandler{
	return &HelloHandler{"Hello"}
}

func (hl *HelloHandler)  HelloString(para string) (r string, err error){
	//job ++
	//c := strconv.Itoa(job)
	//st := c + para
	//jobs <- st
	//fmt.Println(para)
	ExtractData(para)
	return para, nil
}

func ExtractData(para string) (r string, err error){
	data = make([]string,0)
	temp := make([]string,0)
	strs := strings.Split(para,",")
	for _,str := range strs{
		for _,top := range dict{
			if strings.Contains(str,top){
				arr := strings.Split(str,top)
				temp = append(temp, arr[1])
			}
		}
	}
	fmt.Println(temp)
	return
}

func worker(id int) {
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
	}
}

func main() {
	dict = []string{" is ", " am ", " like ", " love ", " want ", " and ", " buy "}
	jobs = make(chan string, 30)
	job = 0
	hel := NewHelloHandler()
	processor := hello.NewHelloProcessor(hel)
	transport,_ := thrift.NewTServerSocket("127.0.0.1:1996")
	tfactory := thrift.NewTTransportFactory()
	pfactory := thrift.NewTBinaryProtocolFactoryDefault()
	server := thrift.NewTSimpleServer4(processor, transport, tfactory, pfactory)
	fmt.Println("Server run at 127.0.0.1:1996")
	//for w := 1; w <= 3; w++ {
	//	go worker(w)
	//}

	server.Serve()

}
