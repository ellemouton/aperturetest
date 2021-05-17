# Aperture Dynamic Price Test

This is a very basic server that implements the `Prices` gRPC server defined in
`pricesrpc`. 

The purpose of this server is to test https://github.com/lightninglabs/aperture/pull/53 which adds the ability to specify for each service an endpoint to hit (via gRPC) to get price info for different resources of that service.

So in the  `aperture.yml` file, you will have something like this:

```
# With no tls:
...
services:
       ...
       usepricerserver: true
       pricerconfig:
            insecure:    true
            grpcaddress: "127.0.0.1:9001"

# With tls:
...
services:
       ...
       usepricerserver: true
       pricerconfig:
            grpcaddress: "127.0.0.1:9001"
            tlscertpath: "/path/to/cert/tls.cert" 
```

This allows the backend service to have different prices for the various
resources it provides. It can also change the prices without touching aperture. 

An example use case: If you have a blog service then you can define this service
in the aperture.yml file so that any request with path `/blog.*` will match.
Then you can freely add new blogs with incrementing indexes (`/blog/1`, `/blog/2`
etc). In the aperture.yml file you set `dynamicprices` to true and
`dynamicpriceendpoint` to the address where the blog service will serve the
price info service. Aperture will then call the gRPC service for price info and
the blog service can look in its DB (or do whatever) to get the price info for
the given resource. 

## Test it out:

- build `aperture` using https://github.com/ellemouton/aperture/pull/1
- see `aperture.yml` for how the service should be defined.
- run lnd, etcd and aperture
- then run this server: `go run cmd/server/main.go`
- hit the various endpoints (and pay the invoice to see it work):   
    - `localhost:8000/book/1` should cost 10 sats
    - `localhost:8000/book/2` should cost 20 sats
    - `localhost:8000/book/3` should be free
        
### TLS:
by default, this server will not use tls. To use it with TLS do the following:
- Run the genCerts script (use cmd-line options to set desired path):
        ```
        $ go run cmd/genCerts/main.go 
        ```
- Run the server  with the `--tls=true` flag
