package main

import (
	"database/sql"
	"log"

	"github.com/bootcamp-go/consignas-go-db.git/cmd/server/handler"
	"github.com/bootcamp-go/consignas-go-db.git/internal/product"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	dataSource := "root:@tcp(localhost:3306)/my_db"
	//Open inicia un pool de conexiones. Sólo abrir una vez
	StorageDB, err := sql.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}

	if err = StorageDB.Ping(); err != nil {
		panic(err)
	}
	log.Println("database Configured")

	//storage := store.NewJsonStore("./products.json")
	//repo := product.NewRepository(storage)

	mySqlRepo := product.NewMySQLRepository(StorageDB)
	service := product.NewService(mySqlRepo)
	productHandler := handler.NewProductHandler(service)

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) { c.String(200, "pong") })
	products := r.Group("/products")
	{
		products.GET("/:id", productHandler.GetByID())
		products.POST("", productHandler.Post())
		// products.DELETE(":id", productHandler.Delete())
		// products.PATCH(":id", productHandler.Patch())
		// products.PUT(":id", productHandler.Put())
	}

	r.Run(":8080")
}
