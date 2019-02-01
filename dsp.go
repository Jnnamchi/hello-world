package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/jnnamchi/hello-world/advertisers"
)

func main() {

	log.Printf("Starting up Demand Side Platform ...")

	http.HandleFunc("/ixrtb", handleAdRequest)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleAdRequest(w http.ResponseWriter, r *http.Request) {

	log.Printf("Received new bid request")

	switch r.Method {
	case "POST":
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}

		bidRequest := make(map[string]string)
		bidRequest[advertisers.ShapeParameter] = r.FormValue(advertisers.ShapeParameter)
		bidRequest[advertisers.LikesParameter] = r.FormValue(advertisers.LikesParameter)
		bidRequest[advertisers.DislikesParameter] = r.FormValue(advertisers.DislikesParameter)
		bidRequest[advertisers.AgeParameter] = r.FormValue(advertisers.AgeParameter)

		topBid := advertisers.GetTopBidForDSP(bidRequest)

		returnBidResponses(&w, topBid)
	default:
		fmt.Fprintf(w, "Sorry, only POST methods are supported.")
	}
}

func returnBidResponses(w *http.ResponseWriter, result advertisers.BidObj) {
	writer := *w

	marshalled, err := json.Marshal(result)

	if err != nil {
		log.Printf("Error unmarshalling bid request")
		writer.WriteHeader(403)
	}

	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(200)
	writer.Write(marshalled)
}
