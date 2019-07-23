package main

import (
	"fmt"
	"strconv"
	"strings"
)

func shortenPath(p string) string {
	switch facetType {
	case Station:
		return strings.Replace(p, "station", "stn", -1)
	}
	return p
}

func processOutput(record []string, forSale bool) {

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
	suburb, endingUrl, titledArea := record[0], record[1], strings.Title(record[0])

	dashedArea := strings.Replace(record[0], " ", "-", -1)
	saleRentLower := strings.ToLower(saleRent)
	inverseSaleRentLower := strings.ToLower(inverseSaleRent)

	campaignName := country + ` | Traffic - PPC | ` + facetType + ` - ` + saleRent
	adGroupName := country + ` | ` + titledArea + ` | For ` + saleRent

	// Create AdGroup Details Row
	detailsMap := map[int]string{
		col[Campaign]:                    campaignName,
		col[FlexibleReach]:               "Audiences",
		col[AdGroup]:                     adGroupName,
		col[MaxCPC]:                      "0.9",
		col[MaxCPM]:                      "0.01",
		col[TargetCPA]:                   "0",
		col[DisplayNetworkCustomBidType]: "None",
		col[TargetingOptimization]:       "Disabled",
		col[ContentKeywords]:             "Disabled",
		col[AdGroupType]:                 "Default",
	}

	writeRow(detailsMap)
	fmt.Println(titledArea + ` - Details row writen to CSV.`)

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
			col[Campaign]:              campaignName,
			col[AdGroup]:               adGroupName,
			col[Keyword]:               keywords[index],
			col[CriterionType]:         "Exact",
			col[FirstPageBid]:          "0",
			col[TopOfPageBid]:          "0",
			col[FirstPositionBid]:      "0",
			col[LandingPageExperience]: "-",
			col[ExpectedCTR]:           "-",
			col[AdRelevance]:           "-",
		}

		writeRow(keywordsMap)
		fmt.Println(titledArea + ` - Keyword ` + strconv.Itoa(index) + ` writen to CSV.`)
	}

	// PREPARING SITELINKS ROWS

	//Preposition for different facet type
	var preposition, condoApt string

	if facetType == Station {
		preposition = "around"
	} else {
		preposition = "in"
	}

	switch country {
	case Australia:
		condoApt = "apartments"
	case Singapore:
		condoApt = "condominiums"
	}

	saleRentAroundUrl := sohoSearch + `for-` + saleRentLower + `-properties-` + preposition + `-` + endingUrl
	inverseSaleRentAroundUrl := sohoSearch + `for-` + inverseSaleRentLower + `-properties-` + preposition + `-` + endingUrl
	ForNewProjectAroundUrl := sohoSearch + `new-properties-` + preposition + `-` + endingUrl
	HouseAroundUrl := sohoSearch + `for-` + saleRentLower + `-houses-` + preposition + `-` + endingUrl
	ApartmentAroundUrl := sohoSearch + `for-` + saleRentLower + `-` + condoApt + `-` + preposition + `-` + endingUrl

	siteLinksSlice := []sitelink{
		{
			finalUrl:         ForNewProjectAroundUrl,
			linkText:         `New Homes - ` + titledArea,
			descriptionLine1: `Find New Homes For ` + saleRent,
		},
		{
			finalUrl:         HouseAroundUrl,
			linkText:         `Houses - ` + titledArea,
			descriptionLine1: `Find Houses For ` + saleRent,
		},
		{
			finalUrl:         ApartmentAroundUrl,
			linkText:         strings.TrimSuffix(strings.Title(condoApt), "miniums") + ` - ` + titledArea,
			descriptionLine1: `Find ` + strings.Title(condoApt) + ` For ` + saleRent,
		},
		{
			finalUrl:         inverseSaleRentAroundUrl,
			linkText:         `For ` + inverseSaleRent + ` - ` + titledArea,
			descriptionLine1: `Find Properties For ` + inverseSaleRent,
		},
	}

	for index, sitelink := range siteLinksSlice {
		// Create AdGroup SiteLinks Rows
		siteLinksMap := map[int]string{
			col[Campaign]:          campaignName,
			col[StartDate]:         "[]",
			col[EndDate]:           "[]",
			col[AdSchedule]:        "[]",
			col[AdGroup]:           adGroupName,
			col[FinalURL]:          sitelink.finalUrl,
			col[FeedName]:          "Main sitelink feed",
			col[PlatformTargeting]: "All",
			col[DevicePreference]:  "All",
			col[LinkText]:          sitelink.linkText,
			col[DescriptionLine1]:  sitelink.descriptionLine1,
			col[DescriptionLine2]:  `Around ` + titledArea,
		}

		writeRow(siteLinksMap)
		fmt.Println(titledArea + ` - SiteLink ` + strconv.Itoa(index) + ` writen to CSV.`)
	}

	// PREPARING ADS ROWS

	adDescriptionLine1 := `Search properties for ` + saleRent + ` in ` + titledArea + `. Find your dream home on Soho.`
	adDescriptionLine2 := "The fastest growing property network used by more than 45,000 property seekers."

	if len(dashedArea) > 15 {
		dashedArea = shortenPath(dashedArea)
	}

	adsSlice := []ad{
		{
			finalUrl:         saleRentAroundUrl,
			descriptionLine1: adDescriptionLine1,
			descriptionLine2: adDescriptionLine2,
			headline1:        `Find Properties for ` + saleRent,
			headline2:        `Around ` + titledArea,
			headline3:        "Soho",
			path1:            "search",
			path2:            dashedArea,
		},
		{
			finalUrl:         saleRentAroundUrl,
			descriptionLine1: adDescriptionLine1,
			descriptionLine2: adDescriptionLine2,
			headline1:        buyRent + ` Properties`,
			headline2:        `Around ` + titledArea,
			headline3:        "Soho",
			path1:            "search",
			path2:            dashedArea,
		},
		{
			finalUrl:         saleRentAroundUrl,
			descriptionLine1: adDescriptionLine1,
			descriptionLine2: adDescriptionLine2,
			headline1:        titledArea,
			headline2:        `Properties for ` + saleRent,
			headline3:        "Soho",
			path1:            "search",
			path2:            dashedArea,
		},
	}

	for index, ad := range adsSlice {
		// Create AdGroup Ads Rows
		adsMap := map[int]string{
			col[Campaign]:         campaignName,
			col[AdGroup]:          adGroupName,
			col[FinalURL]:         ad.finalUrl,
			col[DescriptionLine1]: ad.descriptionLine1,
			col[DescriptionLine2]: ad.descriptionLine2,
			col[Headline1]:        ad.headline1,
			col[Headline2]:        ad.headline2,
			col[Headline3]:        ad.headline3,
			col[Path1]:            ad.path1,
			col[Path2]:            ad.path2,
		}

		writeRow(adsMap)
		fmt.Println(titledArea + ` - Ad ` + strconv.Itoa(index) + ` writen to CSV.`)
	}
}
