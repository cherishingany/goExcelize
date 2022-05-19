package main

import (
	"excelize/controller"
	"fmt"
	"github.com/xuri/excelize/v2"
	"strings"
	"sync"
)

func main() {

	fmt.Println("main 线程开始 -----")
	file := OpenFile()
	rows := Rows(file)

	GOSearch(file, rows)

	//---------------------------------------------------

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
		s := strings.Join(col[:1], "")
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
				controller.Case0(f, v) // 0.申购申请成功，确认部分成功 - 修改文件的交易份额,交易金额，返回代码
			case 1:
				controller.Case1(f, v) // 1.申购申请成功，确认失败 - 修改文件的交易份额,交易金额，返回代码
			case 2:
				controller.Case2(f, v) // 2.赎回申请成功，确认失败
			case 3:
				controller.Case3(f, v) // 3.转托管申请成功，确认失败 - 需要填写是否注册登记人 BG
			case 4:
				controller.Case4(f, v) // 4.基金转换申请成功，确认失败
			}
			defer wg.Done()
		}(index, v)
	}
	wg.Wait()
}
