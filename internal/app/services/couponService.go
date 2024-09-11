package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"
	"tonx-assignment/internal/app/appRedis"
	"tonx-assignment/internal/app/models"
	"tonx-assignment/internal/app/repositories"
)

type couponService struct{}

var CouponService couponService

func (r *couponService) GetCoupons() (any, error) {
	coupons := make([]models.Coupon, 0)
	if err := repositories.CouponRepository.GetCoupons(&coupons); err != nil {
		return nil, err
	}
	return coupons, nil
}

func (r *couponService) Reserve(req models.ReserveCouponRequest, userId int64) (any, error) {
	now := time.Now()
	nowDate := now.Format(time.DateOnly)
	rdb := appRedis.New()
	reserveKey := fmt.Sprintf("reserveCoupon:%s", nowDate)
	field := fmt.Sprintf("userId:%d", userId)

	isReserved, err := rdb.Exists(reserveKey, field)
	if err != nil {
		return nil, err
	}

	if isReserved {
		// if reserved
		return nil, errors.New("reserved coupon")
	}

	couponData, err := rdb.Get("coupon", nowDate)
	if err != nil {
		return nil, err
	}

	var couponModel models.Coupon
	if err := json.Unmarshal([]byte(couponData), &couponModel); err != nil {
		return nil, err
	}

	if couponModel.Id != req.CouponId {
		return nil, errors.New("invalid coupon id")
	}

	if couponModel.ReserveStartedAt.After(now) || couponModel.ReserveEndedAt.Before(now) {
		return nil, errors.New("invalid reserve time")
	}

	var reserve models.UserReservation
	reserve.UserId = userId
	reserve.ReservedAt = now
	reserve.CouponId = couponModel.Id
	b, _ := json.Marshal(reserve)
	if err := rdb.Set(reserveKey, field, string(b)); err != nil {
		return nil, err
	}

	return true, nil
}

func (r *couponService) Grab(req models.GrabCouponRequest, userId int64) (any, error) {
	now := time.Now()
	nowDate := now.Format(time.DateOnly)
	rdb := appRedis.New()
	reserveKey := fmt.Sprintf("reserveCoupon:%s", nowDate)
	field := fmt.Sprintf("userId:%d", userId)
	isReserved, err := rdb.Exists(reserveKey, field)
	if err != nil {
		return nil, err
	}

	if !isReserved {
		return nil, errors.New("not reserved coupon")
	}

	couponData, err := rdb.Get("coupon", nowDate)
	if err != nil {
		return nil, err
	}
	var couponModel models.Coupon
	if err := json.Unmarshal([]byte(couponData), &couponModel); err != nil {
		return nil, err
	}

	if couponModel.Id != req.CouponId {
		return nil, errors.New("invalid coupon id")
	}

	if couponModel.GrabStartedAt.After(now) || couponModel.GrabEndedAt.Before(now) {
		return nil, errors.New("invalid reserve time")
	}

	grabKey := fmt.Sprintf("grabCoupon:%s", nowDate)
	var grabCoupon models.GrabCoupon
	grabCoupon.UserId = userId
	grabCoupon.CouponId = couponModel.Id
	b, _ := json.Marshal(grabCoupon)

	if err := rdb.Set(grabKey, field, string(b)); err != nil {
		return nil, err
	}

	return true, nil
}

func (r *couponService) UseCoupon(req models.UseCouponRequest) (any, error) {
	// check coupon is available
	var couponStatus models.Coupon
	if err := repositories.CouponRepository.GetCouponStatus(&couponStatus, req.CouponId); err != nil {
		return nil, err
	}

	if couponStatus.Id == 0 || !couponStatus.IsAvailable {
		return nil, errors.New("coupon it does not exist or not available")
	}

	// update coupon status
	if err := repositories.CouponRepository.UseCoupon(req.CouponId); err != nil {
		return nil, err
	}

	return true, nil
}
func (r *couponService) AddCoupon(req models.AddCouponRequest) (any, error) {
	var model models.Coupon
	model.CouponType = req.CouponType
	model.ReserveStartedAt = req.ReserveStartedAt
	model.ReserveEndedAt = req.ReserveEndedAt
	model.GrabStartedAt = req.GrabStartedAt
	model.GrabEndedAt = req.GrabEndedAt

	if err := repositories.CouponRepository.AddCoupon(&model); err != nil {
		return nil, err
	}

	return model.Id, nil
}

func (r *couponService) DeleteCoupon(req models.DeleteCouponRequest) (any, error) {
	if err := repositories.CouponRepository.DeleteCoupon(req.CouponId); err != nil {
		return nil, err
	}

	return true, nil
}
