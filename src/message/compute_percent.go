package message

import (
	"time"
)

func GetParcent() (int, bool) {
	sliceParcents := getArrParcentDays()
	parcent, isFlg := getKeyFromSliceParcentDays(sliceParcents)
	return parcent, isFlg
}

func getArrParcentDays() []int64 {
	sliceParcents := []int64{}

	loc, _ := time.LoadLocation("Asia/Tokyo")
	start := time.Date(time.Now().Year(), 1, 1, 0, 0, 0, 0, loc)
	next := start.AddDate(1, 0, 0)
	oneYearTimestamp := next.Unix() - start.Unix()
	oneParcentSec := oneYearTimestamp / 100

	for i := 0; i <= 100; i++ {
		sliceParcents = append(sliceParcents, (start.Unix() + (oneParcentSec * int64(i))))
	}

	return sliceParcents
}

func getKeyFromSliceParcentDays(sliceParcents []int64) (int, bool) {
	nowUnix := time.Now().Unix()
	for i := 0; i < len(sliceParcents); i++ {
		if sliceParcents[i]+60 >= nowUnix && sliceParcents[i]-60 <= nowUnix {
			return i, true
		}
	}
	return 0, false
}
