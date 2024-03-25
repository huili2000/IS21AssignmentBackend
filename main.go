//TODO: Role Based Access Control will be supported by TBD service (future new user stories)

package main

import (
	"localhost/paints-api/models"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func main() {

	db, err := connectDb()
	if err != nil {
		panic(err)
	}

	r := setupRoutes(db)
	if r == nil {
		return
	}

	//TODO: configure the PORT (future new user story)
	r.Run(":8080")
}

func setupRoutes(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	r.Use(cors.New(config))
	//r.Run()

	//users endpoints
	//TODO: manage Users using IAM (future new user stories)

	//auth user by name and password
	r.POST("/auth-user", func(c *gin.Context) {

		var user models.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		users, err := models.AuthUser(db, user.Password, user.Name)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, users)
	})

	//create a user
	r.POST("/users", func(c *gin.Context) {
		var user models.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		if err := models.CreateUser(db, &user); err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(201, user)
	})

	//get list of users
	r.GET("/users", func(c *gin.Context) {
		users, err := models.GetAllUsers(db)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, users)
	})

	//update user by name
	r.PUT("/users/:name", func(c *gin.Context) {
		name := c.Param("name")
		var user models.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		if err := models.UpdateUser(db, name, user.Permission, user.Role, user.Password); err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		updatedUser, err := models.GetUserByName(db, name)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, updatedUser)
	})

	//delete user by name
	r.DELETE("/users/:name", func(c *gin.Context) {
		name := c.Param("name")
		err := models.DeleteUserByName(db, name)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"message": "user deleted"})
	})

	//Paints endpoints
	// get list of paints
	r.GET("/paints", func(c *gin.Context) {

		paints, err := models.GetAllPaints(db)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, paints)
	})

	// get paint by color
	r.GET("/paints/color/:color", func(c *gin.Context) {

		color := c.Param("color")
		paint, err := models.GetPaintByColor(db, color)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		if paint == nil {
			c.JSON(404, gin.H{"error": "paint not found"})
			return
		}
		c.JSON(200, paint)
	})

	//provision paints by color and quantity
	r.PUT("/paints/provision", func(c *gin.Context) {

		var paint models.Paint
		if err := c.ShouldBindJSON(&paint); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		if err := models.ProvisionPaint(db, paint.Color, paint.Quantity); err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		updatedPaint, err := models.GetPaintByColor(db, paint.Color)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, updatedPaint)
	})

	//consume paints by color and quantity
	r.PUT("/paints/consume", func(c *gin.Context) {

		var paint models.Paint
		if err := c.ShouldBindJSON(&paint); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		if err := models.ConsumePaint(db, paint.Color, paint.Quantity); err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		updatedPaint, err := models.GetPaintByColor(db, paint.Color)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, updatedPaint)
	})

	return r
}

func connectDb() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("paints.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
