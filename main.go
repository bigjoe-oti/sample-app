package main

import (
    "image"
    "image/color"
    "image/draw"
    "image/png"
    "net/http"
)

// blueHandler serves a simple blue page or image
func blueHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/html")
    w.Write([]byte("<h1 style='color:blue;'>Blue Page</h1>"))
}

// redHandler serves a 100x100 red PNG image
func redHandler(w http.ResponseWriter, r *http.Request) {
    img := image.NewRGBA(image.Rect(0, 0, 100, 100))
    draw.Draw(img, img.Bounds(), &image.Uniform{color.RGBA{255, 0, 0, 255}}, image.Point{}, draw.Src)
    w.Header().Set("Content-Type", "image/png")
    png.Encode(w, img)
}

func main() {
    http.HandleFunc("/blue", blueHandler)
    http.HandleFunc("/red", redHandler)
    http.ListenAndServe(":8080", nil)
}
