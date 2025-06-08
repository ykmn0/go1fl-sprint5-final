package trainings

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	parts := strings.Split(datastring, ",")
	if len(parts) != 3 {
		return errors.New("invalid data format: expected 3 parts")
	}

	stepsStr := strings.TrimSpace(parts[0])
	stepsStr = strings.TrimPrefix(stepsStr, "+")
	steps, err := strconv.Atoi(stepsStr)
	if err != nil {
		return fmt.Errorf("error converting steps: %w", err)
	}
	if steps <= 0 {
		return errors.New("steps must be greater than zero")
	}
	t.Steps = steps

	t.TrainingType = parts[1]

	duration, err := time.ParseDuration(parts[2])
	if err != nil {
		return fmt.Errorf("error parsing duration: %w", err)
	}
	if duration <= 0 {
		return errors.New("duration must be greater than zero")
	}
	t.Duration = duration

	return nil
}

func (t Training) ActionInfo() (string, error) {
	distance := spentenergy.Distance(t.Steps, t.Height)
	meanSpeed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)

	var calories float64
	var err error

	switch strings.ToLower(t.TrainingType) {
	case "бег":
		calories, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	case "ходьба":
		calories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	default:
		return "", errors.New("unknown training type")
	}

	if err != nil {
		return "", err
	}

	if t.Steps <= 0 {
		return "", errors.New("steps must be greater than zero")
	}

	if t.Duration <= 0 {
		return "", errors.New("duration must be greater than zero")
	}

	result := fmt.Sprintf(
		"Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
		t.TrainingType,
		t.Duration.Hours(),
		distance,
		meanSpeed,
		calories,
	)
	return result, nil
}
