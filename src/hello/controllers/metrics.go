package controllers

import (
	"github.com/astaxie/beego"
	"hello/models"
	"log"
	"strconv"
)


type MetricsController  struct {
	beego.Controller
}

func (c *MetricsController) Get() {
	metrics := models.GetAllMetrics();
	c.Data["Metrics"] = metrics
	c.TplName = "metrics/list.tpl"
}

func (this *MetricsController) Add() {
	this.TplName = "metrics/add.tpl"
	if this.Ctx.Input.Method() == "GET" {
		log.Println("Showing empty form");
		// render the empty form and exit
		return
	}

	log.Println(this.Ctx.Input.RequestBody)
        //id,_ := this.GetInt("Id")
	metric_name := this.GetString("Metric_name")
	value := this.GetString("Value")
        lon,_ := this.GetFloat("Lon")
        lon32 := float32(lon)
        timestamp,_ := this.GetInt("Timestamp")
        lat,_ := this.GetFloat("Lat")
        lat32 := float32(lat) 
        driver_id := this.GetString("Driver_id")
	mtr:= models.MetricStrct {Driver_id:driver_id, Metric_name:metric_name, Lon:lon32, Timestamp:timestamp, Lat:lat32, Value:value}
	log.Println("Going to insert %s", mtr)
	models.InsertMetric(mtr);
	this.Redirect("/metrics",302)
}

func (this *MetricsController) Edit() {
	this.TplName = "metrics/edit.tpl"
	if this.Ctx.Input.Method() == "GET" {
		log.Println("Showing empty form");
		sid := this.Ctx.Input.Param(":id")
		did, err := strconv.Atoi(sid)
                
		if err != nil {
			log.Fatal(err)
		}
		metric := models.GetMetric(did)
		log.Println("metric=",metric)
		this.Data["Metric"] = metric
		return
	}
	//log.Println("--->",this.Ctx.Input.RequestBody)
	metric_name := this.GetString("Metric_name")
	value := this.GetString("Value")
        lon,_ := this.GetFloat("Lon")
        lon32 := float32(lon)
	timestamp,_ := this.GetInt("Timestamp")
	lat,_ := this.GetFloat("Lat")
	lat32 := float32(lat)
        driver_id := this.GetString("Driver_id")
	id,_ := this.GetInt("Id")


	mtr:= models.MetricStrct {Id:id, Metric_name:metric_name, Value:value, Lon:lon32, Timestamp:timestamp, Lat:lat32, Driver_id:driver_id}
	log.Println("Going to UPDATE[id,metric_name,value,lon,timestamp,lat,driver_id] %s", mtr)
	models.UpdateMetric(mtr);
	this.Redirect("/metrics",302)
}

func (this *MetricsController) Delete() {
	sid := this.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(sid)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("got " , id , " to delete")
	models.DeleteMetric(id);
	this.Redirect("/metrics",302)
}
