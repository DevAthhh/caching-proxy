# Caching Proxy
This CLI program allows you to send GET requests to various urls, receive responses and cache them, with access from localhost:port, where the port is what you specified. All responses are cached for 24 hours thanks to Redis.

# Basic commands
- `caching-proxy --port 3030 --origin http://example.com` - allows you to run localhost:3030
- `caching-proxy --cls-cache` - clears the cache
