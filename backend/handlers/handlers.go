package handlers

import (
	"strings"

	"gorm.io/gorm"
)

var Subs map[*chan string]chan string

func Invalidate(resource string) {
	for k := range Subs {
		Subs[k] <- resource
	}
}

func New() {
	Subs = make(map[*chan string]chan string)
}

func ParseManyQuery(tx *gorm.DB, q *ManyQuery) *gorm.DB {
	if q.Deleted != nil {
		tx = tx.Unscoped().Where("deleted_at != ?", 0)
	}

	if q.Select != nil {
		tx = tx.Select(strings.Split(*q.Select, ","))
	}

	if q.Preload != nil {
		preloads := strings.Split(*q.Preload, ",")
		for kp := range preloads {
			tx = tx.Preload(preloads[kp])
		}
	}

	if q.Limit != nil {
		tx = tx.Limit(*q.Limit)
	}

	if q.Like != nil {
		args := strings.Split(*q.Like, ":")
		tx.Where(args[0]+" like ?", args[1])
	}

	return tx
}

func ParseOneQuery(tx *gorm.DB, q *OneQuery) *gorm.DB {
	if q.Deleted != nil {
		tx = tx.Unscoped().Where("deleted_at != ?", 0)
	}

	if q.Select != nil {
		tx = tx.Select(strings.Split(*q.Select, ","))
	}

	if q.Preload != nil {
		preloads := strings.Split(*q.Preload, ",")
		for kp := range preloads {
			tx = tx.Preload(preloads[kp])
		}
	}

	return tx
}
