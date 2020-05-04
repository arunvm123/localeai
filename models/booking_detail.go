package models

import (
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

// BookingDetail model
type BookingDetail struct {
	ID                int      `json:"id" gorm:"primary_key;auto_increment:false" validate:"required"`
	UserID            int      `json:"userID" validate:"required"`
	VehicleModelID    int      `json:"vehicleModelID" validate:"required"`
	PackageID         *int     `json:"packageID"`
	TravelTypeID      int      `json:"travelTypeID" validate:"required"`
	FromAreaID        int      `json:"fromAreaID" validate:"required"`
	ToAreaID          *int     `json:"toAreaID"`
	FromCityID        *int     `json:"fromCityID"`
	ToCityID          *int     `json:"toCityID"`
	FromDate          int64    `json:"fromDate" validate:"required"`
	ToDate            int64    `json:"toDate" validate:"required"`
	MobileSiteBooking bool     `json:"mobileSiteBooking"`
	OnlineBooking     bool     `json:"onlineBooking"`
	BookingCreated    int64    `json:"bookingCreated" validate:"required"`
	FromLat           float64  `json:"fromLat" validate:"required"`
	FromLong          float64  `json:"fromLong" validate:"required"`
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

// SaveBookingDetail creates a new row in table. If one with the ID already exists,
// it updates it with the provided values.
func SaveBookingDetail(db *gorm.DB, bd *BookingDetail) error {
	err := bd.Save(db)
	if err != nil {
		log.WithFields(log.Fields{
			"func":          "CreateBookingDetail",
			"subFunc":       "bd.Save",
			"bookingDetail": *bd,
		}).Error(err)
		return err
	}

	return nil
}
