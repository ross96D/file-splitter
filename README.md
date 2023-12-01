The simplest tool to split a file and join it again.

# Motivation
Sending files through the network in a splitted way and joining them on the other side. 
For this reason, this is both a package and a cli tool.

# Getting started
Install with `go install github.com/ross96D/file-splitter/cmd/file-splitter@latest`

If wanna use it as package `go get github.com/ross96D/file-splitter`

## Usage
To split a file:

`file-splitter split -o <output-file-path> -s <size-of-the-file> <path-of-the-file-to-be-splitted>`

To join a splitted file:

`file-splitter split -o <output-file-path> <path-of-the-file-to-be-join.part0>`