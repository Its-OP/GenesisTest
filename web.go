package main

import (
	"btcRate/docs"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

// @title GSES2 BTC application API
// @version 1.0.0
// @description This is a sample server for a BTC to UAH rate application.
// @host localhost:8080
// @BasePath /api

func main() {
	r := gin.Default()

	docs.SwaggerInfo.BasePath = "/api/v1"
	api := r.Group("/api/v1")
	{
		api.GET("/rate", GetRate)
		api.POST("/subscribe", Subscribe)
		api.POST("/sendEmails", SendEmails)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run(":8080") // Run on port 8080
}

// @Summary Get current BTC to UAH rate
// @Description Get the current rate of BTC to UAH using any third-party service with public API
// @Tags rate
// @Produce  json
// @Success 200 {number} number "Successful operation"
// @Failure 400 {object} string "Invalid status value"
// @Router /rate [get]
func GetRate(c *gin.Context) {
	// TODO: Add logic to fetch BTC to UAH rate and return it

	c.JSON(http.StatusOK, gin.H{
		"rate": 0, // Replace this with your actual rate
	})
}

// @Summary Subscribe email to get BTC rate
// @Description Add an email to the database if it does not exist already
// @Tags subscription
// @Accept  x-www-form-urlencoded
// @Produce  json
// @Param email formData string true "Email to be subscribed"
// @Success 200 {object} string "E-mail added"
// @Failure 409 {object} string "E-mail already exists in the database"
// @Router /subscribe [post]
func Subscribe(c *gin.Context) {
	// TODO: Add logic to subscribe an email to the mailing list

	email := c.PostForm("email")
	if email == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Email is required",
		})
		return
	}

	// TODO: Check if email is already in the database
	// If it is, return a 409 status code

	c.JSON(http.StatusOK, gin.H{
		"message": "E-mail added",
	})
}

// @Summary Send email with BTC rate
// @Description Send the current BTC to UAH rate to all subscribed emails
// @Tags subscription
// @Produce  json
// @Success 200 {object} string "E-mails sent"
// @Router /sendEmails [post]
func SendEmails(c *gin.Context) {
	// TODO: Add logic to send the current BTC to UAH rate to all subscribed emails

	c.JSON(http.StatusOK, gin.H{
		"message": "E-mails sent",
	})
}
