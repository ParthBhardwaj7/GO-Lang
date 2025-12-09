# Simple Go HTTP Server

A minimal Go web server that serves static files and exposes two endpoints:

* `GET /hello` — returns a friendly "Hello!" message.
* `POST /form` — accepts a form with `name` and `address` fields .

This repository contains a `main.go` with the server code you provided and a `static/` directory for static assets (HTML form, CSS, JS).

--

Server will listen on port `8080`. Open `http://localhost:8080/` in your browser.

---


Create `static/index.html` with a simple form (example below) so `http.FileServer` serves it at root.

---

## `main.go` — file overview and why imports are used

```go
import (
    "fmt"
    "log"
    "net/http"
)
```

* `fmt` — used to format and write plain text responses to `http.ResponseWriter` (`Fprintln`, `Fprintf`).
* `log` — used to print fatal errors and for simple logging (e.g., `log.Fatal(err)` when `ListenAndServe` fails).
* `net/http` — core package that provides HTTP server and client implementations, request/response types, handlers, and utilities.

These are all standard library packages; no external dependencies are required for this small server.

---

## What each handler does (line-by-line explanation)

### `helloHandler`

* Checks the requested path: if not exactly `/hello`, returns `404 Not Found`.
* Ensures method is `GET`; otherwise returns `405 Method Not Allowed`.
* Writes `Hello!` to the response body on success.

Why check `r.URL.Path`? If you mount this handler at a pattern that matches more paths, the explicit path check prevents accidental handling of other routes.

### `formHandler`

* Ensures method is `POST`. Rejects other methods with `405`.
* Calls `r.ParseForm()` to parse form values from body (or URL for GET) and handles parse errors with `400 Bad Request`.
* Reads `name` and `address` via `r.FormValue("...")` and writes back a confirmation along with those values.

`FormValue` returns the first value for a key or an empty string if missing — consider validating required fields before using them.

### `main()`

* `http.FileServer(http.Dir("./static"))` serves files from the `static/` folder. Requests to `/` will serve `static/index.html` if present.
* `http.HandleFunc("/form", formHandler)` registers the form `POST` endpoint.
* `http.HandleFunc("/hello", helloHandler)` registers the `/hello` `GET` endpoint.
* `http.ListenAndServe(":8080", nil)` starts the HTTP server on port `8080`. `nil` uses the default `http.DefaultServeMux` where handlers were registered.

---


---

## Test with `curl`

```bash
# Test hello
curl -i http://localhost:8080/hello

# Test form POST
curl -i -X POST -d "name=Jul&address=Gurgaon" http://localhost:8080/form
```

Expected response for POST includes `POST request successful` and the echoed `Name` and `Address` lines.

---
