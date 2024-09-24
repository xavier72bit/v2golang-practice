/*
在tiny_server_04中，返回的李萨如图形的参数是固定的
在这个例子里，我们支持通过表单传递生成李萨如图形的参数，并且尝试使用结构体来描述李萨如图形的生成参数
*/
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand/v2"
	"net/http"
	"strconv"
)

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0
	blackIndex = 1
)

type lissajousConfigS struct {
	cycles  int
	res     float64
	size    int
	nframes int
	delay   int
}

/*
form query:
cycles: X轴振荡器的完整旋转次数
res: 分辨率，越小越清晰
size: 画布大小
nframes: 动画帧数
delay: 帧之间的延迟，以10ms为单位
*/
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// 初始化一个默认的李萨如图形配置
		lissajousConfig := lissajousConfigS{5, 0.0001, 200, 60, 1}

		// 通过表单获取配置并覆盖
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "解析表单失败: %s <br>", err)
		} else {
			// 这个写法是，从Form中寻找cycles，如果存在则覆盖lissajousConfig中对应的字段
			if values, ok := r.Form["cycles"]; ok {
				// 处理values
				value := values[0]
				// 处理value类型转换，没问题，写入lissajousConfig
				if resulr, err := strconv.Atoi(value); err == nil {
					lissajousConfig.cycles = resulr
				}
			}
			if values, ok := r.Form["res"]; ok {
				value := values[0]
				if result, err := strconv.ParseFloat(value, 64); err == nil {
					lissajousConfig.res = result
				}
			}
			if values, ok := r.Form["size"]; ok {
				value := values[0]
				if result, err := strconv.Atoi(value); err == nil {
					lissajousConfig.size = result
				}
			}
			if values, ok := r.Form["nframes"]; ok {
				value := values[0]
				if result, err := strconv.Atoi(value); err == nil {
					lissajousConfig.nframes = result
				}
			}
			if values, ok := r.Form["delay"]; ok {
				value := values[0]
				if result, err := strconv.Atoi(value); err == nil {
					lissajousConfig.delay = result
				}
			}
		}

		// 生成李萨如图形
		lissajous(w, &lissajousConfig)
	})

	log.Fatal(http.ListenAndServe("localhost:6578", nil))
}

func lissajous(out io.Writer, config *lissajousConfigS) {
	cycles := config.cycles
	res := config.res
	size := config.size
	nframes := config.nframes
	delay := config.delay

	freqent := rand.Float64() * 3.0
	anime := gif.GIF{LoopCount: nframes}

	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(cycles*2)*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freqent + phase)
			img.SetColorIndex(size+int(x*float64(size)+0.5), size+int(y*float64(size)+0.5), blackIndex)
		}
		phase += 0.1
		anime.Delay = append(anime.Delay, delay)
		anime.Image = append(anime.Image, img)
	}
	gif.EncodeAll(out, &anime)
}
