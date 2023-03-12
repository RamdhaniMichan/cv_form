package book

import (
	"template/entity"
	pkg "template/pkg/pagination"
)

type InterfaceBookService interface {
	Create(book *entity.Book) error
	FindAll() (*pkg.Pagination, error)
	Update(id int, book *entity.Book) error
	Detail(id int) (*entity.Book, error)
	Delete(id int) error
	Validate(book *entity.Book) error
}
