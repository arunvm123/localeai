package models

import (
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

// BookingDetail model
type BookingDetail struct {
	ID                int      `json:"id" gorm:"primary_key;auto_increment:false"`
	UserID            int      `json:"userID"`
	VehicleModelID    int      `json:"vehicleModelID"`
	PackageID         *int     `json:"packageID"`
	TravelTypeID      int      `json:"travelTypeID"`
	FromAreaID        int      `json:"fromAreaID"`
	ToAreaID          *int     `json:"toAreaID"`
	FromCityID        *int     `json:"fromCityID"`
	ToCityID          *int     `json:"toCityID"`
	FromDate          int64    `json:"fromDate"`
	ToDate            int64    `json:"toDate"`
	OnlineBooking     bool     `json:"onlineBooking"`
	MobileSiteBooking bool     `json:"mobileSiteBooking"`
	BookingCreated    int64    `json:"bookingCreated"`
	FromLat           float64  `json:"fromLat"`
	FromLong          float64  `json:"fromLong"`
	ToLat             *float64 `json:"toLat"`
	ToLong            *float64 `json:"toLong"`
	CarCancellation   bool     `json:"carCancellation"`
}

// Create is a helper function to create a new booking detail
func (bd *BookingDetail) Create(db *gorm.DB) error {
	return db.Create(&bd).Error
}

// Save is a helper function to update booking detail
func (bd *BookingDetail) Save(db *gorm.DB) error {
	return db.Save(&bd).Error
}

// CreateBookingDetail creates a new row in table. If one with the ID already exists,
// it updates it with the provided values.
func CreateBookingDetail(db *gorm.DB, bd *BookingDetail) error {
	err := bd.Save(db)
	if err != nil {
		log.WithFields(log.Fields{
			"func":          "CreateBookingDetail",
			"subFunc":       "CreateBookingDetail",
			"bookingDetail": *bd,
		}).Error(err)
		return err
	}

	return nil
}
