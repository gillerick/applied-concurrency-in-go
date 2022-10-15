## Applied Concurrency in Go
### HTTP request/response cycle
On the internet, information is mostly exchanged via HTTP. The exchange begins with the client sending out an HTTP
request to the server to ask for information. If the server recognizes the request, it begins processing it by invoking
any backend code required to fulfill the request. The invoked functions are typically called _request handlers_. Once
the backend processing completes, the server finishes the information exchange by wrapping the information into an HTTP
response and sending it back to the client.