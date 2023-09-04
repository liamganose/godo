package todo

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"text/tabwriter"
)

type Item struct {
	ID       string
	Text     string
	Priority int
	Status   string
}

func SaveItems(path string, filename string, items []Item) error {
	data, err := json.Marshal(items)
	if err != nil {
		return err
	}
	if err = os.MkdirAll(path, 0777); err != nil {
		return err
	}
	fp := filepath.Join(path, filename)
	err = os.WriteFile(fp, data, 0644)
	return err
}

func ReadItems(path string, filename string) ([]Item, error) {
	fp := filepath.Join(path, filename)
	data, err := os.ReadFile(fp)
	if err != nil {
		return []Item{}, err
	}
	var items []Item
	if err = json.Unmarshal(data, &items); err != nil {
		return []Item{}, err
	}
	return items, nil
}

var prioNameMap = map[int]string{
	1: "High",
	2: "Medium",
	3: "Low",
}

func PrettyPrint(items []Item) {
	w := tabwriter.NewWriter(os.Stdout, 1, 1, 5, ' ', tabwriter.AlignRight)
	fmt.Fprintln(w, "ID\tPriority\tText\tStatus\t")
	for _, item := range items {
		fmt.Fprintf(w, "%v\t%v\t%v\t%v\t\n",
			item.ID, prioNameMap[item.Priority], item.Text, item.Status,
		)
	}
	w.Flush()
}

func SetStatus(items []Item, ID string, status string) []Item {
	for i := 0; i < len(items); i++ {
		if items[i].ID == ID {
			items[i].Status = status
			PrettyPrint(items[i : i+1])
			return items
		}
	}
	fmt.Printf("ID: %v is not in the todo list.", ID)
	return items
}
