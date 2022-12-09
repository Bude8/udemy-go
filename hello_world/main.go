package main
// Package == Project == Workspace
// First line of each file must contain package it belongs to
// Excecutable package - generates a file that we can run
// Reusable package - code used as 'helpers' - good place to put reusable logic
// How do we know which we're making?
// 'main' specifically makes executable package, must have a func called main
// Any other package name gives nothing

import "fmt";

func main() {
    fmt.Println("Hi there!")
}
