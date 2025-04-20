package repository

type CategoriesRepo interface{
	GetCategories() ([]string, error)
}
