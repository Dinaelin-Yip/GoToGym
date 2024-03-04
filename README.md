## Beginning

# Building local environment

To create a virtual environment by python,

```shell
$ python -m venv .venv
$ source .venv/bin/activate
$ pip install -r requirements.txt
```

To install go dependencies,

```shell
$ go mod tidy
```

# Variable Declaration

To create a virtual environment by python,

```go
var x int = 3 // Explicit type definition: global or local
var x = 3 // Implicit type definition: sys infers that x is an integer
x := 3 // local only; sys infers that x is an integer

var (
    a = 3
    b float = 4.4
)

const var c string = "Runoob" // cannot be changed
var a, b int = 1, 2
```

`:=` is the short variable declaration operator.

use `" "` instead of `' '`.


# Print

```go
func func2() {
	name := "Tony Yip"
	num := 1004
	fmt.Print(name, "'s brithday is", num) 
	fmt.Printf("1\n") // no space, no newline
	fmt.Printf("%s 's brithday is %d", name, num)
	fmt.Printf("1\n") // no space, no newline, formatted output
	fmt.Println(name, "'s brithday is", num)
	fmt.Printf("1\n") // space, newline
}
```

the output is
```
Tony Yip's brithday is10041
Tony Yip 's brithday is 10041
Tony Yip 's brithday is 1004
1
```
