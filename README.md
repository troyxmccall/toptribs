# toptribs
Show the top contributors for specified files.

## Installation

```
$ go get github.com/troyxmccall/toptribs
```

## Usage

```
$ toptribs main.go
Troy McCall
```

```
$ toptribs main.go hello.go
Troy McCall
```

`-n` option specifies the number of top contributors.

```
$ toptribs -n 3 main.go hello.go
Troy McCall
Brett Buddin
Joe Biden
```
