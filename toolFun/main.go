package main

import (
	"encoding/json"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"os"
	"strconv"
)

type info struct {
	index int
	name string
	attack int
	attackTime int
	attackRange int
	bloodCount int
	skillInfo string
}


func main() {
	xlsx, err := excelize.OpenFile("./toolFun/wolfWar.xlsx")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// Get value from cell by given sheet index and axis.
	//cell := xlsx.GetCellValue("Sheet1", "")
	//fmt.Println(xlsx.GetRows("狼群技能"))
	getRows := xlsx.GetRows("狼群技能")
	//fmt.Println(getRows)
	result:= make([]info,0)
	
	for _,row:=range getRows{
		index , _:= strconv.Atoi(row[0])
		attack , _:= strconv.Atoi(row[3])
		attackTime , _:= strconv.Atoi(row[4])
		attackRange , _:= strconv.Atoi(row[5])
		bloodCount , _:= strconv.Atoi(row[6])
		wolfInfo := info{index: index, name: row[1],attack: attack, attackTime: attackTime, attackRange: attackRange, bloodCount:bloodCount,skillInfo:row[6]}
		result = append(result,wolfInfo)
		
		if b, err := json.Marshal(wolfInfo); err == nil {
			fmt.Println("================struct 到json str==")
			fmt.Println(string(b))
		}
	}
	fmt.Println(result)
	
	marshal, err := json.Marshal(result)
	fmt.Println(string(marshal))
	// Get sheet index.
	index := xlsx.GetSheetIndex("Sheet2")
	// Get all the rows in a sheet.
	rows := xlsx.GetRows("sheet" + strconv.Itoa(index))
	for _, row := range rows {
		for _, colCell := range row {
			fmt.Print(colCell, "\t")
		}
		fmt.Println()
	}
}