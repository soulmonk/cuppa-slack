package commands

type Command interface {
	Run()
}

type Commands struct {
	WOL Command
}

func Init() Commands {
	return Commands{
		newWol(),
	}
}
