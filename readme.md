# Honey-Do

Usage: `honey-do filename [ pull (default) | undo | swap | add _"task"_ ]`

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