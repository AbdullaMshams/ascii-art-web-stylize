package arts

var Fpath string

func Banner(iput string) string {
	if iput == "standard" {
		Fpath = "arts/banners/standard.txt"
	} else if iput == "shadow" {
		Fpath = "arts/banners/shadow.txt"
	} else if iput == "thinkertoy" {
		Fpath = "arts/banners/thinkertoy.txt"
	} else {
		Fpath = "arts/banners/standard.txt"
	}
	return Fpath
}
