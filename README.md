# alt:V Debug Mode Switcher
I made this script until my feature request [#1500](https://github.com/altmp/altv-issues/issues/1500) is accepted.

## Build
```
git clone https://github.com/vicelo23/altv-debug-mode-switcher.git
cd altv-debug-mode-switcher
go build
```

## Arguments

alt:V Debug Mode Switcher has two CLI arguments.

`-debug` argument is optional it converts `debug: 'false'` to `debug: 'true'` in your altv.cfg. If it's not applied it converts `debug: 'true'` to `debug: 'false'`.

`-noautorun` argument is optional and it doesn't start altv.exe after debug mode switch process completed. It just makes the program to exit, you usually don't need this.

## How to use
Install the exe file from releases page or build it yourself and throw it to your alt:V folder where altv.exe and altv.cfg files are located.



1. You most likely use shortcut to run alt:V, after throwing the exe file to the alt:V location, right click on it and send a shortcut to desktop.
2. On desktop, duplicate this shortcut.
3. Rename the shortcuts to different things, in this tutorial they will be named as `altV (Debug On)` and `altV (Debug Off)`.
4. Right click on `altV (Debug On)` shortcut and select properties.
5. On the properties window you will see a textbox next to the Target. Add -debug to the end. (example: `C:\Games\altv\altv-debug-mode-switcher.exe -debug`)
6. You don't need to do something for other shortcut, it doesn't need any argument.

When you run altV (Debug On) shortcut, the game will run with debug mode enabled and when you run altV(Debug Off) shortcut, it will run with debug mode disabled. If you don't want these shortcuts to be on your desktop and you swill want to access them from Windows Search you can throw these shortcuts to `%appdata%/Microsoft/Windows/Start Menu/Programs` or if you want these shortcuts available globally, you can throw them to `C:/ProgramData/Microsoft/Windows/Start Menu/Programs`. You can also set icons to these shortcuts if you want them look better.

## Errors
In some cases you might get errors and to see them, you can run altv-debug-mode-switcher.exe from your terminal with the arguments given. Example
```
PS location\to\altv>.\altv-debug-mode-switcher.exe -debug
panic: multiple debug lines found in altv.cfg

PS location\to\altv>.\altv-debug-mode-switcher.exe -debug
panic: no debug line found
```
