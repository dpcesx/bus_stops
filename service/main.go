package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"example/bus_stop"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	log.SetPrefix("server: ")
	log.SetFlags(0)

	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:8000"}

	router.Use(cors.New(config))

	router.GET("/stop/:id", getEstimatedArrivalTimes)

	router.Run("localhost:8080")
}

func getEstimatedArrivalTimes(c *gin.Context) {
	id_param := c.Param("id")

	id, err := strconv.Atoi(id_param)
	if err != nil {
		log.Fatal(err)
	}

	times, err := bus_stop.EstimateArrival(id)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(times)

	c.IndentedJSON(http.StatusOK, times)
}
