package gps

import (
	"fmt"
	"math"
)

type world struct {
	radius float64
}

func rad(deg float64) float64 {
	return deg * math.Pi / 180
}

func (w world) distance(p1, p2 location) float64 {
	s1, c1 := math.Sincos(rad(p1.lat))
	s2, c2 := math.Sincos(rad(p2.lat))
	clong := math.Cos(rad(p1.long - p2.long))
	return w.radius * math.Acos(s1*s2+c1*c2*clong)
}

func (l location) description() string {
	return fmt.Sprintf("%v: %.1f°, %.1f°", l.name, l.lat, l.long)
}

type location struct {
	name      string
	lat, long float64
}

type gps struct {
	world       world
	current     location
	destination location
}

func (g gps) distance() float64 {
	return g.world.distance(g.current, g.destination)
}

func (g gps) message() string {
	return fmt.Sprintf("%v まで、%.1f km", g.destination.description(), g.distance())
}

type rover struct {
	gps
}

func main() {
	mars := world{radius: 3389.5}
	BradburyLanding := location{name: "BradburyLanding", lat: -4.5895, long: 137.4417}
	ElysiumPlanitia := location{name: "BradburyLanding", lat: 4.5, long: 135.9}

	marsgps := gps{
		world:       mars,
		current:     BradburyLanding,
		destination: ElysiumPlanitia,
	}

	curiosity := rover{
		gps: marsgps,
	}

	fmt.Println(curiosity.message())
}
