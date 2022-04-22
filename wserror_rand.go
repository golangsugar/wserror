package wserror

import (
	"math/rand"
	"time"
)

var rnd *rand.Rand

func init() {
	ns := rand.NewSource(time.Now().UnixNano())
	rnd = rand.New(ns)
}
