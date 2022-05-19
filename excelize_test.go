package main

import (
	"fmt"
	"testing"
)

func TestSearchToError(t *testing.T) {
	f := OpenFile()
	sli := []string{"9202204279252544401"}
	activeSheet := "63"
	var newslice []string
	for _, value := range sli {
		values, _ := f.SearchSheet(activeSheet, value)
		if values == nil {
			fmt.Printf("订单编号未找到,请手动确认  -  %s\n", value)
			fmt.Println()
		}
		newslice = append(newslice, values[:]...)
	}
	fmt.Println(newslice)
}

func TestSetSheetRow63(t *testing.T) { //

	setValues := [][]string{
		{"0.00", "0.00", "9999"},
		{"0.00", "0.00", "9999"},
	}

	coordinates := [][]string{
		{"Q1", "R1", "AI1"},
		{"Q2", "R2", "AI2"},
		{"Q3", "R3", "AI3"},
	}

	if len(setValues) != 1 {
		panic("0 必须")
	}

	for _, coordinate := range coordinates {
		for i := 0; i < len(coordinate); i++ {
			fmt.Print(coordinate[i], "+ ")
			fmt.Println(setValues[0][i])
		}
	}
}

//从申请数据中遍历申请单编号,根据业务类型写入到不同的slice中。此处case slice 具有很多种，TODO silce初始化放在函数中
func TestRows(t *testing.T) {

	var rowSlice [3][]string

	clos := [][]string{
		{"case0", "case0"},
		{"case1"},
		{"case2"},
	}

	for s, col := range clos {
		s := fmt.Sprintf("%d", s)
		switch s {
		case "0":
			rowSlice[0] = append(rowSlice[0], col[:]...)
		case "1":
			rowSlice[1] = append(rowSlice[1], col[:]...)
		case "2":
			rowSlice[2] = append(rowSlice[2], col[:]...)
		default:
			fmt.Println("Rows default .... error------")
		}
	}

	fmt.Println("case 分组情况:", rowSlice)

}
