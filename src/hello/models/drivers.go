package models

import (
	"database/sql"
	"log"
	_ "github.com/go-sql-driver/mysql"
	_ "database/sql/driver"
)

var db *sql.DB;

func init() {
	var err error;
	db, err = sql.Open("mysql",
		"root:1234509876@tcp(127.0.0.1:3306)/hello")
	if err != nil {
		log.Fatal(err)
	}
}

type DriverStrct struct {
	Id             int
	Name           string
	License_number string
}

func GetAllDrivers() [] DriverStrct {
	var drivers []DriverStrct;
	rows, err := db.Query("select id, name, licenseNumber from driver")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var (
		id            int
		name          string
		licenseNumber string
	)
	for rows.Next() {
		err := rows.Scan(&id, &name, &licenseNumber)
		var driver = DriverStrct{Id: id, Name: name, License_number: licenseNumber};
		if err != nil {
			log.Fatal(err)
		}
		log.Println(driver)
		drivers = append(drivers, driver)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return drivers;
}

func InsertDriver(driver DriverStrct) {
	stmt, err := db.Prepare("INSERT INTO driver(id,name,licenseNumber) VALUES(?,?,?)")
	if err != nil {
		log.Fatal(err)
	}
	//defer stmtInsertDriver.Close()
	_, err = stmt.Exec(driver.Id, driver.Name, driver.License_number)
	if err != nil {
		log.Fatal(err)
	}
}


func UpdateDriver(driver DriverStrct) {
	stmt, err := db.Prepare("UPDATE driver SET name=?, licenseNumber= ? WHERE id=?")
	if err != nil {
		log.Fatal(err)
	}
	_, err = stmt.Exec( driver.Name, driver.License_number, driver.Id,)
	if err != nil {
		log.Fatal(err)
	}
}

func DeleteDriver(id int) {
	log.Println("Going to delete " , id)
	stmt, err := db.Prepare("DELETE FROM driver WHERE id=?")
	if err != nil {
		log.Fatal(err)
	}
	_, err = stmt.Exec( id)
	if err != nil {
		log.Fatal(err)
	}
}


func GetDriver(driverId int) DriverStrct {
	var driver DriverStrct
	rows, err := db.Query("select id, name, licenseNumber from driver where id=?", driverId)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var (
		id            int
		name          string
		licenseNumber string
	)
	for rows.Next() {
		err := rows.Scan(&id, &name, &licenseNumber)
		driver = DriverStrct{Id: id, Name: name, License_number: licenseNumber};
		if err != nil {
			log.Fatal(err)
		}
		log.Println(driver)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return driver;
}
