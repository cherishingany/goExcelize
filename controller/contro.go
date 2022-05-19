package controller

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"strconv"
)

// SearchTo 根据遍历的slice 去文件中查找
// [9202204279252544801 9202204279252544401]
func searchTo(f *excelize.File, s []string, activeSheet string) []string {

	var newslice []string
	for _, value := range s {
		values, _ := f.SearchSheet(activeSheet, value)
		if values == nil {
			fmt.Printf("订单编号未找到,请手动确认  -  %s\n", value)
			fmt.Println()
		}
		newslice = append(newslice, values[:]...)
	}
	return newslice
}

func getSheetRow53(f *excelize.File, rows, cols []string) [][]string {

	//先处理rows 除去列数据
	for i, row := range rows {
		rows = append(rows[:i], row[1:])
	}
	fmt.Println(rows)

	//rows := []string{"399", "400"}
	//cols := []string{"G", "AI"}
	var valueNum [][]string
	//[[A399,b399],[s399,g400]]
	rowcols := rowCol(rows, cols)

	for _, firstSlice := range rowcols {
		var row1 []string
		for _, value := range firstSlice {
			cellValue, _ := f.GetCellValue("53", value)
			row1 = append(row1, cellValue)
		}
		valueNum = append(valueNum, row1)
	}
	return valueNum
}

func setSheetRow63(f *excelize.File, setValues [][]string, rows, cols []string, flag int8) { //

	for i, row := range rows {
		rows = append(rows[:i], row[1:])
	}
	coordinates := rowCol(rows, cols)

	fmt.Println(coordinates, " -----")

	switch flag {
	case 0:
		// flag为0则setValues只能有一组数据
		if len(setValues) != 1 {
			panic("0 必须只能有一组数据")
		}
		for _, coordinate := range coordinates {
			for i := 0; i < len(coordinate); i++ {
				err := f.SetCellStr("63", coordinate[i], setValues[0][i])
				if err != nil {
					panic(err)
				}
			}
		}
	case 1:
		for index, coordinate := range coordinates {
			for i := 0; i < len(coordinate); i++ {
				err := f.SetCellStr("63", coordinate[i], setValues[index][i])
				//TODO 如果此处发生panic,原因为：63文件中查找到多个申请单编号，而取值时候没有在53文件中取。如果写成按照申请单赋值，那么按照那个申请单赋值？
				if err != nil {
					panic(err)
				}
			}
		}
	}

	f.Save()
}

// rowCol 返回行,列重新组合的[][]string
func rowCol(rows, cols []string) [][]string {

	var v1 string
	var valueNum [][]string
	for _, row := range rows {
		var row1 []string
		for _, col := range cols {
			v1 = col + row
			row1 = append(row1, v1)
		}
		valueNum = append(valueNum, row1)
	}
	return valueNum
}

// ChooseNum 函数接收到num 为被除数
func ChooseNum(slices [][]string, num float64) [][]string {
	for i := 0; i < len(slices); i++ {
		for index, value := range slices[i] {
			distFloat, _ := strconv.ParseFloat(value, 32)
			s := strconv.FormatFloat(distFloat/num, 'f', 2, 64)
			slices[i] = append(slices[i][:index], s)
		}
	}
	return slices
}

// caseDefault 该函数提供 默认 0.00 ,0.00,0000 代码使用
func caseDefault(f *excelize.File, s []string) {
	getValue := [][]string{
		{"0.00", "0.00", "9999"},
	}
	row63 := searchTo(f, s, "63")
	col63 := []string{"Q", "R", "AI"}
	setSheetRow63(f, getValue, row63, col63, 0)
}

// case1  确认部分成功    -  调整 Q 交易确认份数
//					    	R 交易确认金额
//					    	返回代码9999
// col63 := []string{"AO", "PO"}
//  [9202204279252544801 9202204279252544401]

func Case0(f *excelize.File, s []string) {

	//行
	row53 := searchTo(f, s, "53") //[A399 A400]
	//列
	col := []string{"O", "P"}

	//[[ZD0001 0.00 218] [ZD0001 0.00 218]]
	getValue := getSheetRow53(f, row53, col)
	fmt.Println("从53文件中获取到的数据", getValue)
	getValue = ChooseNum(getValue, 2)
	fmt.Println("/2的数据", getValue)

	// 中间处理string append 问题

	for i := 0; i < len(getValue); i++ {
		getValue[i] = append(getValue[i], "0000")
	}

	fmt.Println("加完 0000 的数据", getValue)

	//往63放数据了
	row63 := searchTo(f, s, "63")
	col63 := []string{"Q", "R", "AI"}
	setSheetRow63(f, getValue, row63, col63, 1)

}

func Case1(f *excelize.File, s []string) {
	caseDefault(f, s)
}

func Case2(f *excelize.File, s []string) {
	caseDefault(f, s)
}

func Case3(f *excelize.File, s []string) {
	caseDefault(f, s)
}
func Case4(f *excelize.File, s []string) {
	caseDefault(f, s)
}
