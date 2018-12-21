package main

import (
	"fmt"

	"github.com/satori/go.uuid"
)

//Tesla Struct holds the data for a tesla vehicle
type Tesla struct {
	ID       string `json:"ID,omitempty"`
	Make     string `json:"Make,omitempty"`
	Model    string `json:"Model,omitempty"`
	ImageURL string `json:"image_url,omitempty"`
	Owned    bool   `json:"owned,omitempty"`
	Cost     int64  `json:"Cost,omitempty"`
}

/*
TeslaDB interface is used to make sure all the interactions with the database are
thread safe
*/
type TeslaDB interface {
	//ListVehicles lists all the Teslas in the Vehicles table
	ListVehicles() ([]*Tesla, error)

	//AddTesla adds a Tesla to the database and returns a tesla object and an error
	AddTesla(t *Tesla) (id string, err error)

	//UpdateTesla edits info of an existing Tesla and returns an object and an error
	UpdateTesla(t *Tesla) error

	//GetTesla searches for tesla in database and returns the object and an error
	GetTesla(id string) (*Tesla, error)

	//Closes the database and frees up resources
	Close() error
}

/*
Utility function to set the tesla ID for a given tesla vehicle, this is the unique identifier for
each object
*/

func (t *Tesla) setTeslaID() string {
	u1, err := uuid.NewV1()
	if err != nil {
		fmt.Printf("Error Setting the TeslaID: %s", err.Error())
	}
	t.ID = u1.String()

	return t.ID
}

func main() {
	fmt.Println("Initializing project")
}
