package dojo

func insertSort(arr []int) []int {
	var insertPosition int
	lengthArray := len(arr) - 1

	for index := 1; index <= lengthArray; index++ {
		valueInit := arr[index]
		positionInit := index - 1

		insertPosition, arr = canChangePosition(positionInit, valueInit, arr)

		arr[insertPosition+1] = valueInit
	}

	return arr
}

func canChangePosition(position, valueInit int, arr []int) (int, []int) {

	for position >= 0 && valueInit < arr[position] {

		arr[position+1] = arr[position]
		position--
	}

	return position, arr
}
