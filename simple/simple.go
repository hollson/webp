// Copyright 2020 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package simple

import (
    "bytes"
    "fmt"
    "image"
    "image/gif"
    "image/jpeg"
    "image/png"
    "io"
    "io/ioutil"
    "net/http"
    "strings"

    "github.com/hollson/webp"
    // "golang.org/x/image/bmp"
)

// 将webp转换成其他格式
//  reader: 源文件，仅支持：webp
//  quality: 0 ~ 100
//
// Usage example:
//
// func main() {
//    source, err := os.Open("./asset/a.webp")
//    if err != nil {
//        log.Fatalln(err)
//    }
//    data, err := Webp2Others(source, "png", 90)
//    if err != nil {
//        log.Fatalln(err)
//    }
//
//    wp, err := os.Create("./asset/b.png")
//    if err != nil {
//        log.Fatalln(err)
//    }
//    wp.Write(data)
// }
func Webp2Others(reader io.Reader, targetType string, quality int) ([]byte, error) {
    data, err := ioutil.ReadAll(reader)
    if err != nil || len(data) < 512 {
        return nil, fmt.Errorf("source data is incomplete")
    }
    contentType := strings.ToLower(http.DetectContentType(data[:512]))
    if !strings.Contains(contentType, "webp") {
        return nil, fmt.Errorf("only support webp file")
    }
    t, err := webp.Decode(bytes.NewReader(data))
    if err != nil {
        return nil, fmt.Errorf("%v", err)
    }
    var buf = new(bytes.Buffer)
    switch strings.ToLower(targetType) {
    case "jpg", "jpeg":
        jpeg.Encode(buf, t, &jpeg.Options{quality})
    case "png":
        png.Encode(buf, t)
    // case "bmp":
    //     bmp.Encode(buf, t)
    case "gif":
        gif.Encode(buf, t, &gif.Options{NumColors: 256})
    }
    if buf == nil {
        return nil, fmt.Errorf("only jpeg、png、bmp、gif be supported")
    }

    return buf.Bytes(), nil
}

// 将图片转换成webp格式
//  reader: 源文件，支持：jpeg、png、bmp、gif
//  quality: 0 ~ 100
//
// Usage example:
//
// func main() {
//    img,err:= os.Open("./asset/a.png")
//    if err!=nil{
//        return
//    }
//
//    res,err:=Convert2Webp(img,30)
//    if err!=nil{
//        return
//    }
//    ioutil.WriteFile("./asset/a.webp",res,0666)
// }
func Convert2Webp(reader io.Reader, quality float32) (target []byte, err error) {
    var (
        buf bytes.Buffer
        img image.Image
    )
    data, err := ioutil.ReadAll(reader)
    if err != nil || len(data) < 512 {
        return nil, fmt.Errorf("source data is incomplete")
    }

    contentType := http.DetectContentType(data[:512])
    switch {
    case strings.Contains(contentType, "jpeg"):
        img, _ = jpeg.Decode(bytes.NewReader(data))
    case strings.Contains(contentType, "png"):
        img, _ = png.Decode(bytes.NewReader(data))
    // case strings.Contains(contentType, "bmp"):
    //     img, _ = bmp.Decode(bytes.NewReader(data))
    case strings.Contains(contentType, "gif"):
        img, _ = gif.Decode(bytes.NewReader(data))
    }

    if img == nil {
        return nil, fmt.Errorf("not image file or type not supported")
    }

    if err = webp.Encode(&buf, img, &webp.Options{Lossless: false, Quality: quality}); err != nil {
        return
    }
    return buf.Bytes(), nil
}