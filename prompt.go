package main

import (
	"fmt"
	"os"

	"gopkg.in/AlecAivazis/survey.v1"
)

// AddPrompt ... prompt for add logic
func (d *Data) AddPrompt(kmsAliasList []string) error {
	questions := []*survey.Question{
		{
			Name: "type",
			Prompt: &survey.Select{
				Message: "Choose a type:",
				Default: "SecureString",
				Options: []string{"SecureString", "String"},
			},
		},
		{
			Name:   "name",
			Prompt: &survey.Input{Message: "Name:"},
		},
		{
			Name:   "value",
			Prompt: &survey.Password{Message: "Value:"},
		},
		{
			Name:   "description",
			Prompt: &survey.Input{Message: "Description:"},
		},
		{
			Name: "kmsAlias",
			Prompt: &survey.Select{
				Message: "Choose a alias:",
				Options: kmsAliasList,
			},
		},
	}

	if err := survey.Ask(questions, d); err != nil {
		return err
	}

	fmt.Println()
	confirm := false
	prompt := &survey.Confirm{Message: "Are you sure?:"}
	if err := survey.AskOne(prompt, &confirm, nil); err != nil {
		return err
	}

	if !confirm {
		fmt.Println("cancel")
		os.Exit(0)
	}

	return nil
}

// DeletePrompt ... prompt for delete logic
func (r *DeleteData) DeletePrompt(c *Config) error {
	var list []string

	for _, v := range c.FileData {
		list = append(list, v.Name)
	}

	questions := []*survey.Question{
		{
			Name: "name",
			Prompt: &survey.Select{
				Message: "Choose a name:",
				Options: list,
			},
		},
	}

	if err := survey.Ask(questions, r); err != nil {
		return err
	}

	fmt.Println()
	confirm := false
	prompt := &survey.Confirm{Message: "Are you sure?:"}
	if err := survey.AskOne(prompt, &confirm, nil); err != nil {
		return err
	}

	if !confirm {
		fmt.Println("cancel")
		os.Exit(0)
	}

	return nil
}

// ForceDeletePrompt ... prompt for force delete logic.
// force delete : Can it be deleted from the file because it does not exist in the parameter store?
func (r *DeleteData) ForceDeletePrompt(c *Config) error {
	fmt.Println()
	confirm := false
	prompt := &survey.Confirm{Message: "Does not exist in the ParameterStore. Are you sure delete from file?:"}
	if err := survey.AskOne(prompt, &confirm, nil); err != nil {
		return err
	}

	if !confirm {
		fmt.Println("cancel")
		os.Exit(0)
	}

	return nil
}
