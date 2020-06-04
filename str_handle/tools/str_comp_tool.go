package tools

// rate=comp_str/base
func Compstr(_base_str string, _comp_str string) float64 {
	base_str := []rune(_base_str)
	comp_str := []rune(_comp_str)
	mached := 0
	org_base_size := len(base_str)
	for comp_index := 0; comp_index < len(comp_str); comp_index++ {
		comp_char := comp_str[comp_index]
		for base_index := 0; base_index < len(base_str); base_index++ {
			base_char := base_str[base_index]
			if comp_char == base_char {
				mached += 1
				base_str = base_str[base_index+1:]
				break
			}
		}
	}
	return float64(mached) / float64(org_base_size)
}
