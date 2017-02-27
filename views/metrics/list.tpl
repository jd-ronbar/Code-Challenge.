<!DOCTYPE html>

<html>
<head>
    <title>Beego</title>
    This is metricsList.tpl
    <br/><br/><br/>
    <a href="/metrics/add"> <input type="button" value="Create new Metric" /></a>
    <table border=1>
        <thead>
        <tr>
            <td>ID</td>
            <td>Metric_name</td>
            <td>Value</td>
            <td>Lon</td>
            <td>Timestamp</td>
            <td>Lat</td>
            <td>Driver_id</td>
        </tr>
        </thead>
        <tbody>
        {{ range $mtr := .Metrics }}
        <tr>
            <td> <a href="/metrics/edit/{{$mtr.Id}}"> {{$mtr.Id}}</a> </td>
            <td> {{$mtr.Metric_name}} </td>
            <td> {{$mtr.Value}}</td>
	    <td> {{$mtr.Lon}} </td>
            <td> {{$mtr.Timestamp}} </td>
            <td> {{$mtr.Lat}} </td>
            <td> {{$mtr.Driver_id}} </td>
        </tr>
        {{end}}
        </tbody>
    </table>
</head>
</html>

