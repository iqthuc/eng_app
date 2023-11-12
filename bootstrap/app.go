package bootstrap

import (
	"database/sql"
)

type Application struct {
	MySql *sql.DB
}

func App() Application {
	app := &Application{}
	return *app
}

