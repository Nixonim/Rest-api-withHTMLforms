package server

import (
	"database/sql"
	"log"
	"myprod/controllers"
	"myprod/html"
	"myprod/repositories"
	"myprod/services"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type HttpServer struct {
	config            *viper.Viper
	router            *gin.Engine
	runnersController *controllers.RunnersController
	resultsController *controllers.ResultsController
}

func InitHttpServer(config *viper.Viper,
	dbHandler *sql.DB) HttpServer {
	runnersRepository := repositories.NewRunnersRepository(
		dbHandler)
	resultRepository := repositories.NewResultsRepository(
		dbHandler)
	runnersService := services.NewRunnersService(
		runnersRepository, resultRepository)
	resultsService := services.NewResultsService(
		resultRepository, runnersRepository)
	runnersController := controllers.NewRunnersController(
		runnersService)
	resultsController := controllers.NewResultsController(
		resultsService)
	router := gin.Default()
	router.GET("/rundelete", controllers.DeleteGet)
	router.POST("/rundelete", controllers.DeletePost)
	router.GET("/runone", controllers.GetRunOneGet)
	router.POST("/runone", controllers.GetRunOnePost)
	router.GET("/", html.GetMain)
	router.GET("/createRunner", runnersController.CreateRunnerGet)
	router.POST("/createRunner", runnersController.CreateRunnerPost)
	router.GET("/runnerUp", runnersController.UpdateRunnerGet)
	router.POST("/runnerUp", runnersController.UpdateRunnerPut)
	router.GET("/rundelete/:id", runnersController.DeleteRunner)
	router.GET("/runner/:id", runnersController.GetRunner)
	router.GET("/runner", runnersController.GetRunnersBatch)
	router.GET("/result", resultsController.CreateResultGet)
	router.POST("/result", resultsController.CreateResultPost)
	router.DELETE("/result/:id",
		resultsController.DeleteResult)
	return HttpServer{
		config:            config,
		router:            router,
		runnersController: runnersController,
		resultsController: resultsController,
	}
}

func (hs HttpServer) Start() {
	err := hs.router.Run(hs.config.GetString(
		"http.server_address"))
	if err != nil {
		log.Fatalf("Error while starting HTTP server: %v",
			err)
	}
}
