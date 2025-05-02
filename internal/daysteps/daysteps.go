package daysteps

import (
	"errors"
	"fmt"
	"spentcalories"
	"strconv"
	"strings"
	"time"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {
	// TODO: реализовать функцию
	newData := strings.Split(data, ",")
	if len(newData) != 2 {
		return 0, 0, errors.New("Нужно 2 значения!")
	}
	steps, err := strconv.Atoi(newData[0])
	if err != nil {
		return 0, 0, fmt.Errorf("Неверный формат количества шагов: %w", err)
	}
	if steps <= 0 {
		return 0, 0, errors.New("количество шагов не может быть отрицательным")
	}

	duration, err := time.ParseDuration(newData[1])
	if err != nil {
		return 0, 0, fmt.Errorf("Неверный формат продолжительности: %w", err)
	}
	if duration <= 0 {
		return 0, 0, errors.New("Продолжительность не может быть <= 0!")
	}
	return steps, duration, nil
}

func DayActionInfo(data string, weight, height float64) string {
	// TODO: реализовать функцию
	steps, duration, err := parsePackage(data)
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
		return ""
	}

	distanceMeters := float64(steps) * stepLength
	distanceKm := distanceMeters / mInKm

	calories := spentcalories.WalkingSpentCalories(steps, weight, height, duration)

	return fmt.Sprintf(
		"Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.",
		steps,
		distanceKm,
		calories,
	)
}
