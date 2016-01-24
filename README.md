# toptribs
Show the top contributor(s) for specified files.

This is helpful if you want to know who's the expert(s) on whatever you're working on.

## Installation

```
$ go get github.com/troyxmccall/toptribs
```

## Usage

```
$ toptribs hello.go
Troy McCall
```

```
$ toptribs main.go hello.go
Go Gopher
```

`-n` option specifies the number of top contributors.

```
$ toptribs -n 3 gopher.go
Brett Buddin
Chris Gutierrez
Dave Cheney
```
