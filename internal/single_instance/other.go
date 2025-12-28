//go:build !windows

package singleinstance

func CheckSingleInstance(onSecondInstance func()) bool {
	// На не-Windows просто разрешаем запуск
	return true
}
