// Lissajous generates GIF animations of random Lissajous figures.
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

var (
	cycles  = 5     // number of complete x oscillator revolutions
	res     = 0.001 // angular resolution
	size    = 100   // image canvas covers [-size..+size]
	nframes = 64    // number of animation frames
	delay   = 8     // delay between frames in 10ms units
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		param := r.URL.Query().Get("cycles")
		cycles, _ = strconv.Atoi(param)
		lissajous(w)
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))

}

func lissajous(out io.Writer) {

	rand.New(rand.NewSource(time.Now().UTC().UnixNano()))

	palette := make([]color.Color, 0, nframes)
	palette = append(palette, color.RGBA{A: 255})

	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference

	r := float64(rand.Intn(200))
	g := float64(rand.Intn(200))
	b := float64(rand.Intn(200))

	for i := 0; i < nframes; i++ {
		scale := float64(i) / float64(nframes)
		c := color.RGBA{R: uint8(55 + r*scale), G: uint8(55 + g*scale), B: uint8(55 + b*scale), A: 255}
		palette = append(palette, c)
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*float64(size)+0.5), size+int(y*float64(size)+0.5),
				uint8((i%(len(palette)-1))+1))
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
