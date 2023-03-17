# Unified Err Handling
Unify error handling methods
## example
```go
package example

import "fmt"

func ResErr(m string) error{
	return fmt.Errorf("%s",m)
}

func majorityPattern1() {
	if err := ResErr("test"); err != nil {
		fmt.Printf("%v",err)
	}
}

func majorityPattern2() {
	if err := ResErr("test"); err != nil {
		fmt.Printf("%v",err)
	}
}

func minorityPattern1() {
	err := ResErr("test")
	if err != nil { //want "separated notation"
		fmt.Printf("%v",err)
	}
}
```

# install
```sh
go install github.com/mercy34mercy/unifiedErrHandling/cmd/unifiedErrHandling@latest
```

## Useage
```sh
go vet -vettool=`which unifiedErrHandling` pkgname
```