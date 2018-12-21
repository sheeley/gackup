# gackup
Mackup-like go tool to provide basic file move/linking.

## Configuration file

Loads from `~/.gackup` or `$HOME/{target dir}/.gackup`

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

## Usage

Run
```
> go get github.com/sheeley/gackup/...
> gackup -h
Usage of gackup:
  -relink
    	Force re-linking of all files
  -source string
    	Set source directory (default "$HOME")
  -target string
    	Set directory to store synced files in (default "Documents/config")
  -verbose
> gackup
MOVE: /Users/sheeley/.gackup -> /Users/sheeley/Documents/config/.gackup
LINK: /Users/sheeley/.gackup -> /Users/sheeley/Documents/config/.gackup
LINK: /Users/sheeley/Library/Preferences/com.surteesstudios.Bartender.plist -> /Users/sheeley/Documents/config/Library/Preferences/com.surteesstudios.Bartender.plist
LINK: /Users/sheeley/Library/Preferences/info.marcel-dierkes.KeepingYouAwake.plist -> /Users/sheeley/Documents/config/Library/Preferences/info.marcel-dierkes.KeepingYouAwake.plist

Confirm [y/N]:y
```