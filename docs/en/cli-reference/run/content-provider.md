# Start a content provider that serves retrieval requests

{% code fullWidth="true" %}
```
NAME:
   singularity run content-provider - Start a content provider that serves retrieval requests

USAGE:
   singularity run content-provider [command options]

OPTIONS:
   --help, -h        show help
   --no-automigrate  skip automatic database migration and correctness checks on startup; only use if you run 'admin init' on every upgrade or manually before starting daemons (default: false)

   HTTP IPFS Gateway

   --enable-http-ipfs  Enable trustless IPFS gateway on /ipfs/ (default: true)

   HTTP Piece Metadata Retrieval

   --enable-http-piece-metadata  Enable HTTP Piece Metadata, this is to be used with the download server (default: true)

   HTTP Piece Retrieval

   --enable-http-piece, --enable-http  Enable HTTP Piece retrieval (default: true)

   HTTP Retrieval

   --http-bind value  Address to bind the HTTP server to (default: "127.0.0.1:7777")

```
{% endcode %}
