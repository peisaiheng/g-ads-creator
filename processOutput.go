package main

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	ForRentAround       = "https://search.sohoapp.com/marketplace/for-rent-around-"
	ForSaleAround       = "https://search.sohoapp.com/marketplace/for-sale-around-"
	ForNewProjectAround = "https://search.sohoapp.com/marketplace/for-new_project-around-"
	HouseForSale        = "https://search.sohoapp.com/marketplace/house-for-sale-around-"
	ApartmentForSale    = "https://search.sohoapp.com/marketplace/apartment-for-sale-around-"
)

type sitelink struct {
	finalUrl,
	linkText,
	descriptionLine1 string
}

type ad struct {
	finalUrl,
	descriptionLine1,
	descriptionLine2,
	headline1,
	headline2,
	headline3,
	path1,
	path2 string
}

func writeRow(rowMap map[int]string) {
	rowArray := make([]string, 35, 35)
	for key, value := range rowMap {
		rowArray[key] = value
	}
	wOutput.Write(rowArray)
}

func processOutput(record []string, forSale bool) {

	var saleOrRent string
	if forSale {
		saleOrRent = "Sale"
	} else {
		saleOrRent = "Rent"
	}
	// Commonly used values
	suburb, endingUrl, titledSuburb := record[0], record[1], strings.Title(record[0])
	dashedSuburb := strings.Replace(record[0]," ","-",-1)
	saleOrRentToLower := strings.ToLower(saleOrRent)
	campaignName := `AU | Traffic - PPC | Areas - `+saleOrRent
	adGroupName := `AU | ` + titledSuburb + ` | For `+saleOrRent

	// Create AdGroup Details Row
	detailsMap := map[int]string{
		0:  campaignName,
		4:  "Audiences",
		5:  adGroupName,
		6:  "0.9",
		7:  "0.01",
		8:  "0",
		9:  "None",
		10: "Disabled",
		11: "Disabled",
		12: "Default",
	}

	writeRow(detailsMap)
	fmt.Println(titledSuburb+` - Details row writen to CSV.`)

	// PREPARING KEYWORD ROWS

	keywords := []string{
		` for `+saleOrRentToLower,
		`housing for `+saleOrRentToLower+` `,
		`apartment for `+saleOrRentToLower+` `,
		`home for `+saleOrRentToLower+` `,
		`private property `+saleOrRentToLower+` `,
		`property for `+saleOrRentToLower+` `,
		`house for `+saleOrRentToLower+` `,
		`real estate for `+saleOrRentToLower+` `,
	}

	for index, keyword := range keywords {

		// Adding the suburb for each keyword phrase
		if index == 0 {
			keywords[index] = suburb + keywords[index]
		} else {
			keywords[index] = keyword + suburb
		}

		// Create AdGroup Keywords Rows
		keywordsMap := map[int]string{
			0:  campaignName,
			5:  adGroupName,
			13: keywords[index],
			14: "Exact",
			15: "0",
			16: "0",
			17: "0",
			19: "-",
			20: "-",
			21: "-",
		}

		writeRow(keywordsMap)
		fmt.Println(titledSuburb+` - Keyword ` + strconv.Itoa(index) + ` writen to CSV.`)
	}

	// PREPARING SITELINKS ROWS

	siteLinksSlice := []sitelink{
		{finalUrl: ForNewProjectAround + endingUrl, linkText: `New Homes - ` + titledSuburb, descriptionLine1: "Find New Homes For Sale"},
		{finalUrl: ForSaleAround + endingUrl, linkText: titledSuburb + ` Houses`, descriptionLine1: "Find Houses For Sale"},
		{finalUrl: ApartmentForSale + endingUrl, linkText: titledSuburb + ` Apartments`, descriptionLine1: "Find Apartments For Sale"},
		{finalUrl: ForRentAround + endingUrl, linkText: `For Rent - ` + titledSuburb, descriptionLine1: "Find Properties For Rent"},
	}

	for index, siteLinkStruct := range siteLinksSlice {
		// Create AdGroup SiteLinks Rows
		siteLinksMap := map[int]string{
			0:  campaignName,
			1:  "[]",
			2:  "[]",
			3:  "[]",
			5:  adGroupName,
			22: siteLinkStruct.finalUrl,
			24: "Main sitelink feed",
			25: "All",
			26: "All",
			27: siteLinkStruct.linkText,
			28: siteLinkStruct.descriptionLine1,
			29: `in ` + titledSuburb,
		}

		writeRow(siteLinksMap)
		fmt.Println(titledSuburb+` - SiteLink ` + strconv.Itoa(index) + ` writen to CSV.`)
	}

	// PREPARING ADS ROWS

	adDescriptionLine1 := `Search properties for sale around `+titledSuburb+`. Find your dream home on Soho.`
	adDescriptionLine2 := "The fastest growing property network used by more than 18,000 property seekers."

	adsSlice := []ad{
		{finalUrl: ForSaleAround+endingUrl, descriptionLine1: adDescriptionLine1, descriptionLine2: adDescriptionLine2, headline1: titledSuburb+` Properties`, headline2:`For `+saleOrRent+` | Soho`, path1:"marketplace", path2:dashedSuburb},
		{finalUrl: ForSaleAround+endingUrl, descriptionLine1: adDescriptionLine1, descriptionLine2: adDescriptionLine2, headline1: `Homes For Sale `+titledSuburb, headline2: "Find Properties For Sale", path1:"marketplace", path2:dashedSuburb},
		{finalUrl: ForSaleAround+endingUrl, descriptionLine1: adDescriptionLine1, descriptionLine2: adDescriptionLine2, headline1: "Find Properties For Sale", headline2: titledSuburb+`, Australia`, path1:"marketplace", path2:dashedSuburb},


	}

	for index, ad := range adsSlice {
		// Create AdGroup Ads Rows
		siteLinksMap := map[int]string{
			0:  campaignName,
			5:  adGroupName,
			22: ad.finalUrl,
			28: ad.descriptionLine1,
			29: ad.descriptionLine2,
			30: ad.headline1,
			31: ad.headline2,
			33: ad.path1,
			34: ad.path2,
		}

		writeRow(siteLinksMap)
		fmt.Println(titledSuburb+` - Ad ` + strconv.Itoa(index) + ` writen to CSV.`)
	}
}
