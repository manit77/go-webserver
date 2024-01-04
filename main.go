package main

import (
	"fmt"
	"net/http"
	"os"
	"reflect"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	CONFIG_WEBSERVER_PORT := "8004"

	fmt.Println("intializing webserver")

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use((middleware.Recover()))

	var webc webController

	//get all routes from controller
	routes := webc.GetRoutes()

	//serve static files
	e.GET("/public/*", serveFiles)

	//loop through routes and handle requests
	for _, r := range routes {
		fmt.Printf("Route: %v\n", r.routePath)

		if r.method == GET {

			//bind the route to the e.GET
			func(route webRoute) {
				e.GET(route.routePath, func(c echo.Context) error {

					addCorsHeaders(c)

					if c.Request().Method == "OPTIONS" {
						return c.NoContent(http.StatusOK)
					}

					fmt.Printf("GET: %v\n", route.routePath)

					//call the function and passquery string
					var err error
					returnData, err := route.funcCall(c.QueryString())

					if err != nil {
						return err
					} else {
						//convert to string
						str := fmt.Sprintf("%v", returnData)

						//return to client
						return c.String(http.StatusOK, str)
					}
				})
			}(r)

		} else if r.method == POST {

			//bind the route to the e.POST
			func(route webRoute) {
				e.POST(route.routePath, func(c echo.Context) error {
					fmt.Printf("POST: %v\n", route.routePath)

					addCorsHeaders(c)

					if c.Request().Method == "OPTIONS" {
						return c.NoContent(http.StatusOK)
					}

					//create a new object
					var postData = reflect.New(route.postDataType).Interface()

					//bind the object from the body of the request
					if err := c.Bind(postData); err != nil {
						return err
					}

					//call the function
					returnData, err := r.funcCall(postData)
					if err != nil {
						return err
					}

					//return json data
					return c.JSON(http.StatusOK, returnData)

				})
			}(r)

		}
	}

	fmt.Println("starting webserver")

	go func() {
		e.Logger.Fatal(e.Start(":" + CONFIG_WEBSERVER_PORT))
	}()

	fmt.Println("*** press any key to exit ***")
	var input string
	fmt.Scanln(&input)
	fmt.Println("exiting webserver")

}

func bindRoute(route webRoute, c echo.Context) {

}

func serveFiles(c echo.Context) error {

	path := c.Param("*")
	dir := "./public/" + path

	file, err := os.Open(dir)
	if err != nil {
		return err
	}
	defer file.Close()

	info, err := file.Stat()
	if err != nil {
		return err
	}

	if info.IsDir() {

		files, err := os.ReadDir(dir)
		if err != nil {
			return err
		}

		var fileList string
		for _, f := range files {
			fileList += f.Name() + "\n"
		}
		return c.String(http.StatusOK, fileList)
	}

	return c.Stream(http.StatusOK, info.Name(), file)
}

func addCorsHeaders(c echo.Context) {

	if origin := c.Request().Header.Get("Origin"); origin != "" {
		c.Response().Header().Set("Access-Control-Allow-Origin", origin)
		c.Response().Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Response().Header().Set("Access-Control-Allow-Headers", "Accept, Accept-Language, Content-Type, Authorization")
	}
}
