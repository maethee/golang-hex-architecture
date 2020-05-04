package myaxa

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type ProductTableRow struct {
	PlanCode             string
	ProductCategory_EN   string
	ProductCategory_TH   string
	ProductNameAgency_EN string
	ProductNameAgency_TH string
	ProductNameBank_EN   string
	ProductNameBank_TH   string
}

func ConvertProductTable(input string, out string) {
	// I remove this
}

func generateTestcase(product ProductTableRow) string {
	// I remove this
	return ""
}
