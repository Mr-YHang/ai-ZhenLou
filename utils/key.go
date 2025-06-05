package utils

import (
	"fmt"
	"time"
)

func CreateKeyByTime(uniqID any) string {
	return fmt.Sprintf("%v_%v", uniqID, time.Now().UnixMilli())
}
