# Notes

Go aims to be simple to write, read and understand. It does this very well, but one thing that I think goes against it is that it has 4 different ways to declare and initialize variables:

```golang
var x string
var x string = ""
var x = ""
x := ""
```

Maps intentionally don't guarant any order.

`fallthrough` may be used in a switch's case block to transfer control to the next case even though the current case might have matched.

[`iota`](https://yourbasic.org/golang/iota/) is awesome!

`fmt.Println` will use the `String()` method if the type has one.  
e.g
```golang
func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }
...
fmt.Println(c) // prints: "100°C"
```

## Questions

*Q1: Why do we need the `make` command to create a map? Why can't we just declare it as a slice e.g `x := []string{}`. Same thing goes for channels or any other data structure that needs to use `make` to be initialized.*

A1: TODO

*Q2: Related with latest question, go has several ways of memory allocation and value initialization. What's the difference beyween initializing with `make`, `new` or just with regular initialization methods?*

A2: TODO
