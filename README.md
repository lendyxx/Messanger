## Messenger
Mobile messenger on golang with java.

Backend: https://github.com/gopher5/messenger

## Build

### Set go env
```go env -w CGO_CFLAGS="-O2 -g -fdeclspec"```

### Use custom fyne with java support:
```fyne package -os android -appID com.messenger.app -javaSource javaApp\java```

