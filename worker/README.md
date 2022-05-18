# Worker project

This folder will include standard Worker project code structure template.

## project-root

|- :folder: go.mod // Golang module file

|- main.go // entry file, main package

|- main_test.go // unit test for entry file

|= global //global folder contains all cross package shared code files

- |- const.go // this file contains all global constants definitions

- |- setting.go //[optional] this file contains all configuration public vairables and loader functions, if setting is not complex

- |- utils.go // this file contains all utility functions definitions

|= worker //handler folder contains all webapi hanlder code files

- |- common.go // this file contains all common local functions and vairables in the package

- |- const.go // this file contains all local constants in the package

- |- model.go // this file contains all model type definitions in the handler package

- |- queue.go // a sample worker, this file name should indicate worker purpose, for example "queue" will process queue message, it will contains QueueWorker struct declaration and its methods and functions

- |- processor.go // If process logic is very complex, can add one or mutiple processor files contains process logic code

- |- interface.go // all interface definitions of processors

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

|= handler // [optional] handler folder contains all webapi hanlder code files like webapi project

## model file

If there are some model types, you can create a model.go or xxx_model.go (ex: auth_model.go, user_model.go) and put it into specific folder.

## dependency

### worker

- -> manager

- -> service

- -> setting

- -> global

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

### setting

- -> global