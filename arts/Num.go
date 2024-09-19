package arts

func Num(str string, Nums []int) []int {

	for _, char := range str {
		{
			num := int(char)
			num = ((num - 32) * 9) + 1
			Nums = append(Nums, num)
		}
	}
	return Nums
}

func Increment(Nums []int) []int {
	for i := 0; i < len(Nums); i++ {
		Nums[i]++
	}
	return Nums
}