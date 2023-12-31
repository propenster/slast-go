/****************************************************************************
* Filename: slast.go
* Original Author: Faith E. Olusegun (propenster)
* File Creation Date: December 2nd, 2023
* Description: SLAST functions, definitions and implementation details
* LICENSE: MIT
*************************************************************************************/

package main

import (
	"errors"
	"fmt"
)

const (
	MATCH_SCORE    = 2
	MISMATCH_SCORE = -1
	GAP_PENALTY    = -1
)

//returns max of three int real numbers
//@param a
//@param b
//@param c
func maxOf(a, b, c int) int {
	if a > b && a > c {
		return a
	} else if b > c {
		return b
	} else {
		return c
	}
}

//@brief calculates match/alignment score
// @param a first nucleotide (a) to be compared to another (b)
// @param b second nucleotide (b) to be compared with a
// @return match or mismatch score
func calculateScore(a, b byte) int {
	if a == b {
		return MATCH_SCORE
	}
	return MISMATCH_SCORE
}

// @brief performs the local alignment
// @param seq1 sequence 1
// @param seq2 sequence 2
func simpleLocalAlignment(seq1, seq2 string) error {
	len1 := len(seq1)
	len2 := len(seq2)

	fmt.Printf("Length of nucleotides to be aligned Seq1: %v \t Seq2: %v\n", len1, len2)

	//create our dp alignment matrix and allocate memory for it...
	dp := make([][]int, len1)
	for i := 0; i < len1; i++ {
		dp[i] = make([]int, len2+1)
	}

	maxScore, endRow, endCol := 0, 0, 0

	//initialise the alignment matrix
	for row := 0; row < len1; row++ {
		for col := 0; col < len2; col++ {

			if row == 0 || col == 0 {
				if row <= len1 && col <= len2 {
					dp[row][col] = 0 //init with zeros....
				} else {
					return errors.New("alignment failed")
				}
			} else {
				match := dp[row-1][col-1] + calculateScore(seq1[row-1], seq2[col-1])
				delete := dp[row-1][col] + GAP_PENALTY
				insert := dp[row][col-1] + GAP_PENALTY
				dp[row][col] = maxOf(match, delete, insert)

				if dp[row][col] > maxScore {
					maxScore = dp[row][col]
					endRow = row
					endCol = col
				}

			}

		}
	}

	//find and print new alignment
	i := endRow
	j := endCol

	for i > 0 && j > 0 && dp[i][j] != 0 {
		score := dp[i][j]
		diagonal := dp[i-1][j-1]
		//left := dp[i][j-1]
		up := dp[i-1][j]

		if score == diagonal+calculateScore(seq1[i-1], seq2[j-1]) {
			fmt.Printf("%c", seq1[i-1])
			i--
			j--

		} else if score == up+GAP_PENALTY {
			fmt.Printf("-")
			i--
		} else {
			fmt.Printf("-")
			j--
		}
	}

	return nil

}

func main() {
	fmt.Println("Welcome to SLAST go (golang) implementation")
	fmt.Println("Welcome to SLAST (Simple Local Alignment Search Tool) in Golang")

	seq1 := "ACTCCGAT"
	seq2 := "GCTAAGAT"

	fmt.Printf("Sequence 1 to be aligned: %s\n", seq1)
	fmt.Printf("Sequence 2 to be aligned: %s\n", seq2)

	fmt.Println("Alignment:")

	simpleLocalAlignment(seq1, seq2)

}
