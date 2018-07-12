// 产生随机的 lissajous gif
package main

import (
	"image/color"
	"math/rand"
	"time"
	"os"
	"net/http"
	"io"
	"image/gif"
	"image"
	"math"
	"log"
)

var palette = []color.Color{color.White,color.Black} // 色板,可改变参数让图形变成其他颜色

const (
	whiteIndex = 0 // 画板中第一种,此处为白色
	blackIndex = 1 // 第二种,此处为黑
)

func main()  {
	rand.Seed(time.Now().UTC().UnixNano())  // 获取随机种子
	if len(os.Args) > 1 && os.Args[1] == "web" {
		handler := func(w http.ResponseWriter, r *http.Request) {
			lissajous(w)

		}
		http.HandleFunc("/" , handler)
		log.Fatal(http.ListenAndServe("localhost : 8000",nil))
		return
	}
	lissajous(os.Stdout)

}

func lissajous(out io.Writer) {
	const (
		cycles = 5  // 完整的 x 振荡器变化的个数
		res = 0.001	// 角度分辨率
		size = 100	// 图像画布包含 [ -size..+size]
		nframes = 64	// 帧数
		delay = 8	// 帧间延迟
	)
	freq := rand.Float64() * 3.0	// y 振荡器的相对频率
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0 , 0 ,2*size+1 , 2*size +1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2* math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size + int( x * size + 0.5), size + int( y * size +0.5),
				blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay,delay)
		anim.Image = append(anim.Image, img)

	}
	gif.EncodeAll(out, &anim )

}

