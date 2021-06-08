package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/jun2900/refactoryTest2/RefactoryAuth/database"
	"github.com/jun2900/refactoryTest2/RefactoryAuth/models"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	_                 = godotenv.Load()
	googleOauthConfig = &oauth2.Config{
		RedirectURL:  "http://localhost:3000/callback",
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}
	randomState = "random"
)

func initMainDatabase() {
	var err error
	mysqlUser := os.Getenv("MYSQL_USER")
	mysqlPassword := os.Getenv("MYSQL_PASSWORD")
	mysqlPort := os.Getenv("MYSQL_PORT")
	mysqlDatabase := os.Getenv("MYSQL_DATABASE")
	dsn := fmt.Sprintf("%s:%s@tcp(localhost:%s)/%s?charset=utf8&parseTime=True&loc=Local", mysqlUser, mysqlPassword, mysqlPort, mysqlDatabase)
	database.DBConn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("connection open")
	database.DBConn.AutoMigrate(&models.User{})
}

func main() {
	initMainDatabase()

	app := fiber.New()

	app.Static("/", "./home.html")

	app.Get("/login", handleLogin)
	app.Get("/callback", handleCallback)
	app.Get("/users", getUsers)

	app.Listen(":3000")
}

func handleLogin(c *fiber.Ctx) error {
	url := googleOauthConfig.AuthCodeURL(randomState)
	return c.Redirect(url, fiber.StatusTemporaryRedirect)
}

func handleCallback(c *fiber.Ctx) error {
	if c.FormValue("state") != randomState {
		fmt.Println("state is not valid")
		return c.Redirect("/", fiber.StatusTemporaryRedirect)
	}
	token, err := googleOauthConfig.Exchange(oauth2.NoContext, c.FormValue("code"))
	if err != nil {
		fmt.Println("could not get token")
		return c.Redirect("/", fiber.StatusTemporaryRedirect)
	}

	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		fmt.Println("could not get request")
		return c.Redirect("/", fiber.StatusTemporaryRedirect)
	}

	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("could not parse response")
		return c.Redirect("/", fiber.StatusTemporaryRedirect)
	}

	var user = new(models.User)
	err = json.Unmarshal(content, &user)
	if err != nil {
		panic(err)
	}
	db := database.DBConn
	db.Create(&user)

	return c.Send(content)
}

func getUsers(c *fiber.Ctx) error {
	db := database.DBConn
	var users []models.User
	db.Find(&users)
	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{"status": "success", "Users": users})
}
