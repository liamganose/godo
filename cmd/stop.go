package cmd

import (
	"errors"

	"github.com/liamganose/godo/todo"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var stopCmd = &cobra.Command{
	Use:   "stop [flags] ID",
	Short: "Marks a specified todo item as to do",
	Long: `The 'stop' command allows you to mark a task 
in your todo list as to do. Provide the ID of the task as 
an argument to set its status to to do.`,
	RunE: stopItem,
}

func stopItem(cmd *cobra.Command, args []string) error {
	if len(args) != 1 {
		return errors.New("one ID argument is required")
	}
	items, _ := todo.ReadItems(viper.GetString("datadir"), "godo.json")
	items = todo.SetStatus(items, args[0], "To Do")
	if err := todo.SaveItems(viper.GetString("datadir"), "godo.json", items); err != nil {
		return err
	}
	return nil
}

func init() {
	rootCmd.AddCommand(stopCmd)
}
