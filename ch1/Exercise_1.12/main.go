// Exercise 1.12: Modify the Lissajous server to read parameter values from the URL. For example, you might arrange it so that
// a URL like http://localhost:8000/?cycles=20 sets the number of cycles to 20 instead of the default 5. Use the strconv.Atoi function to convert the
// string parameter into an integer. You can see its documentation with go doc strconv.Atoi.

// Exercise 1.5: Change the Lissajous program’s color palette to green on black, for added authenticity.
// To create the web color #RRGGBB, use color.RGBA{0xRR, 0xGG, 0xBB, 0xff},  where each pair of hexadecimal digits represents
// the intensity of the red, green, or blue component of the pixel.
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
)

var palette = []color.Color{color.Black, color.RGBA{40, 200, 20, 255}}

const (
	blackIndex = 0 // first color in palette
	greenIndex = 1 // next color in palette
)

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		var cycles int
		var err error
		cyclesParam, ok := r.URL.Query()["cycles"]
		if !ok || len(cyclesParam) < 1 {
			cycles = 5
		} else {
			cycles, err = strconv.Atoi(cyclesParam[0])
		}
		if err != nil {
			cycles = 5
		}
		lissajous(w, cycles)
	}
	http.HandleFunc("/", handler)
	//!-http
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
	return
}
func lissajous(out io.Writer, cycles int) {
	const ( // number of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				greenIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
