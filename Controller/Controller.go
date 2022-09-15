package Controller

import(
	"main/services"
	"main/Data"
	"github.com/gin-gonic/gin"
	"net/http"

)

type Controller struct{
	Platform *services.Services
}

type Response struct{
	Status int			`json:"status"`
	Message interface{}	`json:"message"`
}

func NewRouter(platform *services.Services) *gin.Engine{
	router := gin.Default()
	control := Controller{Platform : platform}
	book := router.Group("/book")
	book.GET("/",control.GetAllBooks)
	book.GET("/:id",control.GetBook)
	book.POST("/",)
	book.PUT("/:id", )
	book.DELETE("/:id", )
	router.GET("/" , control.HelloGuys)

	return router
}

func (c *Controller) HelloGuys(ctx *gin.Context){
	ctx.JSON(http.StatusOK ,"Hi guys ;)")
}

func (c *Controller) GetAllBooks(ctx *gin.Context){
	var response Response

	books, err := c.Platform.GetAllBooks()

	if err != nil {
		response.Message = err.Error()
		response.Status = http.StatusInternalServerError
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	response.Message = books
	response.Status = http.StatusOK
	ctx.JSON(http.StatusOK, response)
}

func (c *Controller) GetBook(ctx *gin.Context)  {
	var response Response
	id := ctx.Param("id")
	book, err := c.Platform.GetBook(id)

	if err != nil {
		response.Message = err.Error()
		response.Status = http.StatusInternalServerError
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	response.Message = book
	response.Status = http.StatusOK
	ctx.JSON(http.StatusOK, response)

}

func (c *Controller) AddBook(ctx *gin.Context)  {
	var response Response
	var request Data.Book
	ctx.Header("Content-Type", "application/json")
	if err := ctx.BindJSON(&request) ; err != nil{
		response.Message = err.Error()
		response.Status = http.StatusBadRequest
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	book, status, err := c.Platform.AddBook(request)

	if err != nil{
		response.Message = err.Error()
		response.Status = status
		ctx.JSON(status, response)
		return
	}

	response.Message = book
	response.Status = status
	ctx.JSON(status, response)
}



