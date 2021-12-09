package main

import (
	"context"
	"github.com/sirupsen/logrus"
	"time"
)

func main()  {
	ctx := context.Background()
	ctx , cancel := context.WithCancel(ctx)
	t := time.NewTicker(2 * time.Second)
		for {
			select {
			case <-ctx.Done():
				t.Stop()
				logrus.Infof("updating movies context cancelled")
				return
			case <-t.C:
				printer(cancel)
			}
		}
}

func printer(cnc context.CancelFunc){
	logrus.Infof("Hiiii")
	cnc()
}

