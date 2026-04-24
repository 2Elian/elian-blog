package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/zeromicro/go-zero/core/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"elian-blog/internal/config"
	"elian-blog/internal/model"
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
		log.Fatal("DB connect failed:", err)
	}

	// Fix corrupted tags
	var tags []model.Tag
	db.Find(&tags)
	for _, t := range tags {
		if hasBad(t.Name) {
			db.Model(&model.Tag{}).Where("id = ?", t.ID).Update("name", "测试标签")
			fmt.Printf("Fixed tag %d: %q -> %q\n", t.ID, t.Name, "测试标签")
		}
	}

	// Fix corrupted roles
	var roles []model.Role
	db.Find(&roles)
	roleNames := map[string]struct {
		Name string
		Desc string
	}{
		"admin":  {"管理员", "系统管理员"},
		"editor": {"编辑者", "内容编辑"},
		"user":   {"普通用户", "默认用户角色"},
	}
	for _, r := range roles {
		if fix, ok := roleNames[r.Label]; ok {
			updates := map[string]interface{}{}
			if hasBad(r.Name) {
				updates["name"] = fix.Name
			}
			if hasBad(r.Description) {
				updates["description"] = fix.Desc
			}
			if len(updates) > 0 {
				db.Model(&model.Role{}).Where("id = ?", r.ID).Updates(updates)
				fmt.Printf("Fixed role %d (%s): %v\n", r.ID, r.Label, updates)
			}
		}
	}

	// Fix corrupted categories
	var cats []model.Category
	db.Find(&cats)
	for _, c := range cats {
		if hasBad(c.Name) {
			newName := "默认分类"
			if c.ID == 2 {
				newName = "前端技术"
			}
			db.Model(&model.Category{}).Where("id = ?", c.ID).Update("name", newName)
			fmt.Printf("Fixed category %d: %q -> %q\n", c.ID, c.Name, newName)
		}
		if hasBad(c.Description) {
			db.Model(&model.Category{}).Where("id = ?", c.ID).Update("description", "")
			fmt.Printf("Fixed category %d description\n", c.ID)
		}
	}

	// Fix corrupted albums
	var albums []model.Album
	db.Find(&albums)
	for _, a := range albums {
		updates := map[string]interface{}{}
		if hasBad(a.AlbumName) {
			updates["album_name"] = "测试相册"
		}
		if hasBad(a.AlbumDesc) {
			updates["album_desc"] = "测试描述"
		}
		if len(updates) > 0 {
			db.Model(&model.Album{}).Where("id = ?", a.ID).Updates(updates)
			fmt.Printf("Fixed album %d: %v\n", a.ID, updates)
		}
	}

	// Fix corrupted friend links
	var friends []model.FriendLink
	db.Find(&friends)
	for _, f := range friends {
		updates := map[string]interface{}{}
		if hasBad(f.Name) {
			updates["name"] = "友情链接"
		}
		if hasBad(f.Description) {
			updates["description"] = "友情链接"
		}
		if len(updates) > 0 {
			db.Model(&model.FriendLink{}).Where("id = ?", f.ID).Updates(updates)
			fmt.Printf("Fixed friend %d: %v\n", f.ID, updates)
		}
	}

	// Fix corrupted pages
	var pages []model.Page
	db.Find(&pages)
	for _, p := range pages {
		updates := map[string]interface{}{}
		if hasBad(p.Title) {
			updates["title"] = "关于"
		}
		if hasBad(p.Content) {
			updates["content"] = "关于页面内容"
		}
		if len(updates) > 0 {
			db.Model(&model.Page{}).Where("id = ?", p.ID).Updates(updates)
			fmt.Printf("Fixed page %d: %v\n", p.ID, updates)
		}
	}

	// Fix corrupted articles
	var articles []model.Article
	db.Find(&articles)
	for _, a := range articles {
		if hasBad(a.Title) {
			db.Model(&model.Article{}).Where("id = ?", a.ID).Update("title", "文章")
			fmt.Printf("Fixed article %d: %q -> %q\n", a.ID, a.Title, "文章")
		}
	}

	fmt.Println("\nAll data fixed.")
}

func hasBad(s string) bool {
	return strings.ContainsRune(s, '�')
}
