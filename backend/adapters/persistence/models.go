package persistence

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type RegularPerson struct {
	Name  string
	Phone string
	TaxId string
	Email string
}

type EmbeddedCreatedAt struct {
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

type User struct {
	gorm.Model
	Id                uuid.UUID         `gorm:"primaryKey"`
	EmbeddedCreatedAt EmbeddedCreatedAt `gorm:"embedded"`
	RegularPerson     RegularPerson     `gorm:"embedded"`
	HashPass          string
	Persons           []Person `gorm:"foreignKey:UserId"`
	Places            []Place  `gorm:"foreignKey:UserId"`
	Bills             []Bill   `gorm:"foreignKey:UserId"`
}

type Person struct {
	gorm.Model
	Id                uuid.UUID         `gorm:"primaryKey"`
	EmbeddedCreatedAt EmbeddedCreatedAt `gorm:"embedded"`
	RegularPerson     RegularPerson     `gorm:"embedded"`
	Description       string
	UserId            uuid.UUID
	Places            []Place `gorm:"foreignKey:PersonId"`
	Bills             []Bill  `gorm:"foreignKey:PersonId"`
}

type Place struct {
	gorm.Model
	Id                uuid.UUID         `gorm:"primaryKey"`
	EmbeddedCreatedAt EmbeddedCreatedAt `gorm:"embedded"`
	Name              string
	Address           string
	Description       string
	PersonId          uuid.UUID
	UserId            uuid.UUID
}

type Bill struct {
	gorm.Model
	Id                uuid.UUID         `gorm:"primaryKey"`
	EmbeddedCreatedAt EmbeddedCreatedAt `gorm:"embedded"`
	Name              string
	BillTypeId        uuid.UUID
	PersonId          uuid.UUID
	UserId            uuid.UUID
}

type BillType struct {
	gorm.Model
	Id                uuid.UUID         `gorm:"primaryKey"`
	EmbeddedCreatedAt EmbeddedCreatedAt `gorm:"embedded"`
	Name              string
	Description       string
	Bills             []Bill `gorm:"foreignKey:BillTypeId"`
}

type Alert struct {
	gorm.Model
	Id                uuid.UUID         `gorm:"primaryKey"`
	EmbeddedCreatedAt EmbeddedCreatedAt `gorm:"embedded"`
	Name              string
	AlertTypeId       uuid.UUID
}

type AlertType struct {
	gorm.Model
	Id                uuid.UUID         `gorm:"primaryKey"`
	EmbeddedCreatedAt EmbeddedCreatedAt `gorm:"embedded"`
	Name              string
	Description       string
	Alerts            []Alert `gorm:"foreignKey:AlertTypeId"`
}

func GetModels() []interface{} {
	return []interface{}{
		&User{},
		&Person{},
		&Place{},
		&AlertType{},
		&Alert{},
		&BillType{},
		&Bill{},
	}
}
