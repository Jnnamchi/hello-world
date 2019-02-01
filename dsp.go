package main

import (
	"encoding/json"
	"fmt"
	"hello-world/advertisers"
	"hello-world/blink"
	"log"
	"net"
	"net/http"
	"strconv"
	"strings"
)

const port = 8080

func main() {

	log.Printf("Starting up Demand Side Platform ...")

	ip := getOutboundIP()
	address := strings.Join([]string{ip.String(), strconv.Itoa(port)}, ":")

	log.Printf("Listening on: %v", address)

	confirmLEDsBlink()

	http.HandleFunc("/ixrtb", handleAdRequest)
	log.Fatal(http.ListenAndServe(address, nil))
}

// Get preferred outbound ip of this machine
func getOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
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

		blink.Blink(advertisers.AdvertiserColors[topBid.Advertiser])

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

func confirmLEDsBlink() {

	for i := 0; i < 5; i++ {
		blink.Blink("color1")
		blink.Blink("color2")
	}
	for i := 0; i < 3; i++ {
		blink.Blink("color1")
	}
}
