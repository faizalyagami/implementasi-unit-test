package service

import (
	"encoding/json"
	"errors"
	"strings"
	"unit-test-implementasi/model"
	"unit-test-implementasi/utils"

	"github.com/google/uuid"
)

const bookingFile = "data/booking.json"
const ticketFile = "data/master_ticket.json"

type BookingService struct {
	StoreFile utils.StoreFileInterface
}

func NewBookingService(storeFile utils.StoreFileInterface) BookingService {
	return BookingService{
		StoreFile: storeFile,
	}
}

func (booking *BookingService) Create(newBooking model.Booking) error {
	// Step 1: Read ticket data
	ticketData, err := booking.StoreFile.ReadFile(ticketFile)
	if err != nil {
		return err
	}

	var tickets []model.Ticket
	if err := json.Unmarshal([]byte(ticketData), &tickets); err != nil {
		return err
	}

	// Step 2: Find the ticket based on destination
	found := false
	for i, t := range tickets {
		if t.Destination == newBooking.Destination {
			found = true
			if t.Stock <= 0 {
				return errors.New("ticket stock is not available")
			}

			// Step 3: Decrease ticket stock
			tickets[i].Stock--

			// Set booking price from ticket to ensure consistency
			newBooking.Price = t.Price
			break
		}
	}

	if !found {
		return errors.New("ticket for destination not found")
	}

	// Step 4: Save updated ticket stock back to file
	updatedTicketData, err := json.MarshalIndent(tickets, "", "  ")
	if err != nil {
		return err
	}
	if err := booking.StoreFile.WriteFile(ticketFile, string(updatedTicketData)); err != nil {
		return err
	}

	// Step 5: Read existing bookings
	bookings, err := booking.List()
	if err != nil {
		return err
	}

	// Step 6: Add new booking
	newBooking.ID = uuid.New().String()
	bookings = append(bookings, &newBooking)

	// Step 7: Save new bookings to file
	data, err := json.MarshalIndent(bookings, "", "  ")
	if err != nil {
		return err
	}

	return booking.StoreFile.WriteFile(bookingFile, string(data))
}

// List returns all bookings from the file
func (booking *BookingService) List() ([]*model.Booking, error) {
	exists := booking.StoreFile.CheckFile(bookingFile)
	if !exists {
		return []*model.Booking{}, nil
	}

	content, err := booking.StoreFile.ReadFile(bookingFile)
	if err != nil {
		return nil, err
	}

	// Handle empty content
	if strings.TrimSpace(content) == "" {
		return []*model.Booking{}, nil
	}

	var bookings []*model.Booking
	if err := json.Unmarshal([]byte(content), &bookings); err != nil {
		return nil, err
	}

	return bookings, nil
}

// Update modifies an existing booking by ID
func (booking *BookingService) Update(updatedBooking model.Booking) error {
	bookings, err := booking.List()
	if err != nil {
		return err
	}

	updated := false
	for i, b := range bookings {
		if b.ID == updatedBooking.ID {
			bookings[i] = &updatedBooking
			updated = true
			break
		}
	}

	if !updated {
		return errors.New("booking not found")
	}

	data, err := json.MarshalIndent(bookings, "", "  ")
	if err != nil {
		return err
	}

	return booking.StoreFile.WriteFile(bookingFile, string(data))
}

// Delete removes a booking by ID
func (booking *BookingService) Delete(id string) error {
	bookings, err := booking.List()
	if err != nil {
		return err
	}

	index := -1
	for i, b := range bookings {
		if b.ID == id {
			index = i
			break
		}
	}

	if index == -1 {
		return errors.New("booking not found")
	}

	// Remove the booking
	bookings = append(bookings[:index], bookings[index+1:]...)

	data, err := json.MarshalIndent(bookings, "", "  ")
	if err != nil {
		return err
	}

	return booking.StoreFile.WriteFile(bookingFile, string(data))
}
