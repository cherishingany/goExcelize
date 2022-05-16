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
