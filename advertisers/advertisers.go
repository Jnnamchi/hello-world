package advertisers

import (
	"log"
)

var (
	// ShapeParameter is the IXRTB Paramter for shape of the available ad-placement
	ShapeParameter = "s"

	// LikesParameter is the IXRTB Paramter for the ad-viewer's top area of interest
	LikesParameter = "l"

	// DislikesParameter is the IXRTB Paramter for the ad-viewer's top area of disinterest
	DislikesParameter = "d"

	// AgeParameter is the IXRTB Paramter for the ad-viewer's age
	AgeParameter = "a"
)

// List advertisers enabled for bidding
var advertisers = []advertiser{
	cookingAdvertiser,
	moviesAdvertiser,
	sportsAdvertiser,
	workingAdvertiser,
}

// Core Advertiser Unit
type advertiser struct {
	name     string
	bidConfs []bidConf
}

// Core Ad Unit
type bidConf struct {
	adURL         string
	shape         string
	adDescription string
	rules         []rule
}

// Core rule unit for determining bid value
type rule struct {
	param     string
	condition func(string) bool
	impact    int
}

// BidObj contains the information of the top bid and returns it to the DSP
type BidObj struct {
	Advertiser    string
	BidPrice      int
	AdDescription string
	AdURL         string
}

// GetTopBidForDSP gets the top bid for all advertisers and returns the highest bid
func GetTopBidForDSP(adRequest map[string]string) BidObj {

	log.Printf("Processing request: %v", adRequest)

	topBid := BidObj{}

	for _, advertiser := range advertisers {

		bid := getTopBidForAdvertiser(adRequest, advertiser)

		if bid.BidPrice > topBid.BidPrice {

			topBid = bid
		}
	}

	return topBid
}

// getTopBidForAdvertiser gets the top bid for an advertiser
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

// getBid gets one of the bids of an advertiser, an advertiser can have multiple bidding configurations
func getBid(adRequest map[string]string, bidConfiguarion bidConf) (bidConf, int) {

	bid := 0

	if bidConfiguarion.shape != adRequest[ShapeParameter] {
		return bidConfiguarion, 0
	}

	for _, rule := range bidConfiguarion.rules {
		if rule.condition(adRequest[rule.param]) {
			bid += rule.impact
		}
	}

	return bidConfiguarion, bid
}
