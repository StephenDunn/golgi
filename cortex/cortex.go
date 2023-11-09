package cortex

import (
	"context"
	"fmt"
	"golgi/cortex/css"
	"golgi/cortex/js"
	"golgi/cortex/view"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
)

func Startup() {
	e := echo.New()

	e.GET("/", index)
	e.GET("/main", main)
	e.GET("/secondPage", secondPage)
	e.GET("/time", timeNow)

	e.GET("/css/layout", cssLayout)
	e.GET("/js/ui", jsUi)

	e.GET("/dothing/:id", doThing)

	e.Logger.Fatal(e.Start(":8091"))

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}

func index(c echo.Context) error {
	return c.Render(http.StatusOK, "index", view.Shared)
}

func main(c echo.Context) error {
	return c.String(http.StatusOK, view.Main)
}

func secondPage(c echo.Context) error {
	return c.String(http.StatusOK, view.SecondPage)
}

func timeNow(c echo.Context) error {
	return c.String(http.StatusOK, time.Now().String())
}

func cssLayout(c echo.Context) error {
	return c.String(http.StatusOK, css.Layout)
}

func jsUi(c echo.Context) error {
	return c.String(http.StatusOK, js.Ui)
}

func doThing(c echo.Context) error {
	id := string(c.Param("id"))

	fmt.Println(id)
	output := id + id

	return c.String(http.StatusOK, output)
}
