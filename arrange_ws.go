package main

import (
	"fmt"
	"github.com/proxypoke/i3ipc"
)
func main() {
	ipc, _  := i3ipc.GetIPCSocket()
	defer  ipc.Close()
	var nums [10]string
	nums[0] = "zero"
	nums[1] = "one"
	nums[2] = "two"
	nums[3] = "three"
	nums[4] = "four"
	nums[5] = "five"
	nums[6] = "six"
	nums[7] = "seven"
	nums[8] = "eight"
	nums[9] = "nine"
	ws_events, _ := i3ipc.Subscribe(i3ipc.I3WorkspaceEvent)
	for {
		event := <-ws_events
		if event.Change != "" {
			wkspcs, _ := ipc.GetWorkspaces()
			for i := range wkspcs {
				command := fmt.Sprintf(
					"rename workspace %s to %s",
					wkspcs[i].Name, nums[i])
				ipc.Command(command)
			}
		}
	}
}
