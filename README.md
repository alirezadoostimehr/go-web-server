# go-web-server
Trying to warm my hands for coding bigger projects with GO!

I want this server to respond only to 4 subdomains

1. "/hello" for this address it should only prints "Hello" and returns ok status code

2. "/bye" for this address it should returns bad gateway status code

3. "/name1" which gets name as url variable and returns "hello {name}"

4. "/name2" which gets name in request body and returns "hello {name}"