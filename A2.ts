/*
 * @author Marco Cantusci
 * @verion 1.0.0
 * @date 2025-11-10
 * @fileoverview Calculate and display the side length of a cube with a volume of 1000 mm³.
 */

// stores volume as 1000
const VOLUME: number = 1000;

// formula/calculation of side length
let sideLength: number = VOLUME / 100;

// print volume
console.log("Volume:" + VOLUME + "mm³");

// print side length
console.log("The side length of a cube that has a volume of " + VOLUME + " mm³ is " + sideLength + " mm");

console.log("\nDone.");
