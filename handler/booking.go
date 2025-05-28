package handler

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unit-test-implementasi/model"
	"unit-test-implementasi/service"

)

type BookingHandler struct {
	ServiceBooking service.BookingService
}

func NewBookingHandler(serviceBooking service.BookingService) BookingHandler {
	return BookingHandler{
		ServiceBooking: serviceBooking,
	}
}

// Create prompts user input and creates a booking
func (bookingHandler *BookingHandler) Create() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter Name: ")
	name, _ := reader.ReadString('\n')

	fmt.Print("Enter Destination: ")
	dest, _ := reader.ReadString('\n')

	booking := model.Booking{
		Name:        strings.TrimSpace(name),
		Destination: strings.TrimSpace(dest),
	}

	err := bookingHandler.ServiceBooking.Create(booking)
	if err != nil {
		fmt.Println("❌ Failed to create booking:", err)
		return
	}
	fmt.Println("✅ Booking created successfully!")
}

// List displays all bookings
func (bookingHandler *BookingHandler) List() {
	bookings, err := bookingHandler.ServiceBooking.List()
	if err != nil {
		fmt.Println("❌ Failed to list bookings:", err)
		return
	}

	if len(bookings) == 0 {
		fmt.Println("📭 No bookings found.")
		return
	}

	fmt.Println("📋 List of Bookings:")
	for _, b := range bookings {
		fmt.Printf("- ID: %s | Name: %s | Destination: %s | Price: %.2f\n", b.ID, b.Name, b.Destination, b.Price)
	}
}

// Update allows user to update a booking by ID
func (bookingHandler *BookingHandler) Update() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter Booking ID to update: ")
	id, _ := reader.ReadString('\n')

	fmt.Print("Enter New Name: ")
	name, _ := reader.ReadString('\n')

	fmt.Print("Enter New Destination: ")
	dest, _ := reader.ReadString('\n')

	booking := model.Booking{
		ID:          strings.TrimSpace(id),
		Name:        strings.TrimSpace(name),
		Destination: strings.TrimSpace(dest),
	}

	err := bookingHandler.ServiceBooking.Update(booking)
	if err != nil {
		fmt.Println("❌ Failed to update booking:", err)
		return
	}
	fmt.Println("✅ Booking updated successfully!")
}

// Delete deletes a booking by ID
func (bookingHandler *BookingHandler) Delete() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter Booking ID to delete: ")
	id, _ := reader.ReadString('\n')

	err := bookingHandler.ServiceBooking.Delete(strings.TrimSpace(id))
	if err != nil {
		fmt.Println("❌ Failed to delete booking:", err)
		return
	}
	fmt.Println("✅ Booking deleted successfully!")
}
