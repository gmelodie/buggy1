package main

import (
	"encoding/json"
	"net/http"

	_ "github.com/gmelodie/buggy1/docs"

	logging "github.com/ipfs/go-log/v2"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Person struct {
	ID        uint   `gorm:"primarykey" json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Age       uint   `json:"age"`
}

type API struct {
	DB *gorm.DB
}

// handleIndex godoc
// @Summary  API greeting message
// @Produce  json
// @Success  200
// @Router   / [get]
func (a *API) handleIndex(c echo.Context) error {
	return c.JSON(http.StatusOK, "API v2 Operational")
}

// handlePersonCreate godoc
// @Summary  Create a new Person
// @Accept   json
// @Produce  json
// @Param    firstName  body      string  true  "first name"
// @Param    lastName   body      string  true  "last name"
// @Param    age        body      int     true  "age"
// @Success  202        {object}  Person
// @Router   /person [post]
func (a *API) handlePersonCreate(c echo.Context) error {
	newPerson := new(Person)
	err := json.NewDecoder(c.Request().Body).Decode(&newPerson)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	a.DB.Save(&newPerson)

	return c.JSON(http.StatusCreated, newPerson)
}

func createDatabase() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("people.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&Person{})

	return db, nil
}

// @title People API
// @version 2.0

// @BasePath /
// @schemes http
func main() {
	var log = logging.Logger("api")
	var api API
	var err error

	api.DB, err = createDatabase()
	if err != nil {
		log.Fatal("could not create or open database: %s", err)
	}

	e := echo.New()
	e.GET("/", api.handleIndex)
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.POST("/person", api.handlePersonCreate)
	e.Logger.Fatal(e.Start(":8080"))
}
