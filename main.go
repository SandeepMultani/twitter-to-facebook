package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/SandeepMultani/twitter-to-facebook/twittershot"
)

func main() {
	twitterHandle := "engineersinghuk"
	tweetId := 1415660091924168704
	path := fmt.Sprintf("https://twitter.com/%s/status/%d", twitterHandle, tweetId)
	sel := "#react-root article"

	log.Println(path)
	log.Println(sel)

	imgBytes, err := twittershot.Screenshot(path, sel)
	if err != nil {
		log.Fatalln(err)
		return
	}

	err = saveImage("test", imgBytes)
	if err != nil {
		log.Fatalln(err)
		return
	}
}

func saveImage(filename string, imgByte []byte) error {
	if err := ioutil.WriteFile(fmt.Sprintf("images/%s.png", filename), imgByte, 0o644); err != nil {
		return err
	}

	return nil
}
