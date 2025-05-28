package service_test

import (
	"testing"
	"unit-test-implementasi/model"
	"unit-test-implementasi/service"
	"unit-test-implementasi/utils"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

)

func TestCreateBooking_Success(t *testing.T) {

	mockStore := new(utils.MockFileUtils)

	//setup mock expectation
	ticketJson := `[{"destination": "Bali", "stock": 2, "price": 1000000}]`
	bookingJson := `[]`

	mockStore.On("CheckFile", "data/booking.json").Return(true)
	mockStore.On("ReadFile", "data/booking.json").Return(bookingJson, nil)
	mockStore.On("ReadFile", "data/master_ticket.json").Return(ticketJson, nil)
	mockStore.On("WriteFile", "data/master_ticket.json", mock.Anything).Return(nil)
	mockStore.On("WriteFile", "data/booking.json", mock.Anything).Return(nil)

	serviceBooking := service.NewBookingService(mockStore)

	booking := model.Booking{
		Name:        "Faisal",
		Destination: "Bali",
	}

	err := serviceBooking.Create(booking)
	assert.NoError(t, err)
}

func TestCreateBooking_DestNotFound(t *testing.T) {
	mockStore := new(utils.MockFileUtils)

	//ticket list
	ticketJson := `[{"destination": "Medan", "stock": 5, "Price": 5000000}]`
	// bookingJson := `[]`

	// mockStore.On("CheckFile", "data/booking.json").Return(true)
	// mockStore.On("ReadFile", "data/booking.json").Return(bookingJson, nil)
	mockStore.On("ReadFile", "data/master_ticket.json").Return(ticketJson, nil)

	serviceBooking := service.NewBookingService(mockStore)

	booking := model.Booking{
		Name:        "Faisal",
		Destination: "Arab Saudi",
	}

	err := serviceBooking.Create(booking)

	assert.Error(t, err)
	assert.EqualError(t, err, "ticket for destination not found")
}

func TestCreateBooking_StockHabis(t *testing.T) {
	mockStore := new(utils.MockFileUtils)

	ticketJson := `[{"destination": "Bandung", "stock": 0, "Price": 5000000}]`
	bookingJson := `[]`

	mockStore.On("CheckFile", "data/booking.json").Return(true)
	mockStore.On("ReadFile", "data/booking.json").Return(bookingJson, nil)
	mockStore.On("ReadFile", "data/master_ticket.json").Return(ticketJson, nil)

	serviceBooking := service.NewBookingService(mockStore)

	booking := model.Booking{
		Name:        "Faisal",
		Destination: "Bandung",
	}

	err := serviceBooking.Create(booking)

	assert.Error(t, err)
	assert.EqualError(t, err, "ticket stock is not available")

}
