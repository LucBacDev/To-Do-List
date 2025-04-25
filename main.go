package main

import (
	"fmt"
	"log"
	"os"
	authHandler "source-base-go/api/handler/auth"
	todoHandler "source-base-go/api/handler/todo"
	tagHandler "source-base-go/api/handler/tag"
	todoTagHandler "source-base-go/api/handler/todotag"

	"source-base-go/config"
	"source-base-go/infrastructure/repository"
	"source-base-go/infrastructure/repository/util"

	"source-base-go/usecase/user"
	"source-base-go/usecase/todo"
	"source-base-go/usecase/tag"
	"source-base-go/usecase/todotag"


	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	config.SetConfigFile("config")
	os.Setenv("TZ", "Etc/GMT")
}

func main() {
	envConfig := getConfig()

	//Database Connect
	db, err := repository.ConnectDB(envConfig.Sql)
	if err != nil {
		log.Println(err)
		return
	}
	app := gin.New()
	crs := cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowCredentials: true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Set-Cookie", "Authorization"},
	})

	app.Use(crs)
		//Define Repository

	// userRepo := repository.NewUserRepostory(db)
	//Handler
	userRepo := repository.NewUserRepository(db) 
	todoRepo := repository.NewTodoRepository(db)
	tagRepo := repository.NewTagRepository(db)
	todoTagRepo := repository.NewTodoTagRepository(db)


	var verifier util.Verifier = util.NewAccessTokenRepository(db)

	userService := user.NewService(userRepo)
	todoService := todo.NewService(todoRepo, verifier)
	tagService := tag.NewService(tagRepo, verifier)
	todoTagService := todotag.NewService(todoTagRepo, verifier)




	authHandler.MakeHandlers(app, userService)
	todoHandler.MakeHandlers(app, todoService,verifier)
	tagHandler.MakeHandlers(app, tagService,verifier)
	todoTagHandler.MakeHandlers(app, todoTagService,verifier)






	addr := fmt.Sprintf("%s:%d", envConfig.Host, envConfig.Port)
	log.Printf(" Server đang chạy tại http://%s", addr)
	for _, route := range app.Routes() {
		log.Printf("API: %s %s", route.Method, route.Path)
	}
	if err := app.Run(addr); err != nil {
		log.Fatalf(" Lỗi khi chạy server: %v", err)
	}
	
}

func getConfig() config.EnvConfig {
	return config.EnvConfig{
		Host: config.GetString("host.address"),
		Port: config.GetInt("host.port"),
		Sql: config.SqlConfig{
			Timeout:  config.GetInt("database.sql.timeout"),
			DBName:   config.GetString("database.sql.dbname"),
			Username: config.GetString("database.sql.user"),
			Password: config.GetString("database.sql.password"),
			Host:     config.GetString("database.sql.host"),
			Port:     config.GetString("database.sql.port"),
		},
	}
}
