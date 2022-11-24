package main

import (
	"errors"
	"fmt"
	"io/fs"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"

	_ "github.com/climactivity/cy-cms-2/migrations"
)

func defaultPublicDir() string {
	if strings.HasPrefix(os.Args[0], os.TempDir()) {
		// most likely ran with go run
		return "./pb_public"
	}

	return filepath.Join(os.Args[0], "../pb_public")
}

type MessageResponse struct {
	Message string `json:"message"`
}

func CurrentMonthlyChallegne(app *pocketbase.PocketBase) echo.HandlerFunc {

	return func(c echo.Context) error {
		// log.Default().Println("HELLo")

		// return c.JSON(200, MessageResponse{Message: "Hello"})
		collection, err := app.Dao().FindCollectionByNameOrId("monthly_challenges")

		if err != nil {
			res := MessageResponse{
				Message: "Server unhealthy",
			}
			return c.JSON(500, res)
		}

		currentTime := time.Now().Format("2006-01-02")
		log.Default().Println("Find:", currentTime)

		from_expr := dbx.NewExp(`"from" < {:now}`, dbx.Params{"now": currentTime})
		to_expr := dbx.NewExp(`"to" > {:now}`, dbx.Params{"now": currentTime})

		expr := dbx.And(
			from_expr,
			to_expr,
		)

		records, err := app.Dao().FindRecordsByExpr(collection, expr)

		if err != nil {
			res := MessageResponse{
				Message: err.Error(),
			}
			return c.JSON(500, res)

		}
		if len(records) < 1 {
			res := MessageResponse{
				Message: "Not found!",
			}
			return c.JSON(404, res)
		}

		return c.JSON(200, records[0])
	}
}

func ClientChallengeController(app *pocketbase.PocketBase) echo.HandlerFunc {

	return func(c echo.Context) error {

		collection, err := app.Dao().FindCollectionByNameOrId("challenges")

		if err != nil {
			res := MessageResponse{
				Message: "Server unhealthy",
			}
			return c.JSON(500, res)
		}

		p := c.PathParam("*")
		log.Default().Println("Find:", p)

		expr := dbx.HashExp{"slug": p}
		records, err := app.Dao().FindRecordsByExpr(collection, expr)

		if err != nil || len(records) < 1 {
			res := MessageResponse{
				Message: "Not found!",
			}
			return c.JSON(404, res)
		}

		return c.JSON(200, records[0])
	}
}

func ServeCMS(fileSystem fs.FS, disablePathUnescaping bool) echo.HandlerFunc {
	return func(c echo.Context) error {
		p := c.PathParam("*")
		if !disablePathUnescaping { // when router is already unescaping we do not want to do is twice
			tmpPath, err := url.PathUnescape(p)
			if err != nil {
				return fmt.Errorf("failed to unescape path variable: %w", err)
			}
			p = tmpPath
		}

		// fs.FS.Open() already assumes that file names are relative to FS root path and considers name with prefix `/` as invalid
		name := filepath.ToSlash(filepath.Clean(strings.TrimPrefix(p, "/")))
		// log.Default().Println("GET", name)

		if _, err := fileSystem.Open(name); errors.Is(err, os.ErrNotExist) {
			log.Default().Println("ERR:", err)
			// log.Default().Println()

			return c.FileFS("index.html", fileSystem)
		}

		return c.FileFS(name, fileSystem)
	}

}

func main() {
	app := pocketbase.New()

	var publicDirFlag string

	// add "--publicDir" option flag
	app.RootCmd.PersistentFlags().StringVar(
		&publicDirFlag,
		"publicDir",
		defaultPublicDir(),
		"the directory to serve static files",
	)

	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		// serves static files from the provided public dir (if exists)
		e.Router.GET("/current-monthly-challenge", CurrentMonthlyChallegne(app))
		e.Router.GET("/find/*", ClientChallengeController(app))
		e.Router.GET("/*", ServeCMS(os.DirFS(publicDirFlag), false))

		return nil
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
