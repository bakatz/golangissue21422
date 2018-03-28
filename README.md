# golangissue21422
## What is this?
This repository is used to demonstrate that https://github.com/golang/go/issues/24122 is a thing.
## Steps to reproduce https://github.com/golang/go/issues/24122 using this repository
0. Install `go@1.10`:
```
ben@BEN-DESKTOP MINGW64 C:/Users/ben/dev/golangissue21422
$ go version
go version go1.10 windows/amd64
```
1. Run `go test -v ./...`
2. Observe the output:
```
ben@BEN-DESKTOP MINGW64 C:/Users/ben/dev/golangissue21422
$ go test -v ./...
Hey, I ran! The current time is 2018-03-27 17:34:08.1723502 -0700 PDT m=+0.002000201 and the contents of the sql file is: select 1
=== RUN   TestSomethingElseReturnsTrueWhenTrueIsPassed
--- PASS: TestSomethingElseReturnsTrueWhenTrueIsPassed (0.00s)
=== RUN   TestSomethingElseReturnsFalseWhenFalseIsPassed
--- PASS: TestSomethingElseReturnsFalseWhenFalseIsPassed (0.00s)
=== RUN   TestSomethingReturnsTrueWhenTrueIsPassed
--- PASS: TestSomethingReturnsTrueWhenTrueIsPassed (0.00s)
=== RUN   TestSomethingReturnsFalseWhenFalseIsPassed
--- PASS: TestSomethingReturnsFalseWhenFalseIsPassed (0.00s)
PASS
ok      _/C_/Users/ben/dev/golangissue21422     0.047s
```

3. Open `some_migration.sql` and change `select 1` to `select 2`; save the file.
4. Re-run `go test -v ./...`
5. Observe that the output is exactly the same as before i.e. the wrong `select` query is displayed and the code in `main_test.go` has not actually executed again:
```
$ go test -v ./...
Hey, I ran! The current time is 2018-03-27 17:34:08.1723502 -0700 PDT m=+0.002000201 and the contents of the sql file is: select 1
=== RUN   TestSomethingElseReturnsTrueWhenTrueIsPassed
--- PASS: TestSomethingElseReturnsTrueWhenTrueIsPassed (0.00s)
=== RUN   TestSomethingElseReturnsFalseWhenFalseIsPassed
--- PASS: TestSomethingElseReturnsFalseWhenFalseIsPassed (0.00s)
=== RUN   TestSomethingReturnsTrueWhenTrueIsPassed
--- PASS: TestSomethingReturnsTrueWhenTrueIsPassed (0.00s)
=== RUN   TestSomethingReturnsFalseWhenFalseIsPassed
--- PASS: TestSomethingReturnsFalseWhenFalseIsPassed (0.00s)
PASS
ok      _/C_/Users/ben/dev/golangissue21422     (cached)
```
