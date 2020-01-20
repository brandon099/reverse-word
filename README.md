## Reverse Words

Given a list of space separated words, reverse the order of the words. Each line of text contains L letters and W words. A line will only consist of letters and space characters. There will be exactly one space character between each pair of consecutive words.

### Input

The first line of input gives the number of cases, N.
N test cases follow. For each test case there will a line of letters and space characters indicating a list of space separated words. Spaces will not appear at the start or end of a line.

### Output

For each test case, output one line containing "Case #x: " followed by the list of words in reverse order.

### Limits

Small dataset

```
N = 5
1 ≤ L ≤ 25
```

Large dataset

```
N = 100
1 ≤ L ≤ 1000
```

### Sample

| Input           | Output                   |
| --------------- | ------------------------ |
| 3               |                          |
| this is a test  | Case #1: test a is this  |
| foobar          | Case #2: foobar          |
| all your base   | Case #3: base your all   |

### Running
This program can be run with `go run reverse-word.go`, or compliled and run with `go build && ./reverse-word`.
It is basic in functionality to explicitly meet the requirements as outlined in the challenge. It could be made more robust with logging, better error handling, and better content detection/file parsing.
