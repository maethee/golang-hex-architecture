package validation

import (
	"github.com/tealeg/xlsx"
)

type Validation struct {
	Rules []Rule
}

func (v *Validation) Add(rule Rule) {
	v.Rules = append(v.Rules, rule)
}

func (v *Validation) FormatSheet(file *xlsx.File, options ...FormatSheetOptions) {
	//each rule to go routine and wainting done
	validationSheet := NewValidateSheet(options...)
	for _, row := range validationSheet.GetRows() {
		for i, cell := range row.Cells {
			if rule := v.FindRuleByIndex(i); rule != nil {
				if match := rule.Match(cell); match == true {
					rule.CorrectCellValue(cell)
				}
			}
		}
	}
}

func (v *Validation) FindRuleByIndex(i int) Rule {
	for _, rule := range v.Rules {
		if has := rule.HasRule(i); has == true {
			return rule
		}
	}
	return nil
}

type ValidationSheet struct {
	File     *xlsx.File
	Sheet    int
	StartRow int
}

type FormatSheetOptions func(f *ValidationSheet)

func Sheet(n int) FormatSheetOptions {
	return func(v *ValidationSheet) {
		v.Sheet = n
	}
}

func File(n *xlsx.File) FormatSheetOptions {
	return func(v *ValidationSheet) {
		v.File = n
	}
}

func NewValidateSheet(options ...FormatSheetOptions) *ValidationSheet {
	v := &ValidationSheet{
		Sheet: 0,
	}
	for _, opt := range options {
		opt(v)
	}
	return v
}

func (v *ValidationSheet) GetSheet() *xlsx.Sheet {
	return v.File.Sheets[v.Sheet]
}

func (v *ValidationSheet) GetRows() []*xlsx.Row {
	return v.GetSheet().Rows
}
