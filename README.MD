# Fib in Go [![Build Status](https://travis-ci.org/lloydmeta/fibber.svg?branch=master)](https://travis-ci.org/lloydmeta/fibber)

Obligatory Fib project to test the waters in Golang.

Provides:

- Standard "count up wards" iterative `i`th Fibonacci finder
- `ForEach` version that allows you to pass a callback function to work with each fibonacci number generated
- `Memoed` threadsafe Fibonacci number generator that, well, memoises.
- Works as a lib and as a command-line executable

## Areas explored

- Library packaging
    - With binary mode
- Testing
    - Travis
- Documentation
    - Example Docs
    - Godocs
- Visibility
    - Private/Public members, structs, methods
- Basic types, structs
    - Implemented simple integer linked list
    - Implemented a memoized Fibonacci generator with private members
- Sync package (used Mutex, switched to RWLock for Memoed Fib gen)
    - Use of `defer` for unlocking (see `fib.Memoed`) to avoid raciness
    - How to avoid using `defer` for unlocking *and* avoid raciness

## Installing and Using

```bash
$ go get -u github.com/lloydmeta/fibber/cmd/fib

// Command line mode
$ fib 10
Fib of 10 is 55

// Repl mode (uses memoised Fibonacci generator)

$ fib

**************** Interactive Mode ********************
 ctrl+c or ctrl+d or enter any invalid number to exit
******************************************************

Which Fibonacci number would you like to know?
10
Fibonacci[10] is 55

Which Fibonacci number would you like to know?
1000
Fibonacci[1000] is 43466557686937456435688527675040625802564660517371780402481729089536555417949051890403879840079255169295922593080322634775209689623239873322471161642996440906533187938298969649928516003704476137795166849228875

Which Fibonacci number would you like to know?
```