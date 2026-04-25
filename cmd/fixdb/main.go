package main

import (
	"fmt"
	"log"

	"github.com/zeromicro/go-zero/core/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"elian-blog/internal/config"
)

func main() {
	var c config.Config
	conf.MustLoad("configs/config.yaml", &c)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		c.Database.Username, c.Database.Password,
		c.Database.Host, c.Database.Port,
		c.Database.DBName, c.Database.Charset,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// Drop FK constraints on article table
	constraints := []string{"fk_article_category", "fk_article_author"}
	for _, fk := range constraints {
		sql := fmt.Sprintf("ALTER TABLE article DROP FOREIGN KEY %s", fk)
		if err := db.Exec(sql).Error; err != nil {
			fmt.Printf("Drop %s: %v\n", fk, err)
		} else {
			fmt.Printf("Dropped FK: %s\n", fk)
		}
	}

	// Set category_id to NULL where it's 0
	db.Exec("UPDATE article SET category_id = NULL WHERE category_id = 0")
	fmt.Println("Set NULL category_id")

	fmt.Println("Done!")
}
