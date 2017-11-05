
package main

import (
	"git.apache.org/thrift.git/lib/go/thrift"
	"gen-go/hello"
	"fmt"
	"time"
	"strings"
	"github.com/tealeg/xlsx"
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
	xlsxFile := "Luat_hecosotrithuc.xlsx"
	var response string = ""
	extract , err := ExtractData(para, xlsxFile)
	for _, str := range extract{
		response = str + "," + response
	}
	return response, err
}

func ExtractData(para string, xlsxFile string) (r []string, err error){
	data = make([]string,0)
	var diction map[string]string
	strs := strings.Split(para,",")
	for _,str := range strs{
		var check = ""
		var label string
		if strings.Contains(str, "tall"){
			check = "high"
			label = "high"
		}else if strings.Contains(str, "weight"){
			check = "weight"
			label = "weight"
		}
		for _,top := range dict{
			if strings.Contains(str,top) {
				arr := strings.Split(str, top)
				str = " " + arr[1]
			}
		}
		if check == "tall"{
			str = str + "_" + label
		}else if check == "weight"{
			str = str + "_" + label
		}else{
			diction = CreateDictionary(xlsxFile)
			str = strings.Replace(str, " ", "", -1)
			label = diction[str]
			str = str + "_" + label
		}
		data = append(data, str)
	}
	fmt.Println(data)
	return data, nil
}

func AppendIfMissing(slice []string, i string) []string {
	for _, ele := range slice {
		if ele == i {
			return slice
		}
	}
	return append(slice, i)
}

func CreateDictionary(xlsxFile string) map[string]string{
	excel, _ := xlsx.FileToSlice(xlsxFile)
	diction := make(map[string]string)

	var label string
	for i := 0;i< len(excel[0][0]);i++{
		var coll []string

		label = excel[0][0][i]
		for j := 1;j< len(excel[0]);j++ {
			coll = AppendIfMissing(coll, excel[0][j][i])
		}
		for j := 0;j<len(coll);j++{
			diction[coll[j]] = label
		}
	}
	return diction
}

func worker(id int) {
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
	}
}

func main() {
	dict = []string{" is ", " am ", " like ", " love ", " want ", " and ", " buy ", " a "}
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