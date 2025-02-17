# enigma_golang
An emulator of the Enigma cryptographic machine, implemented in the Go language. 

#How does it work?

The program emulates a “civilian” version of the Enigma machine. It differs from the “military” version in that it does not have idempotency in encryption, which is realized by the reflector

Inside, 3 rotors and one stator are realized, each with 26 elements.

#How to run 

Install the golang compiler in your operating system

## Linux
For example, on Debian-base distributions, just type the terminal command:

sudo apt-get install golang
go run Enigma.go

## MacOS

You can install the golang compiler from Homebrew

brew install golang
go run Enigma.go
