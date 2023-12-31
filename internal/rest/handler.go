package rest

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"expenses/internal/model"
	"fmt"
	"time"
	"expenses/internal/db"
	"log"
	l "expenses/internal/log"
)

type HandlerInterface interface {
	GetCity(*gin.Context)
	GetCat(*gin.Context)
	GetSubcat(*gin.Context)
	GetSupplier(*gin.Context)
	GetExpensesNames(*gin.Context)
	GetExpenses(*gin.Context)
	GetStatistics(*gin.Context)
	GetStuff(*gin.Context)
	AddDate(*gin.Context)
	AddExpense(*gin.Context)
}

type Handler struct {
	db     *dblayer.DB
	logger *log.Logger
	keys   map[string]any
}

func NewHandler() (*Handler, error) {
	db, _ := dblayer.NewDB()
	logger, err := l.NewLog()
	if err != nil {
		log.Fatal("couldn't create logger for handler")
	}
	return &Handler{
		db:     db,
		logger: logger,
		keys:   make(map[string]any, 2),
	}, nil
}

func (h *Handler) GetCity(c *gin.Context) {
	if h.db == nil {
		return
	}
	city, err := h.db.GetCity()
	h.errToLog(c, http.StatusInternalServerError, err)
	c.JSON(http.StatusOK, city)
}

func (h *Handler) GetCat(c *gin.Context) {
	if h.db == nil {
		return
	}
	cat, err := h.db.GetCat()
	h.errToLog(c, http.StatusInternalServerError, err)
	c.JSON(http.StatusOK, cat)
}

func (h *Handler) GetSubcat(c *gin.Context) {
	if h.db == nil {
		return
	}
	subcat, err := h.db.GetSubcat()
	h.errToLog(c, http.StatusInternalServerError, err)
	c.JSON(http.StatusOK, subcat)
}

func (h *Handler) GetSupplier(c *gin.Context) {
	if h.db == nil {
		return
	}
	supplier, err := h.db.GetSupplier()
	h.errToLog(c, http.StatusInternalServerError, err)
	c.JSON(http.StatusOK, supplier)
}

func (h *Handler) GetExpensesNames(c *gin.Context) {
	if h.db == nil {
		return
	}
	expenses_name, err := h.db.GetExpensesNames()
	h.errToLog(c, http.StatusInternalServerError, err)
	c.JSON(http.StatusOK, expenses_name)
}

func (h *Handler) GetExpenses(c *gin.Context) {
	if h.db == nil {
		return
	}
	var bef, af string
	var b, a any
	var exist bool
	if b, exist = h.keys["before"]; !exist {
		bef = "2020-01-01"
	} else {
		bef = b.(string)
	}
	if a, exist = h.keys["after"]; !exist {
		af = "2030-01-01"
	} else {
		af = a.(string)
	}
	before, err := time.Parse("2006-01-02", bef)
	h.errToLog(c, http.StatusInternalServerError, err)
	after, err := time.Parse("2006-01-02", af)
	h.errToLog(c, http.StatusInternalServerError, err)
	
	expenses, err := h.db.GetExpenses(before, after)
	h.errToLog(c, http.StatusInternalServerError, err)
	
	c.JSON(http.StatusOK, expenses)
}

func (h *Handler) GetStatistics(c *gin.Context) {
	if h.db == nil {
		return
	}
	var bef, af string
	var b, a any
	var exist bool
	if b, exist = h.keys["before"]; !exist {
		bef = "2020-01-01"
	} else {
		bef = b.(string)
	}
	if a, exist = h.keys["after"]; !exist {
		af = "2030-01-01"
	} else {
		af = a.(string)
	}
	before, err := time.Parse("2006-01-02", bef)
	h.errToLog(c, http.StatusInternalServerError, err)
	after, err := time.Parse("2006-01-02", af)
	h.errToLog(c, http.StatusInternalServerError, err)
	statistics, err := h.db.GetStatistics(before, after)
	h.errToLog(c, http.StatusInternalServerError, err)
	
	c.JSON(http.StatusOK, statistics)
}

func (h *Handler) GetStuff(c *gin.Context) {
	if h.db == nil {
		return
	}
	name := struct {
		Name string `json:"name"`
	}{}
	// name = c.GetHeader("Name")
	err := c.ShouldBind(&name)
	h.errToLog(c, http.StatusBadRequest, err)
	stuff, err := h.db.GetStuff(name.Name)
	fmt.Println("error:", err)
	
	h.errToLog(c, http.StatusInternalServerError, err)
	c.JSON(http.StatusOK, stuff)
}

func (h *Handler) AddDate(c *gin.Context) {
	if h.db == nil {
		return
	}
	type date struct {
		Before, After string
	}
	dateFromPost := date{}
	err := c.ShouldBind(&dateFromPost)
	h.errToLog(c, http.StatusBadRequest, err)
	
	if dateFromPost.Before == "" {
		h.keys["before"] = "2020-01-01"
	} else {
		h.keys["before"] = dateFromPost.Before
	}
	if dateFromPost.After == "" {
		h.keys["after"] = "2030-01-01"
	} else {
		h.keys["after"] = dateFromPost.After
	}
}

func (h *Handler) AddExpense(c *gin.Context) {
	if h.db == nil {
		h.logger.SetPrefix("[ERR] ")
		h.logger.Println(fmt.Errorf("h.db == nil in AddExpense"))
		c.JSON(http.StatusBadRequest, gin.H{"error": "h.db == nil in AddExpense"})
		return
	}
	expense := model.ExpenseAdd{}
	err := c.ShouldBindJSON(&expense)
	h.errToLog(c, http.StatusBadRequest, err)
	
	err = h.db.AddExpense(&expense)
	h.errToLog(c, http.StatusInternalServerError, err)
	
	c.JSON(http.StatusOK, gin.H{"response": "added"})
}

func (h *Handler) errToLog(c *gin.Context, status int, err error) {
	if err != nil {
		h.logger.SetPrefix("[ERR] ")
		h.logger.Println(err)
		c.JSON(status, gin.H{"error": err.Error()})
		return
	}
}
