package cmd

import (
	"errors"
	"strconv"

	"github.com/liamganose/godo/todo"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var deleteCmd = &cobra.Command{
	Use:     "delete [flags] ID(s)",
	Aliases: []string{"del", "rm"},
	Short:   "Removes a specified todo item from the list",
	Long: `The 'delete' command allows you to remove a task 
from your todo list. Provide the ID or a list of IDs of 
the task(s) as an argument to remove it. Once a task is deleted, 
it cannot be retrieved.`,
	RunE: runDelete,
}

var all bool

func runDelete(cmd *cobra.Command, args []string) error {
	if (len(args) < 1) && !all {
		return errors.New("at least one ID is required")
	}

	newItems := []todo.Item{}
	toDelete := make(map[string]bool)
	for _, id := range args {
		toDelete[id] = true
	}

	if !all {
		items, _ := todo.ReadItems(viper.GetString("datadir"), "godo.json")
		for _, item := range items {
			if _, exists := toDelete[item.ID]; !exists {
				item.ID = strconv.Itoa(len(newItems)) // re-index the todos after deleting
				newItems = append(newItems, item)
			}
		}
	}
	if err := todo.SaveItems(viper.GetString("datadir"), "godo.json", newItems); err != nil {
		return err
	}
	return nil
}

func init() {
	rootCmd.AddCommand(deleteCmd)
	deleteCmd.Flags().BoolVarP(&all, "all", "a", false, "deletes all todo items")
}
