package utils

import (
	"Everydo/internal/models"
	"strconv"
	"strings"
	"time"
)

func CalcNextReset(task models.Task, from time.Time) time.Time {
	switch task.ReloadType {

	case "daily":
		return nextDaily(from, task.ResetTime)

	case "weekly":
		if task.ResetWeekday == nil {
			panic("weekly task without reset_weekday")
		}
		return nextWeekly(from, *task.ResetWeekday, task.ResetTime)

	case "custom":
		return nextCustom(from, task.ReloadEvery, task.ResetTime)

	default:
		panic("unknown reload_type")
	}
}

func nextDaily(from time.Time, resetTime string) time.Time {
	h, m := parseHM(resetTime)

	next := time.Date(
		from.Year(), from.Month(), from.Day(),
		h, m, 0, 0, from.Location(),
	)

	if !next.After(from) {
		next = next.AddDate(0, 0, 1)
	}

	return next
}

func nextWeekly(from time.Time, weekday int, resetTime string) time.Time {
	h, m := parseHM(resetTime)

	days := (weekday - int(from.Weekday()) + 7) % 7
	if days == 0 {
		days = 7
	}

	nextDay := from.AddDate(0, 0, days)

	return time.Date(
		nextDay.Year(), nextDay.Month(), nextDay.Day(),
		h, m, 0, 0, from.Location(),
	)
}

func nextCustom(from time.Time, reloadEvery int, resetTime string) time.Time {
	h, m := parseHM(resetTime)

	next := time.Date(
		from.Year(), from.Month(), from.Day(),
		h, m, 0, 0, from.Location(),
	)

	next = next.AddDate(0, 0, reloadEvery)

	return next
}

func parseHM(timeStr string) (int, int) {
	parts := strings.Split(timeStr, ":")
	h, _ := strconv.Atoi(parts[0])
	m, _ := strconv.Atoi(parts[1])
	return h, m
}
