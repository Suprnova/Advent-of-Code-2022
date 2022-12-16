package main

import (
	"bufio"
	"fmt"
	"os"
)

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
}

// function to determine if a beacon could exist at a given x, y coordinate
func (s *Sensor) couldHaveBeacon(x, y int) bool {
	return abs(x-s.x)+abs(y-s.y) > s.beaconRange
}

// function to determine the offset necessary to skip the range of a beacon
func (s *Sensor) findXOffset(x, y int) int {
	if s.couldHaveBeacon(x, y) {
		return 0
	}
	return s.beaconRange - abs(y-s.y) + (s.x - x)
}

func main() {
	// open the input file
	input, err := os.Open("../input.txt")
	if err != nil {
		panic(err)
	}
	defer input.Close()

	// initialize the scanner, sensor slice, and beacon map
	sc := bufio.NewScanner(input)
	sensors := make([]Sensor, 0)
	beacons := make(map[string]bool)

	for sc.Scan() {
		// read the input line, format it into coordinates
		var x, y, bX, bY int
		fmt.Sscanf(sc.Text(), "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &x, &y, &bX, &bY)
		// create a new sensor, calculate its range, and add it to the slice
		sensor := Sensor{x, y, bX, bY, 0}
		sensor.calculateRange()
		sensors = append(sensors, sensor)
		// add the beacon to the map, to keep track of which beacons have been found
		beacons[fmt.Sprintf("%d,%d", bX, bY)] = true
	}

	// iterate through the x and y coordinates from 0 to 4000000
	for y := 0; y <= 4000000; y++ {
		for x := 0; x < 4000000; x++ {
			skipped := false
			for _, sensor := range sensors {
				// find the offset from the sensor necessary to skip the range of the beacon
				offset := sensor.findXOffset(x, y)
				// if the range isn't zero, we need to skip some x values
				if offset != 0 {
					skipped = true
					x += offset
				}
			}
			// if we didn't skip any x values, we found the distress beacon
			if !skipped {
				frequency := int64(x)*int64(4000000) + int64(y)
				println(frequency)
				return
			}
		}
	}
}

// function to return the absolute value of an integer
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
