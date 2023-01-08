package main

import (
    "encoding/json"
    "flag"
    "fmt"
    "io/ioutil"
    "os"

    "github.com/mmcdole/gofeed"
)

// JSON list of feeds to read from with hashtags
type feedConfig struct {
	Name     string   `json:"name"`
	URL      string   `json:"url"`
	Hashtags []string `json:"hashtags"`
}

func main() {
    configFile := flag.String("config", "config.json", "Path to the config file")
    flag.Parse()

    configData, err := ioutil.ReadFile(*configFile)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    var feeds []feedConfig
    err = json.Unmarshal(configData, &feeds)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    fp := gofeed.NewParser()
    for _, feed := range feeds {
		fmt.Println("Name:", feed.Name)
		fmt.Println("Feed:", feed.URL)
		fmt.Println("Hashtags:", feed.Hashtags)
        
        parsedFeed, err := fp.ParseURL(feed.URL)
        if err != nil {
            fmt.Println(err)
            continue
        }

        fmt.Println(parsedFeed.Title)
        for _, item := range parsedFeed.Items {
            fmt.Println(item.Title)
        }
        fmt.Println()
    }
}

