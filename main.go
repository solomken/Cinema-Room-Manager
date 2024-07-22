package main

import (
	"errors"
	"fmt"
)

const (
	minSeatsCount = 60
	priceHigh     = 10
	priceLow      = 8
)

func printTheaterMap(theater [][]bool) {
	fmt.Printf("\nCinema:\n  ")
	for seatsNumber := 1; seatsNumber <= len(theater[0]); seatsNumber++ {
		fmt.Print(seatsNumber, " ")
	}
	fmt.Println()

	for i := 0; i < len(theater); i++ {
		fmt.Print(i+1, " ")

		for j := 0; j < len(theater[i]); j++ {
			if theater[i][j] {
				fmt.Print("B ")

				continue
			}
			fmt.Print("S ")
		}
		fmt.Println()
	}
	fmt.Println()
}

func ticketPrice(rows int, seats int, selectedRow int) int {
	var price int

	if rows*seats <= minSeatsCount {
		price = priceHigh
	} else {
		frontRows := rows / 2
		if selectedRow < frontRows {
			price = priceHigh
		} else {
			price = priceLow
		}
	}

	return price
}

func scanTheaterSize() (int, int) {
	var rows, seats int

	fmt.Println("Enter the number of rows:")
	fmt.Scan(&rows)
	fmt.Println("Enter the number of seats in each row:")
	fmt.Scan(&seats)

	return rows, seats
}

func getReservationDetails() (int, int) {
	var selectedRow, selectedSeat int

	fmt.Println("\nEnter a row number:")
	fmt.Scan(&selectedRow)
	fmt.Println("Enter a seat number in that row:")
	fmt.Scan(&selectedSeat)

	return selectedRow - 1, selectedSeat - 1
}

func menuNavigation(rows int, seats int, theater [][]bool) {
	var menuChoice, selectedRow, selectedSeat, currentIncome int

	for {
		fmt.Printf("\n1. Show the seats\n2. Buy a ticket\n3. Statistics\n0. Exit\n")
		fmt.Scan(&menuChoice)

		switch menuChoice {
		case 1:
			printTheaterMap(theater)
			continue
		case 2:
			for {
				selectedRow, selectedSeat = getReservationDetails()
				currentPrice := ticketPrice(rows, seats, selectedRow)

				currentIncome += currentPrice

				_, err := makeReservation(rows, seats, selectedRow, selectedSeat, theater)
				if err != nil {
					fmt.Println(err)
					continue
				}
				fmt.Printf("\nTicket price: $%d\n", currentPrice)
				break
			}
		case 3:
			purchasedTickets, purchasedTicketsInPercent := statistic(theater)
			totalIncome := calculateTotalIncome(rows, seats)
			printStatistics(purchasedTickets, purchasedTicketsInPercent, currentIncome, totalIncome)
		case 0:
			return
		default:
			fmt.Println("incorrect choice")
			continue
		}
	}
}

func initTheater(rows int, seats int) [][]bool {
	theater := make([][]bool, rows)
	for i := range theater {
		theater[i] = make([]bool, seats)
	}

	return theater
}

func makeReservation(
	rows int,
	seats int,
	selectedRow,
	selectedSeat int,
	theater [][]bool,
) ([][]bool, error) {
	if selectedRow+1 > rows || selectedRow+1 <= 0 || selectedSeat+1 > seats || selectedSeat+1 <= 0 {
		return theater, errors.New("\nWrong input!")
	}

	if theater[selectedRow][selectedSeat] {
		return theater, errors.New("\nThat ticket has already been purchased!")
	}

	theater[selectedRow][selectedSeat] = true

	return theater, nil
}

func statistic(theater [][]bool) (int, float64) {
	if theater == nil || len(theater) == 0 {
		return 0, 0
	}

	var purchasedTickets int

	for i := range theater {
		for j := range theater[i] {
			if theater[i][j] {
				purchasedTickets++
			}
		}
	}

	rows := len(theater)
	seats := len(theater[0])

	purchasedTicketsInPercent := (float64(purchasedTickets)) / (float64(rows * seats)) * 100

	return purchasedTickets, purchasedTicketsInPercent
}

func printStatistics(purchasedTickets int, purchasedTicketsInPercent float64, currentIncome, totalIncome int) {
	fmt.Println("\nNumber of purchased tickets: ", purchasedTickets)
	fmt.Printf("Percentage: %.2f%%\n", purchasedTicketsInPercent)
	fmt.Printf("Current income: $%d\n", currentIncome)
	fmt.Printf("Total income: $%d\n", totalIncome)
}

func calculateTotalIncome(rows int, seats int) int {
	var totalIncome int

	if rows*seats <= minSeatsCount {
		totalIncome = (rows * seats) * priceHigh
	} else {
		frontRows := rows / 2
		totalIncome += (frontRows * seats) * priceHigh
		totalIncome += ((rows - frontRows) * seats) * priceLow
	}
	return totalIncome
}

func main() {
	rows, seats := scanTheaterSize()
	theater := initTheater(rows, seats)
	menuNavigation(rows, seats, theater)
}
