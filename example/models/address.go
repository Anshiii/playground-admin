package models

import (
	"time"

	"github.com/anshiii/playground-admin/media/media_library"
	"github.com/anshiii/playground-admin/publish"
)

type Customer struct {
	ID             uint `gorm:"primarykey"`
	Name           string
	Addresses      []*Address
	MembershipCard *MembershipCard
}

type Address struct {
	ID         uint `gorm:"primarykey"`
	CustomerID uint

	Street    string
	HomeImage media_library.MediaBox `sql:"type:text;"`
	UpdatedAt time.Time
	CreatedAt time.Time

	publish.Status
	Phones []*Phone
}

type Phone struct {
	ID        uint `gorm:"primarykey"`
	AddressID uint
	Number    int
}

type MembershipCard struct {
	ID          uint `gorm:"primarykey"`
	CustomerID  uint
	Number      int
	ValidBefore *time.Time
}
