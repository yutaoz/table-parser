package parser

import (
	"fmt"
	"regexp"
	"strings"
)

type Table struct {
	Tab  [][]string
	Rows int
	Cols int
}

func TableCount(html string) int {
	count := strings.Count(html, "<table")

	return count
}

func filter2DBlanks(slice [][]string) [][]string {
	var filtered [][]string
	for _, innerSlice := range slice {
		if len(innerSlice) > 1 && strings.TrimSpace(innerSlice[1]) != "" {
			filtered = append(filtered, innerSlice)
		}
	}
	return filtered
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// given a <table> raw html, generates a table
func generateTable(html string) Table {
	rows := 0
	cols := 0

	table := Table{
		Rows: rows,
		Cols: cols,
		Tab:  make([][]string, rows),
	}
	start := strings.Index(html, "<thead")
	end := strings.Index(html, "</thead") + len("</thead")
	txt := html[start:end]
	rows += strings.Count(txt, "<tr")

	// iterate over rows and count number of cells for header (usually only 1 row, but just in case xD)
	for i := 0; i < rows; i++ {
		rstart := strings.Index(txt, "<tr")
		rend := strings.Index(txt, "</tr>") + len("</tr>")
		rtxt := txt[rstart:rend]
		ccnt := strings.Count(rtxt, "<th")
		rw := txt[rstart:rend]
		rwslice := []string{}
		// iterate thru cells, regex to find value within a cell
		for j := 0; j < ccnt; j++ {
			cstart := strings.Index(rw, "<th")
			cend := strings.Index(rw, "</th>") + len("</th>")
			rwtx := rw[cstart:cend]
			fmt.Println(rw[cstart:cend])
			re := regexp.MustCompile(`>([^<>]*)<`)

			// find all matches
			matches := re.FindAllStringSubmatch(rwtx, -1)
			fmt.Println(matches[0][1])
			rwslice = append(rwslice, matches[0][1])
			rw = rw[cend:]
		}
		fmt.Println(rwslice)
		table.Rows += 1
		table.Cols = maxInt(table.Cols, len(rwslice))
		table.Tab = append(table.Tab, rwslice)

		fmt.Println(table)

		fmt.Println(txt[rstart:rend])
		txt = txt[rend:]
	}
	fmt.Println(cols)

	start = strings.Index(html, "<tbody")
	end = strings.Index(html, "</tbody>") + len("</tbody")
	txt = html[start:end]
	rows += strings.Count(txt, "<tr")

	for i := 0; i < rows; i++ {
		rstart := strings.Index(txt, "<tr")
		rend := strings.Index(txt, "</tr>") + len("</tr>")
		if rstart == -1 {
			break
		}
		rtxt := txt[rstart:rend]
		ccnt := strings.Count(rtxt, "<td")
		rw := txt[rstart:rend]
		rwslice := []string{}
		// iterate thru cells, regex to find value within a cell
		for j := 0; j < ccnt; j++ {
			cstart := strings.Index(rw, "<td")
			cend := strings.Index(rw, "</td>") + len("</td>")
			rwtx := rw[cstart:cend]
			fmt.Println(rw[cstart:cend])
			re := regexp.MustCompile(`>([^<>]*)<`)

			// find all matches
			matches := re.FindAllStringSubmatch(rwtx, -1)
			// fmt.Println("Alert")
			// fmt.Println(filter2DBlanks(matches))
			matches = filter2DBlanks(matches)
			dat := ""
			for _, lv := range matches {
				if len(lv) > 0 {
					//fmt.Println("XD " + lv[1])
					dat = dat + " " + lv[1]
				}

			}
			rwslice = append(rwslice, dat)
			rw = rw[cend:]
		}
		fmt.Println(rwslice)
		table.Rows += 1
		table.Cols = maxInt(table.Cols, len(rwslice))
		table.Tab = append(table.Tab, rwslice)

		fmt.Println(table)

		// fmt.Println(txt[rstart:rend])
		txt = txt[rend:]
	}

	return table

}

// prarses tables from html, index doesnt currently work, just returns all tables
func ParseTables(html string, idx int) []Table {
	tables := []Table{}
	raw := []string{}
	curraw := html

	tc := TableCount(html)
	if tc <= 0 {
		return nil
	}

	for i := 0; i < tc; i++ {
		start := strings.Index(curraw, "<table")
		fmt.Println(start)

		end := strings.Index(curraw, "</table>") + len("</table>")
		raw = append(raw, curraw[start:end])

		curraw = curraw[end:]
	}
	// fmt.Println(raw[2])
	//newtab := generateTable(raw[2])
	for i := 0; i < len(raw); i++ {
		tables = append(tables, generateTable(raw[i]))
	}
	return tables
}
