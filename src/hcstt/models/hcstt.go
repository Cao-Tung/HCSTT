package models

import (
	"fmt"
	"strings"
	"github.com/tealeg/xlsx"
)

var dict []string
var data []string

func init() {

}


type DataError struct {
	Code 		int  `json:"code,omitempty"`
	Message 	string  `json:"message,omitempty"`
}

type DataResponse struct {
	Error DataError `json:"error,omitempty"`
	Data string `json:"data,omitempty"`
}

func HandleData(input string)(datre DataResponse){
	dict = []string{" is ", " am ", " like ", " love ", " want ", " and ", " buy ", " a "}
	xlsxFile := "Luat_hecosotrithuc.xlsx"
	var response string = ""
	extract , _ := ExtractData(input, xlsxFile)
	fmt.Println("HI", extract)
	for _, str := range extract{
		response = str + "," + response
	}
	fmt.Println(response)
	return DataResponse{Error: DataError{Code:200, Message:"Success"}, Data: response}
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
		if check == "high"{
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









