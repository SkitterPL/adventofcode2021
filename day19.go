package main

import (
	"math"
	"strconv"
	"strings"
)

//https://adventofcode.com/2021/day/19

type Beacon struct {
	x int
	y int
	z int
}

type Scanner struct {
	beacons  []*Beacon
	position *Beacon
}

func (scanner *Scanner) hasBeacon(x int, y int, z int) bool {
	for _, beaconToCheck := range scanner.beacons {
		if beaconToCheck.equals(x, y, z) {
			return true
		}
	}
	return false
}

func (beacon *Beacon) equals(x int, y int, z int) bool {
	return beacon.x == x && beacon.y == y && beacon.z == z
}

func day19() (int, int) {
	data := fileToStringArray("input/19/input.txt")
	scanners := []*Scanner{}
	scanners2 := []*Scanner{}
	scannerCounter := 0
	beaconCounter := 0
	sum := 0
	for _, line := range data {
		if line == "" {
			continue
		}
		if strings.Contains(line, "scanner") {
			scanners = append(scanners, &Scanner{[]*Beacon{}, nil})
			scanners2 = append(scanners2, &Scanner{[]*Beacon{}, nil})
			sum += beaconCounter
			beaconCounter = 0
			scannerCounter++
			continue
		}
		scanners[scannerCounter-1].beacons = append(scanners[scannerCounter-1].beacons, newBeacon(line))
		scanners2[scannerCounter-1].beacons = append(scanners2[scannerCounter-1].beacons, newBeacon(line))
		beaconCounter++
	}
	scanners[0].position = &Beacon{0, 0, 0}
	scanners2[0].position = &Beacon{0, 0, 0}
	beacons := []*Beacon{}
	beacons = appendBeacons(beacons, scanners[0].beacons)
	scannersToMatch := true
	for scannersToMatch {
		scannersToMatch = false
		for key0, scanner1 := range scanners {
			scanner1 = scanners[key0]
			for key, scanner2 := range scanners {
				scanner2 = scanners[key]
				if key0 == key || scanner2.position != nil || scanner1.position == nil {
					continue
				}
				possibleScanner, beaconsToAdd := compare(scanner1, scanner2)
				if possibleScanner != nil {
					scanners[key] = possibleScanner
					beacons = appendBeacons(beacons, beaconsToAdd)
					break
				}
			}
		}

		for _, v := range scanners {
			if v.position == nil {
				scannersToMatch = true
			}
		}
	}

	//Part 2
	scannersToMatch = true
	for scannersToMatch {
		scannersToMatch = false
		for key0, scanner1 := range scanners2 {
			scanner1 = scanners2[key0]
			for key, scanner2 := range scanners2 {
				scanner2 = scanners2[key]
				if key0 == key || scanner2.position != nil || scanner1.position == nil {
					continue
				}
				possibleScanner := calcualteScannerPosition(scanner1, scanner2)
				if possibleScanner != nil {
					scanners2[key] = possibleScanner
					break
				}
			}
		}

		for _, v := range scanners2 {
			if v.position == nil {
				scannersToMatch = true
			}
		}
	}
	max := 0
	for key1, v1 := range scanners2 {
		for key2, v2 := range scanners2 {
			if key1 == key2 {
				continue
			}
			x := math.Abs(float64(v1.position.x - v2.position.x))
			y := math.Abs(float64(v1.position.y - v2.position.y))
			z := math.Abs(float64(v1.position.z - v2.position.z))
			if int(x+y+z) > max {
				max = int(x + y + z)
			}
		}
	}
	return len(beacons), max
}

func newBeacon(line string) *Beacon {
	points := strings.Split(line, ",")
	x, _ := strconv.Atoi(points[0])
	y, _ := strconv.Atoi(points[1])
	z, _ := strconv.Atoi(points[2])
	return &Beacon{x, y, z}
}

func compare(zeroScanner *Scanner, scannerToCompare *Scanner) (*Scanner, []*Beacon) {
	for _, beacon := range zeroScanner.beacons {
		for i := 0; i < 24; i++ {
			scanner := scannerToCompare.rotate(i)
			for _, comparedBeacon := range scanner.beacons {
				diff := &Beacon{beacon.x - comparedBeacon.x, beacon.y - comparedBeacon.y, beacon.z - comparedBeacon.z}
				commonBeacons := calculateBeacons(diff, zeroScanner, scanner)
				if commonBeacons >= 12 {
					scanner.position = &Beacon{zeroScanner.position.x + diff.x, zeroScanner.position.y + diff.y, zeroScanner.position.z + diff.z}
					for key, v := range scanner.beacons {
						scanner.beacons[key] = &Beacon{v.x + diff.x, v.y + diff.y, v.z + diff.z}
					}

					return scanner, scanner.beacons
				}
			}
		}

	}
	return nil, nil
}

