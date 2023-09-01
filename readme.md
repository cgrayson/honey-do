[![Go package](https://github.com/cgrayson/honey-do/actions/workflows/test.yml/badge.svg)](https://github.com/cgrayson/honey-do/actions/workflows/test.yml)

# Honey-Do

Usage: `honey-do [ pull (default) | unpull | swap | add _"task"_ ] filename`

A default honey-do file can be set via the environment variable `HONEY_DO_FILE`. 
When set, no `filename` is needed on the command line. If one is given, it will
override the environment variable.

## Development

- `go test` to run test suite

---

original notes (4/17/23)

> honey.do
> 
> a random, limited-view to-do list
> 
> - local file stored in `/Library/Mobile Documents/com~apple~CloudDocs/`
> - markdown with checkbox style, one item per line
> - metadata, like added-date & completed-date, stored in curly brackets, JSON-style (`{added_date: "Mon. Apr. 17", etc.}`)
> - Go app
> - make library for using local file as simple database
> - **pull** (default): reads the unfinished items & returns one random one, marking it done (add `x` & move to top of bottom)
> - **add**: add a new item
> - **undo**: find latest done item & undo it
> - **re-pull**: “undo”, plus picks another (excluding the just undone)
