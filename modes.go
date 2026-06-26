package main

type ModeDef struct {
	Key   string
	Label string
}

var modes = []ModeDef{
	{Key: "auto", Label: "Auto"},
	{Key: "manual", Label: "Manual"},
}

func getNextMode(current string) ModeDef {
	for i, m := range modes {
		if m.Key == current {
			return modes[(i+1)%len(modes)]
		}
	}
	return modes[0]
}

func isManualMode(key string) bool {
	return key == "manual"
}
