package crud

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type Controller[T any] interface {
	ReplaceFetchAPI() gin.HandlerFunc
	ReplaceFetchOneAPI() gin.HandlerFunc
	ReplaceCreateAPI() gin.HandlerFunc
	ReplaceUpdateAPI() gin.HandlerFunc
	ReplaceDeleteAPI() gin.HandlerFunc

	//Service[T]
}

var _ Controller[any] = (*BaseController[any])(nil)

type BaseController[T any] struct {
	BaseService[T]
}

func (bm *BaseController[T]) ReplaceFetchAPI() gin.HandlerFunc {
	return func(c *gin.Context) {
		books := bm.Fetch()
		c.JSON(200, gin.H{
			"data": books,
		})
	}
}

func (bm *BaseController[T]) ReplaceFetchOneAPI() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := bm.handleFetchById(c)
		if t == nil {
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"data": t,
		})
	}
}

func (bm *BaseController[T]) ReplaceCreateAPI() gin.HandlerFunc {
	return func(c *gin.Context) {
		var t T
		err := c.ShouldBind(&t)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		err = bm.Create(&t)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, gin.H{"data": t})
	}
}

func (bm *BaseController[T]) ReplaceUpdateAPI() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := bm.handleFetchById(c)
		if t == nil {
			return
		}

		err := c.Bind(t)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		err = bm.Update(t)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}

		c.JSON(200, gin.H{"data": t})
	}
}

func (bm *BaseController[T]) ReplaceDeleteAPI() gin.HandlerFunc {
	return func(c *gin.Context) {

		t := bm.handleFetchById(c)
		if t == nil {
			return
		}

		id := c.Param(getIdParam[T]())
		err := bm.Delete(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": t})
	}
}

func (bm *BaseController[T]) handleFetchById(c *gin.Context) *T {
	id := c.Param(getIdParam[T]())
	t, err := bm.FetchOne(id)
	if err != nil {
		code := http.StatusInternalServerError
		msg := http.StatusText(http.StatusInternalServerError)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			code = http.StatusBadRequest
			msg = err.Error()
		}

		c.JSON(code, gin.H{"message": msg})
		return nil
	}
	return t
}
