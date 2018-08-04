package main

import (
	"bytes"
	"io"
	"os"
	"strings"

	"github.com/moshee/go-4chan-api/api"
	"golang.org/x/net/html"
)

func scrapp(board string) {
	var buffer bytes.Buffer
	allThreadIdsPerBoard, err := api.GetThreads(board)
	check(err)
	for _, allThreadIdsPerPage := range allThreadIdsPerBoard[0:4] {
		for _, threadID := range allThreadIdsPerPage {

			thread, err := api.GetThread(board, threadID)
			check(err)
			for _, post := range thread.Posts {
				//fmt.Println(parser(post.Comment))
				buffer.WriteString(parser(post.Comment))
				buffer.WriteString("\n")
			}

		}
	}

	e := WriteStringToFile("./data/"+board+".txt", buffer.String())
	check(e)
	//fmt.Println(buffer.String())

}

func parser(raw string) string {
	var s string
	domDocTest := html.NewTokenizer(strings.NewReader(raw))
	for tokenType := domDocTest.Next(); tokenType != html.ErrorToken; {
		TxtContent := strings.TrimSpace(html.UnescapeString(string(domDocTest.Text())))
		if len(TxtContent) > 5 && TxtContent[:2] != ">>" {
			s = TxtContent
		}
		tokenType = domDocTest.Next()
	}
	return s

}

//WriteStringToFile dsdsdsd
func WriteStringToFile(filepath, s string) error {
	fo, err := os.Create(filepath)
	check(err)
	defer fo.Close()

	_, err = io.Copy(fo, strings.NewReader(s))
	check(err)
	return nil
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
