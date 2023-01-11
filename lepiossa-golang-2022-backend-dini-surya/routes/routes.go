package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"lepiosa/controller"
	"lepiosa/middlewares"

	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowHeaders = []string{"Content-Type", "X-XSRF-TOKEN", "Accept", "Origin", "X-Requested-With", "Authorization"}

	// To be able to send tokens to the server.
	corsConfig.AllowCredentials = true
	// OPTIONS method for ReactJS
	corsConfig.AddAllowMethods("OPTIONS")

	r.Use(cors.New(corsConfig))

	// set db to gin context
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})

	r.POST("/register", controller.Register)
	r.POST("/login", controller.Login)

	r.GET("/laptop", controller.GetAllLaptop)
	//r.POST("/laptop", controller.CreateLaptop)
	r.GET("/laptop/:id", controller.GetLaptopById)
	//r.PATCH("/laptop/:id", controller.UpdateLaptop)
	//r.DELETE("laptop/:id", controller.DeleteLaptop)

	laptopMiddlewareRoute := r.Group("/laptop")
	laptopMiddlewareRoute.Use(middlewares.JwtAuthMiddleware())
	laptopMiddlewareRoute.POST("/", controller.CreateLaptop)
	laptopMiddlewareRoute.PATCH("/:id", controller.UpdateLaptop)
	laptopMiddlewareRoute.DELETE("/:id", controller.DeleteLaptop)

	r.GET("/user", controller.GetAllUser)
	r.POST("/user", controller.CreateUser)
	r.GET("/user/:id", controller.GetUserById)
	r.PATCH("/user/:id", controller.UpdateUser)
	r.DELETE("user/:id", controller.DeleteUser)

	r.GET("/rating", controller.GetAllRating)
	r.POST("/rating", controller.CreateRating)
	r.GET("/rating/:id", controller.GetRatingById)
	r.GET("/rating/:id/laptop", controller.GetLaptopsByRatingId)
	r.PATCH("/rating/:id", controller.UpdateRating)
	r.DELETE("rating/:id", controller.DeleteRating)

	r.GET("/brand", controller.GetAllBrand)
	r.POST("/brand", controller.CreateBrand)
	r.GET("/brand/:id", controller.GetBrandById)
	r.GET("/brand/:id/laptop", controller.GetLaptopByBrandId)
	r.PATCH("/brand/:id", controller.UpdateBrand)
	r.DELETE("brand/:id", controller.DeleteBrand)

	r.GET("/os", controller.GetAllOs)
	r.POST("/os", controller.CreateOperasiSistem)
	r.GET("/os/:id", controller.GetOperasiSistemById)
	r.GET("/os/:id/laptop", controller.GetLaptopsByOperasiSistemId)
	r.PATCH("/os/:id", controller.UpdateOperasiSistem)
	r.DELETE("os/:id", controller.DeleteOperasiSistem)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
