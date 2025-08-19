package main

import (
	"encoding/json"
	"fmt"
	"gmenu/cli"
	"os"
	"os/exec"
	"strings"

	gcolors "gmenu/internal/gmenu_colors"
)

func handleMenu(conf *cli.MenuConf) {
	f := conf.MenuConfPath

	jsonData, err := os.ReadFile(*f)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading config file '%v': %v\n", f, err)
		os.Exit(1)
	}

	var tasks []cli.MenuTask
	if err := json.Unmarshal(jsonData, &tasks); err != nil {
		fmt.Fprintf(os.Stderr, "Error unmarshaling JSON from '%v': %v\n", f, err)
		os.Exit(1)
	}

	if *conf.Result == "" {
		for _, task := range tasks {
			if task.Name == "" {
				continue
			}

			if *conf.IsJoin {
				task.Name = fmt.Sprintf("%s %s", task.Value, task.Name)
			}

			icon := task.Icon
			if icon == "" {
				fmt.Printf("%s\n", task.Name)
			} else {
				fmt.Printf("%s\x00icon\x1f%s\n", task.Name, icon)
			}
		}
		os.Exit(0)
	} else {
		selectedTaskName := *conf.Result

		var selectedTask *cli.MenuTask
		for i := range tasks {
			if strings.Contains(selectedTaskName, tasks[i].Name) {
				selectedTask = &tasks[i]
				break
			}
		}

		if selectedTask == nil {
			fmt.Fprintf(os.Stderr, "Error: Task '%s' not found in config file. No command to execute.\n", selectedTaskName)
			os.Exit(1)
		}

		var commandToExecute string
		if *conf.DefaultExec != "" {
			commandToExecute = fmt.Sprintf("%s %s", *conf.DefaultExec, selectedTask.Value)
		} else {
			commandToExecute = selectedTask.Value
		}

		cmd := exec.Command("bash", "-c", commandToExecute)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err = cmd.Start()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error starting command '%s': %v\n", commandToExecute, err)
			os.Exit(1)
		}
		os.Exit(0)
	}
}

func handlePick(conf *cli.PickConf) {
	colorStrings, err := gcolors.GatherColorStrings(*conf)
	if err != nil {
		fmt.Printf("Error collecting color strings: %v", err)
		os.Exit(1)
	}
	if *conf.DrawThumb {
		err = gcolors.DrawTmpThumbnail(colorStrings.HEX)
		if err != nil {
			fmt.Printf("Error writing thumbnail: %v", err)
			os.Exit(1)
		}
	}

	gcolors.OutputAsLines(colorStrings)
}
