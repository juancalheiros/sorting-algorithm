package dojo

func bubbleSorted(arr []int) []int {
	pointer := &arr
	lenght := len(arr) - 1

	for i := 0; i <= lenght; i++ {
		for index := 0; index <= lenght-1; index++ {
			canChangePosition(pointer, index)
		}
	}

	return arr
}

func canChangePosition(arr *[]int, index int) *[]int {

	if (*arr)[index] > (*arr)[index+1] {
		aux := (*arr)[index+1]
		(*arr)[index+1] = (*arr)[index]
		(*arr)[index] = aux
	}
	return arr
}
