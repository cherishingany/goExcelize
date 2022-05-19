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

	var rowSlice = make([][]string, 6)

	for _, col := range cols[3:] {
		s := strings.Join(col[:1], "")
		switch s {
		case "0":
			//代表该基金申编号确认部分成功
			rowSlice[0] = append(rowSlice[0], col[1:2]...)
		case "1":
			rowSlice[1] = append(rowSlice[1], col[1:2]...)
		case "2":
			rowSlice[2] = append(rowSlice[2], col[1:2]...)
		case "3":
			rowSlice[3] = append(rowSlice[3], col[1:2]...)
		case "4":
			rowSlice[4] = append(rowSlice[4], col[1:2]...)
		case "5":
			rowSlice[5] = append(rowSlice[5], col[1:2]...)
		default:
			fmt.Println("Rows default .... error------")
		}
	}

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
