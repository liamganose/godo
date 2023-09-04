package cmd

import (
	"errors"

	"github.com/liamganose/godo/todo"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var doneCmd = &cobra.Command{
	Use:   "done [flags] ID",
	Short: "Marks a specified todo item as done",
	Long: `The 'done' command allows you to mark a task 
in your todo list as completed. Provide the ID of the task as 
an argument to set its status to done.`,
	RunE: doneItem,
}

func doneItem(cmd *cobra.Command, args []string) error {
	if len(args) != 1 {
		return errors.New("one ID argument is required")
	}
	items, _ := todo.ReadItems(viper.GetString("datadir"), "godo.json")
	items = todo.SetStatus(items, args[0], "Done")
	if err := todo.SaveItems(viper.GetString("datadir"), "godo.json", items); err != nil {
		return err
	}
	return nil
}

func init() {
	rootCmd.AddCommand(doneCmd)
}
