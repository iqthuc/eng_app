package main

import (
	"eng_app_module/api/route"
	"eng_app_module/database"
)

func main() {
	/// đoạn code đóng server khi bấm ctrl + C
	// c := make(chan os.Signal, 1)
	// signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// go func() {
	// 	<-c
	// 	os.Exit(0)
	// }()
	///
	database.GetDB()
	defer database.CloseMysqlConnection()
	route.Setup()
	select {}
}
