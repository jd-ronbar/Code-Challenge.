package main

import (
	"encoding/json"
	"log"
	"io/ioutil"
	"bytes"
	//"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type MetricStrct struct {
	Metric_name string
	Value       string
	Lon         float32
	Timestamp   int
	Lat         float32
	Driver_id   string
}

func getMetrics() []MetricStrct {
	var res [] MetricStrct

	b, err := ioutil.ReadFile("./input/metrics.json")
	if err != nil {
		log.Fatal(err)
	}
	jsonStream := bytes.NewReader(b);
	dec := json.NewDecoder((jsonStream))

	// while the array contains values
	for dec.More() {
		var m MetricStrct
		// decode an array value (Message)
		err := dec.Decode(&m)
		if err != nil {
			log.Println(err)
			continue
		}
		if "" == m.Driver_id  {
			log.Println("ignoring metric without driver " , m)
			continue
		}
		//fmt.Printf("%v: %v\n", m.Metric_name, m.Driver_id)
		res = append(res, m)
	}

	log.Print(dec)
	return res
}

type DriverStrct struct {
	Id             int
	Name           string
	License_number string
}

func getDrivers() [] DriverStrct {
	var drivers []DriverStrct
	b, err := ioutil.ReadFile("./input/drivers.json")
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(b, &drivers)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(drivers)
	return drivers;
}

func insertDriversIntoDb(drivers [] DriverStrct){
	db, err := sql.Open("mysql",
		"root:1234509876@tcp(127.0.0.1:3306)/hello")
                defer db.Close()
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := db.Prepare("INSERT INTO driver(id,name,licenseNumber) VALUES(?,?,?)")
	if err != nil {
		log.Fatal(err)
	}
	for _,element := range drivers {
		//log.Printf("[%d]=%s",index,element)
		_, err := stmt.Exec(element.Id, element.Name, element.License_number)
		if err != nil {
			log.Println(err)
		}
	}
	defer db.Close()
}

func insertMetricsIntoDb(metrics [] MetricStrct){
	db, err := sql.Open("mysql",
		"root:1234509876@tcp(127.0.0.1:3306)/hello")
        defer db.Close()
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := db.Prepare("INSERT INTO metric(metric_name, value, lon, timestamp, lat, driver_id) VALUES(?,?,?,?,?,?)")
	if err != nil {
		log.Fatal(err)
	}
	for _,element := range metrics {
		//log.Printf("[%d]=%s",index,element)
		_, err := stmt.Exec(element.Metric_name, element.Value, element.Lon, element.Timestamp, element.Lat, element.Driver_id)
		if err != nil {
			log.Println(err)
		}
	}
}

func main() {
	metrics := getMetrics()
	insertMetricsIntoDb(metrics)
	drivers:= getDrivers()
	insertDriversIntoDb(drivers)
}
