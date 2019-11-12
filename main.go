package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

type introspectIn struct {
	Token string `form:"token"`
}

func Bootstrap() {
	log.Println("Ping module bootstrap.")
}

func Router(r *gin.Engine) {
	r.POST("/introspect", PostIntrospectHandler)
}

func PostIntrospectHandler(c *gin.Context) {
	var input introspectIn
	c.Bind(&input)
	fmt.Println(input.Token)
	c.JSON(200, gin.H{"token": input.Token})
}
