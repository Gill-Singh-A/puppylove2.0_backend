package main

import (
	"os"

	// "github.com/gin-contrib/cors"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/pclubiitk/puppylove2.0_backend/db"
	"github.com/pclubiitk/puppylove2.0_backend/router"
	"github.com/pclubiitk/puppylove2.0_backend/utils"
)

func main() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		panic("Error loading .env file")
	}
	var CfgAdminPass = os.Getenv("CFG_ADMIN_PASS")
	Db := db.InitDB()

	utils.Randinit()
	store := cookie.NewStore([]byte(CfgAdminPass))
	r := gin.Default()
	r.Use(cors.New(cors.Config{AllowCredentials: true, AllowOrigins: []string{"http://localhost:3000"}, AllowHeaders: []string{"content-type"}}))
	// r.Use(cors.New(cors.Config{AllowCredentials: true, AllowOriginFunc: func(origin string) bool {return true}, AllowHeaders: []string{"content-type"}}))
	r.Use(sessions.Sessions("adminsession", store))
	router.PuppyRoute(r, *Db)

	r.Run(":8080")

	// if err := r.Run(config.CfgAddr); err != nil {
	// 	fmt.Println("[Error] " + err.Error())
	// }
}
