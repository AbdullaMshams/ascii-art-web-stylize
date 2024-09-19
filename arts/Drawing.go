package arts

import (
	"bufio"
	"os"
	"fmt"
)

func PrintLines(Fpath string, NumSlice [][]int) string {
	var result string
	file, _ := os.Open(Fpath)
	// if err != nil {
	// 	return err
	// }
	defer file.Close()

	for _, row := range NumSlice {
		for _, lineNumber := range row {

			file.Seek(0, 0)
			scanner := bufio.NewScanner(file)

			currentLine := 0
			lineFound := false

			for scanner.Scan() {
				currentLine++
				if currentLine == lineNumber {
					fmt.Print(scanner.Text())
					result += scanner.Text()
					lineFound = true
					break
				}
			}
			
			if !lineFound {
				// fmt.Println("Line number", lineNumber, "not found.")
			}
		}
		fmt.Print("\n")
		result += ("\n")
	}

	return result
}
