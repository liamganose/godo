package cmd

import (
	"fmt"
	"sort"

	"github.com/liamganose/godo/todo"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var listCmd = &cobra.Command{
	Use:     "list [flags]",
	Aliases: []string{"ls"},
	Short:   "Lists all current todo items",
	Long: `The 'list' command displays all the tasks 
that you have currently added to your todo list. Tasks 
are presented in the order they were added, allowing you to 
quickly review what needs to be done.`,
	Run: listItems,
}

var showDone bool
var prio int

func listItems(cmd *cobra.Command, args []string) {
	items, _ := todo.ReadItems(viper.GetString("datadir"), "godo.json")
	items = filterDone(items)
	items = filterPriority(items)
	sort.Slice(items, func(i, j int) bool { return items[i].Priority < items[j].Priority })
	if len(items) > 0 {
		todo.PrettyPrint(items)
	} else {
		fmt.Println("Nothing todo!")
	}
}

func filterPriority(items []todo.Item) []todo.Item {
	if prio == 0 {
		return items
	}
	newItems := []todo.Item{}
	for _, item := range items {
		if item.Priority == prio {
			newItems = append(newItems, item)
		}
	}
	return newItems
}

func filterDone(items []todo.Item) []todo.Item {
	if showDone {
		return items
	}
	newItems := []todo.Item{}
	for _, item := range items {
		if item.Status != "Done" {
			newItems = append(newItems, item)
		}
	}
	return newItems
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().IntVarP(&prio, "priority", "p", 0, "only list todo items with certain priority. use 0 for all priorities")
	listCmd.Flags().BoolVar(&showDone, "done", false, "includes Done todo items when listing")
}
