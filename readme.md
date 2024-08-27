# Go cli typing game 
## Description
This is a simple typing game that you can play in your terminal. 
It fetches random quotes with optional min or max length parameters from a public API.

I built this app to practice my Go skills and to have a fun typing game that I can play in my terminal.

It uses the [Bubbletea](https://github.com/charmbracelet/bubbletea) library to create the interactive UI and the [lipgloss](https://github.com/charmbracelet/lipgloss) library to style the text and position the text in the screen.


## Installation
To install the app you can run the following command in your terminal:
```bash
curl https://raw.githubusercontent.com/ProductionPanic/go-typingtest/main/install.sh | bash
```

This will download and install the cli app in `/usr/local/bin/type-test` and you can run the app by typing `type-test` in your terminal.

This requires that you have `curl`, `go` and `git` installed on your system.

## Usage

To start the game you can run the following command in your terminal:
```bash
type-test
```

This will fetch a random quote from the API and then you can start typing. You can also specify a minimum or maximum length of the quote by passing the `--min` or `--max` flag followed by the length you want. For example:
```bash
type-test --min 50 # This will fetch a quote with a minimum length of 50 characters
type-test --max 100 # This will fetch a quote with a maximum length of 100 characters
```

A `--help` flag is also available to show the help message.

## API
This app uses the free and public API at [https://api.quotable.io/](https://api.quotable.io/) to fetch random quotes. For more information on the API click the link.




