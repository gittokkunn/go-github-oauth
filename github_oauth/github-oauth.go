package github_oauth

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)


type CredentialInfo struct {
	AccessToken string `json:"access_token"`
	Scope      string `json:"scope"`
	TokenType string `json:"token_type"`
}

func EnvLoad() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	ClientID = os.Getenv("CLIENT_ID")
	ClientSecret = os.Getenv("CLIENT_SECRET")
}

var(
	ClientID string
	ClientSecret string
	AccessToken string
)

// Githubで認証済みか判定
func LoginHome(c *gin.Context) {
	if AccessToken != "" {
		c.HTML(http.StatusOK, "index.html", nil)
	}else{
		c.Redirect(http.StatusMovedPermanently, "/login")
	}
}

// Github認証画面にリダイレクト
func RedirectAuthrize(c *gin.Context) {
	EnvLoad()
	Scope := ""
	State := "ss"
	RedirectAuthrizeClient(c, ClientID, Scope, State)
}

// クライアントIDを指定してリダイレクト
func RedirectAuthrizeClient(c *gin.Context, clientID string, scope string, state string) {
	authURL := "https://github.com/login/oauth/authorize?client_id=" + clientID + "&scope=" + scope + "&state" + state
	c.Redirect(http.StatusMovedPermanently, authURL)
}

// アクセストークンを取得
func GetAccessToken(c *gin.Context) {
	EnvLoad()
	GetAccessTokenClient(c, ClientID, ClientSecret)
	c.Redirect(http.StatusMovedPermanently, "/")
}

// クライアントID, クライアントパスをしていしてアクセストークンを取得
func GetAccessTokenClient(c *gin.Context, clientID string, clientSecret string) (*CredentialInfo){
	code := c.Request.URL.Query().Get("code")
	state := c.Request.URL.Query().Get("state")
	if state == "" {
		fmt.Println("state is empty")
	}
	values := url.Values{}
	values.Add("code", code)
	values.Add("client_id", clientID)
	values.Add("client_secret", clientSecret)
	req, err := http.NewRequest(
		"POST",
		"https://github.com/login/oauth/access_token",
		strings.NewReader(values.Encode()),
	)
	if err != nil {
		panic(err)
	}

	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)

	var cre *CredentialInfo
	json.Unmarshal(byteArray, &cre)
	err = setAccessToken(cre)
	if err != nil {
		panic(err)
	}
	return cre
}

func setAccessToken(cre *CredentialInfo) error {
	AccessToken = cre.AccessToken
	if AccessToken == "" {
		err := errors.New("accessToken is empty")
		return err
	}
	return nil
}