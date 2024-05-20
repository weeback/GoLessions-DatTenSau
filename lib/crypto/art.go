package crypto

import "crypto/sha256"

const (
	ImageWidth  = 17
	ImageHeight = 9
)

func generateRandomArt(data []byte) string {

	hash := sha256.Sum256(data)
	y := int(hash[0]) % ImageHeight
	x := int(hash[1]) % ImageWidth
	image := make([][]int, ImageHeight)
	for i := range image {
		image[i] = make([]int, ImageWidth)
	}
	for _, b := range hash[2:] {
		dir := int(b) % 4
		switch dir {
		case 0:
			if y > 0 {
				y--
			}
		case 1:
			if x < ImageWidth-1 {
				x++
			}
		case 2:
			if y < ImageHeight-1 {
				y++
			}
		case 3:
			if x > 0 {
				x--
			}
		}
		image[y][x]++
	}
	result := "+---[RSA 2048]----+\n"
	for _, row := range image {
		result += "| "
		for _, cell := range row {
			switch {
			case cell == 0:
				result += " "
			case cell == 1:
				result += "."
			case cell == 2:
				result += "o"
			case cell == 3:
				result += "+"
			case cell == 4:
				result += "="
			case cell == 5:
				result += "*"
			case cell == 6:
				result += "B"
			case cell == 7:
				result += "O"
			case cell == 8:
				result += "X"
			case cell >= 9:
				result += "@"
			}
		}
		result += " |\n"
	}
	result += "+----[SHA256]-----+"
	return result
}
