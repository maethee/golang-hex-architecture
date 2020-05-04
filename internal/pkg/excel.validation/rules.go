package validation

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"

	"github.com/tealeg/xlsx"
)

type Rule interface {
	CorrectCellValue(*xlsx.Cell)
	Match(*xlsx.Cell) bool
	HasRule(int) bool
}

type RuleProps struct {
	Rule     string
	ColIndex int
	Options  map[string]string
}

// Support NullRule, DigitNumberRule
func CreateRules(ruleconfs []RuleProps) (error, []Rule) {
	var rules []Rule
	for _, v := range ruleconfs {
		switch v.Rule {
		case "NullRule":
			newRule := NullRule{
				ColIndex:      v.ColIndex,
				OverrideValue: v.Options["OverrideValue"],
			}
			rules = append(rules, &newRule)
		case "DigitNumberRule":
			newRule := DigitNumberRule{
				ColIndex: v.ColIndex,
			}
			rules = append(rules, &newRule)
		case "LocalFileExistRule":
			newRule := LocalFileExistRule{
				ColIndex: v.ColIndex,
				RootPath: v.Options["RootPath"],
			}
			rules = append(rules, &newRule)
		default:
			fmt.Printf("Rule %d not found", v.Rule)
		}
	}
	return nil, rules
}

type NullRule struct {
	ColIndex      int
	OverrideValue string
}

func (c *NullRule) CorrectCellValue(cell *xlsx.Cell) {
	cell.Value = c.OverrideValue
}

func (c *NullRule) Match(cell *xlsx.Cell) bool {
	exp := "NULL"
	byteValue := []byte(cell.Value)

	r, err := regexp.Match(exp, byteValue)
	if err != nil {
		log.Panic(err)
	}
	return r
}

func (c *NullRule) HasRule(i int) bool {
	return c.ColIndex == i
}

type DigitNumberRule struct {
	ColIndex int
}

func (c *DigitNumberRule) CorrectCellValue(cell *xlsx.Cell) {
	if len(cell.Value) == 1 {
		cell.SetString(fmt.Sprintf("0%v", cell.Value))
	}
	cell.SetString(cell.Value)
}

func (c *DigitNumberRule) Match(cell *xlsx.Cell) bool {
	return len(cell.Value) <= 2
}

func (c *DigitNumberRule) HasRule(i int) bool {
	return c.ColIndex == i
}

type LocalFileExistRule struct {
	ColIndex int
	RootPath string
	files    []string
}

func (c *LocalFileExistRule) CorrectCellValue(cell *xlsx.Cell) {
	fmt.Println("LocalFileExistRule CorrectCellValue", cell.Value, cell.GetStyle())
	cell.Value = "Not found: " + cell.Value

}

func (c *LocalFileExistRule) Match(cell *xlsx.Cell) bool {
	if c.ColIndex == 0 {
		return false
	}

	root := c.RootPath

	if len(c.files) == 0 {
		fmt.Println("WalkTo")
		err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
			c.files = append(c.files, path)
			return nil
		})
		if err != nil {
			panic(err)
		}
	}

	for _, file := range c.files {
		byteValue := []byte(file)
		r, err := regexp.Match(cell.Value, byteValue)
		if err != nil {
			log.Panic(err)
		}

		if r == true {
			return false
		}
	}

	return true
}

func (c *LocalFileExistRule) HasRule(i int) bool {
	return c.ColIndex == i
}
