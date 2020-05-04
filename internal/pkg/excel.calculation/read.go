package main

import (
	"fmt"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func main() {

	f, err := excelize.OpenFile("./assets/12PL.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}

	cell := f.GetCellFormula("กรอกข้อมูล", "F17")
	fmt.Println(cell)
}
