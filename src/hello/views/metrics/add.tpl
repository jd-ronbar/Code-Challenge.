<!DOCTYPE html>

<html>
<head>
    <title>Beego</title>
    This is metricsInsert.tpl


    <form action="/metrics/add" method="post">
        <div>
            <label for="Metric_name">Metric_name:</label>
            <input id="Metric_name" type='text' name="Metric_name">
        </div>
        <div>
            <label for="Value">Value:</label>
            <input id="Value" type='text' name="Value">
        </div>
        <div>

            <label for="Lon">Lon:</label>
            <input id="Lon" type='number' name="Lon">
        </div>
	<div>
            <label for="Timestample">Timestample:</label>
            <input id="Timestample" type='number' name="Timestample">
        </div>
	<div>
            <label for="Lat">Lat:</label>
            <input id="Lat" type='number' name="Lat">
        </div>
	<div>
            <label for="Driver_id">Driver_id:</label>
            <input id="Driver_id" type='text' name="Driver_id">
        </div>

        <button type="submit">Submit</button>
    </form>


    <!--<table>-->
    <!--<thead>-->
    <!--<tr>-->
    <!--<td>Id</td>-->
    <!--<td>Name</td>-->
    <!--<td>License</td>-->
    <!--</tr>-->
    <!--</thead>-->
    <!--<tbody>-->
    <!--{{ range $drv := .Drivers }}-->
    <!--<tr>-->
    <!--<td>  {{$drv.Id}} </td>-->
    <!--<td> {{$drv.Name}} </td>-->
    <!--<td> {{ $drv.License_number }}</td>-->
    <!--</tr>-->
    <!--{{end}}-->
    <!--</tbody>-->
    <!--</table>-->
</head>
</html>

