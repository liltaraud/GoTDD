package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

func TestCountdown(t *testing.T) {
	t.Run("output text correctness test", func(t *testing.T) {
		buf := bytes.Buffer{}
		want := "3\n2\n1\nGo!"
		sleeper := &SpySleeper{}

		Countdown(&buf, sleeper)
		got := buf.String()

		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}

		if sleeper.calls != 4 {
			t.Errorf("got %d calls to sleeper, want 4", sleeper.calls)
		}
	})

	t.Run("sleep before every print", func(t *testing.T) {
		spySleepPrinter := &CountdownOperationsSpy{}
		Countdown(spySleepPrinter, spySleepPrinter)

		want := []string{
			opSleep,
			opWrite,
			opSleep,
			opWrite,
			opSleep,
			opWrite,
			opSleep,
			opWrite,
		}

		if !reflect.DeepEqual(want, spySleepPrinter.Calls) {
			t.Errorf("\nWant :%v\nGot  :%v", want, spySleepPrinter.Calls)
		}

	})

}

func TestConfigurableSleeper(t *testing.T) {
	expectedDuration := 5 * time.Second

	spyTime := &spyTime{}
	sleeper := ConfigurableSleeper{expectedDuration, spyTime.Sleep}
	sleeper.Sleep()

	if spyTime.durationSlept != expectedDuration {
		t.Errorf("\nExpected sleep duration : %v\nEffective sleep duration : %v", expectedDuration, spyTime.durationSlept)
	}
}
