package main

import "fmt"

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

const (
	DisplayNetworkCustomBidType = "Display Network Custom Bid Type"
)

func indexOf (header string) (int) {

}

func main() {
	fmt.Println(indexOf(DisplayNetworkCustomBidType))
}
