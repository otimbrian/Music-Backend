package main

import (
	"fmt"
	"music/database"
	"music/handlers"

	"github.com/gin-gonic/gin"
)

// var db_Instance *gorm.DB
// var err error

// func connect() {

// 	handlers.Db_Instance, err = gorm.Open(postgres.Open(db_credentials), &gorm.Config{})
// 	if err != nil {
// 		panic("Failed to Connect to database")
// 	}

// 	fmt.Println("Connected to database.....")
// 	err = handlers.Db_Instance.AutoMigrate(&models.User{})

// 	if err != nil {
// 		fmt.Println("Migration failed ....." + err.Error())
// 	}
// 	fmt.Println("Migration Succesful")
// }

func main() {
	db_credentials := "host=localhost user=postgres password=1234 dbname=music_app_dev_db port=5432 sslmode=disable TimeZone=Africa/Nairobi"
	database.Connect(db_credentials)
	database.Migrate()

	router := gin.Default()
	router.Use(handlers.CORS())

	user := router.Group("/user")

	{
		user.POST("/login", handlers.LoginHandler)
		user.POST("/create", handlers.CreateUser)
		user.GET("/users", handlers.GetAllUsers)
		user.GET("/user/1", handlers.GetUserWithId)
		user.PATCH("/updateuser", handlers.UpdateUser)
	}

	music := router.Group("/music")
	{
		create := music.Group("/create")
		{
			create.POST("/music", handlers.UploadMusic)
			create.POST("/artist", handlers.CreateArtist)
			create.POST("/album", handlers.CreateMusicAlbum)
		}
		get := music.Group("/get")
		{
			get.GET("/artists", handlers.GetAllArtists)
			get.GET("/music", handlers.GetAllMusic)
		}

		music.PATCH("/like", handlers.LikeOrUnlikeMusic)
		music.DELETE("/delete/:id", handlers.DeleteMusic)
		// music.GET("/", handlers.GetAllMusic)
		// music.GET("/artists", handlers.GetAllArtists)
		// music.POST("/")
	}

	router.GET("/hompage", handlers.HandleGet)

	err := router.Run()
	if err != nil {
		fmt.Println("Error ocuured during run. " + err.Error())
	}
}
