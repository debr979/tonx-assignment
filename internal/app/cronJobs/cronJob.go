package cronJobs

import (
	"encoding/json"
	"fmt"
	"github.com/go-co-op/gocron/v2"
	"math"
	"time"
	"tonx-assignment/internal/app/appRedis"
	"tonx-assignment/internal/app/models"
	"tonx-assignment/internal/app/repositories"
	"tonx-assignment/pkg/utils"
)

type cronJob struct {
}

var Cron cronJob

func (r *cronJob) Run() {

	scheduler, err := gocron.NewScheduler()
	if err != nil {
		utils.Logger.LogOutput(err)
	}

	defer func() {
		scheduler.Start()
	}()
	today := time.Now()

	rdb := appRedis.New()
	// auto create a new daily coupon
	_, err = scheduler.NewJob(
		gocron.DailyJob(1, gocron.NewAtTimes(gocron.NewAtTime(13, 14, 0))),
		gocron.NewTask(func(todayTime time.Time) {
			model := models.Coupon{
				CouponType:       1,
				ReserveStartedAt: time.Date(todayTime.Year(), todayTime.Month(), todayTime.Day(), 22, 55, 0, 0, todayTime.Location()),
				ReserveEndedAt:   time.Date(todayTime.Year(), todayTime.Month(), todayTime.Day(), 22, 58, 0, 0, todayTime.Location()),
				GrabStartedAt:    time.Date(todayTime.Year(), todayTime.Month(), todayTime.Day(), 22, 59, 0, 0, todayTime.Location()),
				GrabEndedAt:      time.Date(todayTime.Year(), todayTime.Month(), todayTime.Day(), 23, 0, 0, 0, todayTime.Location()),
			}

			_ = repositories.CouponRepository.AddCoupon(&model)
			key := todayTime.Format(time.DateOnly)
			b, _ := json.Marshal(model)

			if err := rdb.Set("coupon", key, string(b)); err != nil {
				utils.Logger.LogOutput(err)
			}
		}, today),
	)
	if err != nil {
		utils.Logger.LogOutput(err)
	}

	_, err = scheduler.NewJob(
		gocron.DailyJob(1, gocron.NewAtTimes(gocron.NewAtTime(22, 58, 1))),
		gocron.NewTask(func(todayTime time.Time) {
			key := todayTime.Format(time.DateOnly)
			keyCount, err := rdb.Count(fmt.Sprintf("reserveCoupon:%s", key), "user_id:*")
			if err != nil {
				utils.Logger.LogOutput(err)
			}

			if keyCount > 0 {
				field := "totalCount"
				partField := "partField"
				if err := rdb.Set("reserveCouponCount", field, keyCount); err != nil {
					utils.Logger.LogOutput(err)
				}

				partCount := math.Ceil(float64(keyCount) * 0.2)
				if err := rdb.Set("reserveCouponCount", partField, partCount); err != nil {
					utils.Logger.LogOutput(err)
				}
			}
		}, today),
	)
	if err != nil {
		utils.Logger.LogOutput(err)
	}

}
