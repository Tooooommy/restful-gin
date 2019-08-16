package helpers

import (
	"github.com/jinzhu/now"
	"time"
)

func GetEndOfWeek() time.Time {
	return now.EndOfWeek()
}
