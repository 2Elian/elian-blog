package model

import "gorm.io/gorm"

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&User{},
		&Role{},
		&Menu{},
		&Article{},
		&Category{},
		&Tag{},
		&Comment{},
		&FriendLink{},
		&Message{},
		&Page{},
		&SiteConfig{},
		&OperationLog{},
	)
}
