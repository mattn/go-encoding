# go-encoding

[![Build Status](https://travis-ci.org/mattn/go-encoding.png?branch=master)](https://travis-ci.org/mattn/go-encoding)
[![GoDoc](https://godoc.org/github.com/mattn/go-encoding?status.svg)](http://godoc.org/github.com/mattn/go-encoding)
[![Go Report Card](https://goreportcard.com/badge/github.com/mattn/go-encoding)](https://goreportcard.com/report/github.com/mattn/go-encoding)

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
