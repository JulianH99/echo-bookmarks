package main

import (

	"github.com/JulianH99/gomarks/api"
)

func main() {

	api.NewApp(api.AppConfig{Port: 3000})

}
