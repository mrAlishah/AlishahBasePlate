package services

import (
	"time"
)

type Config struct {
	Service        string        `yaml:"SERVICE"`
	DefaultTTL     int64         `yaml:"DEFAULT_TTL"`
	CacheDisabled  bool          `yaml:"CACHE_DISABLED"`
	SleepTime      time.Duration `yaml:"SLEEP_TIME"`
	FetchBatchSize int           `yaml:"FETCH_BATCH_SIZE"`
}

type Request struct {
	Email string `form:"email" binding:"required"`
	Pass  string `form:"pass" binding:"required"`
}

type Video struct {
	ID    uint64 `json:"id" bson:"-" gorm:"primaryKey,autoIncrement"`
	Title string `json:"title" bson:"title" gorm:"type:varchar(100)"`
	Desc  string `json:"desc" bson:"desc" gorm:"type:varchar(100)"`
	Path  string `json:"path" bson:"path" gorm:"type:varchar(100)"`
}
