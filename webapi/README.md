# Web API project

This folder will include standard Web API project code structure template.

## project-root

|- :folder: go.mod // Golang module file

|- main.go // entry file, main package

|- main_test.go // unit test for entry file

|= global //global folder contains all cross package shared code files

- |- const.go // this file contains all global constants definitions

- |- setting.go //[optional] this file contains all configuration public vairables and loader functions, if setting is not complex

- |- utils.go // this file contains all utility functions definitions

|= handler //handler folder contains all webapi hanlder code files

- |- common.go // this file contains all common local functions and vairables in the package

- |- const.go // this file contains all local constants in the package

- |- model.go // this file contains all model type definitions in the handler package

- |- home.go // this file contains HomeHandler struct declaration, NewHomeHandler function and IndexAction method (route path will be definited in a const RoutePath_Home_Index "/")

- |- hello.go // this file contains HelloHandler struct declaration and NewHelloHandler function, also contains some common func of HelloHandler

- |- hello_index.go // this file contains IndexAction method of HelloHandler (route path will be definited in a const RoutePath_Hello_Index "/hello")

- |- hello_say.go // this file contains SayAction method of HelloHandler (route path will be definited in a const RoutePath_Hello_Say "/hello/say")

- |- hello_request.go // this file contains some functions related to request processing

- |- hello_response.go // this file contains some functions related to response processing

|= manager //manager folder contains all webapi biz logic code files that the handler depend on

- |- common.go // this file contains all common local functions and vairables in the package

- |- const.go // this file contains all local constants in the package

- |- interface.go // all interface definitions of managers

- |- say.go // a sample biz logic for access token process

|= dbaccess //dbaccess folder contains all database access code files that the manager depend on

- |- common.go // this file contains all common local functions and vairables in the package

- |- const.go // this file contains all local constants in the package

- |- interface.go // all interface definitions of dbaccess

- |- user.go // a sample database access for user data

|= middleware //middleware folder contains all webapi middleware code files

- |- common.go // this file contains all common local functions and vairables in the package

- |- const.go // this file contains all local constants in the package

- |- logger.go // a sample middleware for log all request, with -er/-or postfix

|= service //service folder contains the client code for all the external services that the handler, manager and middleware depend on

- |- common.go // this file contains all common local functions and vairables in the package

- |- const.go // this file contains all local constants in the package

- |- interface.go // all interface definitions of service clients

- |- auth.go // a sample service client for auth

|= setting //setting folder contains all code files related to configuration or setting

- |- common.go // this file contains all common local functions and vairables in the package

- |- config.go // this file contains all public vairables for config

- |- const.go // this file contains all constants in the package

- |- loader.go // this file contains load config public functions

## model file

If there are some model types, you can create a model.go or xxx_model.go (ex: auth_model.go, user_model.go) and put it into specific folder.

## dependency

### handler

- -> manager

- -> service

- -> setting

- -> global

### manager

- -> dbaccess

- -> setting

- -> global

### service

- -> setting

- -> global

### dbaccess

- -> setting

- -> global

### setting

- -> global

### middleware

- -> manager

- -> service

- -> setting

- -> global
