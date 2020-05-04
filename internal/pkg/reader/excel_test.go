package reader

import (
	"fmt"
	"testing"
)

func TestSample(t *testing.T) {
	path := "../../../assets/TemplateConfig_Plan_S10P02_20200305.xlsx"
	want := 2
	file, _ := Read(path)
	if len(file.Sheets) != want {
		t.Errorf("Read(path) = %v, want %v", len(file.Sheets), want)
	}

	fmt.Println(file.Sheets[1].Cols[4].Width)
	
	err := file.Save("../../../assets/New.xlsx")
	fmt.Println(err)
}
