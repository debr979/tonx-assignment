package repositories

import (
	"tonx-assignment/internal/app/db"
	"tonx-assignment/internal/app/models"
)

type couponRepository struct{}

var CouponRepository couponRepository

const (
	BATCH_SIZE = 2000
)

func (r *couponRepository) GetCoupons(model any) error {
	dbConn := db.Conn()
	return dbConn.Where("is_available = true and grab_ended_at > now()").Order("id desc").Find(model).Error
}

func (r *couponRepository) Reserve(model any) error {
	dbConn := db.Conn()
	return dbConn.CreateInBatches(model, BATCH_SIZE).Error
}

func (r *couponRepository) Grab(model any) error {
	dbConn := db.Conn()

	return dbConn.CreateInBatches(model, BATCH_SIZE).Error
}

func (r *couponRepository) UseCoupon(couponId int64) error {
	dbConn := db.Conn()
	return dbConn.Where("coupon_id = ? and is_used = false", couponId).Model(&models.UserCoupon{}).Update("is_used", true).Error
}

func (r *couponRepository) AddCoupon(model any) error {
	dbConn := db.Conn()
	return dbConn.Create(model).Error
}

func (r *couponRepository) DeleteCoupon(couponId int64) error {
	dbConn := db.Conn()
	return dbConn.Where("id = ?", couponId).Model(&models.Coupon{}).Update("is_available", false).Error
}

func (r *couponRepository) GetCouponStatus(model any, couponId int64) error {
	dbConn := db.Conn()
	return dbConn.Where("coupon_id = ?", couponId).Model(model).Error
}
