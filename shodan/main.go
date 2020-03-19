package main

import (
	"fmt"
	"log"
	"os"

	shodan "."
)

func main() {

	apiKey := os.Getenv("SHODAN_API_KEY")
	s := shodan.New(apiKey)
	info, err := s.APIInfo()
	if err != nil {
		log.Panicln(err)
	}
	fmt.Printf(
		"Query Credits: %d\nScan Credits: %d\n\n", info.QueryCredits,
		info.ScanCredits)
	hostSearch, err := s.HostSearch(os.Args[1])
	if err != nil {
		log.Panicln(err)
	}
	for _, host := range hostSearch.Matches {
		fmt.Printf("%18s%8d\n", host.IPString, host.Port)
	}
}
