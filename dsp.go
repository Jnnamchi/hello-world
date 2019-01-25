package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"playground/advertisers"
	// "playground/advertisers"
)

func main() {

	fmt.Println("Running auction")

	http.HandleFunc("/ixrtb", handleAdRequest)
	log.Fatal(http.ListenAndServe("10.65.106.14:8080", nil))

	// topBid := advertisers.SubmitBids(buildBidRequest())

	// fmt.Println("Final bid:", topBid)
}

func handleAdRequest(w http.ResponseWriter, r *http.Request) {

	log.Printf("handler called")

	switch r.Method {
	case "POST":
		// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}

		// fmt.Fprintf(w, "Post from website! r.PostFrom = %v\n", r.PostForm)
		bidRequest := make(map[string]string)
		bidRequest["l"] = r.FormValue("l")
		bidRequest["s"] = r.FormValue("s")
		bidRequest["d"] = r.FormValue("d")
		bidRequest["a"] = r.FormValue("a")

		fmt.Println("Bid built:", bidRequest)
		topBid := advertisers.SubmitBids(bidRequest)

		// fmt.Println("Top bid:", topBid)
		// fmt.Println("Top bid: %v", []advertisers.BidObj{topBid})
		handleSuccess(&w, []advertisers.BidObj{topBid})
	default:
		fmt.Fprintf(w, "Sorry, only POST methods are supported.")
	}
}

func handleSuccess(w *http.ResponseWriter, result interface{}) {
	writer := *w

	marshalled, err := json.Marshal(result)

	if err != nil {
		fmt.Println("Error")
		return
	}

	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(200)
	writer.Write(marshalled)
}

func buildBidRequest() map[string]string {

	bidRequest := make(map[string]string)

	bidRequest["s"] = "tall"
	bidRequest["l"] = "movies"
	bidRequest["d"] = "sports"
	bidRequest["a"] = "22"

	return bidRequest
}
