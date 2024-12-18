package route

import (
	"github.com/gin-gonic/gin"
	"gotest/service"
)

func InitRouter() *gin.Engine {

	r := gin.Default()

	api := r.Group("/api")
	{
		api.GET("/ping", service.Ping)
		api.GET("/xhs", service.GetXhs)
		api.GET("/weixin", service.GetWeixin)
		api.GET("/dy", service.GetDy)
		api.GET("/sph", service.GetSph)
		api.GET("/ks", service.GetKs)
		api.GET("/bil", service.GetBill)
		api.GET("/doit", service.Doit)
		api.GET("/deleteAll", service.DeleteAll)
	}

	return r
}
