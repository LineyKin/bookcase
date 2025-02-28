package handlers

import (
	"bookcase/internal/kafka"
	"bookcase/internal/service"
	"bookcase/models/author"
	"bookcase/models/book"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	service service.ServiceInterface
	kp      *kafka.Producer
}

func NewController(service service.ServiceInterface, kp *kafka.Producer) *Controller {
	return &Controller{
		service: service,
		kp:      kp,
	}
}

func (ctrl *Controller) AddBook(c *gin.Context) {
	if c.Request.Method != http.MethodPost {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Метод не поддерживается"})
		return
	}

	var bookData book.BookAdd
	if err := c.BindJSON(&bookData); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ошибка десериализации JSON"})
		return
	}

	userId, err := ctrl.getUserId(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Errorf("невозможно получить id пользователя: %s", err)})
		return
	}

	b, err := ctrl.service.AddBook(bookData, userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"new_book": b})
	c.Set(USER_LOG_KEY, b.NewLog())
	c.Next()
}

func (ctrl *Controller) GetPublishingHouseList(c *gin.Context) {
	if c.Request.Method != http.MethodGet {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Метод не поддерживается"})
		return
	}

	ph, err := ctrl.service.GetPublishingHouseList()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"ph_list": ph})
}

func (ctrl *Controller) GetAuthorList(c *gin.Context) {
	if c.Request.Method != http.MethodGet {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Метод не поддерживается"})
		return
	}

	authorList, err := ctrl.service.GetAuthorList()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"author_list": authorList})
}

func (ctrl *Controller) AddAuthor(c *gin.Context) {
	if c.Request.Method != http.MethodPost {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Метод не поддерживается"})
		return
	}

	var author author.Author
	if err := c.BindJSON(&author); err != nil {
		fmt.Println("handler AddAuthor", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ошибка десериализации JSON"})
		return
	}

	id, err := ctrl.service.AddAuthor(author)
	author.Id = id

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"author_id": id})

	c.Set(USER_LOG_KEY, author.NewLog())
	c.Next()
}

func (ctrl *Controller) GetBookCount(c *gin.Context) {
	if c.Request.Method != http.MethodGet {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Метод не поддерживается"})
		return
	}

	userId, err := ctrl.getUserId(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Errorf("невозможно получить id пользователя: %s", err)})
		return
	}

	count, err := ctrl.service.GetBookCount(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"count": count})
}

func (ctrl *Controller) GetBookCountTotal(c *gin.Context) {
	if c.Request.Method != http.MethodGet {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Метод не поддерживается"})
		return
	}

	count, err := ctrl.service.GetBookCountTotal()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"count_total": count})
}

func (ctrl *Controller) GetBookList(c *gin.Context) {
	if c.Request.Method != http.MethodGet {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Метод не поддерживается"})
		return
	}

	limitString := c.Query("limit")
	limit, err := strconv.Atoi(limitString)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	offsetString := c.Query("offset")
	offset, err := strconv.Atoi(offsetString)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sortedBy := c.Query("sortedBy")
	sortType := c.Query("sortType")

	userId, err := ctrl.getUserId(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Errorf("невозможно получить id пользователя: %s", err)})
		return
	}

	bookList, err := ctrl.service.GetBookList(userId, limit, offset, sortedBy, sortType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"book_list": bookList})
}

func (ctrl *Controller) GetBookListTotal(c *gin.Context) {
	if c.Request.Method != http.MethodGet {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Метод не поддерживается"})
		return
	}

	limitString := c.Query("limit")
	limit, err := strconv.Atoi(limitString)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	offsetString := c.Query("offset")
	offset, err := strconv.Atoi(offsetString)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sortedBy := c.Query("sortedBy")
	sortType := c.Query("sortType")

	bookList, err := ctrl.service.GetBookListTotal(limit, offset, sortedBy, sortType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"book_list_total": bookList})
}
