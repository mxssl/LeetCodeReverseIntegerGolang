# Reverse Integer

Given a 32-bit signed integer, reverse digits of an integer.

`Time Complexity: O(log(x))`

`Space Complexity: O(1)`

```Bash
$ go test -v -cover  ./...
=== RUN   TestCase1
Expected: 654321, Output: 654321
--- PASS: TestCase1 (0.00s)
=== RUN   TestCase2
Expected: -654321, Output: -654321
--- PASS: TestCase2 (0.00s)
=== RUN   TestCase3
Expected: 0, Output: 0
--- PASS: TestCase3 (0.00s)
=== RUN   TestCase4
Expected: 0, Output: 0
--- PASS: TestCase4 (0.00s)
PASS
coverage: 87.5% of statements
ok      github.com/mxssl/reverseInteger 0.148s  coverage: 87.5% of statements
```

