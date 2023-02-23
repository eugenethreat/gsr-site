package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func new_home(c *gin.Context) {

	// generate a list of works
	// https://www.allhandsontech.com/programming/golang/how-to-use-sqlite-with-go/

	works := get_works()

	c.HTML(http.StatusOK, "base.html", gin.H{
		"work_test": works,
	})

}

// return all artists
func get_artists() []Artist {
	artists := []Artist{}

	// LEG
	sharkies_2 := Work{"Sharkies 2", "LEG", "2019", "/assets/album_art/leg/sharkies2.png"}
	johnstown := Work{"Johnstown", "LEG", "2020", "/assets/album_art/leg/johnstown.png"}
	leg3 := Work{"LEG 3", "LEG", "2023", "/assets/album_art/leg/leg3.png"}

	leg_works := []Work{sharkies_2, johnstown, leg3}
	leg_band := Artist{1, "LEG", "/assets/artists/leggy.jpg", leg_works, "LEG produces one of a kind musical adventures that explore the human condition. By mixing progressive electronic music and hip hop, LEG creates a unique textural blend of sounds that brings listeners to surreal re-imaginings of not so iconic locations."}

	// TWB
	goblinrock2 := Work{"Goblin Rock II", "the willard building", "2022", "https://f4.bcbits.com/img/a1430038869_16.jpg"}
	years99 := Work{"99 years of contract hell", "the willard building", "2023", "/assets/album_art/twb/coveridea1.png"}
	poppy := Work{"Poppy", "the willard building", "2023", "/assets/album_art/twb/lilcat.jpg"}

	twb_works := []Work{goblinrock2, years99, poppy}
	twb_band := Artist{2, "the willard building", "/assets/artists/twb artist.jpg", twb_works, "it is unclear if the willard building is a band or a building but somehow eugene, ryan, grant, derek, and andrew are involved."}

	// mango
	gromulet := Work{"Destiny's Gromulet", "mango safari", "2022", "https://f4.bcbits.com/img/a3203506842_16.jpg"}
	momotaro := Work{"Momotaro", "mango safari", "2018", "https://f4.bcbits.com/img/a0956356550_16.jpg"}

	mango_works := []Work{gromulet, momotaro}
	mango_band := Artist{3, "mango safari", "/assets/artists/mango artist.jpg", mango_works, "Mango Safari is a two piece band from Philadelphia PA. The duo, Ryan Keating and Grant Cheung, started playing Beatles covers together in middle school."}

	// h2ohno
	fuckaround := Work{"fuck around and find out", "H2OhNo!", "2022", "https://f4.bcbits.com/img/a2613250794_16.jpg"}

	h2o_works := []Work{fuckaround}
	h2o_band := Artist{4, "H2OhNo!", "/assets/artists/h2o artist.jpg", h2o_works, "Alchemized from the State College diy punk scene, H₂O No! members Brody McElwain and Sean Theoclitus offer an explosive, noisy sound as a bass+drum duo. Their energetic live performances are driven by angsty lyrics directed at elitist white fucks who abuse their authority at the expense of the working class. H₂O No! also enjoys character portraits that are reflective of their own insecurities."}

	artists = append(artists, leg_band, twb_band, mango_band, h2o_band)

	return artists
}

// get all works
func get_works() []Work {
	works := []Work{}

	// get all artists
	artists := get_artists()

	// now for all artists, get all works
	// just for demo purposes...
	for i, s := range artists {

		println(i)
		artist_works := s.Works

		for x, y := range artist_works {
			println(x)
			works = append(works, y)
		}
	}

	return works
}

func all_artists(c *gin.Context) {
	artists := get_artists()

	c.HTML(http.StatusOK, "all_artists.html", gin.H{
		"artists": artists,
	})

}

func artist(c *gin.Context) {

	pk := c.Param("id")
	pk_int, err := strconv.Atoi(pk)

	println(err)

	artists := get_artists()

	if err == nil {
		println("pk2", pk_int)

		c.HTML(http.StatusOK, "artist.html", gin.H{
			"artist": artists[pk_int-1],
		})

	}
}

func main() {
	router := gin.Default()
	// load html
	router.LoadHTMLGlob("templates/*/*")
	// load static
	router.Static("/assets", "./assets")
	// routes
	router.GET("/home/", new_home)
	router.GET("/artists/", all_artists)
	router.GET("/artists/:id/", artist)
	// router.Run("localhost:8080")
	router.Run("localhost:80")
}
