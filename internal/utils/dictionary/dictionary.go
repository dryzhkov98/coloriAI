package dictionary

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
)

type Dictionary struct {
	StartCommand Command `yaml:"start_command"`
	HelpCommand  Command `yaml:"help_command"`
	Buttons      Buttons `yaml:"buttons"`
}

type Command struct {
	Command string `yaml:"command"`
	Message string `yaml:"message"`
}

type Buttons struct {
	MainMenu MainMenu `yaml:"main_menu"`
}

type CallbackButton struct {
	Text           string `yaml:"text"`
	CallbackAction string `yaml:"callback_action"`
}

type MainMenu struct {
	SendFoodPhoto CallbackButton `yaml:"send_food_photo"`
	FoodJournal   CallbackButton `yaml:"food_journal"`
	Statistics    CallbackButton `yaml:"statistics"`
	Help          CallbackButton `yaml:"help"`
}

func MustLoad(lang string) (*Dictionary, error) {
	file := filepath.Join("locales", fmt.Sprintf("%s.yml", lang))
	data, err := os.ReadFile(file)
	if err != nil {
		formattedErr := fmt.Errorf("error while load dictionary (%s): %w", lang, err)
		panic(formattedErr.Error())
	}

	var d Dictionary
	if err := yaml.Unmarshal(data, &d); err != nil {
		formattedErr := fmt.Errorf("error while parse YAML (%s): %w", lang, err)
		panic(formattedErr)
	}

	return &d, nil
}
