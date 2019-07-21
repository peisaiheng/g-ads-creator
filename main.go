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

const (
	Campaign= "Campaign"
	StartDate = "Start Date"
	EndDate = "End Date"
	AdSchedule = "Ad Schedule"
	FlexibleReach = "Flexible Reach"
	AdGroup = "Ad Group"
	MaxCPC = "Max CPC"
	MaxCPM = "Max CPM"
	TargetCPA = "Target CPA"
	DisplayNetworkCustomBidType = "Display Network Custom Bid Type"
	TargetingOptimization = "Targeting optimization"
	ContentKeywords = "Content keywords"
	AdGroupType = "Ad Group Type"
	Keyword = "Keyword"
	CriterionType = "Criterion Type"
	FirstPageBid = "First page bid"
	TopOfPageBid = "Top of page bid"
	FirstPositionBid = "First position bid"
	QualityScore = "Quality score"
	LandingPageExperience = "Landing page experience"
	ExpectedCTR = "Expected CTR"
	AdRelevance = "Ad relevance"
	FinalURL = "Final URL"
	FinalMobileURL = "Final mobile URL"
	FeedName = "Feed Name"
	PlatformTargeting = "Platform Targeting"
	DevicePreference = "Device Preference"
	LinkText = "Link Text"
	DescriptionLine1 = "Description Line 1"
	DescriptionLine2 = "Description Line 2"
	Headline1 = "Headline 1"
	Headline2 = "Headline 2"
	Headline3 = "Headline 3"
	Path1 = "Path 1"
	Path2 = "Path 2"
)

// CSV export Headers
var rowHeader = map[string]int{
	Campaign: 0,
	StartDate: 1,
	EndDate: 2,
	AdSchedule: 3,
	FlexibleReach: 4,
	AdGroup: 5,
	MaxCPC: 6,
	MaxCPM: 7,
	TargetCPA: 8,
	DisplayNetworkCustomBidType: 9,
	TargetingOptimization: 10,
	ContentKeywords: 11,
	AdGroupType: 12,
	Keyword: 13,
	CriterionType: 14,
	FirstPageBid: 15,
	TopOfPageBid: 16,
	FirstPositionBid: 17,
	QualityScore: 18,
	LandingPageExperience: 19,
	ExpectedCTR: 20,
	AdRelevance: 21,
	FinalURL: 22,
	FinalMobileURL: 23,
	FeedName: 24,
	PlatformTargeting: 25,
	DevicePreference: 26,
	LinkText: 27,
	DescriptionLine1: 28,
	DescriptionLine2: 29,
	Headline1: 30,
	Headline2: 31,
	Headline3: 32,
	Path1: 33,
	Path2: 34,
}

var wOutput *csv.Writer

/*func addOutputHeader() error {

	// Create string of Array from rowHeader map
	exportHeaders := make([]string, 35)

	for key, value := range rowHeader {
		exportHeaders[value] = key
	}

	err := wOutput.Write(exportHeaders)
	return err
}*/

func reverseMap (mapStringInt map[string]int) map[int]string {

	var rowHeader map[int]string
	for key, value := range mapStringInt {
		rowHeader[value] = key
	}

	return rowHeader
}

func writeRow(rowMap map[int]string) error {

	rowArray := make([]string, 35, 35)
	for key, value := range rowMap {
		rowArray[key] = value
	}

	err := wOutput.Write(rowArray)
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

	outputFileName := os.Args[2]

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
		fmt.Println(`ERROR: Cannot create file: `, err)
	}

	defer csvOutputFile.Close()

	wOutput = csv.NewWriter(csvOutputFile)
	defer wOutput.Flush()

	// Using a flag instead of a range, I think it is more efficient.
	// We know it's a new file as we created it. If name exist, there will be an err
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
				if writeRow(reverseMap(rowHeader)) != nil {
					fmt.Println(`ERROR ADDING HEADERS: `, err)
				}
			}
			continue
		} else {
			// Do something with each au area row
			processOutput1(record, forSale)
		}

	}
}
