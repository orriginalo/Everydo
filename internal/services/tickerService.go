package services

import "time"

type Ticker struct {
	isRunning bool
	tickFunc  func()

	notifier *Notifier
	ticker   *time.Ticker
	interval time.Duration
}

func NewTicker(notifier *Notifier, interval time.Duration, tickFunc func()) *Ticker {
	return &Ticker{
		notifier: notifier,
		interval: interval,
		tickFunc: tickFunc,

		ticker:    nil,
		isRunning: false,
	}
}

func (t *Ticker) Start() {
	ticker := time.NewTicker(t.interval)
	t.ticker = ticker
	t.isRunning = true

	go func() {
		for range ticker.C {
			t.tickFunc()
		}
	}()
}

func (t *Ticker) Stop() {
	if t.isRunning {
		t.ticker.Stop()
		t.isRunning = false
	}
}
