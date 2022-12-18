package main

import (
	"fmt"
	"log"

	"github.com/martinyonatann/go-invoice/infrastructure/database"
	"github.com/martinyonatann/go-invoice/pkg/metric"
)

// func handleParams() (string, error) {
// 	if len(os.Args) < 2 {
// 		return "", errors.New("invalid query")
// 	}

// 	return os.Args[1], nil
// }

func main() {
	metricService, err := metric.NewPrometheusService()
	if err != nil {
		log.Panic(err.Error())
	}
	appMetric := metric.NewCLI("search")

	appMetric.StartedApp()

	// query, err := handleParams()
	// if err != nil {
	// 	log.Panic(err.Error())
	// }

	db := database.DBConn()

	defer db.Close()

	// userRepository := repository.NewUserRepository(db.DB)
	// userService := user.New(*userRepository)

	all := []string{"a", "b", "c"}

	if err != nil {
		log.Panic(err)
	}
	for i := range all {
		fmt.Printf("%s", all[i])
	}
	appMetric.FinishedApp()

	err = metricService.SaveCLI(appMetric)
	if err != nil {
		log.Panic(err)
	}
}
