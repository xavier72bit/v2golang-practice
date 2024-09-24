/*
之前的3个例子都是返回的字符串内容，这最后一个例子我们返回点别的东西，在1-4MakeGif中，我们输出了一个李萨如图形
这个例子里我们将一个gif传回用户浏览器上
*/
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand/v2"
	"net/http"
)

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0
	blackIndex = 1
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		lissajous(w)
	})
	log.Fatal(http.ListenAndServe("localhost:6578", nil))
}

// 把1-4MakeGif的lissajous_black_and_white_01.go中的lissajous函数copy过来
func lissajous(out io.Writer) {
	const (
		cycles  = 5
		res     = 0.0001
		size    = 200
		nframes = 60
		delay   = 1
	)

	freqent := rand.Float64() * 3.0
	anime := gif.GIF{LoopCount: nframes}

	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freqent + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex)
		}
		phase += 0.1
		anime.Delay = append(anime.Delay, delay)
		anime.Image = append(anime.Image, img)
	}
	gif.EncodeAll(out, &anime)
}
