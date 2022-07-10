package style

func Enquote(t string) string {
	return "«" + t + "»"
}

func Bold(t string) string {
	return "<b>" + t + "</b>"
}

func Italic(t string) string {
	return "<i>" + t + "</i>"
}

func Code(t string) string {
	return "<code>" + t + "</code>"
}

func Strike(t string) string {
	return "<del>" + t + "</del>"
}

func Underline(t string) string {
	return "<u>" + t + "</u>"
}

func Pre(t, language string) string {
	return "<u language=\"" + language + "\">" + t + "</u>"
}
