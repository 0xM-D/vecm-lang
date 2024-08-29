# VecM Language Revision 1.0 - New Syntax

the [current state](current.md) of the language is relevant to this document.


## Memory model
The memory model is probably what motivated this revision the most. So far, we've been relying on Go's garbage collector. We will no longer be doing that.

The language needs to be able to process large amounts of data in large chunks, but general runtime performance is not a priority.

We want safer memory management than C - we will have a garbage collector.

Goal is high interoperability with C libraries, we will have to support all intrinsic c types in some capacity.

Let's use [MPS](https://www.ravenbrook.com/project/mps/) for memory management.

All types will be boxxed - since performance on simple types is not a priority, and we want to avoid the complexity of having both boxed and unboxed types.

All memory objects will be stored in the heap. The stack will only store references to heap objects. The stack will entirely rely on the OS's virtal memory management.

Let's make all our objects word-aligned - will make simd stuff much easier.

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

### Arrays
Arrays will not have fixed capacity, and can have any type.
Multidimensional arrays will be supported.
Depending on the shape of the array, simd operations may be possible.

### Strings
Strings will be implemented as arrays of `u8` with a null terminator. They will be mutable, with variable length.
Will be treated similarly to arrays when it comes to simd operations, but will have additional functions for string manipulation.

### Hashes
Hashes will be implemented as arrays of key-value pairs. The keys will be strings, and the values will be any type.
Hashes will be mutable, with variable size;


