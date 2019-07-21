package main

import (
	"fmt"
	"strconv"
	"strings"
)

//func writeRow(rowMap map[int]string) {
//	rowArray := make([]string, 35, 35)
//	for key, value := range rowMap {
//		rowArray[key] = value
//	}
//	wOutput.Write(rowArray)
//}

func processOutput(record []string, forSale bool) {

	var saleRent, inverseSaleRent string
	if forSale {
		saleRent = "Sale"
		inverseSaleRent = "Rent"
	} else {
		saleRent = "Rent"
		inverseSaleRent = "Sale"
	}
	// Commonly used values
	suburb, endingUrl, titledSuburb := record[0], record[1], strings.Title(record[0])
	dashedSuburb := strings.Replace(record[0], " ", "-", -1)
	saleRentLower := strings.ToLower(saleRent)

	campaignName := `AU | Traffic - PPC | Areas - ` + saleRent
	adGroupName := `AU | ` + titledSuburb + ` | For ` + saleRent

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
	fmt.Println(titledSuburb + ` - Details row writen to CSV.`)

	// PREPARING KEYWORD ROWS

	keywords := []string{
		` for ` + saleRentLower,
		` apartment for ` + saleRentLower,
		` home for ` + saleRentLower,
		` property for ` + saleRentLower,
		` house for ` + saleRentLower,
		` real estate for ` + saleRentLower,
	}

	for index, keyword := range keywords {

		// Adding the suburb for each keyword phrase
		keywords[index] = suburb + keyword

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
		fmt.Println(titledSuburb + ` - Keyword ` + strconv.Itoa(index) + ` writen to CSV.`)
	}

	// PREPARING SITELINKS ROWS

	saleRentAroundUrl := sohoMarketplacePath + `for-` + saleRentLower + `-around-` + endingUrl
	inverseSaleRentAroundUrl := sohoMarketplacePath + `for-` + strings.ToLower(inverseSaleRent) + `-around-` + endingUrl
	ForNewProjectAroundUrl := sohoMarketplacePath + `for-new_project-around-` + endingUrl
	HouseAroundUrl := sohoMarketplacePath + `house-for-` + saleRentLower + `-around-` + endingUrl
	ApartmentAroundUrl := sohoMarketplacePath + `apartment-for-` + saleRentLower + `-around-` + endingUrl

	siteLinksSlice := []sitelink{
		{finalUrl: ForNewProjectAroundUrl, linkText: `New Homes - ` + titledSuburb, descriptionLine1: `Find New Homes For ` + saleRent},
		{finalUrl: HouseAroundUrl, linkText: titledSuburb + ` Houses`, descriptionLine1: `Find Houses For ` + saleRent},
		{finalUrl: ApartmentAroundUrl, linkText: titledSuburb + ` Apartments`, descriptionLine1: `Find Apartments For ` + saleRent},
		{finalUrl: inverseSaleRentAroundUrl, linkText: `For ` + inverseSaleRent + ` - ` + titledSuburb, descriptionLine1: `Find Properties For ` + inverseSaleRent},
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
		fmt.Println(titledSuburb + ` - SiteLink ` + strconv.Itoa(index) + ` writen to CSV.`)
	}

	// PREPARING ADS ROWS

	adDescriptionLine1 := `Search properties for ` + saleRent + ` around ` + titledSuburb + `. Find your dream home on Soho.`
	adDescriptionLine2 := "The fastest growing property network used by more than 18,000 property seekers."

	adsSlice := []ad{
		{finalUrl: saleRentAroundUrl, descriptionLine1: adDescriptionLine1, descriptionLine2: adDescriptionLine2, headline1: titledSuburb + ` Properties`, headline2: `For ` + saleRent + ` | Soho`, path1: "marketplace", path2: dashedSuburb},
		{finalUrl: saleRentAroundUrl, descriptionLine1: adDescriptionLine1, descriptionLine2: adDescriptionLine2, headline1: `Homes For ` + saleRent + ` ` + titledSuburb, headline2: `Find Properties For ` + saleRent, path1: "marketplace", path2: dashedSuburb},
		{finalUrl: saleRentAroundUrl, descriptionLine1: adDescriptionLine1, descriptionLine2: adDescriptionLine2, headline1: `Find Properties For ` + saleRent, headline2: titledSuburb + `, Australia`, path1: "marketplace", path2: dashedSuburb},
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
		fmt.Println(titledSuburb + ` - Ad ` + strconv.Itoa(index) + ` writen to CSV.`)
	}
}
