package server

import (
	"github.axa.com/maethee-chakkuchantorn/psp-improvement/internal/app/uc"
	"github.com/gin-gonic/gin"
)

type RouterHandler struct {
	ucHandler uc.Handler
}

func NewRouter(i uc.Handler) RouterHandler {
	return RouterHandler{
		ucHandler: i,
	}
}

func (rh RouterHandler) SetRoutes(r *gin.Engine) {
	api := r.Group("/api")
	rh.formatRoutes(api)
}

func (rH RouterHandler) websocketRoutes(api *gin.RouterGroup) {
}

func (rH RouterHandler) formatRoutes(api *gin.RouterGroup) {
	profiles := api.Group("/format")
	profiles.GET("/tcp", rH.formatTCPPost)
	profiles.POST("/tcp", rH.importTCPPost)
}