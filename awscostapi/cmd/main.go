package main

import (
	costhandlres "awscostapi/costHandlers"
	datehandlres "awscostapi/dateHandlers"
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Welcome to AWS cost Explorere API \n worked by me with my peers(Anibrata,shivam,rajat,bhavesh)")

	f, err := os.Create("BillingApp.csv")
	defer f.Close()

	if err != nil {
		log.Fatal("File not created", err)
	}

	w := csv.NewWriter(f)
	defer w.Flush()

	if err = w.Write(datehandlres.SetHeading()); err != nil {
		log.Fatal("header setter fail", err)
	}

	for _, record := range costhandlres.GetAwsCost() {
		if err := w.Write(record); err != nil {
			log.Fatalln("error writing record to file", err)
		}
	}

}
