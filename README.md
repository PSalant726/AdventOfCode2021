# Advent of Code 2021

You're minding your own business on a ship at sea when the overboard alarm goes off! You rush to see if you can help. Apparently, one of the Elves tripped and accidentally sent the sleigh keys flying into the ocean!

Before you know it, you're inside a submarine the Elves keep ready for situations like this. It's covered in Christmas lights (because of course it is), and it even has an experimental antenna that should be able to track the keys if you can boost its signal strength high enough; there's a little meter that indicates the antenna's signal strength by displaying 0-50 stars.

Your instincts tell you that in order to save Christmas, you'll need to get all fifty stars by December 25th.

Collect stars by solving puzzles. Two puzzles will be made available on each day in the Advent calendar; the second puzzle is unlocked when you complete the first. Each puzzle grants one star. Good luck!

## Usage

Each package in this repository represents a single day on the advent calendar.

1. Ensure you have [Go](https://go.dev/doc/install) installed locally
    > Alternatively, install via [Homebrew](https://formulae.brew.sh/formula/go)
1. Clone this repository to your local machine
1. In your terminal, navigate to the repository's root directory
1. Run:

    ```sh
    go run ./XX/... # Where XX is a package directory name
    ```

> Each package consumes my personal puzzle input by default. To execute the code for your own puzzle input values, simply replace the contents of `input.go` in the package of concern. Pay close attention to the type(s) and formatting of input, and be sure to maintain these patterns.

### Example

```sh
# With the included puzzle input...
$ go run ./01/...
#=>
# Part 1: 1759
# Part 2: 1805
```
