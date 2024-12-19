```shell
go run [path of file]

go mod tidy

```


# Basics 
Dependency Tracking 
- when code imports packages containing other modules, those dependencies are managed through code's own modules 
	- managed within `go.mod` file 

**Running a command** 
`go run [path of file]`

**Importing packages** 
- Use import command to import the package 
- Run the tidy command: `go mod tidy` -- imports the package
- Run the file 

At compass, need, go is run via the CLI. Need to always append: 
`./scripts/go` and then the command (run, help, etc)
- `./scripts/go run compass.com/workshop/cmd/workshopd`
- `./scripts/go help`


```go
package main  
  
import "fmt"  
  
func main() {  
    fmt.Println("Hello, World!")  
}
```


## Terms 
- **Module** 
	- Module can contain 1 or more related  packages
	- as more functionality is added to the module, new versions of the module are published 
	-  Developers writing code that call functions within the module can import the module's updated packages 
- **Package** 
	- Packages contain a set of functions 
	- eg: a module with packages that have functions related to financial calculations 

## Go Code Hierarchy 
- Go Code is grouped into packages 
- Packages are grouped into Modules 
- Modules specify dependencies needed to run your code 
	- This can include the Go version 
	- And other modules it requires 


### Creating a module 
`go mod init example.com/greetings`
- `go mod init` creates a go.mod file to track the code dependencies 
- The module is created with reference to `example.com/greetings`
	- This is for production 
	- During development, want to point to the package on local 

### Pointing to package on local 
- This is needed when importing a package during development, and don't want to be pointing to the package on production 

```go
package main  
  
import (  
    "example.com/greetings"  //this is pointing to prod 
    "fmt")  
  
func main() {  
    message := greetings.Hello("Alice")  
    fmt.Println(message)  
}
```

- To point to local file, need to run these commands: 
	- `go mod edit -replace example.com/greetings=../greetings
		- This will add a line to the go.mod file
			- `replace example.com/greetings => ../greetings`
	- `go mod tidy`
		- This will add `require example.com/greetings v0.0.0-00010101000000-000000000000`
			- #'s at the end = version number 
	- 

Leaving off here: https://go.dev/doc/tutorial/handle-errors 


