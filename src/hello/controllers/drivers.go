package controllers

import (
	"github.com/astaxie/beego"
	"hello/models"
	"log"
	"strconv"
)


type DriversController  struct {
	beego.Controller
}

func (c *DriversController) Get() {
	drivers := models.GetAllDrivers();
	c.Data["Drivers"] = drivers
	c.TplName = "drivers/list.tpl"
}

func (this *DriversController) Add() {
	this.TplName = "drivers/add.tpl"
	if this.Ctx.Input.Method() == "GET" {
		log.Println("Showing empty form");
		// render the empty form and exit
		return
	}

	log.Println(this.Ctx.Input.RequestBody)
	id,_ := this.GetInt("Id")
	name := this.GetString("Name")
	license_number := this.GetString("License_number")
	drv:= models.DriverStrct {Id:id, Name:name, License_number:license_number	}
	log.Println("Going to insert %s", drv)
	models.InsertDriver(drv);
	this.Redirect("/drivers",302)
}

func (this *DriversController) Edit() {
	this.TplName = "drivers/edit.tpl"
	if this.Ctx.Input.Method() == "GET" {
		log.Println("Showing empty form");
		sid := this.Ctx.Input.Param(":id")
		did, err := strconv.Atoi(sid)
		if err != nil {
			log.Fatal(err)
		}
		driver := models.GetDriver(did)
		this.Data["Driver"] = driver
		return
	}
	log.Println(this.Ctx.Input.RequestBody)
	id,_ := this.GetInt("Id")
	name := this.GetString("Name")
	license_number := this.GetString("License_number")
	drv:= models.DriverStrct {Id:id, Name:name, License_number:license_number	}
	log.Println("Going to UPDATE %s", drv)
	models.UpdateDriver(drv);
	this.Redirect("/drivers",302)
}

func (this *DriversController) Delete() {
	sid := this.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(sid)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("got " , id , " to delete")
	models.DeleteDriver(id);
	this.Redirect("/drivers",302)
}
