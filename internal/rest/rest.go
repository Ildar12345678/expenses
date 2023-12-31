package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"os"
	"fmt"
)

func RunAPI(address string) error {
	h, err := NewHandler()
	if err != nil {
		return err
	}
	return RunAPIWithHandler(address, h)
}
func RunAPIWithHandler(address string, handler HandlerInterface) error {
	r := gin.Default()
	file, err := os.OpenFile("server_log", os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return fmt.Errorf("error while creating log file: %s", err.Error())
	}
	r.Use(cors.Default(), gin.LoggerWithConfig(gin.LoggerConfig{
		Formatter: nil,
		Output:    file,
		SkipPaths: nil,
	}))
	r.GET("/city", handler.GetCity)
	r.GET("/cat", handler.GetCat)
	r.GET("/subcat", handler.GetSubcat)
	r.GET("/supplier", handler.GetSupplier)
	r.GET("/expname", handler.GetExpensesNames)
	r.GET("/expenses", handler.GetExpenses)
	r.GET("/stats", handler.GetStatistics)
	r.POST("/stuff", handler.GetStuff)
	r.POST("/date", handler.AddDate)
	r.POST("/addexpense", handler.AddExpense)
	return r.Run(address)
}
