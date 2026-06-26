package main

import "math/rand/v2"

type DelayRange struct {
	Min int
	Max int
}

var delayConfig = map[string]DelayRange{
	"pageLoad":        {3000, 7000},
	"typing":          {80, 200},
	"betweenStages":   {1500, 4000},
	"betweenAccounts": {10000, 30000},
	"verifyPageLoad":  {3000, 5000},
	"apiDelay":        {2000, 5000},
}

func delay(name string) int {
	r := delayConfig[name]
	if r.Max == 0 {
		return 1000
	}
	return r.Min + rand.IntN(r.Max-r.Min+1)
}
