## What is a Package?
* A focused area of functionality
* Most Go programs are broken into multiple packages
* Each directory in a go program is associated with a unique package

## How to use a Package?
* We use packages by importing them
```Go
import (
    "fmt"  // package1
    "strings" //package2
)
// How we use the functions/variables contained in a imported package
// packagename.function()or variable
fmt.Println("This is how we use the functions from imported packages")
```

* In Go programming language packages can be define in three categories. They are
    1. Own packages (Packages you create)
    2. Third-party packages (Packages created by someone)
    3. Packages that are in Go standard library

* In a package if we want to make variable/function public then the first letter should be capitalized
* Using relative path for importing local packages is considered bad practice

### Resources
* [Golang Packages](https://golang.org/pkg)
* [awesom-go](https://github.com/avelino/awesome-go)
* [Dependency management](https://golang.github.io/dep/)