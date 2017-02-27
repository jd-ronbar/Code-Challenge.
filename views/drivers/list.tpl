<!DOCTYPE html>

<html>
<head>
    <title>Beego</title>
    This is driversList.tpl
    <br/><br/><br/>
    <a href="/drivers/add"> <input type="button" value="Create new Driver" /></a>
    <table>
        <thead>
        <tr>
            <td>Id</td>
            <td>Name</td>
            <td>License</td>
        </tr>
        </thead>
        <tbody>
        {{ range $drv := .Drivers }}
        <tr>
            <td> <a href="/drivers/edit/{{$drv.Id}}"> {{$drv.Id}}</a> </td>
            <td> {{$drv.Name}} </td>
            <td> {{ $drv.License_number }}</td>
        </tr>
        {{end}}
        </tbody>
    </table>
</head>
</html>

