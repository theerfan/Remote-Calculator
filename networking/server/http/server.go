package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo"
	"github.com/theerfan/Remote-Calculator/util"
	"io/ioutil"
	"net/http"
)

type equation = util.Equation
const servAddr = "127.0.0.1:3333"

// CalculateHandler echos back the request as a response
func CalculateHandler(bodyBytes []byte) float64 {
	var data map[string]interface{}
	err := json.Unmarshal(bodyBytes, &data)
	if err != nil {
		fmt.Println(err)
	}
	var eq equation
	err = util.FillStruct(&eq, data)
	if err != nil {
		fmt.Println(err)
	}
	return util.Calculate(eq)
}


//RunServer starts a new server on localhost:3333.
func RunServer() {
	e := echo.New()
	e.POST("/calculate", func(c echo.Context) error {
		var bodyBytes []byte
		if c.Request().Body != nil {
			bodyBytes, _ = ioutil.ReadAll(c.Request().Body)
			fmt.Println(string(bodyBytes))
		}
		// Restore the io.ReadCloser to its original state
		c.Request().Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
		ans := CalculateHandler(bodyBytes)
		return c.String(http.StatusOK, fmt.Sprint(ans))
	})
	e.Logger.Fatal(e.Start(servAddr))
}