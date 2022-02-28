package book

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]BookModel, error)
	FindByID(ID int) (BookModel, error)
	Create(book BookModel) (BookModel, error)
	Update(book BookModel) (BookModel, error)
	Delete(book BookModel) (BookModel, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]BookModel, error) {
	var books []BookModel

	err := r.db.Find(&books).Error

	return books, err
}

func (r *repository) FindByID(ID int) (BookModel, error) {
	var book BookModel

	err := r.db.Find(&book, ID).Error

	return book, err
}

func (r *repository) Create(book BookModel) (BookModel, error) {
	err := r.db.Create(&book).Error

	return book, err
}

func (r *repository) Update(book BookModel) (BookModel, error) {
	err := r.db.Save(&book).Error

	return book, err
}

func (r *repository) Delete(book BookModel) (BookModel, error) {
	err := r.db.Delete(&book).Error

	return book, err
}
