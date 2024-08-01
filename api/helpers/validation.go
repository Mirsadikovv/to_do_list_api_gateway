package helpers

import (
	"errors"
	"fmt"
	"regexp"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v3"
)

func ValidatePhone(phone string) error {
	re := regexp.MustCompile(`^[+][9][9][8]\d{9}$`)
	if !re.MatchString(phone) {
		return errors.New("invalid phone number")
	}

	return nil
}

func Validatetype(types string) error {

	var validTypes = map[string]bool{
		"self_pickup": true,
		"delivery":    true,
	}
	if !validTypes[types] {
		return errors.New("invalid type, you have to input 'self_pickup' or 'delivery'")
	}
	return nil
}

func Validatepayment_type(payment_type string) error {
	if payment_type != "uzum" && payment_type != "cash" && payment_type != "terminal" {

		return errors.New("invalid payment_type  three payment_type 'uzum','cash' and 'terminal'  ")
	}
	return nil
}

func Validstatus(types string) error {

	var validstatus = []string{"waiting_for_payment", "collecting", "delivery", "waiting_on_branch", "finished", "cancelled"}

	for _, validstatus := range validstatus {
		if types == validstatus {
			return nil
		}
	}
	return errors.New("invalid type, you have to input 'waiting_for_payment, collecting, delivery, waiting_on_branch, finished, cancelled' ")
}

// Define valid statuses
var validStatuses = []string{"waiting_for_payment", "collecting", "delivery", "waiting_on_branch", "finished", "cancelled"}

// Define valid transitions
var validTransitions = map[string][]string{
	"waiting_for_payment": {"collecting"},
	"collecting":          {"delivery"},
	"delivery":            {"waiting_on_branch"},
	"waiting_on_branch":   {"finished", "cancelled"},
	"finished":            {},
	"cancelled":           {},
}

// Function to validate status transition
func Validstatusorder(currentStatus, newStatus string) error {
	// Check if the new status is valid
	isValidStatus := false
	for _, status := range validStatuses {
		if newStatus == status {
			isValidStatus = true
			break
		}
	}
	if !isValidStatus {
		return errors.New("invalid type, you have to input 'waiting_for_payment, collecting, delivery, waiting_on_branch, finished, cancelled'")
	}

	// Check if the transition is valid
	validNextStatuses, exists := validTransitions[currentStatus]
	if !exists {
		return errors.New("invalid current status")
	}

	for _, status := range validNextStatuses {
		if newStatus == status {
			return nil
		}
	}

	return fmt.Errorf(fmt.Sprintf("invalid transition from '%s' to '%s'", currentStatus, newStatus))
}

func ValidateEmailAddress(email string) error {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

	if !emailRegex.MatchString(email) {
		return fmt.Errorf("email address %s is not valid", email)
	}

	return nil
}

func ValidateDates(startDate, endDate string) error {
	var layout = "2006-01-02 15:04:05"

	from, err := time.Parse(layout, startDate)
	if err != nil {
		return errors.New("start_date is invalid")
	}

	to, err := time.Parse(layout, endDate)
	if err != nil {
		return errors.New("end_time is invalid")
	}

	if !from.Before(to) {
		return errors.New("start_time can not be greater than end_time")
	}

	return nil
}

func ValidatePassword(password string) error {
	if password == "" {
		return errors.New("password cannot be blank")
	}
	if len(password) < 8 || len(password) > 30 {
		return errors.New("password length should be 8 to 30 characters")
	}
	if validation.Validate(password, validation.Match(regexp.MustCompile("^[A-Za-z0-9$_@.#]+$"))) != nil {
		return errors.New("password should contain only alphabetic characters, numbers and special characters(@, $, _, ., #)")
	}
	if validation.Validate(password, validation.Match(regexp.MustCompile("[0-9]"))) != nil {
		return errors.New("password should contain at least one number")
	}
	if validation.Validate(password, validation.Match(regexp.MustCompile("[A-Za-z]"))) != nil {
		return errors.New("password should contain at least one alphabetic character")
	}
	return nil
}

func ValidateUsername(username string) error {
	if username == "" {
		return errors.New("username cannot be blank")
	}
	if len(username) < 5 || len(username) > 30 {
		return errors.New("username length should be 6 to 30 characters")
	}
	if validation.Validate(username, validation.Match(regexp.MustCompile("^[A-Za-z0-9$@_.#]+$"))) != nil {
		return errors.New("username should contain only alphabetic characters, numbers and special characters(@, $, _, ., #)")
	}
	return nil
}
