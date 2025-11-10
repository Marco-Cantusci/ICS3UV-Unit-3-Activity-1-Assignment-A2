/*
 * @author Marco Cantusci
 * @verion 1.0.0
 * @date 2025-11-10
 * @fileoverview Calculate and display the side length of a cube with a volume of 1000 mm³.
 */

package main

import "fmt"

func main() {

	// Stores volume as 1000
	const VOLUME float64 = 1000

	// formula/calculation of side length
	const sideLength float64 = VOLUME / 100

	// print volume
	fmt.Println("Volume:", VOLUME, "mm³")

	// print side length
	fmt.Println("The side length of a cube that has a volume of", VOLUME, "mm³ is", sideLength, "mm")

	fmt.Println("\nDone.")

	}
