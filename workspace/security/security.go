package security

import (
	"context"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

var oauthConfig = oauth2.Config{
	ClientID:     "your_client_id",
	ClientSecret: "your_client_secret",
	RedirectURL:  "your_redirect_url",
	Endpoint: oauth2.Endpoint{
		AuthURL:  "authorization_endpoint_url",
		TokenURL: "token_endpoint_url",
	},
	Scopes: []string{"scope1", "scope2"},
}

func HandleLogin(c *gin.Context) {
	// OAuth 2.0 yetkilendirme URL'sini oluştur ve istemciyi yönlendir
	url := oauthConfig.AuthCodeURL("state", oauth2.AccessTypeOnline)
	c.Redirect(http.StatusFound, url)
}

func HandleCallback(c *gin.Context) {
	code := c.Query("code")

	// context.WithTimeout(context.Background(), time.Second*30)

	// AuthCodeURL'den aldığımız code ile erişim belirteci al
	token, err := oauthConfig.Exchange(context.Background(), code)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error exchanging code for token")
		return
	}

	// Alınan erişim belirtecini ve refresh token'ı kullanabilirsiniz
	c.String(http.StatusOK, "Access Token: %s\nRefresh Token: %s\nExpiry: %s\n", token.AccessToken, token.RefreshToken, token.Expiry)

	// Alınan erişim belirteciyle kullanıcı adını almak için bir API çağrısı yapalım
	userInfoURL := "https://www.googleapis.com/oauth2/v2/userinfo"
	client := oauthConfig.Client(context.Background(), token)
	resp, err := client.Get(userInfoURL)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error getting user info")
		return
	}
	defer resp.Body.Close()

	// API cevabını işleyelim
	userData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.String(http.StatusInternalServerError, "JSON Parsing Failed")
		return
	}

	// Kullanıcı bilgilerini yanıt olarak döndürelim
	c.String(http.StatusOK, string(userData))
}