func calcualteScannerPosition(zeroScanner *Scanner, scannerToCompare *Scanner) *Scanner {
	for _, beacon := range zeroScanner.beacons {
		for i := 0; i < 24; i++ {
			scanner := scannerToCompare.rotate(i)
			for _, comparedBeacon := range scanner.beacons {
				diff := &Beacon{beacon.x - comparedBeacon.x, beacon.y - comparedBeacon.y, beacon.z - comparedBeacon.z}
				commonBeacons := calculateBeacons(diff, zeroScanner, scanner)
				if commonBeacons >= 12 {
					scanner.position = &Beacon{zeroScanner.position.x + diff.x, zeroScanner.position.y + diff.y, zeroScanner.position.z + diff.z}

					return scanner
				}
			}
		}

	}
	return nil
}

func calculateBeacons(diff *Beacon, scanner1 *Scanner, scanner2 *Scanner) int {
	counter := 0
	for _, beacon := range scanner1.beacons {
		if scanner2.hasBeacon(beacon.x-diff.x, beacon.y-diff.y, beacon.z-diff.z) {
			counter++
		}
	}
	return counter
}

func (scanner *Scanner) rotate(rotation int) *Scanner {
	rotatedBeacons := make([]*Beacon, len(scanner.beacons))
	for key, beacon := range scanner.beacons {
		rotatedBeacons[key] = beacon.rotate(rotation)
	}
	return &Scanner{rotatedBeacons, nil}
}

func (beacon *Beacon) rotate(rotation int) *Beacon {
	switch rotation {
	case 0:
		return &Beacon{beacon.x, beacon.y, beacon.z}
	case 1:
		return &Beacon{beacon.x, -beacon.z, beacon.y}
	case 2:
		return &Beacon{beacon.x, -beacon.y, -beacon.z}
	case 3:
		return &Beacon{beacon.x, beacon.z, -beacon.y}
	case 4:
		return &Beacon{-beacon.x, -beacon.y, beacon.z}
	case 5:
		return &Beacon{-beacon.x, beacon.z, beacon.y}
	case 6:
		return &Beacon{-beacon.x, beacon.y, -beacon.z}
	case 7:
		return &Beacon{-beacon.x, -beacon.z, -beacon.y}
	case 8:
		return &Beacon{beacon.y, beacon.z, beacon.x}
	case 9:
		return &Beacon{beacon.y, beacon.x, -beacon.z}
	case 10:
		return &Beacon{beacon.y, -beacon.z, -beacon.x}
	case 11:
		return &Beacon{beacon.y, -beacon.x, beacon.z}
	case 12:
		return &Beacon{-beacon.y, beacon.z, -beacon.x}
	case 13:
		return &Beacon{-beacon.y, beacon.x, beacon.z}
	case 14:
		return &Beacon{-beacon.y, -beacon.z, beacon.x}
	case 15:
		return &Beacon{-beacon.y, -beacon.x, -beacon.z}
	case 16:
		return &Beacon{beacon.z, beacon.y, -beacon.x}
	case 17:
		return &Beacon{beacon.z, -beacon.y, beacon.x}
	case 18:
		return &Beacon{beacon.z, beacon.x, beacon.y}
	case 19:
		return &Beacon{beacon.z, -beacon.x, -beacon.y}
	case 20:
		return &Beacon{-beacon.z, -beacon.x, beacon.y}
	case 21:
		return &Beacon{-beacon.z, beacon.x, -beacon.y}
	case 22:
		return &Beacon{-beacon.z, beacon.y, beacon.x}
	case 23:
		return &Beacon{-beacon.z, -beacon.y, -beacon.x}
	}
	return nil
}

func appendBeacons(beacons []*Beacon, beaconsToAdd []*Beacon) []*Beacon {
	for _, beaconToAdd := range beaconsToAdd {
		shouldAdd := true
		for _, beacon := range beacons {
			if beacon.equals(beaconToAdd.x, beaconToAdd.y, beaconToAdd.z) {
				shouldAdd = false
				break
			}
		}
		if shouldAdd {
			beacons = append(beacons, beaconToAdd)
		}
	}
	return beacons
}
