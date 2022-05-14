package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"
)

type Sale struct {
	Region        string
	Country       string
	ItemType      string
	SalesChannel  string
	OrderPriority string
	OrderDate     string
	OrderID       string
	ShipDate      string
	UnitsSold     string
	UnitPrice     string
	UnitCost      string
	TotalRevenue  string
	TotalCost     string
	TotalProfit   string
}

func getSales(csvLines [][]string) []Sale {
	sales := []Sale{}
	for _, line := range csvLines {
		sale := Sale{
			Region:        line[0],
			Country:       line[1],
			ItemType:      line[2],
			SalesChannel:  line[3],
			OrderPriority: line[4],
			OrderDate:     line[5],
			OrderID:       line[6],
			ShipDate:      line[7],
			UnitsSold:     line[8],
			UnitPrice:     line[9],
			UnitCost:      line[10],
			TotalRevenue:  line[11],
			TotalCost:     line[12],
			TotalProfit:   line[13],
		}
		sales = append(sales, sale)
	}

	return sales
}

func printSize(list []Sale) {
	fmt.Println(len(list))
}

func main() {
	csvFile, err := os.Open("file/sales.csv")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened CSV file")
	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Println(err)
	}

	sales := getSales(csvLines)

	tam := len(sales)
	half := tam / 2

	salesFirst := sales[0:half]
	salesSecond := sales[half:tam]

	go printSize(salesFirst)
	go printSize(salesSecond)

	time.Sleep(time.Second)

	fmt.Println("Successfully finished")
}
