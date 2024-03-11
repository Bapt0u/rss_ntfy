package main

import (
	"fmt"
	"github.com/mmcdole/gofeed"
	"log"
	"rssnotify/utils"
	"strings"
)

var conf utils.Conf

func main() {
	var i = 1
	var retry int = 3
	var err error

	conf.Config("./conf/feeder.yml")
	feed := make([]*gofeed.Feed, len(conf.Feeds))

	fp := gofeed.NewParser()
	for y, url := range conf.Feeds {

		// Retry X time to get feed info if parsing fails.
		feed[y], err = fp.ParseURL(url)
		for y := y; err != nil && i <= retry; i++ {
			feed[y], err = fp.ParseURL(url)
			if err != nil {
				log.Print(err)
			}
		}
		// Print error if url isn't reachable after X retries
		// Print Generic error.
		if i == (retry - 1) {
			log.Printf("Error: %s couldn't be parsed. Retried %d times (max %d))", url, i, retry)
		}

		// Reinit loop after exiting it
		i = 1
	}
	for _, r := range feed {
		for _, item := range r.Items {

			// If GUID is in cache, do not notify
			res, err := utils.ReadCache(item.GUID)
			if err != nil {
				log.Fatal(err)
			} else if res {
				log.Printf("%s already sent", item.GUID)
			} else {
				if utils.WriteCache(item.GUID) == nil {

					// Format string to be json friendly
					content := strings.Replace(item.Content, "\n", "\\n", -1)
					content = strings.Replace(content, "\"", "\\\"", -1)

					body := fmt.Sprintf("Title: %s\\nPublished on: %s\\nContent: %s", item.Title, item.Published, content)
					utils.SendNotif(body, conf.Notify)
				}
			}
		}
	}
}
