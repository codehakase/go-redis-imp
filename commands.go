package main

type Command struct {
	Fields   []string
	Response chan string
}
