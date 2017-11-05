package main

import (
	"fmt"
	"github.com/tealeg/xlsx"
)

func appendIfMissing(slice []string, i string) []string {
	for _, ele := range slice {
		if ele == i {
			return slice
		}
	}
	return append(slice, i)
}

func main() {
	excel, _ := xlsx.FileToSlice("Luat_hecosotrithuc.xlsx")
	fmt.Println(len(excel[0]))
	fmt.Println(excel[0][1][4])
	fmt.Println(len(excel[0][0]))
	fmt.Println(excel[0][0][0])

	//var diction[]map[string]string


	diction := make(map[string]string)

	var label string
	for i := 0;i< len(excel[0][0]);i++{
		var coll []string

		label = excel[0][0][i]
		for j := 1;j< len(excel[0]);j++ {
			coll = appendIfMissing(coll, excel[0][j][i])
		}
		for j := 0;j<len(coll);j++{
			diction[coll[j]] = label
		}
		//diction = append(diction, diction)

	}


	//label := excel[0][0][0]
	//for i := 1;i< len(excel[0]);i++ {
	//	coll = AppendIfMissing(coll, excel[0][i][0])
	//}
	//for i := 0;i<len(coll);i++{
	//	mapLabel[coll[i]] = label
	//}
	//diction = append(diction, mapLabel)

	fmt.Println(diction["40"])

	//mapLabel := make(map[string]string)
}