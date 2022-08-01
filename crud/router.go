package crud

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"reflect"
)

func getIdParam[T any]() string {
	model := *new(T)
	modelName := reflect.TypeOf(model).Name()
	idParam := modelName + "ID"

	return idParam
}

func Crud[T any](r *gin.Engine, name string, t Modeler[T]) {
	idParam := getIdParam[T]()

	r.GET(fmt.Sprintf("/%s", name), t.ReplaceFetchAPI())
	r.GET(fmt.Sprintf("/%s/:%s", name, idParam), t.ReplaceFetchOneAPI())
	r.POST(fmt.Sprintf("/%s", name), t.ReplaceCreateAPI())
	r.PUT(fmt.Sprintf("/%s/:%s", name, idParam), t.ReplaceUpdateAPI())
	r.DELETE(fmt.Sprintf("/%s/:%s", name, idParam), t.ReplaceDeleteAPI())
}
