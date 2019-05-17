package twitter

import (
	"fmt"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"os"
)

var NewTwClient *twitter.Client

func GetTwClient() (*twitter.Client, error) {
	config := oauth1.NewConfig(os.Getenv("CONSUMER_KEY"), os.Getenv("CONSUMER_SECRET"))
	token := oauth1.NewToken(os.Getenv("ACCESS_TOKEN"), os.Getenv("ACCESS_SECRET"))
	httpClient := config.Client(oauth1.NoContext, token)
	// Twitter client
	client := twitter.NewClient(httpClient)
	NewTwClient = client
	tweet, resp, err := client.Statuses.Update("just setting up my twr", nil)
	fmt.Println(tweet)
	if resp != nil {
	}
	return client, err
}
