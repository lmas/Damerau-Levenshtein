
Damerau-Levenshtein
================================================================================

Calculate and return the true Damerau–Levenshtein distance of string A and B.

See https://en.wikipedia.org/wiki/Damerau%E2%80%93Levenshtein_distance#Distance_with_adjacent_transpositions
for more information.

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

License
--------------------------------------------------------------------------------

MIT License, see the LICENSE file.

