package main

import (
	"gopkg.in/gin-gonic/gin.v1"
	"net/http"
)

func main(){
	router := gin.Default()
	http.ListenAndServe(":8080", router)
}
