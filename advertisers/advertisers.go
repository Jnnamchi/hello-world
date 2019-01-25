package advertisers

import (
	"fmt"
)

// SubmitBids checks the conditions for all advertisers and returns the highest bid
func SubmitBids(adRequest map[string]string) BidObj {

	fmt.Println("Processing request:", adRequest)

	topBid := BidObj{}

	for _, advertiser := range advertisers {
		bid := getTopBidForAdvertiser(adRequest, advertiser)

		if bid.BidPrice > topBid.BidPrice {
			topBid.Advertiser = bid.Advertiser
			topBid.BidPrice = bid.BidPrice
			topBid.AdURL = bid.AdURL
			topBid.AdDescription = bid.AdDescription
		}
	}

	return topBid
}

// Getting the bids
type BidObj struct {
	Advertiser    string
	BidPrice      int
	AdDescription string
	AdURL         string
}

func getTopBidForAdvertiser(adRequest map[string]string, seller advertiser) BidObj {

	topBid := BidObj{
		Advertiser: seller.name,
	}

	for _, bidConfig := range seller.bidConfs {

		bidConfig, bid := getBid(adRequest, bidConfig)

		if bid > topBid.BidPrice {
			topBid.BidPrice = bid
			topBid.AdURL = bidConfig.adURL
			topBid.AdDescription = bidConfig.adDescription
		}
	}

	return topBid
}

func getBid(adRequest map[string]string, bidConfiguarion bidConf) (bidConf, int) {

	bid := 0

	if bidConfiguarion.shape != adRequest["s"] {
		return bidConfiguarion, 0
	}

	for _, rule := range bidConfiguarion.rules {
		if rule.condition(adRequest[rule.param]) {
			bid += rule.impact
		}
	}

	return bidConfiguarion, bid
}

// ADVERTISERS CONFIGURATION HERE

// Core rule unit for determining bid
type rule struct {
	param     string
	condition func(string) bool
	impact    int
}

// Core Ad Unit
type bidConf struct {
	adURL         string
	shape         string
	adDescription string
	rules         []rule
}

// Core Advertise Unit
type advertiser struct {
	name     string
	bidConfs []bidConf
}

// ADVERTISERS DATA HERE

var advertisers = []advertiser{
	cookingAdvertiser,
	moviesAdvertiser,
	sportsAdvertiser,
	workingAdvertiser,
}
