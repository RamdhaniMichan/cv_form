package controller

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	"template/entity"
	"template/function"
	service "template/service/book"
)

type bookController struct {
	bookService service.InterfaceBookService
}

func NewBookController(service service.InterfaceBookService) InterfaceBookController {
	return &bookController{
		bookService: service,
	}
}

func (b *bookController) GetBooks(resp http.ResponseWriter, req *http.Request) {
	books, err := b.bookService.FindAll()

	if err != nil {
		_ = function.SendResponse(resp, 500, "something when wrong", err)
	}

	_ = function.SendResponse(resp, 200, "success", books)
	return
}

func (b *bookController) CreateBook(resp http.ResponseWriter, req *http.Request) {
	var book entity.Book

	book.Title = req.FormValue("title")
	book.Author = req.FormValue("author")
	book.Genre = req.FormValue("genre")
	book.Page = req.FormValue("page")

	err := b.bookService.Create(&book)
	if err != nil {
		_ = function.SendResponse(resp, 400, "something when wrong", err)
		return
	}

	_ = function.SendResponse(resp, 201, "success", book)
	return
}

func (b *bookController) GetDetailBookWithParams(resp http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)

	id, errConv := strconv.ParseInt(params["id"], 10, 64)

	if errConv != nil {
		log.Println("terjadi kesalahan ketika mengkonversi id ke integer", errConv)
	}

	book, err := b.bookService.Detail(int(id))

	if err != nil {
		_ = function.SendResponse(resp, 404, fmt.Sprintf("tidak ada data dengan id %d", id), err)
		return
	}

	_ = function.SendResponse(resp, 200, "success", book)
	return
}

func (b *bookController) GetDetailBookWithQueryParams(resp http.ResponseWriter, req *http.Request) {
	if req.URL.Query().Get("id") == "" {
		_ = function.SendResponse(resp, 422, fmt.Sprint("id tidak boleh kosong"), nil)
		return
	}

	id, errConv := strconv.ParseInt(req.URL.Query().Get("id"), 10, 64)

	if errConv != nil {
		log.Println("terjadi kesalahan ketika mengkonversi id ke integer", errConv)
	}

	book, err := b.bookService.Detail(int(id))

	if err != nil {
		_ = function.SendResponse(resp, 404, fmt.Sprintf("tidak ada data dengan id %d", id), err)
		return
	}

	_ = function.SendResponse(resp, 200, "success", book)
	return
}

func (b *bookController) DeleteBook(resp http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)

	id, errConv := strconv.ParseInt(params["id"], 10, 64)

	if errConv != nil {
		log.Println("terjadi kesalahan ketika mengkonversi id ke integer", errConv)
	}

	err := b.bookService.Delete(int(id))

	if err != nil {
		_ = function.SendResponse(resp, 500, "something when wrong", err)
		return
	}

	_ = function.SendResponse(resp, 200, "success", fmt.Sprintf("suskses menghapus data dengan id %d", id))
	return
}

func (b *bookController) UpdateBook(resp http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)

	id, errConv := strconv.ParseInt(params["id"], 10, 64)

	if errConv != nil {
		log.Println("terjadi kesalahan ketika mengkonversi id ke integer", errConv)
	}

	book, err := b.bookService.Detail(int(id))

	if err != nil {
		_ = function.SendResponse(resp, 404, fmt.Sprintf("tidak ada data dengan id %d", id), err)
		return
	}

	book.Title = req.FormValue("title")
	book.Author = req.FormValue("author")
	book.Genre = req.FormValue("genre")
	book.Page = req.FormValue("page")
	book.ID = uint(id)

	err = b.bookService.Update(int(id), book)
	if err != nil {
		_ = function.SendResponse(resp, 500, "something when wrong", err)
		return
	}

	_ = function.SendResponse(resp, 200, "success", fmt.Sprintf("sukses mengupdate data dengan id %d", id))
	return
}
