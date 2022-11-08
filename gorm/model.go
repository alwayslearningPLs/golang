package main

import "gorm.io/gorm"

type Food struct {
	gorm.Model
	Name  string `gorm:"column:name;unique;not null" json:"name"`
	Value uint   `gorm:"column:value;not null;" json:"value"`
}

type Warehouse struct {
	gorm.Model
	Name     string `gorm:"column:name;unique;not null" json:"name"`
	Location string `gorm:"column:location;not null" json:"location"`

	FoodID uint
	Food   Food
}

// One-to-one relationship. We need the CompanyID and Company fields to fill the data when eager loading
type User struct {
	gorm.Model
	Name      string
	CompanyID uint
	Company   Company
}

type Company struct {
	ID   uint
	Name string
}

// has-many relationship.
type Library struct {
	gorm.Model
	Name string
	Book []Book
}

type Book struct {
	gorm.Model
	Name      string
	LibraryID uint
}
