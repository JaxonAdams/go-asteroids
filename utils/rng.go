package utils

import (
	"math/rand/v2"
	"time"
)

var s = rand.NewPCG(42, uint64(time.Now().Unix()))
var Rng = rand.New(s)
