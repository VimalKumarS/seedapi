# Go lang Seed API 

Project is quick start go lang api project

Project setup consist of
  - Common logging using (user/zap)
  - Common Error handler middleware
  - Pannic handler middleware
  - Authorization middleware  (Todo: need to update the code based on project need)
  - Grace shut down api
  - Appsetting config to setup for different enviornment
  - Modular build (using go mod)
  - Base controller (under folder 'controller')
  - Hearbeat endpoint  
  - Cors Enable
 
### Module 

* [gorilla/mux] - https://github.com/gorilla/mux
* [gorilla/handlers] - https://github.com/gorilla
* [pkg/errors] - github.com/pkg/errors
* [Logger] - go.uber.org/zap
* [filelogger] - gopkg.in/natefinch/lumberjack.v2
* [GOORM] - 

### Installation

Seed Api requires [Go lang](https://golang.org/dl/) v11.1 above to run.

```sh
$ go build
$ seedapi.exe
```

For dev environments using appsetting.dev.json

```sh
$ go build
$ seedapi.exe --env=dev
```

### Code Structure
 - All controller defined under controller folder (eg: SearchContrller) , using based controller struct type
 - Service interface defined under folder serviceinterface of type interface (eg: searchinterface.go)
 - Service implementation under service folder (eg: searchService.go)
 - Route for search controller are defined under route folder (eg: searchroute.go)
 - New route defined should be included in route.go under bundleRoute function
 ```
 	route = append(route, 
		(SearchRoutes(config, Logger))...)  // append all route
 ```
 - All model are defined under model folder
 - In Controller, all function definition should of type 
  ```
  type HttpHandlerWithError func(http.ResponseWriter, *http.Request) error
 ```
 - Error object , refering common custom error module
 ```
    var err error
	err = errors.New("New error created")
	err = errors.BadRequest.Wrapf(err, "Custom error or validation error")
	err = errors.AddErrorContext(err, "id", "wrong id format, should be an integer")
 ```
- Error logged in file appsetting config
```
"logging":{
        "filepath":".\\api-log\\api.log",
        "maxSize": 20  // size in MB for file rotation
    }
    
Sample error log with stack trace
{"level":"error","ts":1541372197.0582957,"msg":"Error","id1":"wrong id format, should be an integer","Error":"Custom error or validation error: New error created","stacktrace":"New error created\nCustom error or validation error\nseedapi/util.ErrorType.Wrapf\n\tC:/Dexter/Practice/goCoursera/SeedApi/util/errors.go:54\nseedapi/controller.(*SearchController).CustomError\n\tC:/Dexter/Practice/goCoursera/SeedApi/controller/searchcontroller.go:45\nseedapi/controller.(*SearchController).CustomError-fm\n\tC:/Dexter/Practice/goCoursera/SeedApi/route/searchroutes.go:38\nseedapi/middleware.ErrorHandlerMiddleware.func1\n\tC:/Dexter/Practice/goCoursera/SeedApi/middleware/errorhandler.go:15\nnet/http.HandlerFunc.ServeHTTP\n\tC:/Go/src/net/http/server.go:1964\nseedapi/middleware.PanicRecoveryMiddleware.func1\n\tC:/Dexter/Practice/goCoursera/SeedApi/middleware/panicrecovery.go:21\nnet/http.HandlerFunc.ServeHTTP\n\tC:/Go/src/net/http/server.go:1964\ngithub.com/gorilla/mux.(*Router).ServeHTTP\n\tC:/Users/vimal/go/pkg/mod/github.com/gorilla/mux@v1.6.2/mux.go:162\ngithub.com/gorilla/handlers.(*cors).ServeHTTP\n\tC:/Users/vimal/go/pkg/mod/github.com/gorilla/handlers@v1.4.0/cors.go:52\nnet/http.serverHandler.ServeHTTP\n\tC:/Go/src/net/http/server.go:2741\nnet/http.(*conn).serve\n\tC:/Go/src/net/http/server.go:1847\nruntime.goexit\n\tC:/Go/src/runtime/asm_amd64.s:1333"}
```
- Appsettting
```
{
    "port":"0.0.0.0:8000", // define port on which to expose
    "env":"",  // define enviornment
    "dbConnection":"connection string",  // define DB connection
    "enablecors":true,  // enable cors 
    "logging":{
        "filepath":".\\api-log\\api.log",  // log file path
        "maxSize": 20  // Max size of file 
    }
}
```

### Api expose part of seed api

```
 GET /search/getall 
 curl -X GET \
  http://localhost:8000/search/getall \
  -H 'auth-token: 123' \
  -H 'cache-control: no-cache' 
```
```
 GET /search/error 
 curl -X GET \
  http://localhost:8000/search/error \
  -H 'auth-token: 123' \
  -H 'cache-control: no-cache' 
```


### Todos
 - Go Docs / Go swagger
- Content negotiation

