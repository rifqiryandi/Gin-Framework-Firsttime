package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"learn/web-gin/model"
	"net/http"
	"encoding/json"
	"log"
)



// getAlbums responds with the list of all albums as JSON.
func GetAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, model.Albums)
}

// postAlbums adds an album from JSON received in the request body.
func PostAlbums(c *gin.Context) {
	var newAlbum model.Album
	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}
	fmt.Println(newAlbum)

	// Add the new album to the slice.
	model.Albums = append(model.Albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func GetAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range model.Albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func TesInterface(c *gin.Context) {
	data := []byte(`[{"href":"/publication/192a7f47-84c1-445e-a615-ff82d92e2eaa/article/b;version=1493756861347"},{"href":"/publication/192a7f47-84c1-445e-a615-ff82d92e2eaa/article/a;version=1493756856398"}]`)
	var objmap []map[string]interface{}
	if err := json.Unmarshal(data, &objmap); err != nil {
		log.Fatal(err)
	}
	fmt.Println(objmap[0]["href"]) // to parse out your value
}

func TestBindURLencode(c *gin.Context)  {
	var form model.LoginForm
	if c.ShouldBind(&form) == nil {
		if form.User == "user" && form.Password == "ryan" {
			c.JSON(200,gin.H{"status":"success"})
		}else{
			c.JSON(400,gin.H{"status":"error"})
		}
	}
	fmt.Println(form)
}
