package handlers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/Kv-062-DevOps/monitoring/metrics"
	"gopkg.in/yaml.v3"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {

	// gather method, time, endpoint for Prometheus
	start := time.Now()
	serName := "get-srv"
	method := r.Method
	endpoint := r.URL.Path
	status := ""

	var err error


	// defer function which collects required metrics
	defer func() {
		metrics.CounterVec.WithLabelValues(serName, method, endpoint, status).Inc()
		metrics.HistogramVec.WithLabelValues(serName, endpoint).Observe(time.Since(start).Seconds())
	}()

	req, err := http.NewRequest("GET", "http://"+os.Getenv("ENDPOINT")+"/list",
		r.Body)
	log.Println("Forming GET request to DB service")
	if err != nil {
		status = "500"
		log.Println("Problems with forming GET request. Execution aborted")
		http.Error(w, "Troubles with forming request to DB service", 500)
		return
	}
	log.Println("GET request formed successfully")
	req.Header = r.Header

	client := &http.Client{}
	resp, err := client.Do(req)
	log.Println("Sending GET request to DB service")

	if err != nil {
		log.Println("Troubles with getting data back from DB service. Execution aborted")
		status = "500"
		http.Error(w, "Troubles with receiving data from DB service", 500)
		return
	}
	log.Println("Successfully got data from DB service")
	defer resp.Body.Close()

	var response []Employee

	data, _ := ioutil.ReadAll(resp.Body)

	err = yaml.Unmarshal(data, &response)

	log.Println("Unmarshalling received data from DB service")

	if err != nil {
		var emp Employee
		err := yaml.Unmarshal(data, &emp)
		if err != nil {
			status = "500"
			log.Println("Unmarshalling was unsuccessful. Execution aborted.")
			http.Error(w, "Error unmarshalling YAML OR No valid data returned by DB", 500)
			return
		}
		emp.Salary, err = calculateSalary(&emp)
		if err != nil {
			status = "500"
			log.Println("Unmarshalling was unsuccessful. Execution aborted.")
			http.Error(w, "Error unmarshalling YAML OR No valid data returned by DB", 500)
			return
		}
		log.Println("Unmarshalling was successful")
		emp.Salary, err = calculateSalary(&emp)
		log.Println("Calculating salary")
		if err != nil {
			status = "500"
			log.Println("Calculating salary was unsuccessful. Execution aborted")
			http.Error(w, "Troubles in calculating salary", 500)
			return
		}
		log.Println("Salary calculated successfully")
		err = json.NewEncoder(w).Encode(emp)
		log.Println("Encoding response")
		if err != nil {
			log.Println("Encoding response was unsuccessful")
			status = "500"
			http.Error(w, "Error encoding data to JSON", 500)
			return
		}
		log.Println("Response was encoded successfully")
		status = "200"
		w.WriteHeader(200)
		return
	}

	log.Println("Calculating salary")
	for i := 0; i < len(response); i++ {
		emp := &response[i]
		emp.Salary, err = calculateSalary(emp)
		if err != nil {
			status = "500"
			log.Println("Calculating salary was unsuccessful. Execution aborted")
			http.Error(w, "Troubles in calculating salary", 500)
			return
		}
	}
	log.Println("Salary was calculated successfully")

	err = json.NewEncoder(w).Encode(response)
	log.Println("Encoding response")

	if err != nil {
		status = "500"
		log.Println("Encoding response was unsuccessful")
		http.Error(w, "Error encoding data to JSON", 500)
		return
	}
	log.Println("Response was encoded successfully")
}

func calculateSalary(emp *Employee) (string, error) {
	exp, err := strconv.Atoi(emp.Experience)
	if err != nil {
		return "", errors.New("Trouble converting data")
	}
	def, err := strconv.Atoi(emp.DefaultSalary)
	if err != nil {
		return "", errors.New("Trouble converting data")
	}
	temp := def * exp
	return strconv.Itoa(temp), nil
}
