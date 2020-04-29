package main

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"os"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"io/ioutil"
	"net/http"
)

type Service struct {
	client *oauth2.Config

}

// Return homepage from templates/
func (s Service) HomepageHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html.tpl", gin.H{})
}

func (s Service) LoginHandler(c *gin.Context) {
	url := s.client.AuthCodeURL("oioioi")
	c.Redirect(http.StatusTemporaryRedirect, url)
}

func (s Service) AuthHandler(c *gin.Context) {
	content, err := s.getUserInfo(c.Query("state"), c.Query("code"))
	if err != nil {
		fmt.Println(err.Error())
		c.Redirect(http.StatusTemporaryRedirect, "/")
		return
	}

	c.Data(200, "application/json; charset=utf-8", content)
}

func (s Service) getUserInfo(state string, code string) ([]byte, error) {
	if state != "oioioi" {
		return nil, fmt.Errorf("invalid oauth state")
	}

	// Exchange converts an authorization code into a token.
	token, err := s.client.Exchange(oauth2.NoContext, code)
	if err != nil {
		return nil, fmt.Errorf("code exchange failed: %s", err.Error())
	}

	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		return nil, fmt.Errorf("failed getting user info: %s", err.Error())
	}

	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed reading response body: %s", err.Error())
	}

	return contents, nil
}

// func AuthRequired(c *gin.Context) {

// }

func main() {
	service := &Service{
		client: &oauth2.Config{
			RedirectURL:  os.Getenv("REDIRECT_URL"),
			ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
			ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
			Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
			Endpoint:     google.Endpoint,
		},
	}


	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	router.GET("/", service.HomepageHandler)
	router.GET("/login", service.LoginHandler)
	router.GET("/auth", service.AuthHandler)

	router.Run(":8080")
}