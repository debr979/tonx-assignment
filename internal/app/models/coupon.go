package models

import "time"

type GrabCouponRequest struct {
	CouponId int64 `json:"coupon_id" binding:"required"`
}
type GrabCouponResponse struct{}
type GetCouponsResponse struct {
	Coupons []Coupon `json:"coupons"`
}

type ReserveCouponRequest struct {
	CouponId int64 `json:"coupon_id" binding:"required"`
}

type AddCouponRequest struct {
	CouponType       int       `json:"coupon_type" binding:"required"`
	ReserveStartedAt time.Time `json:"reserve_started_at" binding:"required"`
	ReserveEndedAt   time.Time `json:"reserve_ended_at" binding:"required"`
	GrabStartedAt    time.Time `json:"grab_started_at" binding:"required"`
	GrabEndedAt      time.Time `json:"grab_ended_at" binding:"required"`
}

type DeleteCouponRequest struct {
	CouponId int64 `json:"coupon_id" binding:"required"`
}

type UseCouponRequest struct {
	CouponId int64 `json:"coupon_id" binding:"required"`
}

type UserReservation struct {
	Id         int64     `gorm:"primaryKey;autoIncrement"`
	UserId     int64     `gorm:"type:bigint;index:idx_user_coupon,unique;not null"`
	CouponId   int64     `gorm:"type:bigint;index:idx_user_coupon,unique;not null"`
	ReservedAt time.Time `gorm:"autoCreateTime"`
}

type ReservedCouponCount struct {
	Count int64 `json:"count"`
}

type GrabbableCouponCount struct {
	Count int `json:"count"`
}

type GrabCoupon struct {
	UserId   int64 `json:"user_id"`
	CouponId int64 `json:"coupon_id"`
}
