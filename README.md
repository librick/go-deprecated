# Go Deprecated Experiment
A simple Go module to determine when deprecated functions are explicitly tagged as `DEPRECATED` on pkg.go.dev.

## Motivation
[go-libp2p](https://github.com/libp2p/go-libp2p/) has some functions that are explicitly marked as deprecated
via a newline followed by a line starting with `// Deprecated:`. E.g.,
```
// EnableAutoRelay configures libp2p to enable the AutoRelay subsystem.
// ...
//
// Deprecated: Use EnableAutoRelayWithStaticRelays or EnableAutoRelayWithPeerSource
func EnableAutoRelay(opts ...autorelay.Option) Option {
    // implementation here…
}
```
pkg.go.dev tags these function as [DEPRECATED](https://pkg.go.dev/github.com/libp2p/go-libp2p#EnableAutoRelay).

But there are other functions that are marked as deprecated in the source code
but are not tagged as `DEPRECATED` on pkg.go.dev. E.g.,
```
// DialRanker configures libp2p to use d as the dial ranker. To enable smart
// dialing use `swarm.DefaultDialRanker`. use `swarm.NoDelayDialRanker` to
// disable smart dialing.
// Deprecated: use SwarmOpts(swarm.WithDialRanker(d)) instead
func DialRanker(d network.DialRanker) Option {
    // implementation here…
}
```
In this case, the function is not tagged as [DEPRECATED](https://pkg.go.dev/github.com/libp2p/go-libp2p#DialRanker),
despite the `// Deprecated:` comment in the source code.

## Hypothesis
I think pkg.go.dev requires a newline before "// Deprecated:".

## Testing
This is a simple Go module that should appear on pkg.go.dev.
It includes a few functions.
- One isn't deprecated at all (used as a control)
- One is deprecated with a newline before the `// Deprecated:` line
- One is deprecated without a newline before the `// Deprecated: ` line

I want to see what functions are actually tagged as deprecated by pkg.go.dev.

If my hypothesis is true, I'll make a PR to fix the deprecated comments in [go-libp2p](https://pkg.go.dev/github.com/libp2p/go-libp2p).
