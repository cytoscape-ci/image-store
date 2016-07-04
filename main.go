package main

import (
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"os"
)

var format = map[string]string{}

func getGinRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func setRoutes(r *gin.Engine) {
	r.POST("/image/:format/:network_id", func(c *gin.Context) {
		picture := c.Request.Body
		id := c.Param("network_id")
		new_format := c.Param("format")
		format[id] = new_format
		out, err := os.Create("./" + id + "." + new_format)
		if err != nil {
			log.Fatal(err)
		}
		defer out.Close()
		_, err = io.Copy(out, picture)
		if err != nil {
			log.Fatal(err)
		}
		c.JSON(200, gin.H{"message": "Image submission successful."})
	})

}

func setImageServer(r *gin.Engine) {
	r.Use(static.Serve("/", static.LocalFile("", true)))
}

func main() {
	router := getGinRouter()
	setRoutes(router)
	setImageServer(router)
	router.Run("0.0.0.0:80")
}
