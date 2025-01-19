package handlers

import (
	"bookcase/internal/kafka"
	"bookcase/internal/service"
	"bookcase/models/author"
	"bookcase/models/book"
	"encoding/json"
	"fmt"
	"log"
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

	b, err := ctrl.service.AddBook(bookData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"new_book": b})
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

func (ctrl *Controller) GetBookCount(c *gin.Context) {
	if c.Request.Method != http.MethodGet {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Метод не поддерживается"})
		return
	}

	count, err := ctrl.service.GetBookCount()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"count": count})
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
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ошибка десериализации JSON"})
		return
	}

	id, err := ctrl.service.AddAuthor(author)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"author_id": id})

	// если продюсер кафки не активен, завершаем хэндлер
	if ctrl.kp == nil {
		log.Println("лог добавления автора не передан в кафку из-за её неактивности")
		return
	}

	// data for kafka
	author.Id = id
	authorInBytes, err := json.Marshal(author)
	if err != nil {
		log.Println("can't prepare data for kafka: ", err)
	}

	// Send the bytes to kafka
	err = ctrl.kp.PushLogToQueue("bookcase_log", authorInBytes)
	if err != nil {
		log.Println("can't send data to kafka: ", err)
	}
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

	bookList, err := ctrl.service.GetBookList(limit, offset, sortedBy, sortType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	isAuto := c.Query("isAuto")

	fmt.Println("isAuto:", isAuto)

	c.JSON(http.StatusOK, gin.H{"book_list": bookList})
}
