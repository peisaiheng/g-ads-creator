package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)

// CSV import Headers
var importHeaders = []string{
	"suburb",
	"url path",
}

// CSV export Headers
var exportHeaders = []string{
	"Campaign",
	"Start Date",
	"End Date",
	"Ad Schedule",
	"Flexible Reach",
	"Ad Group",
	"Max CPC",
	"Max CPM",
	"Target CPA",
	"Display Network Custom Bid Type",
	"Targeting optimization",
	"Content keywords",
	"Ad Group Type",
	"Keyword",
	"Criterion Type",
	"First page bid",
	"Top of page bid",
	"First position bid",
	"Quality score",
	"Landing page experience",
	"Expected CTR",
	"Ad relevance",
	"Final URL",
	"Final mobile URL",
	"Feed Name",
	"Platform Targeting",
	"Device Preference",
	"Link Text",
	"Description Line 1",
	"Description Line 2",
	"Headline 1",
	"Headline 2",
	"Headline 3",
	"Path 1",
	"Path 2",
}

var wOutput *csv.Writer

func addOutputHeader() error {
	err := wOutput.Write(exportHeaders)
	return err
}

func checkFirstRow(record []string) bool {

	// Flag for checking of cols
	rowCorrect := true

	// Range cols to check headers
	for index, colValue := range record {
		if colValue == importHeaders[index] {
			continue
		} else {
			fmt.Println(`ERROR: CSV not in required format - "` + importHeaders[index] + `" not found or incorrectly named.`)
			rowCorrect = false
			break
		}
	}
	return rowCorrect
}

func main() {

	if len(os.Args) < 3 {
		fmt.Println(`Key in the programme name followed by the path to the imported csv and the name you want to export csv as.`)
		return
	}

	outputFileName := os.Args[2]

	// Sale or Rent
	forSale := true

	fmt.Print("Is this list for Sale or Rent?\n\n" + "1) Sale\n" + "2) Rent\n\n" + "Select with 1 or 2:")
	scanner := bufio.NewScanner(os.Stdin)
	counter := 0
	for scanner.Scan() {
		if counter >= 2 {
			fmt.Println(`Please execute the programme again`)
			break
		}
		if scanner.Text() == "1" {
			forSale = true
			break
		} else if scanner.Text() == "2" {
			forSale = false
			break
		}
		if scanner.Text() != "1" || scanner.Text() != "2" {
			counter++
			fmt.Print(`Please type only 1 or 2:`)
		}
	}

	// Opens CSV
	fileName := os.Args[1]
	csvInputFile, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer csvInputFile.Close()
	rInput := csv.NewReader(bufio.NewReader(csvInputFile))

	fmt.Println(`Opening File name: ` + fileName)

	// Create and open output file before ranging input file.
	// This way we won't be checking for file exist and opening and writing to file for each input row.
	// Should be more efficient.
	csvOutputFile, err := os.Create(`./` + outputFileName)
	if err != nil {

		// Error while creating file
		fmt.Println(`ERROR: Cannot create file.`)
		fmt.Println(err)
	}
	defer csvOutputFile.Close()
	wOutput = csv.NewWriter(csvOutputFile)
	defer wOutput.Flush()

	// Using a flag instead of a range, I think it is more efficient.
	isFirstRow := true

	// Range CSV
	for {
		record, err := rInput.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
		}

		// Check for CSV headers, reject if columns length is incorrect.
		if len(record) != 2 {
			fmt.Println(`ERROR: CSV not in required format - ` + strconv.Itoa(len(importHeaders)) + ` cols required.`)
			break
		}

		// Check for CSV headers, reject if columns header are incorrect.
		if isFirstRow {

			fmt.Println(record)
			// if input file headers are correct, add headers to output file
			if checkFirstRow(record) {

				// Change flag to signal first row is checked.
				isFirstRow = false

				// If first row of input file is successfully checked.
				// Add headers for output file.
				if addOutputHeader() != nil {
					fmt.Println(`ERROR ADDING HEADERS`)
					fmt.Println(err)
				}
			}
			continue
		} else {
			// Do something with each au area row
			processOutput1(record, forSale)
		}

	}
}
