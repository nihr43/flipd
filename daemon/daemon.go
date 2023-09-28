package daemon

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

const defaultTick = 1 * time.Second

func run(ctx context.Context, out io.Writer) error {
	log.SetOutput(out)

	for {
		select {
		case <-ctx.Done():
			return nil
		case <-time.Tick(defaultTick):
			log.Printf("tick")
		}
	}
}

func Main(cidr string) error {
	fmt.Println("daemon got cidr", cidr)

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	defer func() {
		cancel()
	}()

	if err := run(ctx, os.Stdout); err != nil {
		return err
	} else {
		return nil
	}
}
