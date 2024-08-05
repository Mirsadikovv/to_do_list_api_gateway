package validator

import (
	"errors"
	"regexp"
	"time"
)

func ValidateGmail(gmail string) bool {
	return regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@gmail.com$`).MatchString(gmail)
}
func ValidatePhone(phone string) bool {
	regex := `^\+998\d{9}$`
	return regexp.MustCompile(regex).MatchString(phone)
}

func ValidatePassword(password string) error {
	if len(password) < 8 {
		return errors.New("password must be at least 8 characters")
	}

	hasLetter := regexp.MustCompile(`[a-zA-Z]`).MatchString(password)
	hasDigit := regexp.MustCompile(`[0-9]`).MatchString(password)
	hasSpecial := regexp.MustCompile(`[^a-zA-Z0-9\s]`).MatchString(password)

	if !hasLetter || !hasDigit || !hasSpecial {
		return errors.New("password must contain letters, number and symbol")
	}

	return nil
}

func CheckDeadline(timestamp string) (float64, error) {
	layout := time.RFC3339

	date, err := time.Parse(layout, timestamp)
	if err != nil {
		return -1, errors.New("wrong timestamp format")
	}

	now := time.Now()

	hoursUntil := date.Sub(now).Hours()

	if hoursUntil < 0 {
		return 0, nil
	}

	return hoursUntil, nil
}

func ValidateBitrthday(birthday string) error {
	layout := "02-01-2006"

	date, err := time.Parse(layout, birthday)
	if err != nil {
		return errors.New("wrong date format")
	}
	today := time.Now()

	duration := today.Sub(date)
	years := int(duration.Hours() / (24 * 365))
	if years < 14 {
		return errors.New("you are younger than 14 ")
	}
	return nil
}

func IsSunday(timestampStr string) (bool, error) {
	timestamp, err := time.Parse(time.RFC3339, timestampStr)
	if err != nil {
		return false, err
	}

	return timestamp.Weekday() == time.Sunday, nil
}
