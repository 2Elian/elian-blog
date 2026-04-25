package model

import (
	"time"

	"gorm.io/gorm"
)

type Model struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

// User 用户表
type User struct {
	Model
	Username  string    `json:"username" gorm:"uniqueIndex;size:50;not null"`
	Password  string    `json:"-" gorm:"size:255;not null"`
	Nickname  string    `json:"nickname" gorm:"size:50"`
	Avatar    string    `json:"avatar" gorm:"size:500"`
	Email     string    `json:"email" gorm:"size:100"`
	Intro     string    `json:"intro" gorm:"size:500"`
	Website   string    `json:"website" gorm:"size:200"`
	Gender    int       `json:"gender" gorm:"default:0;comment:0-保密 1-男 2-女"`
	Status    int       `json:"status" gorm:"default:1;comment:1-正常 0-禁用"`
	LoginType string    `json:"login_type" gorm:"size:20;default:username;comment:登录方式"`
	LastLogin time.Time `json:"last_login"`
	Roles     []Role    `json:"roles" gorm:"many2many:user_roles;"`
}

func (User) TableName() string { return "user" }

// Role 角色表
type Role struct {
	Model
	Name        string `json:"name" gorm:"uniqueIndex;size:50;not null"`
	Label       string `json:"label" gorm:"uniqueIndex;size:50;not null"`
	Description string `json:"description" gorm:"size:200"`
	Sort        int    `json:"sort" gorm:"default:0"`
	Status      int    `json:"status" gorm:"default:1"`
	Menus       []Menu `json:"menus" gorm:"many2many:role_menus;"`
}

func (Role) TableName() string { return "role" }

// Menu 菜单表
type Menu struct {
	Model
	Name       string `json:"name" gorm:"size:50;not null"`
	Title      string `json:"title" gorm:"size:50;comment:菜单标题"`
	Path       string `json:"path" gorm:"size:200"`
	Component  string `json:"component" gorm:"size:200"`
	Redirect   string `json:"redirect" gorm:"size:200"`
	Icon       string `json:"icon" gorm:"size:100"`
	ParentID   uint   `json:"parent_id" gorm:"default:0"`
	Sort       int    `json:"sort" gorm:"default:0"`
	IsHidden   int    `json:"is_hidden" gorm:"default:0"`
	AlwaysShow int    `json:"always_show" gorm:"default:0"`
	KeepAlive  int    `json:"keep_alive" gorm:"default:1"`
	Type       int    `json:"type" gorm:"default:1;comment:1-目录 2-菜单 3-按钮"`
	Children   []Menu `json:"children" gorm:"-"`
}

func (Menu) TableName() string { return "menu" }
