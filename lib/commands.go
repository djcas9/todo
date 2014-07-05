package todo

import (
	"github.com/codegangsta/cli"
)

func Commands() []cli.Command {
	return []cli.Command{
		{
			Name:   "add",
			Usage:  "Add a new todo item.",
			Action: AddAction,
		},
		{
			Name:   "list",
			Usage:  "List all active todo items.",
			Action: ListAction,
		},
		{
			Name:   "show",
			Usage:  "Show a todo item by id",
			Action: ShowAction,
		},
		{
			Name:   "remove",
			Usage:  "Remove a todo item by id",
			Action: RemoveAction,
		},
	}
}
