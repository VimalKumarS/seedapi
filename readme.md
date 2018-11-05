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
- 
