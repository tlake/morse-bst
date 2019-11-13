# Morse BST Decoder

A quick program written in Go based off the idea of structuring Morse code into a BST.
Given a Morse code pattern, the encoded character can be found by traversing the tree based upon whether the given Morse character is a dot or a dash.

Illustration of Morse BST:

```
        <-- DOT                 DASH -->
                     start
                 /           \
           E                         T
        /     \                   /     \
      I         A             N            M
    /   \      / \          /   \         / \
  S       U   R   W       D       K      G   O
 / \     /   /   / \     / \     / \    / \
H   V   F   L   P   J   B   X   C   Y  Z  Q
```

Example: the Morse pattern `.--`.

- From `start`, the first character `.` tells us to move left to the `E`.
- From `E`, the next character `-` tells us to move right to `A`.
- From `A`, the next character `-` tells us to move right to `W`.
- That's the end of the Morse pattern, so we're done at `W`!

## Usage

Download a binary from [releases](https://github.com/tlake/morse-bst/releases), or clone the repo and run from source yourself (`go build && ./morse-bst` or `go run main.go`).
