package handler

import (
	"fmt"

	"github.com/fuadsuleyman/go-auth1/pkg/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)
type Handler struct{
	services *service.Service 
}

// this is dependency injection
func NewHandler(services *service.Service) *Handler{
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine{
	gin.SetMode(gin.ReleaseMode)
	
	// router := gin.New()
	router := gin.Default()


	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowHeaders = []string{"*"}
	// config.AddAllowHeaders("Authorization", "language")
	config.AllowCredentials = true
	fmt.Println("*******************Fuad")


	router.Use(cors.New(config))


	// router.Use(CORSMiddleware())

	// router.Use(cors.Default())

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api", h.userIdentity)
	{
		lists := api.Group("/lists")
		{
			lists.POST("/", h.createItem)
		}		
	}

	return router
}

func CORSMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }

        c.Next()
    }
}


// func CORSMiddleware() gin.HandlerFunc {
//     return func(c *gin.Context) {

//         c.Header("Access-Control-Allow-Origin", "*")
//         c.Header("Access-Control-Allow-Credentials", "true")
//         c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
//         c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")
//         c.Header("Vary", "Accept,Orign")
//         c.Header("Allow", "POST,OPTION")
//         c.Header("Referrer-Policy", "same-orign")
//         c.Header("Connection", "keep-alive")

//         if c.Request.Method == "OPTIONS" {
//             c.AbortWithStatus(204)
//             return
//         }

//         c.Next()
//     }
// }