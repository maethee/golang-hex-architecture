package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (rH RouterHandler) formatTCPPost(c *gin.Context) {

}

func (rH RouterHandler) importTCPPost(c *gin.Context) {
	filePath := "./assets/TemplateConfig_Plan_S10P02_20200305.xlsx"
	importPath := "tempTCP.xlsx"
	excelSlice, err := rH.ucHandler.ImportTCP(filePath, importPath)

	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	c.JSON(200, excelSlice)
}
