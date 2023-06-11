# A silly web app: SplitDim

In the course of this lab we are going to build a Go web app that allows groups of people to keep track of money transfers between themselves and helps clear debs and credits with minimal money transfer. The app is by and large modeled after the excellent [SplitWise](https://www.splitwise.com) app, but it is much dumber so we will call it SplitDim. 

The below tasks walk you through writing a simple web app that implements the barebones SplitDim functionality with a basic local database. Later we will gradually extend the app to implement the 5 cloud native pillars. Each section contains tests that you can run to check whether you successfully completed all tasks in the section.

## Table of Contents

1. [A webapp skeleton]([#a-skeleton])
1. [Database API](#database-api)
1. [Local data layer](#local-data-layer)
1. [Transfer](#transfer)
1. [Clear](#clear)
1. [Reset](#reset)

## A skeleton

SplitDim helps housemates, trips, friends, and family members maintain their internal money transfers and keep track who owns who. Imagine you are at a trip with your friends, you invite one of your friends for a coffee, they pay the taxi fee for the entire group, and then someone else from the group pays your train ticket. After a while, it because practically impossible to keep track of. 

Enter SplitDim, a simple web app that allows friends to register their transfers (e.g., "Joe paid Alice's coffee for 5 USD", and then "Alice paid Joe's train ticket for 3 USD") and see (1) the current balance of each registered user (how much debt or credit they have) and (2) the minimal list of mutual money transfers that would allow them the clear all debts ("Alice would need to pay Joe 2 USD to clear the debt").

We are going to build SplitDim as a Go web app. During this lab we will write only the barebones web service that keeps the balances in memory; later we will extend it into a proper cloud-native app. The web service will implement 4 APIs:
- `POST: /api/transfer`: register a transfer between two users of a given amount (this API uses the POST HTTP method to let users post the transfer's details in JSON format),
- `GET: /api/accounts`: return the list of current balances for each registered user,
- `GET: /api/clear`: return the list of transfers that would allow users to clear their debts between themselves, and
- `GET: /api/reset`: reset all balances to zero.

So let's start, shall we?

1. Initialize a new Go project under `99-labs/code/splitdim`. Make sure you actually use this directory: there are some files placed there for you to help your work. 

   ``` sh
   cd 99-labs/code/splitdim
   go mod init github.com/<my-user>/splitdim
   go get github.com/stretchr/testify/assert
   go mod tidy
   ```

1. Open a new file called `main.go` and declare that you are going to build an executable.

   ``` go
   package main
   ```

1. Import the packages to be used.

   ``` go
   import (
       "log"
       "net/http"
   )
   ```
1. Implement 4 empty HTTP handlers: these will be the placeholders for the SplitDim API.

   ``` go
   // TransferHandler is a HTTP handler that implements the money transfer API.
   func TransferHandler(w http.ResponseWriter, r *http.Request) {}
   
   // AccountListHandler is a HTTP handler that returns the current balance of each registered user.
   func AccountListHandler(w http.ResponseWriter, r *http.Request) {}
   
   // ClearHandler is a HTTP handler that returns a list of transfers to clear the balance of each user.
   func ClearHandler(w http.ResponseWriter, r *http.Request) {}
   
   // ResetHandler is a HTTP handler that allows to zero out all balances.
   func ResetHandler(w http.ResponseWriter, r *http.Request) {}
   ```

1. Start the main function:

   ``` go
   func main() {
       // Set the default logger to a fancier log format.
       log.SetFlags(log.LstdFlags | log.Lshortfile)
   
       ... 
   }
   ```

1. Install a HTTP handler to serve a static HTML file with inline JavaScript to interact with the SplitDim API. This page implements the client GUI of our service so that you can connect from a browser, send transfers and see the current balances (of course real men use `curl`, but anyway). The HTML file is provided as part of this lab in order to allow you to concentrate on the server side, but feel free to modify it according to your liking.

   The next will register a static HTTP handler that will serve the prepackaged HTML file for the default path (`/`).

   ``` go
   http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
       http.ServeFile(w, r, "static/index.html")
   })
   ```

1. Register the 4 empty HTTP handlers for the 4 API endpoints.

   ``` go
   http.HandleFunc("/api/transfer", TransferHandler)
   http.HandleFunc("/api/accounts", AccountListHandler)
   http.HandleFunc("/api/clear", ClearHandler)
   http.HandleFunc("/api/reset", ResetHandler)
   ```

1. And finally start the HTTP server on port 8080. Remember, `http.ListenAndServe` will block until the program exits or an error happens: `log.Fatal` will write the error message to the standard output in the latter case.

   ``` go
   log.Println("Server listening on http://:8080")
   log.Fatal(http.ListenAndServe(":8080", nil))
   ```

Once ready, you can run the program with `go run main.go`: if all goes well you should see the output:

```
20XX/YY/ZZ 19:03:59 main.go:48: Server listening on http://:8080
```

This means your server is ready to accept HTTP requests.

> ✅ **Check**: 
>
> Run the below test to make sure that you have successfully completed the exercise. If all goes well, you should see the output `PASS`.
> ``` sh
> go test ./... --tags=main -run '/TestSkeletonAPIEndpoints'
> PASS
> ```
> Make sure the web service is running: the test issues requests to the HTTP server and checks whether the response is as expected.

At the moment all HTTP handlers respond to all HTTP methods, whereas our goal is for each API handler to accept only one HTTP method: `/api/transfer` should accept on HTTP POST requests, and all the other APIs should respond to `GET` requests only. Every other type of access should result a HTTP 405 error code ("Not Allowed"). This can be achieved by adding the following test to the beginning of your HTTP handlers. 

``` go
// SomeHandler accepts only POST requests ()
func SomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
        // Return HTTP 405
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}
```
Substitute `http.MethodPost` with `http.MethodGet` for handlers that accept only GET requests.

Another subtlety worth noting that `http.HandleFunc("/api/transfer", TransferHandler)` will route *all* HTTP requests whose path starts `/api/transfer` to the `TransferHandler`, e.g., `/api/transfer/random/api` and `/api/transfer/some/malicious/attack`. To make sure *only* the required API is served, add the below to the handler:

``` go
func handler(w http.ResponseWriter, r *http.Request) {
    if req.URL.Path != "/api/transfer" {
        http.NotFound(w, req)
        return
}
```

> ✅ **Check**: 
>
> Run the below test to make sure that you have successfully completed the exercise. If all goes well, you should see the output `PASS`.
> ``` sh
> go test ./... --tags=main -run '/TestSkeletonAPIMethods'
> PASS
> ```
> Make sure the web service is running: the test issues requests to the HTTP server and checks whether the response is as expected.

## Database API

## Local data layer

## Transfer

## Clear

## Reset database

## Deploy


> ✅ **Check**

<!--
Local Variables:
eval: (auto-fill-mode -1)
eval: (visual-line-mode t)
markdown-enable-math: t
End:
-->