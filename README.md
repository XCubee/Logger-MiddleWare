# Logger Middleware in GoLang

This project demonstrates how i created a middleaware logger which can be used to check for the the latency of the program and the working. I had to add a time delay of 2seconds
else the terminal wouldnt show the time properly. 

## Features

- Logs HTTP method (GET, POST, etc.)
- Logs URL path
- Logs the time taken to handle the request (duration)
- Can be used with any Go web server

## How It Works

The `logger` middleware is a function that wraps around an HTTP handler. It performs the following:

1. Logs the start of the request with the HTTP method and URL path.
2. Passes the request to the next handler in the chain.
3. After the request is processed, it logs the time taken to process the request.

## Prerequisites

- [Go](https://golang.org/doc/install) must be installed. You can verify it by running:

  ```bash
  go version


![Screenshot 2025-04-24 060805](https://github.com/user-attachments/assets/e5bbc0c9-6ee4-4b0f-a026-f0f85cc3cfdd)
![Screenshot 2025-04-24 061002](https://github.com/user-attachments/assets/590ffe4e-9442-46d3-8465-332591fc3ee8)
