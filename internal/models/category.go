package models

type Category struct {
	ID   uint   `gorm:"unique"`
	Name string `json:"name"`
}

type SubCategory struct {
	ID         uint   `gorm:"unique"`
	Name       string `json:"name"`
	CategoryID uint   `json:"categoryID"`
}

type CategoryUpsert struct {
	Name string `json:"name"`
}

type SubCategoryUpsert struct {
	Name       string `json:"name"`
	CategoryID uint   `json:"categoryID"`
}
