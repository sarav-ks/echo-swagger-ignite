This is a POC project to show the integration between below go libaries 

**echo** (https://github.com/labstack/echo) : Light wieght Web API framework
**swagger** (https://github.com/deepmap/oapi-codegen) : Echo swagger code generator 
**ignite client** (https://github.com/amsokol/ignite-go-client) : Ignite client for GoLang
**ignite server** running in Docker 

After installing the above go packages , follow the below steps to run the project

1) Generate model objects using 
oapi-codegen --package=model --generate types -o model/gwapp-types.gen.go gwapp.yaml

2) Generate server side code 
oapi-codegen --package=api --generate server,spec -o api/gwapp-server.gen.go gwapp.yaml

3) Write Ignite Store for storing the objects in cache
see [ignite-store.go](ignite-store.go) for example

4) Create the main class to register the handlers 
see [gwapp.go](gwapp.go) for example


**Deploying ignite cluster in OpenShift** 

1) Create PV (gwapp-pv.yaml) & PVC (gwapp-pvc.yaml) to store ignite configs

    

      
