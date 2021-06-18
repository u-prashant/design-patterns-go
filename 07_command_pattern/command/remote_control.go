package command

type SimpleRemoteControl struct {
	onCommands  [4]Command
	offCommands [4]Command
}

func (s *SimpleRemoteControl) SetCommand(slot int, onCommand Command, offCommand Command) {
	s.onCommands[slot] = onCommand
	s.offCommands[slot] = offCommand
}

func (s *SimpleRemoteControl) OnButtonPressed(slot int) {
	s.onCommands[slot].Execute()
}

func (s *SimpleRemoteControl) OffButtonPressed(slot int) {
	s.offCommands[slot].Execute()
}

func NewSimpleRemoteControl() *SimpleRemoteControl {
	s := SimpleRemoteControl{}
	noCommand := &NoCommand{}
	for i := 0; i < 4; i++ {
		s.onCommands[i] = noCommand
		s.offCommands[i] = noCommand
	}
	return &s
}
