package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct{
	ID string `json:"id"`
	Title string `json:"title"`
	Artist string `json:"artist"`
	Price float64 `json:"price"`
}

var albums=[]album{
	{ID:"1", Title:"Blue Train", Artist:"John Coltrane", Price: 54.99},
	{ID:"2", Title:"Jeru", Artist:"Gerry Mulligan", Price: 17.99},
	{ID:"3", Title:"Sarah Varughan", Artist:"Tarkan", Price: 37.99},
}
func getAlbums(c *gin.Context){
	c.IndentedJSON(http.StatusOK, albums)
}
func postAlbums(c *gin.Context){
	var newAlbum album
	if err:=c.BindJSON(&newAlbum);err!=nil{
		return
	}
	albums=append(albums, newAlbum)
	c.IndentedJSON(http.StatusOK, newAlbum)
}
func getAlbumByID(c *gin.Context){
	fmt.Println(c)
	id:=c.Param("id")

	for _, a :=range albums{
		if a.ID==id{
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message":"Album not found!"})
}

func main(){
	router:=gin.Default()
	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.Run("localhost:8080")

}