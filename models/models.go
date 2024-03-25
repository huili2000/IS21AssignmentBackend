package models

import (
	"fmt"

	"gorm.io/gorm"
)

// user struct & related functions
// TODO: manage Users using IAM (future new user stories)
type User struct {
	Id         uint   `json:"id"`
	Name       string `json:"name"`
	Role       string `json:"role"`
	Permission string `json:"permission"`
	Password   string `json:"password"`
}

func AuthUser(db *gorm.DB, password, name string) (*[]User, error) {
	users := &[]User{}
	if err := db.Where("password = ? AND name = ?", password, name).Find(users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func CreateUser(db *gorm.DB, user *User) error {
	err := db.Create(user).Error
	if err != nil {
		return err
	}
	return nil
}

func GetAllUsers(db *gorm.DB) (*[]User, error) {
	users := &[]User{}
	if err := db.Find(users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func GetUserByName(db *gorm.DB, name string) (*User, error) {
	user := &User{}
	if err := db.Where("name = ?", name).First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func UpdateUser(db *gorm.DB, name, permission, role, password string) error {
	if err := db.Model(&User{}).Where("name = ?", name).Updates(User{Permission: permission, Role: role, Password: password}).Error; err != nil {
		return err
	}
	return nil
}

func DeleteUserByName(db *gorm.DB, name string) error {
	if err := db.Where("name = ?", name).Delete(&User{}).Error; err != nil {
		return err
	}
	return nil
}

// Paint struct & related functions
type Paint struct {
	Id       uint   `json:"id"`
	Color    string `json:"color"`
	Quantity uint   `json:"quantity"`
}

func CreatePaint(db *gorm.DB, paint *Paint) error {
	err := db.Create(paint).Error
	if err != nil {
		return err
	}
	return nil
}

func GetAllPaints(db *gorm.DB) (*[]Paint, error) {
	paints := &[]Paint{}
	if err := db.Find(paints).Error; err != nil {
		return nil, err
	}
	return paints, nil
}

func GetPaintByColor(db *gorm.DB, color string) (*Paint, error) {
	paint := &Paint{}
	if err := db.Where("color = ?", color).First(paint).Error; err != nil {
		return nil, err
	}
	return paint, nil
}

func ProvisionPaint(db *gorm.DB, color string, quantity uint) error {
	paint := &Paint{}
	if err := db.Where("color = ?", color).First(paint).Error; err != nil {
		return err
	}
	if err := db.Model(&Paint{}).Where("color = ?", color).Updates(Paint{Quantity: paint.Quantity + quantity}).Error; err != nil {
		return err
	}
	return nil
}

func ConsumePaint(db *gorm.DB, color string, quantity uint) error {
	paint := &Paint{}
	if err := db.Where("color = ?", color).First(paint).Error; err != nil {
		return err
	}
	if paint.Quantity < quantity {
		return fmt.Errorf("cannot ask for %d quantity of %s paint, only %d is available", quantity, color, paint.Quantity)
	}
	if err := db.Model(&Paint{}).Where("color = ?", color).Updates(Paint{Quantity: paint.Quantity - quantity}).Error; err != nil {
		return err
	}
	return nil
}
