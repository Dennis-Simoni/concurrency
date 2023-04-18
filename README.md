# Concurrency with Go #

Making progress on more than one task simultaneously is known as concurrency. 
Go has rich support for concurrency using goroutines and channels.

This is a tutorial repository which aims to help developers new to Go with understanding
how all this works.

The following concepts are covered:
* Goroutines
* Channels
* Chat server - demo concurrent program

To run the chat server, cd in the directory and run:
```
❯ make run
```

Once main is running, you can create N clients
To create a client, cd in the directory and run:
```
❯ make client
```