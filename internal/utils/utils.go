package utils

import (
	"Everydo/internal/models"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"
)

func GetDataPath() string {
	dataPath := ""

	switch runtime.GOOS {
	case "windows":
		username := os.Getenv("USERNAME")
		path := "C:/Users/" + username + "/AppData/Local/Everydo"
		dataPath = path
	case "linux":
		data_home := os.Getenv("XDG_DATA_HOME")
		if data_home == "" {
			data_home = os.Getenv("HOME") + "/.local/share"
		}
		path := data_home + "/Everydo"
		dataPath = path
	default:
		return ""
	}
	cleanPath := filepath.FromSlash(strings.ReplaceAll(dataPath, "\\", "/"))
	os.MkdirAll(cleanPath, os.ModePerm)
	cleanPath = filepath.Join(cleanPath, "data.db")
	fmt.Println(cleanPath)
	return cleanPath
}

func CheckNextResetValid(task models.Task) (time.Time, bool) {
	now := time.Now()
	remainingTime := task.NextResetAt.Sub(now)
	if remainingTime < 0 {
		return CalcNextReset(task, now), false
	}
	return task.NextResetAt, true
}

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
