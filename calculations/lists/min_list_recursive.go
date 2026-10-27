package lists

// Erwartet eine Liste von Zahlen und gibt das Minimum zurück.
// Wenn die Liste leer ist, wird 0 zurückgegeben.
// ZUSATZBEDINGUNG: Diese Funktion darf keine Schleife verwenden.
func MinListRecursive(nums []int) int {
	if len(nums) == 0 {
		return 0

	} else if len(nums) == 1 {
		return nums[0]
	}

	// Die Funktion wird in der Funktion aufgerufen "Schleife"
	// [1:] verkürzt die Liste jedesmal
	min := MinListRecursive(nums[1:])
	if nums[0] < min {
		return nums[0]
	}

	return min
}

// REMARKS
// - Diese Funktion nennt man "rekursiv".
// - Rekursion ist ein größeres Thema, das in der Vorlesung separat behandelt wird.
//   Um die Denkweise frühzeitig zu üben, gibt es aber gelegentlich auch vorab Aufgaben dazu.
