package banner

import (
	"fmt"
	"runtime"
	"strings"
	"time"
)

const banner = `
                                                       
 __  __     __     __   __     ______     ______   __     ______     ______    
/\ \/ /    /\ \   /\ "-.\ \   /\  ___\   /\__  _\ /\ \   /\  ___\   /\  __ \   
\ \  _"-.  \ \ \  \ \ \-.  \  \ \  __\   \/_/\ \/ \ \ \  \ \ \____  \ \  __ \  
 \ \_\ \_\  \ \_\  \ \_\\"\_\  \ \_____\    \ \_\  \ \_\  \ \_____\  \ \_\ \_\ 
  \/_/\/_/   \/_/   \/_/ \/_/   \/_____/     \/_/   \/_/   \/_____/   \/_/\/_/

                   Embedded Sensor Data Server (Go)

`

func LaunchAPP(appVersion string) {
	var build strings.Builder
	build.WriteString(banner)
	build.WriteString(strings.Repeat("-", 80))
	build.WriteString(fmt.Sprintf("\nMode:         %s", "Console"))
	build.WriteString(fmt.Sprintf("\nVersion:      %s", appVersion))
	build.WriteString(fmt.Sprintf("\nStarted at:   %s", time.Now().Format("02-01-2006 15:04:05")))
	build.WriteString(fmt.Sprintf("\nHost runtime: %s", runtime.Version()))
	build.WriteString(fmt.Sprintf("\nInterfaces:   Server ←→ BLE ←→ Controller-Hub\n"))
	build.WriteString(strings.Repeat("-", 80) + "\n")

	fmt.Println(build.String())
}
