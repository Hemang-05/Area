package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"
)

func handleCalculate(w http.ResponseWriter, r *http.Request) {
	radius := r.FormValue("radius")
	side := r.FormValue("side")
	length := r.FormValue("length")
	breadth := r.FormValue("breadth")

	//convert string to float
	radiusF, _ := strconv.ParseFloat(radius, 64)
	sideF, _ := strconv.ParseFloat(side, 64)
	lengthF, _ := strconv.ParseFloat(length, 64)
	breadthF, _ := strconv.ParseFloat(breadth, 64)
	// area calc

	circleAre := radiusF * radiusF * math.Pi
	areaSqr := sideF * sideF
	areaRect := lengthF * breadthF

	// display area

	fmt.Fprintf(w, "<p>Area of Circle: %.2f cm sq</p>", circleAre)
	fmt.Fprintf(w, "<p>Area of Square: %.2f cm sq</p>", areaSqr)
	fmt.Fprintf(w, "<p>Area of Rectangle: %.2f cm sq</p>", areaRect)

}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./html/form.html")
	})

	http.HandleFunc("/calc", handleCalculate)

	er := http.ListenAndServe(":3000", nil)
	if er != nil {
		log.Fatal(er)
	}

}
