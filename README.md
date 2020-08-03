A struct named messageHandler is created. 
To make this type implement Handler, a method with the signature ServeHTTP(http.ResponseWriter, *http.Request) is implemented.
You can implement an interface type into concrete types by providing the methods based on the method signature defined by the interface. 
The receiver method is added as a messageHandler type into the ServeHTTP function to make it a method of the messageHandler struct. 
In the ServeHTTP method, a string message is returned as the HTTP response, taking data from the struct field message.
In the main function, instances of the messageHandler struct (the & symbol is used to yield a pointer) are created and then ServeMux.
Handle is called to register handlers with the messageHandler struct instances.
 If there are requests for the URL path "/message1" and "/message2", the ServeHTTP method of messageHandler does all processing at the server. 
You can also reuse the custom handlers. In the example, messageHandler is used as the handler for two URL paths.