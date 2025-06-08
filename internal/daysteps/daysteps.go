package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	parts := strings.Split(datastring, ",")
	if len(parts) != 2 {
		return errors.New("invalid data format: expected 2 parts")
	}

	stepsStr := parts[0]
	if strings.Contains(stepsStr, " ") {
		return errors.New("steps field contains spaces")
	}

	stepsStr = strings.TrimPrefix(stepsStr, "+")
	steps, err := strconv.Atoi(stepsStr)
	if err != nil {
		return fmt.Errorf("error converting steps: %w", err)
	}
	if steps <= 0 {
		return errors.New("steps must be greater than zero")
	}
	ds.Steps = steps

	durationStr := parts[1]
	if strings.Contains(durationStr, " ") {
		return errors.New("duration field contains spaces")
	}

	duration, err := time.ParseDuration(durationStr)
	if err != nil {
		return fmt.Errorf("error parsing duration: %w", err)
	}
	if duration <= 0 {
		return errors.New("duration must be greater than zero")
	}
	ds.Duration = duration

	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	distance := spentenergy.Distance(ds.Steps, ds.Height)
	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", err
	}

	result := fmt.Sprintf(
		"Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n",
		ds.Steps,
		distance,
		calories,
	)
	return result, nil
}
