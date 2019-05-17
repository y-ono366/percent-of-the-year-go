package message

import (
	"github.com/y-ono366/percent-of-the-year-go/src/common"
	"log"
	"strconv"
	"strings"
	"time"
)

var asciiArtMaxRow = 6
var Log *log.Logger

func init() {
	Log = common.GetLog()
}

func CreateTweetMessage(parcent int) string {
	loc, _ := time.LoadLocation("Asia/Tokyo")
	now := time.Date(time.Now().Year(), 1, 1, 0, 0, 0, 0, loc)

	next := now.AddDate(1, 0, 0)
	message := strconv.Itoa(now.Year()) + "年の" + strconv.Itoa(parcent) + "%が終了しました\n"
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
	return message
}

func createScrollNumOfAsciiArt(key int, parcent int, sliceAsciiArt [][]string) []string {
	yearAsciiArt := []string{}
	oneMemo := 100 / asciiArtMaxRow
	nextYearRow := asciiArtMaxRow - (parcent / oneMemo)
	nowYearRow := 0
	firstFlg := true

	for row := 0; row < asciiArtMaxRow; row++ {
		if row < parcent/oneMemo {
			nextNum := map[bool]int{true: 0, false: key + 1}[key == 9]
			yearAsciiArt = append(yearAsciiArt, sliceAsciiArt[nextNum][nextYearRow])
			nextYearRow++
		} else {
			if firstFlg == true && parcent != 0 {
				yearAsciiArt = append(yearAsciiArt, "")
				firstFlg = false
			}
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
