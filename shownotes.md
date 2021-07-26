eh?

# Protocol Buffers

## What are Protocol Buffers?

* Developed by Google as its open-source Mechanism for serializing structured data
* an alternative to XML and JSON (behaves similarly)
* Used as the de facto Interface Definition Language (IDL) and message interchange format for gRPC
* designed to be language agnostic

## Why use Protocol Buffers? (as opposed to XML or JSON)

* The Size difference
  1. provice concrete examples
* Speed (XML is 20-100 times slower as it gets asymtotically larger)
* its used exclusively in gRPC

## what is gRPC

* 
* In gRPC, a client application can directly call a method on a server application on a different machine as if it were a local object

## How is it implemented?

The " = 1", " = 2" markers on each element identify the unique "tag" that field uses in the binary encoding. Tag numbers 1-15 require one less byte to encode than higher numbers, so as an optimization you can decide to use those tags for the commonly used or repeated elements, leaving tags 16 and higher for less-commonly used optional elements. Each element in a repeated field requires re-encoding the tag number, so repeated fields are particularly good candidates for this optimization.
