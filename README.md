# Chip-8 Interpreter

A chip-8 interpreter written in Go.

This is just a hobby project for me to get more familiar with the go programming language and with hardware emulation.

## Dependencies
Requires Go 1.22 or later

The Frontend for the emulator uses the Ebitengine game engine.
Ebitengine requires a C compiler like gcc or clang when installing on MacOS or Linux.

## Building
Build the executable by navigating to the root directory .../gochip/
and run the command:

> $ go build

## Running the emulator
After building, start the executable by running:
> ./gochip \<delay\> \<rom path\>

\<delay\> determines how fast the emulator runs and can be adjusted to make a program run at an appropriate speed. A higher value means the emulator runs slower.

Example:
> $ ./gochip 5 roms/maze.ch8

## Resources used:
I used [Austin Morlan's](https://austinmorlan.com/posts/chip8_emulator/)
 blogpost as a guide for my own implementaion and it was very helpful!!!

### Roms
I've not written any of the roms available in this repo myself. They are from the following Repos:

[dmatlack](https://github.com/dmatlack/chip8/tree/master)
- Good collection of various chip8 roms!

[corax89](https://github.com/corax89/chip8-test-rom)
- Great for testing the basic instructions work as intended 
