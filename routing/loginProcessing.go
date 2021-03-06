package routing

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"main/commontools"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// LoginProcessing ... GET loginprocessing
func LoginProcessing(ctx *gin.Context) {
	session := sessions.Default(ctx)

	var discordid string
	fmt.Println("**************************************session.Get(ID)")
	fmt.Println(session.Get("ID"))

	code := ctx.Query("code")
	fmt.Println("**************************************code")
	fmt.Println(code)

	if code != "" {
		discordid = getToken(code)
	}
	commontools.LoginFuncion(ctx, discordid)
	ctx.Redirect(303, "/")

}

// Discordtoken ... {"access_token": "_____", "expires_in": 604800, "refresh_token": "_____", "scope": "identify email", "token_type": "Bearer"}
type Discordtoken struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
	TokenType    string `json:"token_type"`
}

func getToken(code string) (userid string) {

	var clientID string = commontools.Fixedparam("clientID")
	var clientSecret string = commontools.Fixedparam("clientSecret")

	var redirectURI string = commontools.Fixedparam("myhostname") + "loginprocessing"
	// make HTTP Client
	client := &http.Client{}

	// make Request
	var datastr string = "client_id=" + clientID + "&client_secret=" + clientSecret + "&grant_type=authorization_code&code=" + code + "&redirect_uri=" + redirectURI + "&scope=identify%20email"

	req, _ := http.NewRequest("POST", "https://discordapp.com/api/oauth2/token", bytes.NewBufferString(datastr))
	header := http.Header{}
	header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header = header

	// do Request
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// read Response
	body, _ := ioutil.ReadAll(resp.Body)

	// return JSON parse
	bytes := []byte(body)
	var discordtoken Discordtoken
	json.Unmarshal(bytes, &discordtoken)

	userdata := getMe(discordtoken.AccessToken, discordtoken.TokenType)
	fmt.Println(userdata.ID)

	return userdata.ID
}

// DiscordMe ... {"id": "_____", "username": "_____", "avatar": "_____", "discriminator": "____", "public_flags": 0, "flags": 0, "email": "_____", "verified": true, "locale": "ja", "mfa_enabled": false}
type DiscordMe struct {
	ID            string `json:"id"`
	UserName      string `json:"username"`
	Avatar        string `json:"avatar"`
	Discriminator string `json:"discriminator"`
	PublicFlags   int    `json:"public_flags"`
	Flags         int    `json:"flags"`
	Email         string `json:"email"`
	Verified      bool   `json:"verified"`
	Locale        string `json:"locale"`
	MfaEnabled    bool   `json:"mfa_enabled"`
}

func getMe(accessToken string, tokenType string) (userdata DiscordMe) {
	// make HTTP Client
	client := &http.Client{}

	// make Request
	req, _ := http.NewRequest("GET", "https://discordapp.com/api/users/@me", nil)
	header := http.Header{}
	var authorizationStr string = tokenType + " " + accessToken
	header.Add("Authorization", authorizationStr)
	req.Header = header

	// do Request
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// read Response
	body, _ := ioutil.ReadAll(resp.Body)

	// return JSON parse
	bytes := []byte(body)
	var discordMe DiscordMe
	json.Unmarshal(bytes, &discordMe)
	// fmt.Println(discordMe)
	return discordMe
}
