Go has built in testing package: https://pkg.go.dev/testing

Naming Conventions 
- Test File: `<packageName>_test.go`
- Test functions: `func TestHelloName`

To write a test: 
1. Import the test package 
2. Add in the testing package as a parameter to the test function 

```go
func TestHelloName(t *testing.T) {  
    name := "Alice"  
    want := regexp.MustCompile(`\b` + name + `\b`)  
    msg, err := Hello("Alice")  
    if !want.MatchString(msg) || err != nil {  
       t.Fatalf(`Hello("Alice") = %q, %v, want match for %#q, nil`, msg, err, want)  
    }  
}
```

Run the Tests: Be in the directory 
- `go test`
- `go test -v`
- `go test help`

