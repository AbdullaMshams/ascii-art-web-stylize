package arts

var Nums []int
var Nsplit []string
var result string

func Generator(input []string, style string) string {
	result = ""
	// for _, str := range input {
	// 	str = strings.Replace(str, "\\n", "\n", -1)
	//  	parts := strings.Split(str, "\n")
	//  	Nsplit = append(Nsplit, parts...)

	// }
	for i := 0; i < len(input); i++ {
		first := Num(input[i], Nums)

		var NumSlice [][]int

		for i := 0; i < 8; i++ {
			first = Increment(first)
			NumSlice = append(NumSlice, append([]int(nil), first...)) // Append a copy of the incremented slice
		}

		Fpath := Banner(style)

		result += PrintLines(Fpath, NumSlice)
	}
	return result
}
