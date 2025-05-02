package spentcalories

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep                    = 0.65 // средняя длина шага.
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе
)

func parseTraining(data string) (int, string, time.Duration, error) {
	// TODO: реализовать функцию
	newData := strings.Split(data, ",")
	if len(newData) != 3 {
		return 0, "", 0, errors.New("Требуются 3 значения!")
	}

	steps, err := strconv.Atoi(newData[0])
	if err != nil {
		return 0, 0, fmt.Errorf("Неверный формат количества шагов: %w", err)
	}
	if steps <= 0 {
		return 0, 0, errors.New("количество шагов не может быть отрицательным")
	}

	activity := strings.TrimSpace(newData[1])

	duration, err := time.ParseDuration(newData[2])
	if err != nil {
		return 0, 0, fmt.Errorf("Неверный формат продолжительности: %w", err)
	}
	if duration <= 0 {
		return 0, 0, errors.New("Продолжительность не может быть <= 0!")
	}

	return steps, activity, duration, nil
}

func distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
	stepLength := height * stepLengthCoefficient
	return (float64(steps) * stepLength) / mInKm
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
	if duration <= 0 {
		return 0
	}
	return distance(steps, height) / duration.Hours()
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	// TODO: реализовать функцию
	steps, activity, duration, err := parseTraining(data)
	if err != nil {
		return "", err
	}

	if weight <= 0 {
		return "", errors.New("Масса тела не может быть <= 0!")
	}
	if height <= 0 {
		return "", errors.New("Рост не может быть <= 0!")
	}

	dist := distance(steps, height)
	speed := meanSpeed(steps, height, duration)
	durationHours := duration.Hours()

	var calories float64

	switch strings.ToLower(activity) {
	case "Ходьба":
		calories, calcErr = WalkingSpentCalories(steps, weight, height, duration)
	case "Бег":
		calories, calcErr = RunningSpentCalories(steps, weight, height, duration)
	default:
		return "", fmt.Errorf("неизвестный тип тренировки: %s", activity)
	}

	if err != nil {
		return "", fmt.Errorf("ошибка расчета калорий: %w", err)
	}

	result := fmt.Sprintf(
		"Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f",
		activity,
		durationHours,
		dist,
		speed,
		calories,
	)
	return result, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 {
		return 0, errors.New("Количество шагов не может быть <= 0!")
	}
	if weight <= 0 {
		return 0, errors.New("Масса тела не может быть <= 0!")
	}
	if height <= 0 {
		return 0, errors.New("Рост не может быть <= 0!")
	}
	if duration <= 0 {
		return 0, errors.New("Продолжительность не может быть <= 0!")
	}
	avgSpeed := meanSpeed(steps, height, duration)
	durInMins := duration.Minutes()
	calories := (weight * avgSpeed * durInMins) / minInH
	return calories, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 {
		return 0, errors.New("Количество шагов не может быть <= 0!")
	}
	if weight <= 0 {
		return 0, errors.New("Масса тела не может быть <= 0!")
	}
	if height <= 0 {
		return 0, errors.New("Рост не может быть <= 0!")
	}
	if duration <= 0 {
		return 0, errors.New("Продолжительность не может быть <= 0!")
	}
	avgSpeed := meanSpeed(steps, height, duration)
	durInMins := duration.Minutes()
	calories := (weight * avgSpeed * durInMins) / minInH
	return calories * walkingCaloriesCoefficient, nil
}
