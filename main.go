package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"learn/web-gin/handler"
	"encoding/json"
	"log"
)

func main() {
	router := gin.Default()
	router.GET("/albums", handler.GetAlbums)
	router.GET("/albums/:id", handler.GetAlbumByID)
	router.POST("/albums", handler.PostAlbums)

	v1 := router.Group("/v1")
	{
		v1.POST("/testUrl",handler.TestBindURLencode)
	}

	router.POST("/test", func(ctx *gin.Context) {
		fmt.Println(ctx.GetRawData())
		data := []byte(`[{"href":"/publication/192a7f47-84c1-445e-a615-ff82d92e2eaa/article/b;version=1493756861347"},{"href":"/publication/192a7f47-84c1-445e-a615-ff82d92e2eaa/article/a;version=1493756856398"}]`)
		var objmap []map[string]interface{}
		if err := json.Unmarshal(data, &objmap); err != nil {
			log.Fatal(err)
		}
		// fmt.Println(objmap[0]["href"])
	})
	router.Run("localhost:8080")
}
