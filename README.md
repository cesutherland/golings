# golings

## Acuity Introduction

We've forked this tool from [golings](https://github.com/mauricioabreu/golings) to use as an interactive introduction to Golang. We've added more information to various 
READMEs and created several new sections. **We strongly recommend doing the exercises in order using the `watch` command
as described below**.

Should you choose not to, here is the recommended order:

* variables
* functions
* if
* switch
* primitive types
* arrays
* slices
* maps
* loops
* structs
* pointers
* methods
* interfaces
* errors
* anonymous functions
* generics

Any feedback or questions, please reach out in [#acuity-golang](https://squarespace.enterprise.slack.com/archives/C08488AM16D)


![gopher](misc/gopher-dance.gif)

> rustlings but for golang this time

You may know [rustlings](https://github.com/rust-lang/rustlings), right? If you don't, please go ahead and check out.
`rustlings` is awesome. It is a CLI app designed to teach the awesome Rust programming language through exercises.

`golings` has the very same idea, but for the [Go programming language](https://go.dev/)

After setting up all the tools required to run `golings` you have the task to fix tiny go programs.

## Installing

You need to have `go` installed. You can install it by visiting the [Go downloads page](https://go.dev/dl/)

## Doing exercises

All the exercises can be found in the directory `golings/exercises/<topic>`. For every topic there is an additional README file with some resources to get you started on the topic. We really recommend that you have a look at them before you start.

Now you have the task to fix all the programs. Some of them don't compile, and you need to fix them. Some of them compile, but have tests and you need to write some code to have them all green (these are the `compile` and `test` modes).

Clone the repository:

```sh
git clone git@github.com:sqsp-scratchpad/acuity-golings.git
```

To run the exercises in the recommended order while taking advantage of fast feedback loop, use the `watch` command:

```sh
make watch
```

This command will run golings in interactive mode. Every time you save a file it will verify if the code is correct
and move on to the next exercise.

If you need a hint, you can use the `hint` command while `watch` is running.

If you want to see the list of exercises, you can use the `list` command.

To exit `watch`, use the `exit` command.

To manually run the next pending exercise:

```sh
golings run next
```

If you want to run a single exercise:

```sh
golings run variables1
```

In case you are stuck and need a hint:

```sh
golings hint variables1
```

To list all exercise while checking your progress:

```sh
golings list
```

To compile and run all the exercises:

```sh
golings verify
```

If you need help with CLI commands:

```sh
golings --help
```

A demo running the command `golings run <exercise name>`

![demo](misc/demo.gif)


## Learning resources

* [Golang official tutorial](https://go.dev/doc/tutorial/)
* [Go by example](https://gobyexample.com)
