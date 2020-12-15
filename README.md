
Install
=======

Install `GCC` or `MinGW` ([download here](http://tdm-gcc.tdragon.net/download)) at first,
and then run these commands:

1. `go get github.com/hollson/webp`
2. `go run hello.go`


Example
=======

This is a simple example:

```go
package main

import (
    "io/ioutil"
    "log"
    "os"

    "github.com/hollson/webp/simple"
)

func main() {
    // png ==> webp
    img, err := os.Open("./asset/a.png")
    if err != nil {
        return
    }

    res, err := simple.Convert2Webp(img, 30)
    if err != nil {
        return
    }
    ioutil.WriteFile("./asset/a.webp", res, 0666)

    // webp ==> png
    source, err := os.Open("./asset/a.webp")
    if err != nil {
        log.Fatalln(err)
    }
    data, err := simple.Webp2Others(source, "png", 90)
    if err != nil {
        log.Fatalln(err)
    }

    wp, err := os.Create("./asset/b.png")
    if err != nil {
        log.Fatalln(err)
    }
    wp.Write(data)
}
```

Decode and Encode as RGB format:

```Go
m, err := webp.DecodeRGB(data)
if err != nil {
	log.Fatal(err)
}

data, err := webp.EncodeRGB(m)
if err != nil {
	log.Fatal(err)
}
```

Notes
=====

Change the libwebp to fast method:

	internal/libwebp/src/enc/config.c
	WebPConfigInitInternal
	config->method = 0; // 4;

REF
====
https://developers.google.com/speed/webp/docs/api

