package controllers

import (
	"fmt"
	"myprod/html"
	"myprod/models"
	"myprod/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RunnersController struct {
	runnersService *services.RunnersService
}

func NewRunnersController(runnersService *services.
	RunnersService) *RunnersController {
	return &RunnersController{
		runnersService: runnersService,
	}
}
func (rh RunnersController) CreateRunnerPost(ctx *gin.Context) {
	preob, _ := strconv.Atoi(ctx.PostForm("age"))
	runner := models.Runner{
		FirstName: ctx.PostForm("first_name"),
		LastName:  ctx.PostForm("last_name"),
		Age:       preob,
		Country:   ctx.PostForm("country"),
	}
	response, responseErr := rh.runnersService.
		CreateRunner(&runner)
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status,
			responseErr)
		return
	}
	html.Templates["thank"].Execute(ctx.Writer, response)
}
func (rh RunnersController) CreateRunnerGet(ctx *gin.Context) {
	var o models.Runner
	html.Templates["crRunner"].Execute(ctx.Writer, o)
}
func (rh RunnersController) UpdateRunnerGet(ctx *gin.Context) {
	var o models.Runner
	html.Templates["upRunner"].Execute(ctx.Writer, o)
}

// func (rh RunnersController) UpdateRunnerPut(ctx *gin.Context) {
// 	preob, _ := strconv.Atoi(ctx.PostForm("age"))
// 	runner := models.Runner{
// 		ID:        ctx.PostForm("id"),
// 		FirstName: ctx.PostForm("first_name"),
// 		LastName:  ctx.PostForm("last_name"),
// 		Age:       preob,
// 		Country:   ctx.PostForm("country"),
// 	}
// 	responseErr := rh.runnersService.UpdateRunner(&runner)
// 	if responseErr != nil {
// 		ctx.AbortWithStatusJSON(responseErr.Status,
// 			responseErr)
// 		return
// 	}
// 	html.Templates["thankUp"].Execute(ctx.Writer, runner)

// }

// func (rh RunnersController) UpdateRunnerPut(ctx *gin.Context) {
// 	var runner models.Runner
// 	if err := ctx.ShouldBind(&runner); err != nil {
// 		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	responseErr := rh.runnersService.UpdateRunner(&runner)
// 	if responseErr != nil {
// 		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
// 		return
// 	}

//		html.Templates["thankUp"].Execute(ctx.Writer, runner)
//	}
func (rh RunnersController) UpdateRunnerPut(ctx *gin.Context) {
	// // Проверяем, является ли запрос POST с параметром "_method" равным "PUT"
	// if ctx.Request.Method == http.MethodPost && ctx.PostForm("_method") == "PUT" {
	// 	// Продолжаем обработку как для PUT запроса

	// Получаем данные бегуна из формы
	preob, _ := strconv.Atoi(ctx.PostForm("age"))
	runner := models.Runner{
		ID:        ctx.PostForm("id"),
		FirstName: ctx.PostForm("first_name"),
		LastName:  ctx.PostForm("last_name"),
		Age:       preob,
		Country:   ctx.PostForm("country"),
	}

	// Обновляем бегуна в базе данных или в другом хранилище
	responseErr := rh.runnersService.UpdateRunner(&runner)
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}

	// Отправляем ответ клиенту
	html.Templates["thankUp"].Execute(ctx.Writer, runner)
}

func (rh RunnersController) DeleteRunner(ctx *gin.Context) {

	runnerId := ctx.Param("id")
	responseErr := rh.runnersService.DeleteRunner(runnerId)
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status,
			responseErr)
		return
	}
	html.Templates["thankdelete"].Execute(ctx.Writer, runnerId)
}
func (rh RunnersController) GetRunner(ctx *gin.Context) {
	runnerId := ctx.Param("id")
	response, responseErr := rh.runnersService.
		GetRunner(runnerId)
	if responseErr != nil {
		ctx.JSON(responseErr.Status, responseErr)
		return
	}
	html.Templates["idrun"].Execute(ctx.Writer, response)
}
func (rh RunnersController) GetRunnersBatch(ctx *gin.Context) {
	params := ctx.Request.URL.Query()
	country := params.Get("country")
	year := params.Get("year")
	response, responseErr := rh.runnersService.GetRunnersBatch(country, year)
	if responseErr != nil {
		ctx.JSON(responseErr.Status, responseErr)
		return
	}
	html.Templates["list"].Execute(ctx.Writer, response)

}

func GetRunOneGet(ctx *gin.Context) {
	o := models.Runner{
		ID: "",
	}
	html.Templates["runonePost"].Execute(ctx.Writer, o)
}

func GetRunOnePost(ctx *gin.Context) {
	id := ctx.PostForm("id")
	redirectURL := fmt.Sprintf("/runner/%s", id)
	ctx.Redirect(http.StatusFound, redirectURL)
}

func DeleteGet(ctx *gin.Context) {
	o := models.Runner{
		ID: "",
	}
	html.Templates["deleteGet"].Execute(ctx.Writer, o)
}
func DeletePost(ctx *gin.Context) {
	id := ctx.PostForm("id")
	redirectURL := fmt.Sprintf("/rundelete/%s", id)
	ctx.Redirect(http.StatusFound, redirectURL)
}
