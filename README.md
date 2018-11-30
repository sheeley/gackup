# gackup
Mackup-like go tool to provide basic file move/linking.

## Usage

Create `~/.gackup` with contents like:

```
.mongorc.js
.gitident-work
.zshrc
.ssh/config
.gitignore
.gitconfig
.zshenv
# .vscode
Library/Preferences/com.surteesstudios.Bartender.plist
Library/Preferences/com.googlecode.iterm2.plist
Library/Preferences/info.marcel-dierkes.KeepingYouAwake.plist
# Library/KeyBindings/DefaultKeyBinding.dict
# Library/Services
# Library/Speech/Speakable Items
# Library/Scripts
# Library/Workflows
# Library/PDF Services
Library/Preferences/com.apple.symbolichotkeys.plist
Library/Preferences/org.shiftitapp.ShiftIt.plist
Library/Application Support/Code/User/settings.json
```

Run
```
go get github.com/sheeley/gackup/...
gackup
```

## Options

```
-base string
    	Set base directory (default "$HOME")
-configDir string
    Set directory to store synced files in (default "Documents/config")
-relink
    Force re-linking of all files
-verbose
```