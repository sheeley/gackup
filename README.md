# gackup
Mackup-like go tool to provide basic file move/linking.

## usage

- right now, you'd need to modify the file list in `cmd/main.go`

```
-base string
    	Set base directory (default "$HOME")
-configDir string
    Set directory to store synced files in (default "Documents/config")
-relink
    Force re-linking of all files
-verbose
```