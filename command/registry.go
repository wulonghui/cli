package command

var Commands = make(map[string]Command)

func Register(cmd Command) {
	m := cmd.MetaData()
	Commands[m.Name] = cmd
}
