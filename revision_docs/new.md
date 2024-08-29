# VecM Language Revision 1.0 - New Syntax

the [current state](current.md) of the language is relevant to this document.

## Types

### Primitive types
The language has the following primitive types:
- `i8` - 8-bit signed integer
- `i16` - 16-bit signed integer
- `i32` - 32-bit signed integer
- `i64` - 64-bit signed integer
- `u8` - 8-bit unsigned integer
- `u16` - 16-bit unsigned integer
- `u32` - 32-bit unsigned integer
- `u64` - 64-bit unsigned integer
- `f32` - 32-bit floating point number
- `f64` - 64-bit floating point number
- `bool` - boolean.

Default values for arrays are 0 for numbers, false for bools, and null for pointers.

### Named types
Structs and enums are not planned in revision 1.0.
They may get partially implemented if I so choose

