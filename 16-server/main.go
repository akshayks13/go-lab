package main

import "github.com/gin-gonic/gin"

func main() {
    r := gin.Default() // Create a Gin router with default middleware (logger and recovery)

    // Simple GET route
    r.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "Hi",
        })
    })

    // GET route with a parameter
    r.GET("/hello/:name", func(c *gin.Context) {
        name := c.Param("name") // Get the 'name' parameter from the URL
        c.JSON(200, gin.H{
            "message": "Hi " + name,
        })
    })

    // POST route
    r.POST("/add", func(c *gin.Context) {
        var user struct
        {
            Id  int    `json:"id"`   // Bind the JSON body to this field
            Name string `json:"name"` // Bind the JSON body to this field
            Age  int    `json:"age"`  // Bind the JSON body to this field
        }
        if err := c.ShouldBindJSON(&user); err != nil {
            c.JSON(400, gin.H{"error": err.Error()}) // Return an error
            return
        }
        c.JSON(200, gin.H{
            "message": "User added",
            "user":    user, // Return the user data
        })
    })

    // PUT route
    r.PUT("/update/:id", func(c *gin.Context) {
        id := c.Param("id") // Get the 'id' parameter from the URL
        var user struct {
            Name string `json:"name"` // Bind the JSON body to this field
            Age  int    `json:"age"`  // Bind the JSON body to this field
        }
        if err := c.ShouldBindJSON(&user); err != nil {
            c.JSON(400, gin.H{"error": err.Error()}) // Return an error
            return
        }
        c.JSON(200, gin.H{
            "message": "User updated",
            "id":      id,
            "user":    user, // Return the updated user data
        })
    })

    // DELETE route
    r.DELETE(("/delete/:id"), func(c *gin.Context) {
        id := c.Param("id") // Get the 'id' parameter from the URL
        c.JSON(200, gin.H{
            "message": "User deleted",
            "id":      id, // Return the deleted user ID
        })
    })

    r.Run(":8080")
}

