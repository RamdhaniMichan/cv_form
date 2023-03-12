package routes

import (
	controller "template/controller/book"
	router "template/http"
	repository "template/repository/book"
	service "template/service/book"
)

var (
	bookRepository repository.InterfaceBookRepository = repository.NewBookRepository()
	bookService    service.InterfaceBookService       = service.NewBookService(bookRepository)
	bookController controller.InterfaceBookController = controller.NewBookController(bookService)
)

type BookRoute struct{}

func (r *BookRoute) Routing(httpRouter router.Router) {
	httpRouter.GET("/books", bookController.GetBooks)
	httpRouter.GET("/book/{id}", bookController.GetDetailBookWithParams)
	httpRouter.GET("/book/", bookController.GetDetailBookWithQueryParams)
	httpRouter.POST("/book", bookController.CreateBook)
	httpRouter.PUT("/book/{id}", bookController.UpdateBook)
	httpRouter.DELETE("/book/{id}", bookController.DeleteBook)
}
