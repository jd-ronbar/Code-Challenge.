Code-Challenge:

Goal: Create a simple web application to which you can import drivers and metrics and add a rest api to view the imported metrics. 

Getting Started:

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. 


Prerequisites:

1 - Create schema (MySql) - Location: (db.txt)
2 - App conf - Location: /src/hello/conf
appname = hello
httpport = 8080
runmode = dev
copyrequestbody = true
3) for your convenience the input files - location: (/src/hello/input)

A step by step series of examples that tell you have to get a development env running:

1 - Build a data model for the metrics and driver - Location: (/src/hello/models)
2 - Build simple GO program to import the drivers’data and metrics into the datastore - Location: (/src/learning)
3 - Build simple CRUD API for drivers and metrics with Beego framework - Location: ( view - /views, model,controller,router - /src/hello)


Running the tests:

Created 1 test (insert new driver) - Location: (/src/hello/tests)
