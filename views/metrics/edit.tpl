<!DOCTYPE html>



<html>
<head>
    <title>Beego</title>
    This is metrics Edit.tpl
<br/><br/><br/><br/>
    <form method="post">
	<div>
           ID: {{.Metric.Id}}
            <input id="Id" type='number' name="Id" hidden="hidden" value="{{.Metric.Id}}">
        </div>
        <div>
            <label for="Metric_name">Metric_name:</label>
            <input id="Metric_name" type='text' name="Metric_name" value="{{.Metric.Metric_name}}">
        </div>
        <div>
            <label for="Value">Value:</label>
            <input id="Value" type='text' name="Value" value="{{.Metric.Value}}">
        </div>
 	<div>
            <label for="Lon">Lon:</label>
            <input id="Lon" type='float' name="Lon" value="{{.Metric.Lon}}">
        </div>
 	<div>
            <label for="Timestamp">Timestamp:</label>
            <input id="Timestamp" type='number' name="Timestamp" value="{{.Metric.Timestamp}}">
        </div>
 	<div>
            <label for="Lat">Lat:</label>
            <input id="Lat" type='float' name="Lat" value="{{.Metric.Lat}}">
        </div>
        <div>
            <label for="Driver_id">Driver_id:</label>
            <input id="Driver_id" type='text' name="Driver_id" value="{{.Metric.Driver_id}}">
        </div>

        <button type="submit">Submit</button>
    </form>
    <a href="/metrics/delete/{{.Metric.Id}}">Delete</a>
</head>
</html>

