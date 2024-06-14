package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"tableparser/parser"
)

func main() {
	resp, err := http.Get("https://www.espn.com/nba/boxscore/_/gameId/401672980")
	if err != nil {
		fmt.Println(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	sb := string(body)

	f, err := os.Create("xd.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	n3, err := f.WriteString(sb)
	if err != nil {
		panic(err)
	}
	fmt.Printf("wrote %d bytes\n", n3)

	fmt.Println(parser.TableCount(sb))
	parser.ParseTables(sb, 0)
}
