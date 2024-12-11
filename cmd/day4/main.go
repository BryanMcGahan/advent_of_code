package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
    // Example grid from the puzzle (hardcoded for demonstration)
    // Feel free to replace with your own input reading logic.
    data, err := os.ReadFile("./test/data/input/day4.txt")
    if err != nil {

    }

    puzzle := string(data)
    grid := strings.Split(puzzle, "\n")
    grid = grid[:len(grid)-1]

    word := "XMAS"
    count := countXMAS(grid)
    fmt.Println("Total occurrences of", word, ":", count)
}

func countOccurrences(grid []string, word string) int {
    // Directions: (dx, dy)
    // dx is row increment, dy is column increment
    directions := [][2]int{
        {0, 1},   // Right
        {0, -1},  // Left
        {1, 0},   // Down
        {-1, 0},  // Up
        {1, 1},   // Down-Right Diagonal
        {1, -1},  // Down-Left Diagonal
        {-1, 1},  // Up-Right Diagonal
        {-1, -1}, // Up-Left Diagonal
    }

    rows := len(grid)
    if rows == 0 {
        return 0
    }
    cols := len(grid[0])

    count := 0
    wordLen := len(word)

    for r := 0; r < rows; r++ {
        for c := 0; c < cols; c++ {
            // If the first letter doesn't match, no need to check directions.
            if grid[r][c] != word[0] {
                continue
            }

            // Check all directions from this starting point
            for _, dir := range directions {
                dx, dy := dir[0], dir[1]
                nr, nc := r, c
                matched := true

                for i := 1; i < wordLen; i++ {
                    nr += dx
                    nc += dy

                    // Check boundaries
                    if nr < 0 || nr >= rows || nc < 0 || nc >= cols {
                        matched = false
                        break
                    }
                    // Check if the character matches
                    if grid[nr][nc] != word[i] {
                        matched = false
                        break
                    }
                }

                if matched {
                    count++
                }
            }
        }
    }

    return count
}



func countXMAS(grid []string) int {
    rows := len(grid)
    if rows == 0 {
        return 0
    }
    cols := len(grid[0])

    count := 0

    // Each X-MAS pattern is a 3x3 area, so we must leave a border of 1 around
    // the grid to avoid out-of-bound checks.
    for r := 1; r < rows-1; r++ {
        for c := 1; c < cols-1; c++ {
            if grid[r][c] != 'A' {
                continue
            }

            // Coordinates of the diagonals:
            topLeft := grid[r-1][c-1]
            topRight := grid[r-1][c+1]
            bottomLeft := grid[r+1][c-1]
            bottomRight := grid[r+1][c+1]

            // We need to check the four configurations:
            // Diagonal 1 (top-left to bottom-right): M-A-S or S-A-M
            // Diagonal 2 (top-right to bottom-left): M-A-S or S-A-M

            // We'll define a helper function to check if two chars and the center 'A' form "M-A-S" or "S-A-M"
            if isMASorSAM(topLeft, bottomRight) && isMASorSAM(topRight, bottomLeft) {
                count++
            }
        }
    }

    return count
}

func isMASorSAM(a, b byte) bool {
    // With 'A' in the center (already guaranteed by the caller),
    // we must have either M and S or S and M.
    return (a == 'M' && b == 'S') || (a == 'S' && b == 'M')
}
