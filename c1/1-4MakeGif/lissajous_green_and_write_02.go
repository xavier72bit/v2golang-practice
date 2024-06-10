/*
在lissajous_black_and_white_01的基础上，将绘制图形的黑色像素点换成绿色像素点。
*/
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"time"
)

// 定义调色板
var palette = []color.Color{color.White, color.RGBA{0x00, 0xff, 0x00, 0xff}}

// 为访问调色板的颜色index，定义一个可读性良好的名字
const (
	whiteIndex = 0
	greenIndex = 1
)

func main() {
	// 定义随机种子
	rand.Seed(time.Now().UTC().UnixNano())

	// 调用绘图方法
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5      // 绘制多少圈
		res     = 0.0001 // 分辨率
		size    = 200    // 画布大小
		nframes = 1024   // gif图片的帧数
		delay   = 1      // gif帧之间的延时，以10ms为单位
	)

	frequent := rand.Float64() * 3.0

	// 初始化一个帧率为nframes的GIF图片
	anime := gif.GIF{LoopCount: nframes}

	// 外层循环绘制GIF
	phase := 0.0
	for i := 0; i < nframes; i++ {
		// 绘制画布
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		// 初始化调色板
		img := image.NewPaletted(rect, palette)

		// 内层循环绘制一帧
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*frequent + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), greenIndex)

		}

		phase += 0.1

		// 将绘制结果添加到GIF帧末尾
		anime.Delay = append(anime.Delay, delay)
		anime.Image = append(anime.Image, img)
	}

	// 循环结束，编码整个GIF图片
	gif.EncodeAll(out, &anime)
}
