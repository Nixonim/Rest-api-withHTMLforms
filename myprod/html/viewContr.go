package html

import (
	"html/template"

	"github.com/gin-gonic/gin"
)

var Templates = make(map[string]*template.Template, 3)

func LoadTemplates() {
	templateNames := [11]string{"mainPage", "list", "runonePost", "idrun", "crRunner", "thank", "upRunner", "thankUp", "deleteGet", "thankdelete", "crResult"}
	for _, name := range templateNames {
		t, err := template.ParseFiles("html/main.html", "html/"+name+".html")
		if err == nil {
			Templates[name] = t
		} else {
			panic(err)
		}
	}
}

func GetMain(ctx *gin.Context) {
	Templates["mainPage"].Execute(ctx.Writer, nil)
}
