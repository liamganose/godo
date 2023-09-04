# godo
a todo app written in Go

## Running it
1. Checkout the code
2. Make sure your gopath is set and is added to the path
e.g.
```
export GOPATH="$HOME/gopath"
export PATH="$GOPATH/bin:$PATH"
```
3. Install locally via `go install /path/to/godo`
4. Run via `godo --help`

# Features (via --help)
```
A simple and efficient CLI-based todo app that
allows you to manage your tasks.

Usage:
  godo [command]

Available Commands:
  add         Adds a new todo item to the list
  completion  Generate the autocompletion script for the specified shell
  delete      Removes a specified todo item from the list
  done        Marks a specified todo item as done
  help        Help about any command
  list        Lists all current todo items
  start       Marks a specified todo item as in progress
  stop        Marks a specified todo item as to do

Flags:
  -c, --config string    godo config file
  -d, --datadir string   the directory to store godo data (default "/home/ganose")
  -h, --help             help for godo

Use "godo [command] --help" for more information about a command.
```
