package middlewares

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
)

func UseMiddlewares(eg *gin.Engine) {
	eg.Use(setCors())
	if err := eg.SetTrustedProxies(nil); err != nil {
		log.Fatalln(err.Error())
	}
}

func setCors() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{`GET`, `POST`, `OPTIONS`, `PATCH`, `DELETE`},
		AllowHeaders:     []string{`Content-Type`, `Origin`, `Authorization`},
		AllowCredentials: true,
	})
}
