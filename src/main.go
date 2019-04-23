package main

import (
  	"net/http"
  	"github.com/gin-gonic/gin"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
)

func SetupRouter() *gin.Engine {
  router := gin.Default()

  v1 := router.Group("api/v1") 
  {
    v1.GET("/books", fetchAllBook)
    v1.GET("/books/:id", fetchSingleBook)
    v1.POST("/books", createBook)
    v1.PUT("/books/:id", updateBook)
    v1.DELETE("/books/:id", deleteBook)  
  }

  return router
}

var db *gorm.DB

func init() {
    //open a db connection
    var err error
    db, err = gorm.Open("mysql", "root:@/demo?charset=utf8&parseTime=True&loc=Local")

    if err != nil {
        panic("failed to connect database")
    }

    //Migrate the schema
    db.AutoMigrate(&bookModel{})
}

func main() {
    router := SetupRouter()
    router.Run(":8080")
}

type (
	bookModel struct {
		gorm.Model
		Title       string  `json:"title"` 
		Description string  `json:"description"`  
	}

	transformedBook struct {
		ID          uint    `json:"id"`
		Title       string  `json:"title"`
		Description string  `json:"description"`
	}
)

func createBook(c *gin.Context) {

    book := bookModel{Title: c.PostForm("title"), Description: c.PostForm("description")}
    c.Bind(&book) 

	db.Create(&book)
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Book created successfully!", "resourceId": book.ID})
}

func fetchAllBook(c *gin.Context) {
	var books []bookModel
	var _books []transformedBook

	db.Find(&books)

	if len(books) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No book found!"})
		return
	}

	for _, item := range books {
        _books = append(_books, transformedBook{ID: item.ID, Title: item.Title, Description: item.Description})
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": _books})
}

func fetchSingleBook(c *gin.Context) {
	var book bookModel
	bookID := c.Param("id")

	db.First(&book, bookID)

	if book.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No book found!"})
		return
	}

	_book := transformedBook{ID: book.ID, Title: book.Title, Description: book.Description}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": _book})
}

func updateBook(c *gin.Context) {
	var book bookModel
	bookID := c.Param("id")

	db.First(&book, bookID)

	if book.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No book found!"})
		return
	}

    c.BindJSON(&book)
    db.Save(&book)

    c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Book updated successfully!"})
}

func deleteBook(c *gin.Context) {
	var book bookModel
    bookID := c.Param("id")

	db.First(&book, bookID)

	if book.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No book found!"})
		return
	}

	db.Delete(&book)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Book deleted successfully!"})
}
