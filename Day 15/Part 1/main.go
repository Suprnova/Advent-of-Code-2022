package main

import (
	"bufio"
	"fmt"
	"os"
)

var minX, maxX, maxRange int

// Sensor represents a sensor in the cave
type Sensor struct {
	x           int
	y           int
	beaconX     int
	beaconY     int
	beaconRange int
}

// function to calculate the range from a sensor in which a beacon cannot exist via taxicab distance
func (s *Sensor) calculateRange() {
	s.beaconRange = abs(s.x-s.beaconX) + abs(s.y-s.beaconY)
	if s.beaconRange > maxRange {
		maxRange = s.beaconRange
	}
}

// function to determine if a beacon could exist at a given x, y coordinate
func (s *Sensor) couldHaveBeacon(x, y int) bool {
	return abs(x-s.x)+abs(y-s.y) > s.beaconRange
}

func main() {
	// open the input file
	input, err := os.Open("../input.txt")
	if err != nil {
		panic(err)
	}
	defer input.Close()

	// initialize the scanner and sensor slice
	sc := bufio.NewScanner(input)
	sensors := make([]Sensor, 0)

	for sc.Scan() {
		// read the input line, format it into coordinates
		var x, y, bX, bY int
		fmt.Sscanf(sc.Text(), "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &x, &y, &bX, &bY)
		// update the min and max x values if applicable
		updateMinMax(x, bX)
		// create a new sensor, calculate its range, and add it to the slice
		sensor := Sensor{x, y, bX, bY, 0}
		sensor.calculateRange()
		sensors = append(sensors, sensor)

	}

	// define the y value to search for beacons on, and a counter for the number of impossible positions
	targetY, impossiblePositions := 2000000, 0
nextX:
	// loop through all possible x values, and check if a beacon could exist at each one
	for i := minX - maxRange; i <= maxX+maxRange; i++ {
		for _, sensor := range sensors {
			// if a beacon could not exist at the given coords, increment the counter and move on to the next x value
			if !sensor.couldHaveBeacon(i, targetY) {
				impossiblePositions++
				continue nextX
			}
		}
	}
	dupeBeacons := make(map[int]bool)
	// remove any known beacons from the impossible positions count
	for _, sensor := range sensors {
		if sensor.beaconY == targetY && !dupeBeacons[sensor.beaconX] {
			impossiblePositions--
			dupeBeacons[sensor.beaconX] = true
		}
	}
	println(impossiblePositions)
}

// function to return the absolute value of an integer
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// function to update the min and max x values if applicable
func updateMinMax(x, bX int) {
	if x < minX {
		minX = x
	} else if x > maxX {
		maxX = x
	}
	if bX < minX {
		minX = bX
	} else if bX > maxX {
		maxX = bX
	}
}
