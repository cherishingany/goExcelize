package main

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"strconv"
	"strings"
	"sync"
)

func main() {

	fmt.Println("main 线程开始 -----")
	file := OpenFile()
	rows := Rows(file)
	//
	fmt.Println(rows)

	GOSearch(file, rows)
	fmt.Println("main 线程结束 -----")
}

func OpenFile() *excelize.File {
	f, _ := excelize.OpenFile("53.xlsx")
	return f
}

//从申请数据中遍历申请单编号,根据业务类型写入到不同的slice中。此处case slice 具有很多种，TODO silce初始化放在函数中
func Rows(f *excelize.File) [][]string {

	cols, err := f.GetRows("x1")
	if err != nil {
		fmt.Println(err)
		return nil
	}

	case1 := make([]string, 0)
	case2 := make([]string, 0)
	case3 := make([]string, 0)
	case4 := make([]string, 0)
	case5 := make([]string, 0)
	case6 := make([]string, 0)

	for _, col := range cols[3:] {

		s := fmt.Sprintf(strings.Join(col[:1], ""))
		switch s {
		case "1":
			//代表该基金申编号确认部分成功
			case1 = append(case1, col[1:2]...)
		case "2":
			case2 = append(case2, col[1:2]...)
		case "3":
			case3 = append(case3, col[1:2]...)
		case "4":
			case4 = append(case4, col[1:2]...)
		case "5":
			case5 = append(case5, col[1:2]...)
		case "6":
			case6 = append(case6, col[1:2]...)
		default:
			fmt.Println("Rows default .... error------")
		}
	}
	var rowSlice [][]string
	//rowSlice = append(rowSlice, string1, string2, string3, string4, string5, string6)
	rowSlice = append(rowSlice, case6)

	fmt.Println("case 分组情况:", rowSlice)
	return rowSlice
}

//[[9202204279252544801 9202204279252544401]]
func GOSearch(f *excelize.File, s [][]string) {
	wg := &sync.WaitGroup{}
	//s [][]string的数量开groutine TODO 看他如果[]string为空不要去开了。
	for index, v := range s {
		if len(v) <= 0 {
			return
		}
		wg.Add(1)
		go func(index int, v []string) {
			switch index {
			case 0:
				Case1(f, v)
			case 2:
			case 3:
			}
			defer wg.Done()
		}(index, v)
	}
	wg.Wait()
}

//根据遍历的slice 去文件中查找
// [9202204279252544801 9202204279252544401]
func SearchTo(f *excelize.File, s []string, activeSheet string) []string {

	var newslice []string
	for _, value := range s {
		values, err := f.SearchSheet(activeSheet, value)
		if err != nil {
			panic(err)
		}
		for _, value := range values {
			newslice = append(newslice, value)
		}
	}
	return newslice
}

func GetSheetRow53(f *excelize.File, rows, cols []string) [][]string {

	//先处理rows 除去列数据
	for i, row := range rows {
		rows = append(rows[:i], row[1:])
	}
	fmt.Println(rows)

	//rows := []string{"399", "400"}
	//cols := []string{"G", "AI"}
	var valueNum [][]string
	//[[A399,b399],[s399,g400]]
	rowcols := RowCol(rows, cols)

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

func SetSheetRow63(f *excelize.File, setValues [][]string, rows, cols []string) { //

	//先处理rows 除去列数据

	//rows := []string{
	//	"B9",
	//	"B8",
	//}
	//cols := []string{"L", "M", "N"}

	//setValues := [][]string{
	//	{"ZD0001", "0.00", "218"},
	//	{"ZD0001", "0.00", "218"},
	//}

	for i, row := range rows {
		rows = append(rows[:i], row[1:])
	}
	coordinates := RowCol(rows, cols)

	//coordinates := [][]string{
	//	{"B156", "B129"},
	//	{"B157", "B127"},
	//}
	//setValues := [][]string{
	//	{"156", "129"},
	//	{"157", "127"},
	//}

	for index, coordinate := range coordinates {
		for i := 0; i < len(coordinate); i++ {
			err := f.SetCellStr("63", coordinate[i], setValues[index][i])
			if err != nil {
				fmt.Println(err)
			}
		}
	}

	f.Save()
}

func RowCol(rows, cols []string) [][]string {

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

// case1  确认部分成功    -  调整 Q 交易确认份数
//					    	R 交易确认金额
//					    	返回代码9999
// col63 := []string{"AO", "PO"}
//  [9202204279252544801 9202204279252544401]
func Case1(f *excelize.File, s []string) {

	//行
	row53 := SearchTo(f, s, "53") //[A399 A400]
	//列
	col := []string{"O", "P"}

	//[[ZD0001 0.00 218] [ZD0001 0.00 218]]
	getValue := GetSheetRow53(f, row53, col)
	fmt.Println("从53文件中获取到的数据", getValue)
	getValue = ChooseNum(getValue, 2)
	fmt.Println("/2的数据", getValue)

	// 中间处理string append 问题

	for i := 0; i < len(getValue); i++ {
		getValue[i] = append(getValue[i], "0000")
	}

	fmt.Println("加完 0000 的数据", getValue)

	//往63放数据了
	row63 := SearchTo(f, s, "63")
	col63 := []string{"Q", "R", "AI"}
	SetSheetRow63(f, getValue, row63, col63)

	// 53文件中的位置
	//s, err := f.GetCellValue("Sheet1", "O"+s) //TODO 根据返回的单元格行进行查找
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//

	//setvalue(f, "63", "4Q", s)       //设置交易份额
	//setvalue(f, "63", "4R", s)       //设置交易份额
	//setvalue(f, "63", "4AI", "0000") //设置返回代码
}

func setvalue(f *excelize.File, sheet string, axis string, value interface{}) {
	f.SetCellValue(sheet, axis, value)
	f.Save()
}

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
