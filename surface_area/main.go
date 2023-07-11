package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	x, y, z int
}

func main() {
	fmt.Println("Enter numeric values separated with (,) in the format x,y,z: E.g (1,1,1) \n ")
	//Collect input from terminal
	var pointsGrid []Point
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := strings.TrimSpace(scanner.Text())

		if len(input) == 0 {
			break
		}
		grid := strings.Split(input, ",")
		//Grid length must be 3 to avoid error
		if len(grid) != 3 {
			panicError(fmt.Errorf("Expected format x,y,z was not provide, enter the correct coordinate format"))
		}

		x, err := strconv.Atoi(grid[0])
		panicError(err)
		y, err := strconv.Atoi(grid[1])
		panicError(err)
		z, err := strconv.Atoi(grid[2])
		panicError(err)

		pointsGrid = append(pointsGrid, Point{x, y, z})
	}
	panicError(scanner.Err())

	surfaceArea := computeUnIntersectedPoint(pointsGrid)
	fmt.Printf("\n After computing the un-intersected points, the total surface area is %v \v",
		surfaceArea)
}

func panicError(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func computeUnIntersectedPoint(grids []Point) int {
	unIntersectedPoint := 0
	for _, p := range grids {
		unIntersectedPoint += pointIntersected(grids, Point{p.x + 1, p.y, p.z}) // RIGHT direction
		unIntersectedPoint += pointIntersected(grids, Point{p.x - 1, p.y, p.z}) // LEFT direction
		unIntersectedPoint += pointIntersected(grids, Point{p.x, p.y + 1, p.z}) // RIGHT direction
		unIntersectedPoint += pointIntersected(grids, Point{p.x, p.y - 1, p.z}) // LEFT direction
		unIntersectedPoint += pointIntersected(grids, Point{p.x, p.y, p.z + 1}) // RIGHT direction
		unIntersectedPoint += pointIntersected(grids, Point{p.x, p.y, p.z - 1}) // LEFT direction
	}

	return unIntersectedPoint
}

func pointIntersected(points []Point, grid Point) int {
	for _, point := range points {
		if point == grid {
			return 0
		}
	}
	return 1
}
