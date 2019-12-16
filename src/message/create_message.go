package message

import (
	"math"
	"strconv"
	"strings"
	"time"
)

var asciiArtMaxRow = 6

func CreateTweetMessage(parcent int) string {
	loc, _ := time.LoadLocation("Asia/Tokyo")
	now := time.Date(time.Now().Year(), 1, 1, 0, 0, 0, 0, loc)

	next := now.AddDate(1, 0, 0)
	var message string = ""

	if parcent == 0 {
		message = "☆☆☆Happy New Year☆☆☆\n"
	} else {
		message = strconv.Itoa(now.Year()) + "年の" + strconv.Itoa(parcent) + "%が終了しました。\n"
	}
	sliceNextYear := strings.Split(strconv.Itoa(next.Year()), "")
	sliceAsciiArt := getYearAsciiArt()

	sliceAsciiArtMsg := [][]string{}

	for key, val := range strings.Split(strconv.Itoa(now.Year()), "") {
		nVal, _ := strconv.Atoi(val)
		if val == sliceNextYear[key] {
			sliceAsciiArtMsg = append(sliceAsciiArtMsg, sliceAsciiArt[nVal])
		} else {
			sliceAsciiArtMsg = append(sliceAsciiArtMsg, createScrollNumOfAsciiArt(nVal, parcent, sliceAsciiArt))
		}
	}

	for i := 0; i < asciiArtMaxRow; i++ {
		message += sliceAsciiArtMsg[0][i] + sliceAsciiArtMsg[1][i] + sliceAsciiArtMsg[2][i] + sliceAsciiArtMsg[3][i] + "\n"
	}
	message += "#一年の進捗bot"
	return message
}

func createScrollNumOfAsciiArt(key int, parcent int, sliceAsciiArt [][]string) []string {
	fParcent := float64(parcent)
	fAsciiArtMaxRow := float64(asciiArtMaxRow)
	yearAsciiArt := []string{}
	oneMemo := 100.0 / fAsciiArtMaxRow
	nextYearRow := fAsciiArtMaxRow - (fParcent / oneMemo)
	nowYearRow := 0

	for row := float64(1); row <= fAsciiArtMaxRow; row++ {
		if row <= fParcent/oneMemo {
			nextNum := map[bool]int{true: 0, false: key + 1}[key == 9]
			yearAsciiArt = append(yearAsciiArt, sliceAsciiArt[nextNum][int(math.Ceil(nextYearRow))])
			nextYearRow++
		} else {
			yearAsciiArt = append(yearAsciiArt, sliceAsciiArt[key][nowYearRow])
			nowYearRow++
		}
	}
	return yearAsciiArt
}

func getYearAsciiArt() [][]string {
	return [][]string{
		{
			"┏━━┓",
			"┃┏┓┃",
			"┃┃┃┃",
			"┃┃┃┃",
			"┃┗┛┃",
			"┗━━┛",
		},
		{
			"┏┓",
			"┃┃",
			"┃┃",
			"┃┃",
			"┃┃",
			"┗┛",
		},
		{
			"┏━━┓",
			"┗━┓┃",
			"┏━┛┃",
			"┃┏━┛",
			"┃┗━┓",
			"┗━━┛",
		},
		{
			"┏━━┓",
			"┗━┓┃",
			"┏━┛┃",
			"┗━┓┃",
			"┏━┛┃",
			"┗━━┛",
		},
		{
			"┏┓┏┓",
			"┃┃┃┃",
			"┃┗┛┃",
			"┗━┓┃",
			"⠀⠀⠀┃┃",
			"⠀⠀⠀┗┛",
		},
		{
			"┏━━┓",
			"┃┏━┛",
			"┃┗━┓",
			"┗━┓┃",
			"┏━┛┃",
			"┗━━┛",
		},
		{
			"┏┓⠀⠀",
			"┃┃⠀⠀",
			"┃┗━┓",
			"┃┏┓┃",
			"┃┗┛┃",
			"┗━━┛",
		},
		{
			"┏━━┓",
			"┗━┓┃",
			"⠀⠀⠀┃┃",
			"⠀⠀⠀┃┃",
			"⠀⠀⠀┃┃",
			"⠀⠀⠀┗┛",
		},
		{
			"┏━━┓",
			"┃┏┓┃",
			"┃┗┛┃",
			"┃┏┓┃",
			"┃┗┛┃",
			"┗━━┛",
		},
		{
			"┏━━┓",
			"┃┏┓┃",
			"┃┗┛┃",
			"┗━┓┃",
			"⠀⠀⠀┃┃",
			"⠀⠀⠀┗┛",
		},
	}
}
