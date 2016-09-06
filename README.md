# go-encoding

Collection of encodings.

## Usage

```go
for _, n := range encoding.EncodingNames() {
    println(n)
}
```

```go
e := encoding.GetEncoding("cp932")
if e == nil {
    panic("unknown encoding")
}
```

This return pointer of golang.org/x/text/encoding#Encoding.

## Installation

```
go get github.com/mattn/go-encoding
```

## License

MIT

## Author

Yasuhiro Matsumoto (a.k.a. mattn)
