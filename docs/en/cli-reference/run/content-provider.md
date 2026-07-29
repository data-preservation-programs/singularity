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

   --enable-http-ipfs              Enable trustless IPFS gateway on /ipfs/ (default: true)
   --ipfs-cache-blocks value       Block cache capacity in blocks (~1MiB each); 0 sizes it to max-backend-reads * span-blocks * (1 + prefetch-spans) (default: derived)
   --ipfs-max-backend-reads value  Max concurrent source storage reads across all requests (default: 64)
   --ipfs-prefetch-spans value     Spans prefetched ahead on sequential access (default: 2)
   --ipfs-read-timeout value       Max duration of a single backend span read once it holds a connection slot (default: 2m0s)
   --ipfs-span-blocks value        Blocks (~1MiB each) fetched per backend read (default: 8)

   HTTP Piece Metadata Retrieval

   --enable-http-piece-metadata  Enable HTTP Piece Metadata, this is to be used with the download server (default: true)

   HTTP Piece Retrieval

   --enable-http-piece, --enable-http  Enable HTTP Piece retrieval (default: true)

   HTTP Retrieval

   --http-bind value  Address to bind the HTTP server to (default: "127.0.0.1:7777")

```
{% endcode %}
