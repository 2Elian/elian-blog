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

	db.Exec("UPDATE article SET category_id = 1 WHERE category_id = 3")
	fmt.Println("Reassigned articles from category 3 to 1")
	db.Exec("DELETE FROM category WHERE id = 3")
	fmt.Println("Deleted category 3")
	db.Exec("DELETE FROM article WHERE id IN (3, 4)")
	fmt.Println("Deleted test articles 3, 4")

	fmt.Println("Done!")
}
