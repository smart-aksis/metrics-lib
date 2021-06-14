package metrics

import "github.com/gin-gonic/gin"

type HealthStatusData struct {
	name string
	data map[string]interface{}
}

func GenerateHealthStatusData(_name string, _data map[string]interface{}) HealthStatusData{
	return HealthStatusData{
		name: _name,
		data: _data,
	}
}

type HealthStatus func() HealthStatusData

func HealthCheck(router *gin.Engine, status... HealthStatus){
	health := router.Group("/health")
	health.GET("check", func(context *gin.Context) {
		var result map[string]interface{}
		result = make(map[string]interface{})

		basic:=make(map[string]interface{})
		basic["status"]="UP"

		result["Basic"]=basic
		if status != nil {
			for _, healthStatus := range status {
				index:=healthStatus()
				result[index.name] = index.data
			}
		}
		context.JSON(200, result)
	})
}

