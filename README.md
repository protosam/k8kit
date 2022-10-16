# k8kit
A single binary containing `kubectl`, `helm`, and `flux` embedded as subcommands.

This is a proof of concept right now, to show what's possible if the `*cobra.Cli` commands were exposed upstream.

## Try it Out
The releases page have a copy of the `v0.0.0-demo`. Download the binary for your OS and architecture and run it to get started.

```
$ k8kit --help
$ k8kit kubectl
$ k8kit helm
$ k8kit flux
```
