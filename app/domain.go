// This file contains types inside the asapp domain.

// Package app contains structs and implementations for this app.
package app

import (
	"github.com/Bitspark/go-bitnode/bitnode"
	"log"
)

// Credentials represents the required information to obtain an access token.
type Credentials struct {
	// ClientID is the identifier of the client.
	ClientID string `json:"clientId"`

	// ClientSecret is the secret key associated with the client.
	ClientSecret string `json:"clientSecret"`
}

// PersonioEmployee represents an employee in Personio.
type PersonioEmployee struct {
	// ID is the unique identifier of the employee.
	ID int `json:"id"`

	// FirstName is the first name of the employee.
	FirstName string `json:"firstName"`

	// LastName is the last name of the employee.
	LastName string `json:"lastName"`
}

// PersonioProject represents a project in Personio.
type PersonioProject struct {
	// ID is the unique identifier of the project.
	ID int `json:"id"`

	// Name is the name of the project.
	Name string `json:"name"`

	// Active indicates this project is ongoing.
	Active bool `json:"active"`

	// CreatedAt is the UNIX timestamp when the project was created.
	CreatedAt float64 `json:"createdAt"`

	// UpdatedAt is the UNIX timestamp when the project was last updated.
	UpdatedAt float64 `json:"updatedAt"`
}

// PersonioAttendance represents an attendance for a Personio employee.
type PersonioAttendance struct {
	// Employee is the ID of the employee.
	Employee int `json:"employee"`

	// Project is the ID of the project.
	Project int `json:"project"`

	// Date is the date formatted as YYYY-MM-DD.
	Date string `json:"date"`

	// StartTime is the start time formatted as HH:MM in 24h format.
	StartTime string `json:"startTime"`

	// EndTime is the end time formatted as HH:MM in 24h format.
	EndTime string `json:"endTime"`

	// DurationNet is the net duration excluding break.
	DurationNet float64 `json:"durationNet"`

	// Break is the break in hours.
	Break float64 `json:"break"`

	// Comment is the description of what has been done during the attendance.
	Comment string `json:"comment"`
}

// Domain containing mainly wrappers for applications.
type Domain struct {
	Domain *bitnode.Domain
	Node   bitnode.Node
}

// NewPersonioAccount creates a new BlankSparkable instance.
func (asapp *Domain) NewPersonioAccount() (*BlankSparkable, error) {
	// Get the BlankSparkable sparkable from the domain.
	accSpark, err := asapp.Domain.GetSparkable("hub.asapp.BlankSparkable")
	if err != nil {
		log.Fatal(err)
	}

	// Prepare the BlankSparkable system.
	accSys, err := asapp.Node.PrepareSystem(bitnode.Credentials{}, *accSpark)
	if err != nil {
		log.Fatal(err)
	}

	// Initialize the BlankSparkable.
	acc := &BlankSparkable{
		System: accSys,
	}

	return acc, nil
}
