package shared

type Book struct {
	Name        string  `json:"name" gorm:"not null;unique"`
	Description string  `json:"description"`
	Author      string  `json:"author" gorm:"not null"`
	Price       float64 `json:"price"`
	UserID      string  `json:"userID" gorm:"not null;unique"`
}

type Common interface {
	Save(book *Book) error
	GetAll(book *Book) ([]Book, error)
	UpdateBook(name string, updatedBook Book) error
}
