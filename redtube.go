package redtube

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

const apiURL = "https://api.redtube.com/"
const APITimeout = 3

func SearchVideos(search string) (RedtubeSearchResult) {
	timeout := time.Duration(APITimeout * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	resp, _ := client.Get(fmt.Sprintf(apiURL+"?data=redtube.Videos.searchVideos&output=json&search=%s&thumbsize=all", url.QueryEscape(search)))
	b, _ := ioutil.ReadAll(resp.Body)
	var result RedtubeSearchResult
	err := json.Unmarshal(b, &result)
	if err != nil {
		log.Fatal(err)
	}
	return result
}

func GetVideoByID(ID string) (RedtubeSingleVideo) {
	resp, _ := http.Get(fmt.Sprintf(apiURL+"?data=redtube.Videos.getVideoById&video_id=%s&output=json", ID))
	b, _ := ioutil.ReadAll(resp.Body)
	var result RedtubeSingleVideo
	err := json.Unmarshal(b, &result)
	if err != nil {
		log.Fatal(err)
	}
	return result
}

func GetVideoEmbedCode(ID string) (RedtubeEmbedCode) {
	resp, _ := http.Get(fmt.Sprintf(apiURL+"?data=redtube.Videos.getVideoEmbedCode&video_id=%s&output=json", ID))
	b, _ := ioutil.ReadAll(resp.Body)
	var result RedtubeEmbedCode
	err := json.Unmarshal(b, &result)
	if err != nil {
		log.Fatal(err)
	}
	return result
}