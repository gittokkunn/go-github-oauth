package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gittokkunn/go-github-oauth/github_oauth"
	//"log"
)

func main() {
	r := gin.Default()
	//r.Static("/img", "./client/img")
	//r.Static("/javascript", "./client/javascript")
	r.LoadHTMLGlob("./index.html")
	r.GET("/", github_oauth.LoginHome)
	r.GET("/login", github_oauth.RedirectAuthrize)
	r.GET("/callback", github_oauth.GetAccessToken)
	r.Run(":3000")

}
