package reader

import (
	"log"

	"github.com/tealeg/xlsx"
)

// Receive string path and return [sheet][row][col]string
func Read(path string) (*xlsx.File, error) {
	file, err := xlsx.OpenFile(path)
	if err != nil {
		log.Panic(err)
	}

	return file, nil
}

func UpdateCell(file *xlsx.File, rowIndex int, colIndex int, value interface{}) error {
	r := file.Sheets[1].Row(0)
	c := r.Cells[0]
	c.SetValue(value)

	return nil
}

func Save(file *xlsx.File) error {
	return file.Save("../../assets/New.xlsx")
}

