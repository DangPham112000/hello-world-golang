# hello-world-golang

To run multiple Go files at the same time on a Macintosh like this:

```console
go run *.go
```

On Windows, though, this will not work unless you have customized your IDE or terminal. Instead, use this command:

```console
go run .
```

(note the dot at the end of this command). Note also that you can use the same command, go run ., on Mac and Linux as well.

#New commandline to run
On windows

```console
go run ./cmd/web/.
```

On mac

```console
go run cmd/web/*.go
```
