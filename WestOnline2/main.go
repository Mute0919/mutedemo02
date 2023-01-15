package main

import (
	"WestOnline2/conf"
	"WestOnline2/routes"
)

func main() {
	conf.Init()
	r := routes.NewRouter()
	_ = r.Run(":9000")
}
