# go-crud V2


```go

package main

import (
	"github.com/gin-gonic/gin"
	"hello/crud"
)

type Book struct {
	crud.BaseModel[Book]
	Name   string `json:"name" binding:"required"`
	Author string `json:"author" binding:"required"`
}

type Product struct {
	crud.BaseModel[Product]
	Name    string `json:"name"`
	Factory string `json:"factory"`
}

type Stu struct {
	crud.BaseModel[Stu]
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	crud.ConnectDB(crud.SQLite, "test.db")
	crud.AutoMigrate(&Book{}, &Product{}, &Stu{})

	r := gin.Default()
	crud.Crud[Book](r, "book", &Book{})
	crud.Crud[Product](r, "product", &Product{})
	crud.Crud[Stu](r, "stu", &Stu{})

	r.Run(":8081") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}


```