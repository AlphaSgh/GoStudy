package sort

func GetMax(arr []int) int {
	for j := 1; j < len(arr); j++ {
		if arr[j-1] > arr[j] {
			arr[j-1], arr[j] = arr[j], arr[j-1]
		}
	}
	return arr[len(arr)-1]

}
func BubbleSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return arr
}
