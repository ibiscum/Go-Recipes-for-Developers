# Chapter 1

## Creating a module

### Page 4

The sequence of commands on page 4 should be:

```
$ cd projects
$ mkdir webform
$ cd webform
$ go mod init github.com/examplecompany/webform
```
This, of course creates the go.mod file:

```
module github.com/examplecompany/webform/
go 1.22.0
```

The go.mod printed in the book points to the actual repository for
this chapter.

# Chapter 2

## Combining strings

### Page 23

The text refers to `string.Builder`. That should be `strings.Builder`.
