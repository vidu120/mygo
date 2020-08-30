package main

import (
	"fmt"
)

func main() {
	var a int = 0
	_, _ = fmt.Scanf("%d\n", &a)
	var activity int
	var origin string
	for i := 0; i < a; i++ {
		_, _ = fmt.Scanf("%d %s\n", &activity, &origin)
		var points int = 0
		var pr string
		var pr0 int = 0
		for j := 0; j < activity; j++ {
			_, _ = fmt.Scanf("%s", &pr)
			if pr == "CONTEST_WON" {
				_, _ = fmt.Scanf("%d\n", &pr0)
				points = points + 300
				if pr0 < 20 {
					points = points + 20 - pr0
				}
			} else if pr == "TOP_CONTRIBUTOR" {
				points = points + 300
			} else if pr == "BUG_FOUND" {
				_, _ = fmt.Scanf("%d\n", &pr0)
				points = points + pr0
			} else {
				points = points + 50
			}
		}
		if origin == "INDIAN" {
			points = int(points / 200)
			fmt.Println(points)
		} else {
			points = int(points / 400)
			fmt.Println(points)
		}
	}
}
