package scheduler_test

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/hyperjiang/scheduler"
	"github.com/stretchr/testify/require"
)

var i = 0

func demoTask() error {
	i++
	fmt.Printf("%d\n", i)
	if i == 3 {
		return errors.New("something wrong happens")
	}
	return nil
}

func TestScheduler(t *testing.T) {
	should := require.New(t)

	i = 0
	s := scheduler.New("DemoTask", demoTask, "100ms")
	s.Start()

	time.Sleep(time.Second)
	s.Stop()

	should.Equal(10, i)
}

func TestInvalidInterval(t *testing.T) {
	should := require.New(t)

	i = 0
	s := scheduler.New("Unhappy", demoTask, "1d")
	s.Start()

	time.Sleep(time.Second)
	s.Stop()

	should.Equal(1, i)
}
