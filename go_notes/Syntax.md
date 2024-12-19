## Functions 
- Capitalization Rules: 
	- Capitalized function name -- can be called by a function not in the same function = "Exported Name"
	- lowercase function name -- can only be called by a function within the same package 
- 
```go

// declare `greetings` package to collect related functions 
package greetings  
  
import "fmt"
  
// Hello returns a greeting for the named person.

func Hello(name string) string {  
    // Return a greeting that embeds the name in a message.  
    message := fmt.Sprintf("Hi, %v. Welcome!", name)  
    //using the Sprintf fn that is part of fmt package: https://pkg.go.dev/fmt#Sprintf
    return message  
}
```



["fmt" package, part of Go Standard Library](https://pkg.go.dev/fmt)
- `message := fmt.Sprintf("Hi, %v. Welcome!", name)`
	- Note: `%v` is standard in Go for representing the value 
	- Can't just use any random letter, since there are ones that are reserved -- eg: `%x` will give the letters as base16 integers 

## Variable Declaration 
- `:=` is a shortcut for declaring and initializing variabe in 1 line 
	- this can only be done within a function 
	- The value on the right determines the variable's type 

## Data Types 

### Default Zero Values 

| Data Type                                               | Zero Value |
| ------------------------------------------------------- | ---------- |
| integer                                                 | 0          |
| floating point                                          | 0.0        |
| boolean                                                 | false      |
| string                                                  | ""         |
| interfaces, slices, channels, maps, pointers, functions | nil        |

### Maps 
-  `map[KeyType]ValueType`

**Need to always Initialize Map**
- During map initialization, use the `make` function, which will allocate and initialize a hash map data structure and returns a map value that points to it
- **WARNING:** WITHOUT initializing, attempts to WRITE to a map will cause runtime panic 
- Initialize a map using `make`:   
	- `m = make(map[string]int)`

**Setting Map Value**
- `m["route"] = 66`

**Retrieves the value stored in the key and assigns it to a new variable**
- `i := m["route"]`
	- NOTE: if the key doesn't exist, it will return the map's key-value's type *default zero value*
		- eg: for 'int' type, it is 0 


## For Loops

==12/14/24 -- Leaving off here: https://go.dev/doc/tutorial/greetings-multiple-people==

```go

// Hellos returns a map that associates ea of the named ppl with a greeting msgfunc 
Hellos(names []string) (map[string]string, error) {  
    //given an array of names  
    //return a map where keys are strings and keyvalues are strings 
     
    //declare a map, named messages to assoc. names with messages    
    messages := make(map[string]string)  
  
    //for each name of all names, call Hello fn on the name  
    for _, name := range names {  
       message, err := Hello(name)  
       //if Hello fn doesn't return err of nil, return nil and err msg  
       if err != nil {  
          return nil, err  
       }  
       //else, set messages[name] = to the msg  
       messages[name] = message  
    }  
    //return the map of messages, with err of nil  
    return messages, nil  
}
```


# Errors 

1. import `errors` package 
2. add the error as a return type 

 **Defining the Errors** 
```go
// Hello returns a greeting for the named person.

func Hello(name string) (string, error) {  
    //fn returns 2 values: string AND error  
  
    if name == "" {  
       return "", errors.New("empty name")  
    }  
    // Return a greeting that embeds the name in a message.  
    message := fmt.Sprintf("Hi, %v. Welcome!", name)  
  
    //nil = no error, indicating a success return  
    return message, nil  
}
```

**Handling the Errors** 
```go
// declare a main package  
// a package is way to group functions  
// it is made up of all the files in the same directory  
package main  
  
import (  
    "example.com/greetings"  
    "fmt"    
    "log"
    )  
  
func main() {  
    log.SetPrefix("greetings: ") //sets the log entry prefix  
    log.SetFlags(0)              //this disables printing time, source file, and line number  
  
    message, err := greetings.Hello("")  
    if err != nil {  
       //if error returned, prints to console and exit program  
       log.Fatal(err)  
    }  
  
    //else, prints the message to console  
    fmt.Println(message)  
}
```


## Operators 

```
!= 
== 

|| 
&& 

nil


```