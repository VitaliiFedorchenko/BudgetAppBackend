package models

import (
	"fmt"
	"github.com/gosimple/slug"
	"gorm.io/gorm"
	"strconv"
	"strings"
)

type Category struct {
	gorm.Model
	Name string `json:"name" gorm:"not null;index" faker:"name"`
	Slug string `json:"slug" gorm:"type:varchar(255);not null;index"`
}

// GenerateUniqueSlug generates a unique slug based on the base slug
func GenerateUniqueSlug(db *gorm.DB, baseSlug string) (string, error) {
	slugFromName := generateSlug(baseSlug)
	fmt.Println(slugFromName)
	for i := 1; ; i++ {
		var exists bool
		if err := db.Model(&Category{}).Where("slug = ?", slugFromName).Select("count(*) > 0").Scan(&exists).Error; err != nil {
			return "", err
		}

		if !exists {
			break
		}

		slugFromName = generateSlug(baseSlug) + "-" + strconv.Itoa(i)
	}

	return slugFromName, nil
}

// generateSlug generates a slug based on the name
func generateSlug(name string) string {
	return slug.Make(strings.ToLower(name))
}

// BeforeCreate hook is called before the database create operation
func (c *Category) BeforeCreate(tx *gorm.DB) (err error) {
	c.Slug, err = GenerateUniqueSlug(tx, c.Name)
	return
}
