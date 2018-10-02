package main

import (
    "github.com/labstack/echo"
    "go-sql/app/routes"
    "go-sql/app/adapters"
    config "github.com/spf13/viper"
	"go-sql/app/repositories"
)

func main() {
    if err := initConfig(); err != nil {
        panic(err)
    }
    // Init DB Conn
    new(adapters.DbAdapter).Init()
    if err := initMigration(); err != nil {
    	panic(err)
	}

    // Init Echo
    e := echo.New()
    // Define Routes
    e = routes.SetRoute(e)
    e.Debug = true

    // Start the server
    e.Logger.Fatal(e.Start(":8000"))
}

func initConfig() error {
    config.SetConfigName("config")
    config.AddConfigPath("./config")
    return config.ReadInConfig()
}
func initMigration() error {
	return repositories.InitMigration()

}