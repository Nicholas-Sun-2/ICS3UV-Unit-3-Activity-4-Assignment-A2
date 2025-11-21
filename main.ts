/**
 * @author Nicholas Sun
 * @version 1.0.0
 * @date 2025-11-21
 * @fileoverview This program tests if you are in the correct weight range to enter the pie-eating contest.
 */

// Constants
const MIN_WEIGHT: number = 77.0;
const MAX_WEIGHT: number = 105.0;

// Input

// parseFloat by https://www.rameshfadatare.com/typescript-tutorial/typescript-parsefloat/
const WEIGHT: number = parseFloat(prompt("How much do you weigh (kg)?") || "0");

// If statement
if (WEIGHT >= MIN_WEIGHT && WEIGHT <= MAX_WEIGHT) {
  console.log("You may enter the contest.");
} else {
  console.log("You may not enter the contest.");
}
76
76.5
77
90
105
105.5
106
210
