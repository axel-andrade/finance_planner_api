package category

type CategoryRepository interface {
	Create(category *Category) error
	Update(category *Category) error
	Delete(id int) error
	FindByID(id int) (*Category, error)
	FindAll() ([]*Category, error)
}
