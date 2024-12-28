package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// album represents data about a record album
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// albums slice to seed record album data
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbums)
	router.GET("/albums/:id", getAlbumByID)

	router.Run("localhost:8080")
}

/*
TODO:
When client makes request to GET /albums,
return all the albums as JSON
*/

// getAlbums responds with list of all albums as JSON
func getAlbums(c *gin.Context) {
	/*
		c = variable name to refer to the context
		* = indicates c is a pointer to gin.Context
		ginContext = struct provided by Gin framework

				gin.Context carries:
				- request details
				- validates (everything right format)
			and serializes JSON (converting obj into str)
	*/

	//c.IndentedJSON(http.StatusOK, albums)
	c.JSON(http.StatusOK, albums)
}

func postAlbums(c *gin.Context) {
	var newAlbum album

	//Bind the received JSON to newAlbum
	if err := c.BindJSON(&newAlbum); err != nil {
		//if err -> if an error occurs during binding, it will assign error to err variable
		//which means err will not be nil and immediately returns the fn
		return
	}

	//add new album to the slice
	albums = append(albums, newAlbum)
	c.JSON(http.StatusCreated, newAlbum)
}

func getAlbumByID(c *gin.Context) {
	//id sent by the client
	id := c.Param("id")

	//look for the album with the id
	for _, album := range albums {
		if album.ID == id {
			c.JSON(http.StatusOK, album)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"msg": "album not found"})
}
