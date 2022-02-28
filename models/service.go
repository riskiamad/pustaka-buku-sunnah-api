package book

type Service interface {
	FindAll() ([]BookModel, error)
	FindByID(ID int) (BookModel, error)
	Create(bookRequest BookRequest) (BookModel, error)
	Update(ID int, book BookRequest) (BookModel, error)
	Delete(ID int) (BookModel, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]BookModel, error) {
	books, err := s.repository.FindAll()
	return books, err
}

func (s *service) FindByID(ID int) (BookModel, error) {
	book, err := s.repository.FindByID(ID)
	return book, err
}

func (s *service) Create(bookRequest BookRequest) (BookModel, error) {
	book := BookModel{
		Title:       bookRequest.Title,
		Price:       int(bookRequest.Price.(float64)),
		Description: bookRequest.Description,
		Rating:      int(bookRequest.Rating.(float64)),
		Discount:    int(bookRequest.Discount.(float64)),
	}

	newBook, err := s.repository.Create(book)
	return newBook, err
}

func (s *service) Update(ID int, bookRequest BookRequest) (BookModel, error) {
	book, err := s.repository.FindByID(ID)

	book.Title = bookRequest.Title
	book.Price = int(bookRequest.Price.(float64))
	book.Description = bookRequest.Description
	book.Rating = int(bookRequest.Rating.(float64))
	book.Discount = int(bookRequest.Discount.(float64))

	newBook, err := s.repository.Update(book)
	return newBook, err
}

func (s *service) Delete(ID int) (BookModel, error) {
	book, err := s.repository.FindByID(ID)

	newBook, err := s.repository.Delete(book)
	return newBook, err
}
