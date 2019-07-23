package main

import (
	"fmt"
	"bufio"
	"os"
)

const (
	Council = "Councils"
	Suburb = "Suburbs"
	District = "Districts"
	PlanningArea = "Planning Areas"
	Neighborhood = "Neighborhoods"
	Station = "Stations"
)

const (
	Singapore = "SG"
	Australia = "AU"
)

func forSaleOrRent () bool {
	fmt.Print("Is this list for Sale or Rent?\n\n" + "1) Sale\n" + "2) Rent\n\n" + "Select with 1 or 2:")
	scanner := bufio.NewScanner(os.Stdin)
	counter := 0
	for scanner.Scan() {
		if counter >= 2 {
			fmt.Println(`Please execute the programme again`)
			break
		}
		if scanner.Text() == "1" {
			return true
			break
		} else if scanner.Text() == "2" {
			return false
			break
		}
		if scanner.Text() != "1" || scanner.Text() != "2" {
			counter++
			fmt.Print(`Please type only 1 or 2:`)
		}
	}
	return true
}

func getFacetType () string {
	fmt.Print("Which facet type?\n\n" + "1) Council(AU)\n" + "2) Suburb(AU)\n" + "3) District(SG)\n" + "4) Planning area(SG)\n" + "5) Neighborhood(SG)\n" + "6) Station\n\n" + "Select one:")
	scanner := bufio.NewScanner(os.Stdin)
	counter := 0
	for scanner.Scan() {
		if counter >= 2 {
			fmt.Println(`Please execute the programme again`)
			break
		}
		if scanner.Text() == "1" {
			return Council
			break
		} else if scanner.Text() == "2" {
			return Suburb
			break
		} else if scanner.Text() == "3" {
			return District
			break
		} else if scanner.Text() == "4" {
			return PlanningArea
			break
		} else if scanner.Text() == "5" {
			return Neighborhood
			break
		} else if scanner.Text() == "5" {
			return Station
			break
		}
		if scanner.Text() != "1" || scanner.Text() != "2" {
			counter++
			fmt.Print(`Please type only 1 or 2:`)
		}
	}
	return ""
}

func getCountry () string {
	fmt.Print("Which country?\n\n" + "1) Australia\n" + "2) Singapore\n\n" + "Select with 1 or 2:")
	scanner := bufio.NewScanner(os.Stdin)
	counter := 0
	for scanner.Scan() {
		if counter >= 2 {
			fmt.Println(`Please execute the programme again`)
			break
		}
		if scanner.Text() == "1" {
			return Australia
			break
		} else if scanner.Text() == "2" {
			return Singapore
			break
		}
		if scanner.Text() != "1" || scanner.Text() != "2" {
			counter++
			fmt.Print(`Please type only 1 or 2:`)
		}
	}
	return ""
}