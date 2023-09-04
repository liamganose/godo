package cmd

import (
	"errors"
	"strconv"

	"github.com/liamganose/godo/todo"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var addCmd = &cobra.Command{
	Use:   "add [flags] text(s)",
	Short: "Adds a new todo item to the list",
	Long: `The 'add' command lets you add a new task to your todo list.
Provide a description of the task as an argument to keep track of it`,
	RunE: addItems,
}

var priority int

func addItems(cmd *cobra.Command, args []string) error {
	if len(args) < 1 {
		return errors.New("at least one todo argument is required")
	}
	items, _ := todo.ReadItems(viper.GetString("datadir"), "godo.json")
	for _, arg := range args {
		item := todo.Item{
			ID:       strconv.Itoa(len(items)),
			Text:     arg,
			Priority: priority,
			Status:   "To Do",
		}
		items = append(items, item)
	}
	if err := todo.SaveItems(viper.GetString("datadir"), "godo.json", items); err != nil {
		return err
	}
	return nil
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().IntVarP(&priority, "priority", "p", 2, "Sets the priority of the todo item.")
}
