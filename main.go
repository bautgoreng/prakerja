/*
package main

import (

	"fmt"

)

	func main() {
		fmt.Println("WAW")
	}
*/
/*package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type article struct {
	ID      int
	Title   string
	Content string
}

var data = []article{
	article{1, "lorem", "totem"},
	article{2, "hoream", "awawaw"},
}

func articles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {
		var result, err = json.Marshal(data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(result)
		return
	}
	http.Error(w, "", http.StatusBadRequest)
}

func main() {
	http.HandleFunc("/articles", articles)
	fmt.Println("Memulai web server di http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
*/
/*
// test echonya
package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func main() {
	//bikin instance echo baru
	e := echo.New()
	//merute / ke handlerfunction
	e.GET("/", HelloController)
	//memulai server dan me-log jika gagal
	e.Start(":8080")
}

// bikin handler simpel
func HelloController(c echo.Context) error {
	//return string "Hello World" sebagai respon body
	//dengan status http.StatusOK (200)
	return c.String(http.StatusOK, "Hello World!")
}
*/
/*
package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type User struct {
	Name  string
	Email string
}

func GetUser(c echo.Context) error {
	user := User{Name: "Fajar Ganteng", Email: "jajayganteng@bautgoreng.jp"}
	return c.JSON(http.StatusOK, user)
}

func main() {
	//baca aja yang atas lah
	e := echo.New()
	e.GET("/user", GetUser)
	e.Start(":8080")
}
*/
/*
//URL PARAM
package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type User struct {
	ID    int
	Name  string
	Email string
}

func GetUserController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	user := User{ID: id, Name: "Isback", Email: "smbbekasi@mangga1.id"}
	// Render data - respon mas json
	return c.JSON(http.StatusOK, map[string]interface{}{
		"user": user,
	})
}

func main() {
	e := echo.New()
	//routing
	e.GET("/users/:id", GetUserController)
	e.Start(":8080")
}
*/

//QUERY PARAM
/*
package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func UserSearchController(c echo.Context) error {
	//ambil data dari query param
	match := c.QueryParam("match")
	return c.JSON(http.StatusOK, map[string]interface{}{
		"match":  match,
		"result": []string{"asep", "abu bakr", "febrina"}, //hc data
	})
}

func main() {
	e := echo.New()
	// routing dengan parameter query
	e.GET("/users", UserSearchController)
	e.Start(":8080")
}
*/

// Form Value
/*
package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type User struct {
	Name  string `json:"name" form:"name"`
	Email string `json:"email" form:"email"`
}

func CreateUser(c echo.Context) error {
	name := c.FormValue("name")
	email := c.FormValue("email")

	var user User
	user.Name = name
	user.Email = email

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success create user",
		"user":     user,
	})
}

func main() {
	e := echo.New()
	//routing dengan query parameter
	e.POST("/users", CreateUser)
	e.Start(":8080")
}
*/

// Binding data
/*package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type User struct {
	Name  string `json:"name" form:"name"`
	Email string `json:"email" form:"email"`
}

func CreateUser(c echo.Context) error {
	//binding data
	user := User{}
	c.Bind(&user)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success create user",
		"user":     user,
	})
}

func main() {
	e := echo.New()
	//routing dengan query parameter
	e.GET("/users", CreateUser)
	e.Start(":8080")
}
agak lain ini*/

// DATABASE KONEK DAN GORM
//Masih error ygy
/*
package main

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
)

func InitDB() {
	//deklarasi struct konfig dan variabel stringkoneksi
	var err error
	DB, err = gorm.Open("mysql", connectionString)
	if err != nil {
		panic(err)
	}
}

var (
	DB *gorm.DB
)

type User struct {
	gorm.Model
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

func InitialMigration() {
	DB.AutoMigrate(&User{})
}

func init() {
	InitDB()
	InitialMigration()
}

// ambil semua user
func GetUsersController(c echo.Context) error {
	var users []User

	if err := DB.Find(&users).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all users",
		"users":   users,
	})
}

// buat user baru
func CreateUserController(c echo.Context) error {
	user := User{}
	c.Bind(&user)

	if err := DB.Save(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success crete new user",
		"user":    user,
	})
}

func main() {
	//buat instance echo baru
	e := echo.New()

	//route ke / handler function
	e.GET("/users", GetUsersController)
	e.POST("/users", CreateUserController)
	//mulai server
	e.Start(":8000")
}
ga beres*/

//ade su nikah
/*package main

import (
	"GO/config"
	"GO/routes"
)

func main() {
	config.InitDB()
	e := routes.New()
	e.Logger.Fatal(e.Start(":8080"))

}
*/

package main

import (
	"log"
	"os"
	"prakerja/config"
	"prakerja/routes"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	loadEnv()
	config.InitDatabase()
	e := echo.New()
	routes.InitRoute(e)
	e.Start(getPort())
}

func getPort() string {
	port := os.Getenv("DATABASE_PORT")
	if port == "" {
		port = "8000"
	}
	return ":" + port
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
