package book

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"template/datasource"
	"template/entity"
	pkg "template/pkg/pagination"
)

type bookRepository struct{}

func NewBookRepository() InterfaceBookRepository {
	return &bookRepository{}
}

func (b *bookRepository) Create(book *entity.Book) error {
	db := datasource.OpenDB()

	err := db.Create(&book).Error

	if err != nil {
		return err
	}

	defer db.Close()

	return nil
}

func (b *bookRepository) FindAll() (*pkg.Pagination, error) {
	db := datasource.OpenDB()
	var books []entity.Book
	var paginate pkg.Pagination

	err := db.Scopes(pkg.Paginate(books, &paginate, db)).Find(&books).Error
	if err != nil {
		return nil, err
	}
	paginate.Rows = &books

	defer db.Close()

	return &paginate, nil
}

func (b *bookRepository) Update(id int, book *entity.Book) error {
	db := datasource.OpenDB()

	fmt.Println("isinya", book)

	err := db.Model(&entity.Book{}).Where("id = ?", id).Updates(&book).Error

	if err != nil {
		return err
	}

	defer db.Close()

	return nil
}

func (b *bookRepository) Detail(id int) (*entity.Book, error) {
	db := datasource.OpenDB()
	var book entity.Book

	err := db.Where("id = ?", id).Take(&book).Error
	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			return nil, err
		default:
			return nil, err
		}
	}

	defer db.Close()

	return &book, nil
}

func (b *bookRepository) Delete(id int) error {
	db := datasource.OpenDB()
	var book entity.Book

	err := db.Where("id = ?", id).Delete(&book).Error
	if err != nil {
		return err
	}

	defer db.Close()

	return nil
}
