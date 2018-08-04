package main

import (
	"fmt"
	"strings"

	api "github.com/moshee/go-4chan-api/api"
	html "golang.org/x/net/html"
)

func main() {
	var board = "a"
	allThreadIdsPerBoard, err := api.GetThreads(board)
	if err != nil {
		panic(err)
	}

	for _, allThreadIdsPerPage := range allThreadIdsPerBoard {
		for _, threadID := range allThreadIdsPerPage {
			thread, err := api.GetThread(board, threadID)
			if err != nil {
				panic(err)
			}
			for _, post := range thread.Posts {
				fmt.Println(parser(post.Comment))
			}

		}
	}

}

func parser(raw string) string {
	var s string
	domDocTest := html.NewTokenizer(strings.NewReader(raw))
	for tokenType := domDocTest.Next(); tokenType != html.ErrorToken; {
		TxtContent := strings.TrimSpace(html.UnescapeString(string(domDocTest.Text())))
		if len(TxtContent) > 0 {
			s = TxtContent
		}
		tokenType = domDocTest.Next()
	}
	return s

}
