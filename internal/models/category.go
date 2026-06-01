package models

type Category struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `json:"name"`
}

type SubCategory struct {
	ID         uint     `gorm:"primaryKey"`
	Name       string   `json:"name"`
	CategoryID uint     `json:"categoryID"`
	Category   Category `gorm:"foreignKey:CategoryID" json:"category"`
}

type CategoryUpsert struct {
	Name string `json:"name"`
}

type SubCategoryUpsert struct {
	Name       string `json:"name"`
	CategoryID uint   `json:"categoryID"`
}
