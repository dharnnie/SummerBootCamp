package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {
	// 1. Get csv file
	// 2. open file
	// 3. read file
	// 4. parse file
	// 5. display file

	// 2. open file
	src, err := os.Open("../resources/table.csv")
	if err != nil {
		panic(err)
	}
	defer src.Close()
	// src is a *File
	// func (f *File) Read(b []byte) (n int, err error)
	// func NewReader(r io.Reader) *Reader

	// 3. read file
	rdr := csv.NewReader(src)
	// rdr is a *Reader
	// func (r *Reader) ReadAll() (records [][]string, err error)

	data, err := rdr.ReadAll()
	if err != nil {
		log.Fatalln("couldn't readall", err.Error())
	}

	// 4. parse file
	fmt.Println(`
			<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title></title>
</head>
<body>
<table>
    <tr>
        <th>DATE</th>
        <th>OPEN</th>
        <th>HIGH</th>
        <th>LOW</th>
        <th>CLOSE</th>
        <th>VOLUME</th>
        <th>ADJ CLOSE</th>
    </tr>`)

	for _, value := range data {
		fmt.Println(`<tr>`)
		for _, v := range value {
			fmt.Println(`<td>` + v + `</td>`)
		}
		fmt.Println(`<tr>`)
	}
	fmt.Println(`
	</table>
</body>
</html>
`)
}

/*
Grab historical financial data from Yahoo
http://finance.yahoo.com/q/hp?s=GOOG+Historical+Prices
as a csv file, read that file and dumps the contents as an HTML table.
EXTRA CHALLENGE: draw a line chart of the data.
*/
