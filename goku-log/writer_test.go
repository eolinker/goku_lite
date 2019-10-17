package goku_log

import (
	"context"
	"fmt"
	"testing"
	"time"
)

type MinPeriod struct {
}

func (p *MinPeriod) String() string {
	return "Minute"
}

func (p *MinPeriod) FormatLayout() string {
	return "200601021504"
}

func TestFileWriterByPeriod(t *testing.T) {
	w := NewFileWriteBytePeriod("/Users/huangmengzhu/test/log", "app.log", new(MinPeriod))
	defer w.Close()
	ctx, _ := context.WithTimeout(context.Background(), time.Minute*3)

	tick := time.NewTicker(time.Millisecond)
	defer tick.Stop()
	index := 0

	for {
		select {
		case <-ctx.Done():
			{
				w.Close()
				return

			}
		case <-tick.C:
			{
				index++
				fmt.Fprintf(w, "line:%d\n", index)
			}
		}
	}
}
