package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unit-test-implementasi/handler"
	"unit-test-implementasi/service"
	"unit-test-implementasi/utils"
)

func main() {
	fileStore := utils.FileUtils{}
	service := service.NewBookingService(&fileStore)
	handler := handler.NewBookingHandler(service)

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("====== Booking Menu ======")
		fmt.Println("1. Create Booking")
		fmt.Println("2. List Bookings")
		fmt.Println("3. Update Booking")
		fmt.Println("4. Delete Booking")
		fmt.Println("5. Exit")
		fmt.Print("Select an option [1-5]: ")

		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		fmt.Println() // for spacing

		switch choice {
		case "1":
			handler.Create()
		case "2":
			handler.List()
		case "3":
			handler.Update()
		case "4":
			handler.Delete()
		case "5":
			fmt.Println("👋 Exiting application.")
			return
		default:
			fmt.Println("⚠️ Invalid option, please choose between 1 and 5.")
		}

		fmt.Println() // spacing between loops
	}
}
