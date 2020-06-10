
Damerau-Levenshtein
================================================================================

[![GoDoc](https://godoc.org/github.com/lmas/Damerau-Levenshtein?status.svg)](https://godoc.org/github.com/lmas/Damerau-Levenshtein)

Calculate and return the True Damerau–Levenshtein distance of string A and B.

See [Issue #2](https://github.com/lmas/Damerau-Levenshtein/issues/2) for a list
of research papers used as a reference when implementing the algorithm and a
minor struggle trying to verify the results.

Example
--------------------------------------------------------------------------------

Install package:

        go get github.com/lmas/Damerau-Levenshtein

Get distance between string A and B:

        package main

        import (
                "fmt"
                tdl "github.com/lmas/Damerau-Levenshtein"
        )

        func main() {
                dist := tdl.Distance("CA", "ABC")
                fmt.Println(dist)
        }

Or if you want more performant code and want to avoid memory allocations:

        func main() {
                // Set a max length of 100 characters for strings A and B
                t := tdl.New(100)
                dist := t.Distance("CA", "ABC")
                fmt.Println(dist)
                dist := t.Distance("AC", "ABC")
                fmt.Println(dist)
        }

Take a look at the tests for more examples.

License
--------------------------------------------------------------------------------

MIT License, see the LICENSE file.

