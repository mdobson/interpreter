# Interpreter

This project is just a simple interpreter for a tiny scripting language. Basically just a way for me to brush up on golang. This is based on the Writing an Interpreter in Go by Thorsten Ball. I'll probably update this language for other things that interest me as time goes on, but for now the language design and structure of code come from them.

## Sample script 

The interpreter will parse and execute a c-like scripting language. The example is below.

```
let five = 5;
let ten = 10;

let add = fn(x, y) {
    x + y;
};

let results = add(five, ten);
```

## REPL

This tool comes with a REPL to test commands out. Currently it only prints the lexed code. Run the REPL with the following command:

```
$ go run main.go
```

The REPL can be exited with 

```
>$ exit
```

## Tests

The book I am working off of also includes tests using the standard way of unit testing in golang. Simply run:

```
$ go test ./<DIRECTORY>
```

to get test outputs. One concrete command to see this can be seen here:

```
$ go test ./lexer/
```

