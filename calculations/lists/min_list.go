package lists

// Erwartet eine Liste von Zahlen und gibt das Minimum zurück.
// Wenn die Liste leer ist, wird 0 zurückgegeben.
func MinList(nums []int) int {

	if len(nums) == 0 {
		return 0
	}

	// Merke das erste Element als vorläufiges minimum (min)
	min := nums[0]

	// Laufe durch die Liste
	for i := range nums {

		// - Vergleiche jedes Element mit min
		if nums[i] < min {
			//  - Wenn das Element kleiner ist als min, update min
			min = nums[i]
		}
	}
	// Am Ende steht in min das kleinste Element, dieses liefern wir als Ergebnis
	return min
}

/* KI Loesung
package lists

// Erwartet eine Liste von Zahlen und gibt das Minimum zurück.
// Wenn die Liste leer ist, wird 0 zurückgegeben.
func MinList(nums []int) int {

	if len(nums) == 0 {
		return 0
	}
	min := nums[0]
	for _, num := range nums {
		if num < min {
			min = num
		}
	}
	return min
}

*/
