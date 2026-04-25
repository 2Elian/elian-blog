package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"elian-blog/internal/config"
	"elian-blog/internal/model"
	"elian-blog/internal/utils"
)

var configFile = "configs/config.yaml"

func main() {
	var c config.Config
	conf.MustLoad(configFile, &c)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		c.Database.Username, c.Database.Password,
		c.Database.Host, c.Database.Port,
		c.Database.DBName, c.Database.Charset,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("连接数据库失败:", err)
	}

	fixAdminPassword(db)
	seedAdminUser(db)
}

// fixAdminPassword 检测并修复明文密码
func fixAdminPassword(db *gorm.DB) {
	var users []model.User
	db.Where("email = ? OR username = ?", "lizimo@elian.net.cn", "admin").Find(&users)

	for _, u := range users {
		if !strings.HasPrefix(u.Password, "$2a$") && !strings.HasPrefix(u.Password, "$2b$") {
			hashed, err := utils.HashPassword(u.Password)
			if err != nil {
				fmt.Printf("密码加密失败 (user=%s): %v\n", u.Username, err)
				continue
			}
			db.Model(&model.User{}).Where("id = ?", u.ID).Update("password", hashed)
			fmt.Printf("已修复用户 %s (email=%s) 的明文密码 -> bcrypt哈希\n", u.Username, u.Email)
		} else {
			fmt.Printf("用户 %s 密码已是bcrypt格式，无需修复\n", u.Username)
		}
	}
}

// seedAdminUser 创建管理员账号（如不存在）
func seedAdminUser(db *gorm.DB) {
	var existingUser model.User
	err := db.Preload("Roles").Where("email = ?", "lizimo@elian.net.cn").First(&existingUser).Error
	if err == nil {
		fmt.Println("管理员账号已存在，ID:", existingUser.ID)

		var adminRole model.Role
		if err := db.Where("label = ?", "admin").First(&adminRole).Error; err == nil {
			var count int64
			db.Table("user_roles").Where("user_id = ? AND role_id = ?", existingUser.ID, adminRole.ID).Count(&count)
			if count == 0 {
				db.Exec("INSERT INTO user_roles (user_id, role_id) VALUES (?, ?)", existingUser.ID, adminRole.ID)
				fmt.Println("已为管理员分配admin角色")
			}
		}
		return
	}

	password, err := utils.HashPassword("lizimo@elian.net.cn")
	if err != nil {
		log.Fatal("密码加密失败:", err)
	}

	adminUser := &model.User{
		Username:  "admin",
		Password:  password,
		Nickname:  "Elian",
		Email:     "lizimo@elian.net.cn",
		Status:    1,
		LastLogin: time.Now(),
	}
	if err := db.Create(adminUser).Error; err != nil {
		log.Fatal("创建管理员失败:", err)
	}

	var adminRole model.Role
	if err := db.Where("label = ?", "admin").First(&adminRole).Error; err != nil {
		adminRole = model.Role{Name: "管理员", Label: "admin", Description: "系统管理员", Status: 1, Sort: 1}
		db.Create(&adminRole)
	}
	db.Exec("INSERT INTO user_roles (user_id, role_id) VALUES (?, ?)", adminUser.ID, adminRole.ID)

	fmt.Println("管理员账号创建成功！用户名: admin, 邮箱: lizimo@elian.net.cn")
}
