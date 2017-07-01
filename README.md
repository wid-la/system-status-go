# System Status

### Environnement Variables

* VERSION_TAG `For reference`
* INTERVAL `Default to 300 ms`
* HTTP_PORT `Default to :8080`
* TIMEOUT `Default to 1800 seconds`
* TOKEN `Header Auth-Token`

### Server Docker Install

```
version: '2'
services:
  system-status:
    image: widla/system-status-go:latest 
    expose:
      - "8080"
    restart: always
    environment:
      TOKEN: "xxx"
```

### Client Usage

```
import "github.com/wid-la/system-status-go/cmd/sys-status-cli"

var sys systemStatus.SystemStatus

func main() {
	sys = systemStatus.NewSystemStatusCli(URL, TOKEN)
	sys.CreateService(KEY, NAME, DESC, TIMEOUT)
}

func process() {
	sys.UpdateService()
}

```