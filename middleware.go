package gallop

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"strings"
)

type IMidHandler func(*Context)

//MidFactory 构造函数
func MidFactory(h IMidHandler) gin.HandlerFunc {
	return func(c *gin.Context) {
		webContext := &Context{Context: c}
		h(webContext)
	}
}

func OpenCors(engine *gin.Engine) {
	if viper.GetBool("cors.open") {
		cfg := cors.DefaultConfig()
		origins := viper.GetString("cors.origin")
		if origins == "" {
			cfg.AllowAllOrigins = true
		} else {
			cfg.AllowOrigins = strings.Split(origins, ",")
		}
		headers := viper.GetString("cors.headers")
		if headers != "" {
			cfg.AllowHeaders = strings.Split(headers, ",")
		}
		engine.Use(cors.New(cfg))
	}
}
