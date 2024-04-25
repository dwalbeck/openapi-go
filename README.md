# Golang OpenAPI POC

### Overview
This covers various methods for implementing OpenAPI 3 (a.k.a. Swagger) for documenting 
API end point using the swagger spec file to generate boilerplate Go code for the server.
Packages that were reviewed include the following:
* oapi-codegen
* ogen-go
* openapi-generator
* openapi-generator-go
* go-swagger
* swaggo

Each of these packages (except swaggo) will take a defined swagger spec file and can 
generate from that spec the boilerplate code for both client and server.  Swaggo uses 
annotations in the code to generate a spec file, which makes it a less than desirable 
method to use, as it does not insure the code and the docs operate the same.  Some packages 
generated a lot of code and others less or not enough code.  Some only support up to 
version 2 and others support of some functionality was hit or miss.  

Ultimately the best choice with the cleanest code, supporting features and capable of using 
version 3 syntax for the spec file, I came to the conclusion of **oapi-codegen**, which will 
the focus of this POC.

### Quick View Steps
* go init
* create server.cfg.yaml
* create generate.go file
* go get github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen
* mkdir api
* go generate
* go mod tidy
* create api_func.go
* create main.go
* go get github.com/oapi-codegen/nethttp-middleware

### Generating the Code
So to start with, you'll need to create a swagger specification file, which can be formatted 
using either JSON or Yaml.  I personally prefer the Yaml format, as it's cleaner and easier to 
read than JSON, due to all the closing curly brackets used by JSON.  The spec file can be 
either 2.0 or 3.0, of which there is a pretty big difference not only with the syntax of the 
file, but also in capabilities.  A lot of improvements were made in version 3 and it's release 
was anxiously waited on for many years until it was released.

With the spec file ready to go, there is a bit of setup that makes it easier to use, the first 
of which is creating a configuration file that defines how the generator should behave. Create 
a Yaml file (it can be named anything, I'll use server.cfg.yaml) and add the following:
```bash
package: api
output: api/api.go
generate:
  embedded-spec: true
  models: true
  chi-server: true
compatibility:
  apply-chi-middleware-first-to-last: true
```
There are other options available, but this is the nitty gitty of what you want.  The **output** 
defines the directory the code will be placed and **package** is the name of the package that 
the code will be part of.  The **embedded-spec** option is whether to include a code version of 
spec file (which you will want) and **models** defines whether to generate models, which is 
absolutely needed.  I've picked to use the **chi-server**, but it also supports other web 
server packages like:  mux and http

Next we create the **generate.go** file, which enables an easy and repeatable way to generate 
the server code.  It also makes things work without having to install a binary file of the 
generator
```bash
package main

//go:generate go run github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen --config server.cfg.yaml swagger.yaml
```
I don't believe that you'll need the following when you execute **go generate**, but it doesn't 
hurt to add it now and it will be needed later anyway, so execute the following to load the main 
library module:
```bash
go get github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen
```
Then create the **api** directory where the code will be generated to, followed by the call 
to actually generate the boilerplate server code.  Additionally, you'll need to install all 
the modules that the generated code now requires, so running **go mod tidy** will install those.
```bash
mkdir ./api

go generate

go mod tidy
```
There is also an alternative way to generate the code, which is to either globally install 
the module, so that it can be called from the command line or download the executable and make 
a call to it directly.  When calling it this way, you need to pass specific values after as flags 
like the following:
```bash
oapi-codegen -generate chi-server,types,spec -config server.cfg.yaml -o api swagger.yaml
```

### Wiring Up the Generated Code
After generating the code, you should have a file in the **api** called api.go.  This file will 
contain all the models (structs), functions and interfaces needed to execute enforcement of 
the swagger spec for all the end points contained in it.  The most important part of that file 
is with a defined interface called **ServerInterface**, as this is how we will wire everything 
to work.  So create a file called **api_func.go** within the **api** directory.  Note that 
you'll create this file with an interface that has all the same functions linked to it as the 
**ServerInterface** mentioned before.  They must match perfectly in order for Golang to treat 
them both as if they were the same type.
