package triangles

import "fmt"

// Erwartet eine Seitenlänge `length`.
// Zeichnet ein gleichschenkliges, rechtwinkliges Dreieck mit diesen Seitenlängen auf der Konsole.
// Der rechte Winkel soll links unten liegen und die Seiten sollen
// vertikal bzw. horizontal verlaufen.
// Die Zeichen für Rand und Füllung des Rechtecks werden als Parameter erwartet.
func DrawTriangle(length int, inner, outer string) {

	row := 0

	for range length + 1 {
		for range row {
			if row == 3 || row == length-1 {
				fmt.Print(inner)
			} else {
				fmt.Print(outer)
			}
		}
		fmt.Println()
		row++
	}
}
