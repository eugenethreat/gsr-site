package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// https://www.allhandsontech.com/programming/golang/how-to-use-sqlite-with-go/

func home(c *gin.Context) {

	// generate a list of works

	goblinrock2 := Work{"Goblin Rock II", "the willard building", "2022", "https://f4.bcbits.com/img/a1430038869_16.jpg"}
	gromulet := Work{"Destiny's Gromulet", "mango safari", "2022", "https://f4.bcbits.com/img/a3203506842_16.jpg"}
	fuckaround := Work{"fuck around and find out", "H2OhNo!", "2022", "https://f4.bcbits.com/img/a2613250794_16.jpg"}

	works := []Work{goblinrock2, gromulet}

	// works = append(works, Work{"LEG 3", "LEG", "2022", ""})
	works = append(works, fuckaround)

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
