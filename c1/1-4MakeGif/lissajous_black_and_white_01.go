/*
在这个程序中，演示Go语言标准库里的image这个package的用法，
我们会用这个包来生成一系列的bit-mapped图，然后将这些图片编码为一个GIF动画，
GIF动画的内容是李萨如图形。
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

// 定义一个调色板，一个包含了颜色常量的数组
/*
当我们import了一个包路径包含有多个单词的package时，比如image/color
通常我们只需要用最后那个单词表示这个包就可以，当我们写color.White时，这个变量指向的是image/color包里的变量

这里的[]color.Color{...}是“复合声明”，这是实例化Go语言里的复合类型的一种写法，此处产生一个slice。
*/
var palette = []color.Color{color.White, color.Black}

// 为访问调色板中的颜色index定义一个可读性良好的名字
/*
	const用来声明常量
*/
const (
	whiteIndex = 0
	blackIndex = 1
)

func main() {
	// 将当前时间作为随机数种子，确保真随机
	rand.Seed(time.Now().UTC().UnixNano())

	/*
		程序的结果将输出到os.Stdout中，配合重定向生成此GIF图片
		go run lissajous.go > ./temp.gif
	*/
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5      // X轴振荡器的完整旋转次数
		res     = 0.0001 // 分辨率，越小越清晰
		size    = 200    // 画布大小
		nframes = 1024   // 动画帧数
		delay   = 1      // 帧之间的延迟，以10ms为单位
	)

	freqent := rand.Float64() * 3.0
	/*
		这里的复合声明生成的gif.GIF是一个struct类型
		struct是一组值（字段）的集合，不同的类型集合在一个struct可以让我们以一个统一的单元进行处理。
		anime是一个gif.GIF类型的struct变量。
		这种写法会生成一个struct变量，并且其内部变量LoopCount字段会被设置为nframes，而其它的字段会被设置为各自类型默认的零值。
		struct内部的变量可以用点号`.`来进行访问。
	*/
	anime := gif.GIF{LoopCount: nframes}

	// 外层循环，一次循环生成一帧画面
	/*
		它生成了一个包含白色和黑色的201*201大小的图片。
		所有像素点都会被默认设置为其零值（也就是调色板palette里的第0个值），这里我们设置的是白色。
		每次外层循环都会生成一张新图片，并将一些像素设置为黑色。其结果会append到之前结果之后。
	*/
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		// 初始化一个调色板
		img := image.NewPaletted(rect, palette)

		// 内层循环，用于绘制这一帧中的内容
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			/*
				内层循环设置两个偏振值。
					x轴偏振使用sin函数。
					y轴偏振也是正弦波，但其相对x轴的偏振是一个0-3的随机值，初始偏振值是一个零值，随着动画的每一帧逐渐增加。
				循环会一直跑到x轴完成cycles值设定的循环。每一步它都会调用SetColorIndex来为(x,y)点来染黑色。
			*/
			x := math.Sin(t)
			y := math.Sin(t*freqent + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex)
		}

		phase += 0.1

		// 将结果append到anim中的帧列表末尾，并设置延迟值。
		anime.Delay = append(anime.Delay, delay)
		anime.Image = append(anime.Image, img)
	}

	// 循环结束后所有的延迟值被编码进了GIF图片中，并将结果写入到输出流
	gif.EncodeAll(out, &anime)
}
