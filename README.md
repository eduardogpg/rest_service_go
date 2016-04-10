# REST service With GO

This is a REST services builded with Go An some library, Some features.

  - Connection with a database
  - reading environment variables
  - Management of routes with params
  - Handle requests and the response in JSON Format.

### Version
Only version

### External libraries

Dillinger uses a number of open source projects to work properly:

* [martini] - Routes
* [mysql] - Driver SQL
* [gorila] - Gorira Web Socket

### Installation

Is necesary create the table user with the user.sql file in the repo.
then just need

```sh
$ go run server.go
```

go to http://localhost:8000 and visualize the result

### Test
Not in all the case you can check the result for these case:

for check the correct operation you can use the test.go file, in this file you can check the response of a some functions, delete, put, push, get.

for this is necessary to change the following lines:

* 9  (What function you can test)
* 12 (Json to send)
* 14 (For the type por request)

Or Change the type of the method in service.go all in Get.

In another terminar run
```sh
$ go run test.go
```

### Contact

* [@eduardo_gpg]
----



   [@eduardo_gpg]: <https://twitter.com/eduardo_gpg>
   [express]: <http://expressjs.com>
   [martini]: <https://github.com/go-martini/martini>
   [mysql]: <https://github.com/go-sql-driver/mysql>
   [gorila]: <http://www.gorillatoolkit.org/pkg/websocket>
   
   
   

