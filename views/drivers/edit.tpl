<!DOCTYPE html>

<html>
<head>
    <title>Beego</title>
    This is drivers Edit.tpl
<br/><br/><br/><br/>
    <form method="post">
        <div>
           ID: {{.Driver.Id}}
            <input id="Id" type='number' name="Id" hidden="hidden" value="{{.Driver.Id}}">
        </div>
        <div>
            <label for="Name">Name:</label>
            <input id="Name" type='text' name="Name" value="{{.Driver.Name}}">
        </div>
        <div>

            <label for="License_number">License:</label>
            <input id="License_number" type='text' name="License_number" value="{{.Driver.License_number}}">
        </div>

        <button type="submit">Submit</button>
    </form>
    <a href="/drivers/delete/{{.Driver.Id}}">Delete</a>
</head>
</html>

