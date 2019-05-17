package twitter

import (
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/y-ono366/percent-of-the-year-go/src/common"
	"os"
)

type Twitter struct {
	Client *twitter.Client
}

func (tw *Twitter) GetTwClient() {
	config := oauth1.NewConfig(os.Getenv("CONSUMER_KEY"), os.Getenv("CONSUMER_SECRET"))
	token := oauth1.NewToken(os.Getenv("ACCESS_TOKEN"), os.Getenv("ACCESS_SECRET"))
	httpClient := config.Client(oauth1.NoContext, token)
	// Twitter client
	tw.Client = twitter.NewClient(httpClient)
}

func (tw *Twitter) Post(msg string) {
	log := common.GetLog()
	env := common.GetEnv()
	if env != "production" {
		log.Info(msg)
	}
	if env == "production" {
		if _, _, err := tw.Client.Statuses.Update(msg, nil); err != nil {
			log.Error(err)
		}
	}
}
