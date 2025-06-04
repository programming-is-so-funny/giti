package contract

type Command interface {
	Name() string
	Description() string
	Run(args []string) error
}
