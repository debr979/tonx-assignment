package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tonx-assignment/internal/app/models"
	"tonx-assignment/internal/app/services"
)

type couponController struct {
	baseControllers
}

var CouponController couponController

// GetCoupons
// @Title GetCoupons
// @Tags coupon
// @Description get current available coupons
// @Summary get current available coupons
// @Success 200 {object} models.GetCouponsResponse ""
// @Failure 400 {object} string "api error"
// @Router /coupons/getCoupons [get]
func (r *couponController) GetCoupons(c *gin.Context) {
	coupons, err := services.CouponService.GetCoupons()
	if err != nil {
		r.Response(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), nil)
	}
	var result models.GetCouponsResponse
	result.Coupons = coupons.([]models.Coupon)
	r.Request(c).Response(http.StatusOK, http.StatusText(http.StatusOK), result)
}

// AddCoupon
// @Title AddCoupon
// @Tags Manager
// @Description create coupon
// @Summary create coupon
// @Param body body models.AddCouponRequest true "AddCouponRequest"
// @Success 200 {bool} bool ""
// @Failure 400 {object} string "api error"
// @Router /mgr/coupons [post]
func (r *couponController) AddCoupon(c *gin.Context) {
	var req models.AddCouponRequest
	if err := r.Request(c).ParseBody(&req); err != nil {
		r.Response(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), nil)
		return
	}

	if req.ReserveStartedAt.After(req.ReserveEndedAt) || req.GrabStartedAt.After(req.ReserveEndedAt) || req.GrabStartedAt.After(req.GrabEndedAt) {
		r.Response(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), nil)
		return
	}

	couponId, err := services.CouponService.AddCoupon(req)
	if err != nil {
		r.Response(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), nil)
		return
	}

	r.Response(http.StatusOK, http.StatusText(http.StatusOK), couponId)
}

// DeleteCoupon
// @Title DeleteCoupon
// @Tags Manager
// @Description delete coupon
// @Summary delete coupon
// @Param body body models.DeleteCouponRequest true "DeleteCouponRequest"
// @Success 200 {bool} bool ""
// @Failure 400 {object} string "api error"
// @Router /mgr/coupons [post]
func (r *couponController) DeleteCoupon(c *gin.Context) {
	var req models.DeleteCouponRequest
	if err := r.Request(c).ParseBody(&req); err != nil {
		r.Response(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), nil)
		return
	}

	if req.CouponId <= 0 {
		r.Response(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), nil)
		return
	}

	isDeleted, err := services.CouponService.DeleteCoupon(req)
	if err != nil {
		r.Response(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), nil)
		return
	}

	r.Response(http.StatusOK, http.StatusText(http.StatusOK), isDeleted)
}

// Reserve
// @Title Reserve
// @Tags coupons
// @Description reserve coupon
// @Summary reserve coupon
// @Param body body models.ReserveCouponRequest true "ReserveCouponRequest"
// @Success 200 {bool} bool ""
// @Failure 400 {object} string "api error"
// @Router /coupons/reserve [post]
func (r *couponController) Reserve(c *gin.Context) {
	var req models.ReserveCouponRequest
	if err := r.Request(c).ParseBody(&req); err != nil {
		r.Response(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), nil)
		return
	}

	if req.CouponId <= 0 {
		r.Response(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), nil)
		return
	}

	userId, isExist := c.Get("user_id")
	if !isExist {
		r.Response(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), nil)
		return
	}

	result, err := services.CouponService.Reserve(req, userId.(int64))
	if err != nil {
		r.Response(http.StatusInternalServerError, err.Error(), nil)
		return
	}

	r.Response(http.StatusOK, http.StatusText(http.StatusOK), result)
}

// Grab
// @Title Grab
// @Tags coupons
// @Description grab coupon
// @Summary grab coupon
// @Param body body models.GrabCouponRequest true "GrabCouponRequest"
// @Success 200 {bool} bool""
// @Failure 400 {object} string "api error"
// @Router /coupons/grab [post]
func (r *couponController) Grab(c *gin.Context) {
	var req models.GrabCouponRequest
	if err := r.Request(c).ParseBody(&req); err != nil {
		r.Response(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), nil)
		return
	}

	if req.CouponId <= 0 {
		r.Response(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), nil)
		return
	}

	userId, isExist := c.Get("user_id")
	if !isExist {
		r.Response(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), nil)
		return
	}

	result, err := services.CouponService.Grab(req, userId.(int64))
	if err != nil {
		r.Response(http.StatusInternalServerError, err.Error(), nil)
		return
	}

	r.Response(http.StatusOK, http.StatusText(http.StatusOK), result)

}

// UseCoupon
// @Title UseCoupon
// @Tags coupon
// @Description use coupon
// @Summary use coupon
// @Param body body models.UseCouponRequest true "UseCouponRequest"
// @Success 200 {bool} bool ""
// @Failure 400 {object} string "api error"
// @Router /coupons/useCoupon [post]
func (r *couponController) UseCoupon(c *gin.Context) {
	var req models.UseCouponRequest
	if err := r.Request(c).ParseBody(&req); err != nil {
		r.Response(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), nil)
		return
	}

	if req.CouponId <= 0 {
		r.Response(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), nil)
		return
	}

	isSuccess, err := services.CouponService.UseCoupon(req)
	if err != nil {
		r.Response(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), nil)
		return
	}

	r.Response(http.StatusOK, http.StatusText(http.StatusOK), isSuccess)
}
