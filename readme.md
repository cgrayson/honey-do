[![Go package](https://github.com/cgrayson/honey-do/actions/workflows/test.yml/badge.svg)](https://github.com/cgrayson/honey-do/actions/workflows/test.yml)

# Honey-Do

Honey-do is a task/to-do list meant to imitate the old-fashioned "job-jar". 
Instead of writing tasks on a slip of paper that's folded up and stuffed
into a jar to be pulled out at random, this simple command-line app manages
lines in a markdown-flavored plain text file of your choosing. You can add
new tasks, pull a random one, swap the last one for another, or just put 
the last one back if you change your mind. 

Usage: `honey-do [ pull (default) | unpull | swap | add "task" ] filename`

A default honey-do file can be set via the environment variable `HONEY_DO_FILE`. 
When set, no `filename` is needed on the command line. If a filename is given
on the command-line, it overrides the environment variable.

## about

I had the brilliant idea for this groundbreaking app all by myself, and 
likewise have implemented it here, from scratch and with no intelligence other
than my own, as a pastime and to apply what I was learning about developing
in Go.