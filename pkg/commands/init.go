package commands

type Command interface {
	Run() error
}

type Commands struct {
	WOL Command
}

func Init() Commands {
	// todo
	// to map
	// {'/commandName': commandInstance()}
	return Commands{
		newWol(),
	}
}
