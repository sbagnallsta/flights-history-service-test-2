package flightsRepository

import (
	"log"
	"github.com/sbagnallsta/flights-history-service-test-2/service/models/flight"
	"github.com/sbagnallsta/flights-history-service-test-2/service/models/dbConnector"
)

func SaveFlight(dbConnector dbConnector.DBData, data []flight.Flight) {

	sessionCopy := dbConnector.Session.Copy()
	defer sessionCopy.Close()

	db := sessionCopy.DB(dbConnector.Database).C(dbConnector.Collection)

	for _, elem := range data {

		err := db.Insert(elem)

		if err != nil {
			log.Printf("Failed to insert %v due to %v", elem, err)
		}

	}

}