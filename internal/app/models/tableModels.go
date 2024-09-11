package models

import "time"

type Coupon struct {
	Id               int64     `gorm:"primaryKey;autoIncrement;index:sort:desc"`
	IsAvailable      bool      `gorm:"type:boolean;not null;default:true"`
	CouponType       int       `gorm:"type:smallint;not null;comment: 1:daily,2:event,3:other"`
	ReserveStartedAt time.Time `gorm:"type:timestamp;not null;check:reserve_started_at > now();comment: start reserve time"`
	ReserveEndedAt   time.Time `gorm:"type:timestamp;not null;check:reserve_ended_at > reserve_started_at;comment: end reserve time should later than reserve_started_at"`
	GrabStartedAt    time.Time `gorm:"type:timestamp;not null;check:grab_started_at > reserve_ended_at"`
	GrabEndedAt      time.Time `gorm:"type:timestamp;not null;check:grab_ended_at > grab_started_at"`
	CreatedAt        time.Time `gorm:"autoCreateTime"`
	UpdatedAt        time.Time `gorm:"autoUpdateTime"`
}

type User struct {
	Id        int64     `gorm:"primaryKey;autoIncrement"`
	Username  string    `gorm:"type:varchar(100);not null;unique"`
	Password  string    `gorm:"type:text;not null"`
	IsDeleted bool      `gorm:"type:boolean;not null;default:false"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

type UserCoupon struct {
	Id        int64     `gorm:"primaryKey;autoIncrement"`
	UserId    int64     `gorm:"type:bigint;index:idx_user_coupon,unique;not null"`
	CouponId  int64     `gorm:"type:bigint;index:idx_user_coupon,unique;not null"`
	User      User      `gorm:"foreignKey:UserId"`
	Coupon    Coupon    `gorm:"foreignKey:CouponId"`
	IsUsed    bool      `gorm:"type:boolean;default:false;comment:coupon is used"`
	UpdatedAt time.Time `gorm:"autoCreateTime"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}

type Manager struct {
	Id          int64     `gorm:"primaryKey;autoIncrement"`
	ManagerName string    `gorm:"type:varchar(100);not null;unique"`
	Password    string    `gorm:"type:text;not null"`
	IsAvailable bool      `gorm:"type:boolean;not null;default:true"`
	UpdatedAt   time.Time `gorm:"autoCreateTime"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
}
