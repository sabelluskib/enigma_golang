# enigma_golang
An emulator of the Enigma encryption machine, implemented in the Go language. 

# How does it work?

The program emulates a “civilian” version of the Enigma machine. It differs from the “military” version in that it does not have idempotency in encryption, which is realized by the reflector

Inside, 3 rotors and one stator are realized, each with 26 elements.

# How to run 

Install the golang compiler in your operating system. The program will request at runtime the starting positions of the rotors, as well as the incoming and outgoing files. 

The incoming file must be in UTF-8 encoding. For example, the program encrypts the book “Martin Eden” by Jack Londan in 7 seconds.

## Linux
For example, on Debian-base distributions, just type the terminal command:

`sudo apt-get install golang`
`go run Enigma.go`

## MacOS

You can install the golang compiler from Homebrew

`brew install golang`
`go run Enigma.go`
