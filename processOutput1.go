package main

import (
	"fmt"
	"strconv"
	"strings"
)

func shortenPath(p string) string {
	return strings.Replace(p, "station", "stn", -1)
}

func processOutput1(record []string, forSale bool) {

	var saleRent, inverseSaleRent, buyRent string
	if forSale {
		saleRent = "Sale"
		inverseSaleRent = "Rent"
		buyRent = "Buy"
	} else {
		saleRent = "Rent"
		inverseSaleRent = "Sale"
		buyRent = "Rent"
	}
	// Commonly used values
	suburb, endingUrl, titledSuburb := record[0], record[1], strings.Title(record[0])

	dashedSuburb := strings.Replace(record[0], " ", "-", -1)
	saleRentLower := strings.ToLower(saleRent)

	campaignName := `AU | Traffic - PPC | Stations - ` + saleRent
	adGroupName := `AU | ` + titledSuburb + ` | For ` + saleRent

	// Create AdGroup Details Row
	detailsMap := map[int]string{
		rowHeader[Campaign]:                    campaignName,
		rowHeader[FlexibleReach]:               "Audiences",
		rowHeader[AdGroup]:                     adGroupName,
		rowHeader[MaxCPC]:                      "0.9",
		rowHeader[MaxCPM]:                      "0.01",
		rowHeader[TargetCPA]:                   "0",
		rowHeader[DisplayNetworkCustomBidType]: "None",
		rowHeader[TargetingOptimization]:       "Disabled",
		rowHeader[ContentKeywords]:             "Disabled",
		rowHeader[AdGroupType]:                 "Default",
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

	saleRentAroundUrl := sohoSearch + `for-` + saleRentLower + `-around-` + endingUrl
	inverseSaleRentAroundUrl := sohoSearch + `for-` + strings.ToLower(inverseSaleRent) + `-around-` + endingUrl
	ForNewProjectAroundUrl := sohoSearch + `for-new_project-around-` + endingUrl
	HouseAroundUrl := sohoSearch + `house-for-` + saleRentLower + `-around-` + endingUrl
	ApartmentAroundUrl := sohoSearch + `for-` + saleRentLower + `-around-` + endingUrl

	siteLinksSlice := []sitelink{
		{
			finalUrl:         ForNewProjectAroundUrl,
			linkText:         `New Homes - ` + titledSuburb,
			descriptionLine1: `Find New Homes For ` + saleRent,
		},
		{
			finalUrl:         HouseAroundUrl,
			linkText:         `Houses - ` + titledSuburb,
			descriptionLine1: `Find Houses For ` + saleRent,
		},
		{
			finalUrl:         ApartmentAroundUrl,
			linkText:         `Apartments - ` + titledSuburb,
			descriptionLine1: `Find Apartments For ` + saleRent,
		},
		{
			finalUrl:         inverseSaleRentAroundUrl,
			linkText:         `For ` + inverseSaleRent + ` - ` + titledSuburb,
			descriptionLine1: `Find Properties For ` + inverseSaleRent,
		},
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
			29: `Around ` + titledSuburb,
		}

		writeRow(siteLinksMap)
		fmt.Println(titledSuburb + ` - SiteLink ` + strconv.Itoa(index) + ` writen to CSV.`)
	}

	// PREPARING ADS ROWS

	adDescriptionLine1 := `Search properties for ` + saleRent + ` in ` + titledSuburb + `. Find your dream home on Soho.`
	adDescriptionLine2 := "The fastest growing property network used by more than 39,000 property seekers."

	if len(dashedSuburb) > 15 {
		dashedSuburb = shortenPath(dashedSuburb)
	}

	adsSlice := []ad{
		{
			finalUrl:         saleRentAroundUrl,
			descriptionLine1: adDescriptionLine1,
			descriptionLine2: adDescriptionLine2,
			headline1:        `Find Properties for ` + saleRent,
			headline2:        `Around ` + titledSuburb,
			headline3:        "Soho",
			path1:            "marketplace",
			path2:            dashedSuburb,
		},
		{
			finalUrl:         saleRentAroundUrl,
			descriptionLine1: adDescriptionLine1,
			descriptionLine2: adDescriptionLine2,
			headline1:        buyRent + ` Properties`,
			headline2:        `Around ` + titledSuburb,
			headline3:        "Soho",
			path1:            "marketplace",
			path2:            dashedSuburb,
		},
		{
			finalUrl:         saleRentAroundUrl,
			descriptionLine1: adDescriptionLine1,
			descriptionLine2: adDescriptionLine2,
			headline1:        titledSuburb,
			headline2:        `Properties for ` + saleRent,
			headline3:        "Soho",
			path1:            "marketplace",
			path2:            dashedSuburb,
		},
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
			32: ad.headline3,
			33: ad.path1,
			34: ad.path2,
		}

		writeRow(siteLinksMap)
		fmt.Println(titledSuburb + ` - Ad ` + strconv.Itoa(index) + ` writen to CSV.`)
	}
}
