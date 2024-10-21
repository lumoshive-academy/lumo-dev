package lumo

var units = []string{
	"", "satu", "dua", "tiga", "empat", "lima", "enam", "tujuh", "delapan", "sembilan",
}

var tens = []string{
	"", "sepuluh", "dua puluh", "tiga puluh", "empat puluh", "lima puluh", "enam puluh", "tujuh puluh", "delapan puluh", "sembilan puluh",
}

var teens = []string{
	"sepuluh", "sebelas", "dua belas", "tiga belas", "empat belas", "lima belas", "enam belas", "tujuh belas", "delapan belas", "sembilan belas",
}

var scale = []string{
	"", "ribu", "juta", "miliar", "triliun",
}

// NumberToIndonesianText converts an integer number to Indonesian words (e.g., 123 -> "seratus dua puluh tiga")
func NumberToIndonesianText(num int) string {
	if num == 0 {
		return "nol"
	}

	result := ""
	scaleIdx := 0

	for num > 0 {
		if num%1000 != 0 {
			result = convertHundreds(num%1000) + " " + scale[scaleIdx] + " " + result
		}
		num /= 1000
		scaleIdx++
	}

	return result
}

// Helper function to convert numbers below 1000
func convertHundreds(num int) string {
	if num == 0 {
		return ""
	} else if num < 10 {
		return units[num]
	} else if num < 20 {
		return teens[num-10]
	} else if num < 100 {
		return tens[num/10] + " " + units[num%10]
	} else {
		if num/100 == 1 {
			return "seratus " + convertHundreds(num%100)
		}
		return units[num/100] + " ratus " + convertHundreds(num%100)
	}
}

func HelloLumoAcademy(name string) string {
	return "Pagi " + name
}
