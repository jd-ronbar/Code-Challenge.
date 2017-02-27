package models

import (
	"database/sql"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

var dbMetric *sql.DB;

func init() {
	var err error;
	dbMetric, err = sql.Open("mysql",
		"root:1234509876@tcp(127.0.0.1:3306)/hello")
	if err != nil {
		log.Fatal(err)
	}
}

type MetricStrct struct {
        Id          int
	Metric_name string
	Value       string
	Lon         float32
	Timestamp   int
	Lat         float32
	Driver_id   string
}

func GetAllMetrics() [] MetricStrct {
	var metrics []MetricStrct;
	rows, err := dbMetric.Query("select id, metric_name, value, lon, timestamp, lat, driver_id from metric")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var (
                id            int
                metric_name   string
		value         string
                lon           float32
                timestamp     int
	        lat           float32
                driver_id     string
	)
	for rows.Next() {
		err := rows.Scan(&id, &metric_name, &value, &lon, &timestamp, &lat, &driver_id)
		var metric = MetricStrct{Id:id, Metric_name:metric_name, Value:value, Lon: lon, Timestamp:timestamp, Lat:lat, Driver_id:driver_id};
		if err != nil {
			log.Fatal(err)
		}
		///???log.Println(metric)
		metrics = append(metrics, metric)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return metrics;
}

func InsertMetric(metric MetricStrct) {
	stmt, err := dbMetric.Prepare("INSERT INTO metric(driver_id,metric_name,lon,timestamp,lat,value) VALUES(?,?,?,?,?,?)")
	if err != nil {
		log.Fatal(err)
	}
	//defer stmtInsertMetric.Close()
	_, err = stmt.Exec(metric.Driver_id, metric.Metric_name, metric.Lon, metric.Timestamp, metric.Lat, metric.Value)
	if err != nil {
		log.Fatal(err)
	}
}


func UpdateMetric(metric MetricStrct) {
	stmt, err := dbMetric.Prepare("UPDATE metric SET driver_id= ?, metric_name=?, lon= ? ,timestamp= ?, lat= ?, value= ? WHERE id=?")
	if err != nil {
		log.Fatal(err)
	}
	_, err = stmt.Exec( metric.Driver_id,  metric.Metric_name, metric.Lon, metric.Timestamp, metric.Lat, metric.Value,    metric.Id,)
	if err != nil {
		log.Fatal(err)
	}
}

func DeleteMetric(id int) {
	log.Println("Going to delete " , id)
	stmt, err := db.Prepare("DELETE FROM metric WHERE id=?")
	if err != nil {
		log.Fatal(err)
	}
	_, err = stmt.Exec( id)
	if err != nil {
		log.Fatal(err)
	}
}

func GetMetric(tid int) MetricStrct {
	var metric MetricStrct
	rows, err := dbMetric.Query("select id, driver_id, metric_name, value, lon, timestamp, lat from metric where id=?", tid)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var (
                id            int
                driver_id     string
                metric_name   string
		value         string
                lon           float32
                timestamp     int
	        lat           float32
                
	)
	for rows.Next() {
		err := rows.Scan(&id, &driver_id, &metric_name, &value, &lon, &timestamp, &lat)
		metric = MetricStrct{Id:id, Driver_id:driver_id, Metric_name:metric_name, Value:value, Lon:lon, Timestamp: timestamp, Lat:lat};
		if err != nil {
			log.Fatal(err)
		}
		log.Println(metric)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return metric;
}
