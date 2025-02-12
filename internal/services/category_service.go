package services

import (
	"BudgetApp/internal/configs"
	"BudgetApp/models"
	"gorm.io/gorm"
	"log"
)

type CategoryService struct {
	db *gorm.DB
}

func NewCategoryService() *CategoryService {
	db, err := configs.ConnectionToDataBase()
	if err != nil {
		log.Fatal(err)
	}

	return &CategoryService{db: db}
}

func (s *CategoryService) GetAllCategories() ([]map[string]string, error) {
	var categories []map[string]string

	rows, err := s.db.Model(&models.Category{}).Select("name, slug").Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var name, slug string
		if err := rows.Scan(&name, &slug); err != nil {
			return nil, err
		}

		categories = append(categories, map[string]string{
			"name": name,
			"slug": slug,
		})
	}

	return categories, nil
}

func (s *CategoryService) CreateCategory(name string) error {
	category := models.Category{Name: name}
	if err := s.db.Create(&category).Error; err != nil {
		return err
	}

	return nil
}

func (s *CategoryService) DeleteCategory(slug string) error {
	if err := s.db.Unscoped().Where("slug = ?", slug).Delete(&models.Category{}).Error; err != nil {
		return err
	}

	return nil
}

func (s *CategoryService) UpdateCategory(slug, newName string) error {
	var category models.Category
	if err := s.db.Where("slug = ?", slug).First(&category).Error; err != nil {
		return err
	}

	category.Name = newName
	category.Slug, _ = models.GenerateUniqueSlug(s.db, newName)

	if err := s.db.Save(&category).Error; err != nil {
		return err
	}

	return nil
}
