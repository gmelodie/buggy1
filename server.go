package main

import (
	"encoding/json"
	"net/http"

	logging "github.com/ipfs/go-log/v2"
	"github.com/labstack/echo/v4"
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

func (a *API) handleIndex(c echo.Context) error {
	return c.JSON(http.StatusOK, "API v2 Operational")
}

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
	e.POST("/person", api.handlePersonCreate)
	e.Logger.Fatal(e.Start(":8080"))
}
