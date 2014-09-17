webp
=====

PkgDoc: [http://godoc.org/github.com/chai2010/webp](http://godoc.org/github.com/chai2010/webp)

Install
=======

Install `GCC` or `MinGW` ([download here](http://tdm-gcc.tdragon.net/download)) at first,
and then run these commands:

1. `go get github.com/chai2010/webp`
2. `go run hello.go`

Example
=======

	package main

	import (
		"fmt"
		"io/ioutil"
		"log"

		"github.com/chai2010/webp"
	)

	func main() {
		var width, height int
		var data []byte
		var err error

		// Load file data
		if data, err = ioutil.ReadFile("./testdata/1_webp_ll.webp"); err != nil {
			log.Fatal(err)
		}

		// GetInfo
		if width, height, _, err = webp.GetInfo(data); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("image size: width = %d, height = %d\n", width, height)

		// Decode webp
		rgba, err := webp.DecodeRGBA(data)
		if err != nil {
			log.Fatal(err)
		}

		// Encode lossless webp
		if data, err = webp.EncodeLosslessRGBA(rgba); err != nil {
			log.Fatalf("saveWebp: webp.EncodeLosslessRGBA, err = %v", err)
		}
		if err = ioutil.WriteFile("output.webp", data, 0666); err != nil {
			log.Fatalf("saveWebp: ioutil.WriteFile, err = %v", err)
		}
		fmt.Printf("Save output.webp ok\n")
	}


BUGS
====

Report bugs to <chaishushan@gmail.com>.

Thanks!
