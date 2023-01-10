package helper

import (
	"errors"
	"net/mail"
	"github.com/manifoldco/promptui"
)

func PromptText(label string) (string, error) {
	validate := func(input string) error {
		_, err := mail.ParseAddress(input)
    if err != nil {
        return errors.New("invalid email address")
    }
    return nil
	}
	prompt := promptui.Prompt{
		Label:    label,
		Validate: validate,
	}

	result, err := prompt.Run()

	if err != nil {
		return "", errors.New("PROMPT: failed to get input - " + label)
	}

	return result, nil
}

func PromptPassword(label string) (string, error) {
	prompt := promptui.Prompt{
		Label:    label,
		Mask: '*',
	}

	result, err := prompt.Run()

	if err != nil {
		return "", errors.New("PROMPT: failed to get input - " + label)
	}

	return result, nil
}
