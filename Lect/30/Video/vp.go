package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {

	arg := "youtube.md"
	if len(os.Args) > 1 {
		arg = os.Args[1]
	}

	buf, err := ioutil.ReadFile(arg)
	if err != nil {
		fmt.Printf("Unable to open v.txt : %s\n", err)
		os.Exit(1)
	}

	lines := strings.Split(string(buf), "\n")
	for _, line := range lines {
		s2 := strings.Split(line, " ")
		if len(s2) == 0 {
		} else if len(s2) == 2 {
			fmt.Printf("[%s - %s](%s)<br>\n", s2[0], s2[1], s2[0])
		}
	}

	fmt.Printf("\nFrom Amazon S3 - for download (same as youtube videos)\n\n")

	for line_no, line := range lines {
		s2 := strings.Split(line, " ")
		if len(s2) <= 1 {
		} else if len(s2) >= 2 {
			// [http://uw-s20-2015.s3.amazonaws.com/Lect-20-4010-pt1-news.mp4](http://uw-s20-2015.s3.amazonaws.com/Lect-20-4010-pt1-news.mp4)<br>
			fmt.Printf("[http://uw-s20-2015.s3.amazonaws.com/%s](http://uw-s20-2015.s3.amazonaws.com/%s)<br>\n", s2[1], s2[1])
		} else {
			fmt.Fprintf(os.Stderr, "Error on line %d - missing field ->%s<- len=%d\n", line_no+1, s2, len(s2))
		}
	}

}
