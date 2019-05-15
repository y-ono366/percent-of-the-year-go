package message

import (
	"fmt"
	"time"
)

func GetParcent() int {
	getParcent := getArrParcentDays()
	fmt.Println(getParcent)
	return 3
}

func getArrParcentDays() []int64 {
	getParcent := []int64{}

	loc, _ := time.LoadLocation("Asia/Tokyo")
	start := time.Date(time.Now().Year(), 1, 1, 0, 0, 0, 0, loc)
	next := start.AddDate(1, 0, 0)
	oneYearTimestamp := next.Unix() - start.Unix()
	oneParcentSec := oneYearTimestamp / 100

	for i := 0; i <= 100; i++ {
		getParcent = append(getParcent, (start.Unix() + (oneParcentSec * int64(i))))
	}

	return getParcent
}
