# VecM Language Revision 1.0 - Current state
## Current state of the language, and plans for it's improvement
So far development of the language has entirely been an exercise in compiler development. No thought was put into the design. Now that I have a pretty good idea how to build a compiler, I can start thinking about the language itself.

### Syntax
I came up with the syntax as I went, I did not much care about anything except implementing individual features. The syntax is a mess, and I need to clean it up.

Misc syntax issues:

- `;` is the statement terminator - it is not defined if it's mandatory or when. The parser is the judge of that, don't make it angry.

#### Types
The language has the following types:
- `int8` or `char` - 8-bit signed integer
- `int16` - 16-bit signed integer
- `int32` or `int` - 32-bit signed integer
- `int64` - 64-bit signed integer
- `uint8` - 8-bit unsigned integer
- `uint16` - 16-bit unsigned integer
- `uint32` - 32-bit unsigned integer
- `uint64` - 64-bit unsigned integer
- `float32` - 32-bit floating point number
- `float64` - 64-bit floating point number
- `bool` - boolean
- `string` - string
- `array` - array
- `hash` - hash
- `null` - null
- `void` - void
- Named types (structs, enums, etc.) were planned but aren't implemented yet.

Issues with types:
- `int` is a 32-bit signed integer, but `int32` is also a 32-bit signed integer. I need to decide which one to keep.
- `int8` and `char` are the same thing, I need to decide which one to keep.
- `string` is a type, but it's not clear how it's implemented. I need to decide if it's a primitive type or a class.
- `array` and `hash` are types, but it's not clear how they're implemented. I need to decide if they're primitive types or classes.
- `null` - I'm not even sure if this is a usable type, I think it's just there cause I invented it

Additional notes on types:
- I think I like shorter type names so `i8` instead of `int8`, `i16` instead of `int16`, etc.

Type annotations are done in several ways, and can be implicit:
- `const int a = 1;`
- `const a: int = 1;`
- `const a = 1;`
- `let a = 1;`
- `a := 1`

### Memory model
The language started with an interpreter, written in go, using Go's garbage collector. The compiler currently only manages local variables on the stack, and doesn't support arrays, strings, hashes, heap allocation, etc. I need to decide on a memory model for my compiler and implement it. It cannot remain as it is now with the interpreter.

### Errors
Parser errors are impossible to understant. They need to be human readable. The issue isn't in just pretty-printing and nicer error messages. The issue is errors like:

```
parser error: Parser error at line 1, column 1:
export const main = fn() -> void {
^
invalid token in declaration statement. expected="=" got="fibSize"
```

Can you guess what the issue here is?
```
export const main = fn() -> void {
    [] // <-- This random token is the issue
    const int fibSize = 50;
```
The message is correct `invalid token` although I'd like to change that to `Unexpected Token`, but the line and column are wrong, and I have no idea why `=` is expected here.


Compiler errors are in a similar state.

New rule: Every time an error is bad, add it to a .md file in this directory, and fix those later.

### Compiler Logs
No compiler logs are generated - this sucks, as it would ease development significantly.

TODO: Add logging system - with different log levels
New rule: every new method has to consider logging

### Standard Library
There is no standard library or ability to call external functions. This is a must-have feature.

### Module/Packages system
A module/package system exists partly in the system but isn't implemented yet. 

### Documentation
There is no documentation, and the code is not well commented.

TODO: Enforce documentation with a linter.

### Code style
There is no coding standards, no linter, and no formatter.

TODO: Add linter and formatter, and apply to entire codebase.

