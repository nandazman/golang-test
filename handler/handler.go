package handler

import (
	"strconv"

	"github.com/labstack/echo"
)

type kelas struct {
	Nama string `json:"nama"`
	Umur int    `json:"umur"`
}

// Bar - untuk ngetes
func Bar(c echo.Context) error {
	param := c.Param("param")
	input := c.QueryParam("input")

	toNumber, err := strconv.Atoi(input)
	if err != nil {
		return echo.NewHTTPError(400, err.Error())
	}
	return c.JSON(200, echo.Map{
		"input":  input,
		"number": toNumber,
		"param":  param,
	})

}

// PostSomething - untuk ngepost
func PostSomething(c echo.Context) error {
	k := new(kelas)

	if err := c.Bind(k); err != nil {
		return echo.NewHTTPError(400, err.Error())
	}

	return c.JSON(200, k)
}

// CreateData - Ini masukin data ke table_xutopia
func CreateData(c echo.Context) error {
	inputData := new(TableXutopia)

	if err := c.Bind(inputData); err != nil {
		return echo.NewHTTPError(400, err.Error())
	}

	db, err := getConnection()
	if err != nil {
		return echo.NewHTTPError(500, err.Error())
	}
	defer db.Close()

	err = inputData.InsertData(db)
	if err != nil {
		return echo.NewHTTPError(500, err.Error())
	}

	return c.JSON(201, inputData)
}

// GetData - ini untuk get data by its name
func GetData(c echo.Context) error {
	// Get input from query param
	name := c.QueryParam("name")

	// Initiate data with type TableXutopia
	data := TableXutopia{
		Nama: name,
	}

	// Get connection to database for passing in query
	db, err := getConnection()
	if err != nil {
		return echo.NewHTTPError(500, err.Error())
	}
	// if function already done then close db
	defer db.Close()

	// Query data to database using method from TableXutopia
	response, err := data.GetDataByName(db)
	if err != nil {
		return echo.NewHTTPError(500, err.Error())
	}

	return c.JSON(200, response)

}

// GetAllData - Ini untuk get all data
func GetAllData(c echo.Context) error {
	var d TableXutopia

	db, err := getConnection()
	if err != nil {
		return echo.NewHTTPError(500, err.Error())
	}
	defer db.Close()

	res, err := d.GetAllData(db)
	if err != nil {
		return echo.NewHTTPError(500, err.Error())
	}

	return c.JSON(200, res)
}
