package worker

import (
	"sync"
	"time"

	"gorm.io/gorm"
)

type QorJob struct {
	gorm.Model

	Job    string
	Status string `sql:"default:'new'"`
}

type QorJobInstance struct {
	gorm.Model

	QorJobID uint `gorm:"index"`

	Job    string
	Status string `sql:"default:'new'"`
	Args   string

	Progress     uint
	ProgressText string
	Log          string `sql:"size:65532"`

	jb          *JobBuilder `sql:"-"`
	mutex       sync.Mutex  `sql:"-"`
	stopRefresh bool        `sql:"-"`
	inRefresh   bool        `sql:"-"`
}

type Scheduler interface {
	GetScheduleTime() *time.Time
}

// Schedule could be embedded as job argument, then the job will get run as scheduled feature
type Schedule struct {
	ScheduleTime *time.Time
}

// GetScheduleTime get scheduled time
func (schedule *Schedule) GetScheduleTime() *time.Time {
	if scheduleTime := schedule.ScheduleTime; scheduleTime != nil {
		if scheduleTime.After(time.Now().Add(time.Minute)) {
			return scheduleTime
		}
	}
	return nil
}