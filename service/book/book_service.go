package book

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"log"
	"template/entity"
	pkg "template/pkg/pagination"
	repository "template/repository/book"
)

type bookService struct {
	bookRepository repository.InterfaceBookRepository
}

func NewBookService(repository repository.InterfaceBookRepository) InterfaceBookService {
	return &bookService{
		bookRepository: repository,
	}
}

func (b *bookService) Create(book *entity.Book) error {

	validateErr := b.Validate(book)

	if validateErr != nil {
		return validateErr
	}

	err := b.bookRepository.Create(book)

	if err != nil {
		log.Printf("terjadi error ketika akan menyimpan data buku : %w", err)
		return err
	}

	return nil
}

func (b *bookService) FindAll() (*pkg.Pagination, error) {
	books, err := b.bookRepository.FindAll()
	if err != nil {
		log.Printf("terjadi error ketika akan mengambil data buku :%s", err)
		return nil, err
	}

	return books, nil
}

func (b *bookService) Update(id int, book *entity.Book) error {
	err := b.bookRepository.Update(id, book)
	if err != nil {
		return err
	}

	return nil
}

func (b *bookService) Detail(id int) (*entity.Book, error) {
	book, err := b.bookRepository.Detail(id)
	if err != nil {
		log.Printf("terjadi error ketika akan mengambil data buku :%s", err)
		return nil, err
	}

	return book, nil

}

func (b *bookService) Delete(id int) error {
	err := b.bookRepository.Delete(id)

	if err != nil {
		return err
	}

	return nil
}

func (b *bookService) Validate(book *entity.Book) error {
	validateErr := validation.ValidateStruct(book,
		validation.Field(&book.Title, validation.Required),
		validation.Field(&book.Author, validation.Required),
		validation.Field(&book.Genre, validation.Required),
		validation.Field(&book.Page, validation.Required),
	)

	if validateErr != nil {
		return validateErr
	}

	return nil
}
