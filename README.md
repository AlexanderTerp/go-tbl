# Go Tbl -- ASCII Table Writer for Go

This is a fork of [olekukonko/tablewriter](https://github.com/olekukonko/tablewriter).

This fork exists because the original is no longer actively maintained, but in developing [Rad](https://github.com/amterp/rad), I've hit several things I've wanted to change. The easiest thing for me is to fork and make the changes I want, with the primary purpose being for use in Rad.

This is not intended as a 'maintained' version of the original for others to rely on -- I might very well strip out features that I don't need for Rad.

## Installation

In your Go project:

```shell
go get github.com/amterp/go-tbl
```

## Dev

### Release Process

1. Commit your changes.
2. `git tag <version>` e.g. `git tag v0.5.0`
3. `git push origin main --tags`
