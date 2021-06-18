package main

import "design-patterns-go/07_command_pattern/command"

func main() {
	remote := command.NewSimpleRemoteControl()

	light := command.Light{}
	lightOnCommand := command.LightOnCommand{Light: light}
	lightOffCommand := command.LightOffCommand{Light: light}
	remote.SetCommand(0, &lightOnCommand, &lightOffCommand)

	fan := command.Fan{}
	fanOnCommand := command.FanOnCommand{Fan: fan}
	fanOffCommand := command.FanOffCommand{Fan: fan}
	remote.SetCommand(1, &fanOnCommand, &fanOffCommand)

	remote.OnButtonPressed(1)
	remote.OffButtonPressed(1)

	remote.OnButtonPressed(0)
	remote.OffButtonPressed(0)

	remote.OnButtonPressed(2)
	remote.OffButtonPressed(2)
}
