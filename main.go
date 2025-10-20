package main

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)

type blog struct {
	Title		string		`title:"string"`
	Content 	string		`content:"string"`
	Category 	string		`category:"string"`
	Tags		[]string	`tags:"[]string"`
}

var blogs = []blog{}

func addBlog(context *gin.Context) {
	var newBlog blog

	if err:= context.BindJSON(&newBlog); err!= nil{
		return 
	}

	blogs = append(blogs, newBlog)

	context.IndentedJSON(http.StatusCreated, newBlog)
	fmt.Printf("ClientIP: %s\n", context.ClientIP())
}

func main() {
	r := gin.Default()
	r.SetTrustedProxies([]string{"192.168.5.61"})
	r.POST("/posts", addBlog)
	fmt.Println("Hosting on localhost:9090")
	r.Run("localhost:9090")
}
