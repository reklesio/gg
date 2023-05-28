package main

import (
	"github.com/surdeus/gox/src/gx"
	"fmt"
)

func main() {
	lines := []gx.Line{
		gx.LineSegment{
			gx.Point{0, 1},
			gx.Point{1, 2},
		}.Line(),
		gx.LineSegment{
			gx.Point{0, 5},
			gx.Point{1, 2},
		}.Line(),
		gx.LineSegment{
			gx.Point{-1, -1},
			gx.Point{1, 50},
		}.Line(),
	}
	
	for _, l := range lines { fmt.Println(l) }
	
	l1 := gx.LineSegment{
		gx.Point{0, 0},
		gx.Point{1, 1},
	}.Line()
	
	l2 := gx.LineSegment{
		gx.Point{0, 1},
		gx.Point{1, 0},
	}.Line()
	fmt.Println(l1.Crosses(l2))
	fmt.Println(l1.ContainsPoint(gx.Point{1, 4}))
	/*t := gx.Triangle{
		gx.Point{0, 0},
		gx.Point{0, 100},
		gx.Point{100, 0},
	}
	
	points := []gx.Point{
		gx.Point{},
		gx.Point{.1, .1},
		gx.Point{-1, -1},
		gx.Point{1, 1},
		gx.Point{101, 1},
		gx.Point{100, 1},
		gx.Point{50, 1},
	}
	
	for _, p := range points {
		fmt.Println(p, t.PointIsIn(p))
	}*/
}



