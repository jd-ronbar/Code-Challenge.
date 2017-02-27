<!DOCTYPE html>

<html>
<head>
    <title>Beego</title>
    This is driversInsert.tpl

    <form action="/drivers/add" method="post">
        <div>
            <label for="Id">Id:</label>
            <input id="Id" type='number' name="Id">
        </div>
        <div>
            <label for="Name">Name:</label>
            <input id="Name" type='text' name="Name">
        </div>
        <div>

            <label for="License_number">License:</label>
            <input id="License_number" type='text' name="License_number">
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

