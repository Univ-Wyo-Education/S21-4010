
Assignment 3
===

200pts


Pt 2: Basic Client Server
==============================

This part of the homework is to add a very simple API call to a client server.  The server is in ./simple-server, the sample client is in
./simple-client.

Modify the sever to have a `/double` that will take a value that it is passed on a `GET` call and double that value.  The return the
doubled value to the client in a JSON format.

So...

```
	http://localhost:3000/double?value=12
```

will return

```
	{"double":24}
```

You will need to set the MIME type of the return to JSON with

```
	www.Header().Set("Content-Type", "application/json; charset=utf-8")
```

Use curl or wget to test this call.

```
	wget 'http://localhost:3000/double?value=100'
```

or

```
	curl 'http://localhost:3000/double?value=100'
```

Test this with a request from a browser and with the `./simple-client` code.


## Submit

1. a copy of the 2 modified .go files in for the client and the server.

