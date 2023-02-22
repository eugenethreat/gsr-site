package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func home(c *gin.Context) {

	// generate a list of works

	goblinrock2 := Work{"Goblin Rock II", "the willard building", "2022", "https://f4.bcbits.com/img/a1430038869_16.jpg"}
	gromulet := Work{"Destiny's Gromulet", "mango safari", "2022", "https://f4.bcbits.com/img/a3203506842_16.jpg"}

	works := [2]Work{goblinrock2, gromulet}

	c.HTML(http.StatusOK, "index.html", gin.H{
		"work_test": works,
	})

}

func main() {
	router := gin.Default()
	// load html
	router.LoadHTMLGlob("templates/*")
	// load static
	router.Static("/assets", "./assets")
	// routes
	router.GET("/home", home)
	router.Run("localhost:8080")
}
