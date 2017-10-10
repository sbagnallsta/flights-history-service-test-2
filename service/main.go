package main 

import (
	"os"
	"fmt"
	"log"
	"time"
	"gopkg.in/mgo.v2"
	"github.com/gin-gonic/gin"
	// "github.com/sbagnallsta/flights-history-service-test-2/service/profiles"
	"github.com/spf13/viper"
	"github.com/sbagnallsta/flights-history-service-test-2/service/models/dbConnector"
	"github.com/sbagnallsta/flights-history-service-test-2/service/controllers/health"
	"github.com/sbagnallsta/flights-history-service-test-2/service/controllers/flights"
)

func main() {
	// config := profiles.LoadConfiguration()
	profile := os.Getenv("FLIGHTS_HISTORY_SERVICE_PROFILE")
	if(profile == "") {
		profile = "local"
		fmt.Println(profile)
	}	
	viper.SetConfigName(profile + string(".json"))
	viper.AddConfigPath("./profiles/")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Error", err)
	}
	router := gin.Default()

	mongoInfo := &mgo.DialInfo {
		// Addrs: []string{config.DataSource.Url},
		Addrs: []string{viper.GetString("DataSource.Url")},
		Timeout: 60 * time.Second,
	}

	session, err := mgo.DialWithInfo(mongoInfo)

	if err != nil {
		log.Fatal("Failed to connect to mongoDB ", err)
	}

	// db := dbConnector.DBData{config.DataSource.Database, config.DataSource.Collection, session}
	db := dbConnector.DBData{viper.GetString("DataSource.Database"), viper.GetString("DataSource.Collection"), session}

	router.GET("/health", healthController.Health)

	flights := router.Group("/flights")
	{
		flights.POST("", flightController.StoreFlights(db))
	}

	router.Run(viper.GetString("Port"))
}
