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

Some cool `fmt` tricks and adverbs we can use:
```golang
x := 0x000f
fmt.Printf("%d %#[1]x %[1]X %#08[1]b", x)
// 15 0xf F 00001111
```
`[1]` tells fmt to use the index argument 1 (x), avoiding the need to repeat x in the arguments list of `Printf`  
`%` indicates an adverb (formatting)
`#` tells `Printf` to print the a leading identifier of the value's type (e.g `0x` for hexadecimal)
`%08b` prints in binary format with 8 bits

`range` performs utf8 decoding implicitly. This means we can iterate through runes (who can occupy more than one byte for non-ASCII) without the hassle to know how many bytes a rune has.

Slices yeld 3 properties: a pointer (to the underlying array), a lenght (current number of elements) and a capacity (number of elements that can be added to the underlying array).

If we know the required size for the underlying array (in slices) up front, it is more efficient to allocate an array of the required size, avoiding reallocations.
```go
ages := make(map[string]int)
...
names := make([]string, 0, len(ages)) // <--
```

Maps will return the zero value for the type they store when you try to access a key that doesn't exist. To know if a key exists inside
the map, use:
```go
age, ok := ages["bob"]
if !ok {
    // bob is not a key in the map
}
```

## Questions

*Q1: Why do we need the `make` command to create a map? Why can't we just declare it as a slice e.g `x := []string{}`. Same thing goes for channels or any other data structure that needs to use `make` to be initialized.*

A1: TODO

*Q2: Related with latest question, go has several ways of memory allocation and value initialization. What's the difference beyween initializing with `make`, `new` or just with regular initialization methods?*

A2: TODO

*Q3: When using a map, Go forces you to allocate the map before appending elements. Why this is true for maps and not for slices? Why
not force you to allocate both, or none (and do the magic of allocations behind the scenes).*

A3: TODO
One possibly reason would be that when appending to a slice we use the `append` built it function, but when appending to a map
we use direct allocation e.g `ages["carol"] = 21`. But I can dig deep into the why, I believe.

## TODO

- Revisit 2.3, 2.4 and 2.5 to learn more about bit operations.
- I've stopped doing the exercises, as it requires to have a computer around while reading. I'll do them all at once when I finish
reading the book.
- Read more about UTF-8 enconding and Unicode.