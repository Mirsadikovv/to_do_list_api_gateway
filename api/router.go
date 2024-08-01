package api

import (
	_ "api_gateway/api/docs" //for swagger
	"api_gateway/api/handler"
	"api_gateway/config"
	"api_gateway/pkg/grpc_client"
	"api_gateway/pkg/logger"

	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Config ...
type Config struct {
	Logger     logger.Logger
	GrpcClient *grpc_client.GrpcClient
	Cfg        config.Config
}

// New ...
// @title           Swagger CRM system API
// @version         1.0
// @description     This is a CRM celler server.
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func New(cnf Config) *gin.Engine {
	r := gin.New()

	// r.Static("/images", "./static/images")

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = append(config.AllowHeaders, "*")
	// config.AllowOrigins = cnf.Cfg.AllowOrigins
	r.Use(cors.New(config))

	handler := handler.New(&handler.HandlerConfig{
		Logger:     cnf.Logger,
		GrpcClient: cnf.GrpcClient,
		Cfg:        cnf.Cfg,
	})

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "Api gateway"})
	})
	// r.Use(authMiddleware)

	///////////////////////// USER_service

	r.GET("/v1/admin/getall", handler.GetAllAdmin)
	r.GET("/v1/admin/get/:id", handler.GetAdminById)
	r.POST("/v1/admin/create", handler.CreateAdmin)
	r.PUT("/v1/admin/update", handler.UpdateAdmin)
	r.DELETE("/v1/admin/delete/:id", handler.DeleteAdmin)
	r.PATCH("/v1/admin/change_password/", handler.AdminChangePassword)
	r.POST("/v1/admin/login", handler.AdminLogin)
	r.POST("/v1/admin/register", handler.AdminRegister)
	r.POST("/v1/admin/register-confirm", handler.AdminRegisterConfirm)

	r.GET("/v1/user/getall", handler.GetAllUser)
	r.GET("/v1/user/get/:id", handler.GetUserById)
	r.POST("/v1/user/create", handler.CreateUser)
	r.PUT("/v1/user/update", handler.UpdateUser)
	r.DELETE("/v1/user/delete/:id", handler.DeleteUser)
	r.PATCH("/v1/user/change_password/", handler.UserChangePassword)
	r.POST("/v1/user/login", handler.UserLogin)
	r.POST("/v1/user/register", handler.UserRegister)
	r.POST("/v1/user/register-confirm", handler.UserRegisterConfirm)

	url := ginSwagger.URL("swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return r

}

// func authMiddleware(c *gin.Context) {
// 	auth := c.GetHeader("Authorization")
// 	if auth == "" {
// 		c.AbortWithError(http.StatusUnauthorized, errors.New("unauthorized"))
// 	}
// 	c.Next()
// }
