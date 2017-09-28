package flightController

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/sbagnallsta/flights-history-service-test-2/service/repository"
	"github.com/sbagnallsta/flights-history-service-test-2/service/models/dbConnector"
	"github.com/sbagnallsta/flights-history-service-test-2/service/models/flight"
)

func StoreFlights(dbConnector dbConnector.DBData) gin.HandlerFunc {

	function := func(context *gin.Context) {

		var data []flight.Flight

		context.BindJSON(&data)

		go flightsRepository.SaveFlight(dbConnector, data) 

		context.Status(http.StatusCreated)

	}

	return gin.HandlerFunc(function)
}

