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
