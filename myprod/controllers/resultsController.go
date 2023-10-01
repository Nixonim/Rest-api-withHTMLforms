package controllers

import (
	"myprod/html"
	"myprod/models"
	"myprod/services"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type ResultsController struct {
	resultsService *services.ResultsService
}

func NewResultsController(resultsService *services.
	ResultsService) *ResultsController {
	return &ResultsController{
		resultsService: resultsService,
	}
}
func (rh ResultsController) CreateResultPost(ctx *gin.Context) {
	preobyt := ctx.PostForm("race_result")
	preobyt = strings.ReplaceAll(preobyt, "\n", "")
	preob, _ := strconv.Atoi(ctx.PostForm("position"))
	preobQ, _ := strconv.Atoi(ctx.PostForm("year"))
	result := models.Result{
		RunnerID:   ctx.PostForm("runner_id"),
		RaceResult: preobyt,
		Position:   preob,
		Year:       preobQ,
		Location:   ctx.PostForm("location"),
	}
	response, responseErr := rh.resultsService.CreateResult(&result)
	if responseErr != nil {
		ctx.JSON(responseErr.Status, responseErr)
		return
	}
	// html.Templates["thank"].Execute(ctx.Writer, response)
	ctx.JSON(http.StatusOK, response)
}
func (rh ResultsController) CreateResultGet(ctx *gin.Context) {
	var o models.Result
	html.Templates["crResult"].Execute(ctx.Writer, o)
}

//	func (rh ResultsController) CreateResult(ctx *gin.Context) {
//		body, err := io.ReadAll(ctx.Request.Body)
//		if err != nil {
//			log.Println("Error while reading "+
//				"create result request body", err)
//			ctx.AbortWithError(http.StatusInternalServerError,
//				err)
//			return
//		}
//		var result models.Result
//		err = json.Unmarshal(body, &result)
//		if err != nil {
//			log.Println("Error while unmarshaling "+
//				"creates result request body", err)
//			ctx.AbortWithError(http.StatusInternalServerError,
//				err)
//			return
//		}
//		response, responseErr := rh.resultsService.CreateResult(&result)
//		if responseErr != nil {
//			ctx.JSON(responseErr.Status, responseErr)
//			return
//		}
//		ctx.JSON(http.StatusOK, response)
//	}
func (rh ResultsController) DeleteResult(ctx *gin.Context) {
	resultId := ctx.Param("id")
	responseErr := rh.resultsService.DeleteResult(resultId)
	if responseErr != nil {
		ctx.JSON(responseErr.Status, responseErr)
		return
	}
	ctx.Status(http.StatusNoContent)
}
