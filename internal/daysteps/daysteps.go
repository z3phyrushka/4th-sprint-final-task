package daysteps

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/spentcalories"
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
		return 0, 0, fmt.Errorf("Нужно 2 значения: шаги и продолжительность")
	}

	steps, err := strconv.Atoi(newData[0])
	if err != nil {
		return 0, 0, fmt.Errorf("Неверный формат количества шагов: %v", err)
	}
	if steps <= 0 {
		return 0, 0, errors.New("Количество шагов не может быть <= 0!")
	}

	duration, err := time.ParseDuration(newData[1])
	if err != nil {
		return 0, 0, fmt.Errorf("Неверный формат продолжительности: %v", err)
	}
	if duration <= 0 {
		return 0, 0, errors.New("Продолжительность не может быть <= 0!")
	}
	return steps, duration, nil
}

func DayActionInfo(data string, weight, height float64) string {
	// TODO: реализовать функцию
	if weight <= 0 {
		log.Println("Ошибка: вес не может быть <= 0!")
		return ""
	}

	if height <= 0 {
		log.Println("Ошибка: рост не может быть <= 0!")
		return ""
	}

	steps, duration, err := parsePackage(data)
	if err != nil {
		log.Printf("Ошибка: %v", err)
		return ""
	}

	distanceMeters := float64(steps) * stepLength
	distanceKm := distanceMeters / mInKm

	calories, err := spentcalories.WalkingSpentCalories(steps, weight, height, duration)
	if err != nil {
		log.Printf("Ошибка расчета калорий: %v", err)
		return ""
	}

	return fmt.Sprintf(
		"Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n",
		steps,
		distanceKm,
		calories,
	)
}
