package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	v1 := r.Group("/api/v1")
	{
		v1.GET("person", getPersons)
		v1.GET("person/:id", getPersonById)
		v1.POST("person", addPerson)
		v1.PUT("person/:id", updatePerson)
		v1.DELETE("person/:id", deletePerson)
		v1.OPTIONS("person", options)
	}
}

func getPersons(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "getPersons Called"})
}

func getPersonById(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"message": "getPersonsById " + id + " Called"})
}

func addPerson(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "addPerson Called"})
}

func updatePerson(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "updatePerson Called"})
}

func deletePerson(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"message": "deletePersons " + id + " Called"})
}

func options(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "options Called"})
}
