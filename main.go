package main

import (
	"fmt"
	"log/slog"
	"math/rand"
	"strings"
	"time"

	"github.com/dchest/uniuri"
)

func main() {
	for {
		nn := rand.Intn(1000)
		sr := strings.ToUpper(fmt.Sprintf("%s-%s", uniuri.NewLen(5), uniuri.NewLen(6)))
		txt := fmt.Sprintf("[%s] Lorem ipsum dolor sit amet, consectetur adipiscing elit", sr)
		if nn < 750 {
			slog.Error(txt)
			time.Sleep(time.Second * 1)
		}
		slog.Info(txt)
		time.Sleep(time.Second * 1)
	}
}
