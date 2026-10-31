package triangles

import "fmt"

// Erwartet eine Seitenlänge `length`.
// Zeichnet ein gleichschenkliges, rechtwinkliges Dreieck mit diesen Seitenlängen auf der Konsole.
// Das Dreieck soll komplett mit `#`-Zeichen gefüllt sein.
// Der rechte Winkel soll links unten liegen und die Seiten sollen
// vertikal bzw. horizontal verlaufen.
func DrawSolidTriangle(length int) {

	row := 0

	for range length + 1 {
		for range row {
			fmt.Print("#")
		}
		fmt.Println()
		row++
	}
}
