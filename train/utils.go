package train

func expandTrainRouteName(abbrRouteName string) string {
	switch abbrRouteName {
	case "P":
		return "Purple"
	case "Y":
		return "Yellow"
	case "Blue":
		return "Blue"
	case "Pink":
		return "Pink"
	case "G":
		return "Green"
	case "Org":
		return "Orange"
	case "Brn":
		return "Brown"
	default:
		return "UNKNOWNLINE"
	}
}
