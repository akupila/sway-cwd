package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func main() {
	out, err := exec.Command("swaymsg", "-t", "get_tree").Output()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not get sway tree: %v\n", err)
	}

	root := &node{}
	if err := json.Unmarshal(out, root); err != nil {
		fmt.Fprintf(os.Stderr, "Could not unmarshal sway tree: %v\n", err)
	}
	pid := focusedPID(root)
	if pid == 0 {
		fmt.Fprintf(os.Stderr, "Could not find focused pid\n")
	}

	child, err := exec.Command("pgrep", "--newest", "--parent", strconv.Itoa(pid)).Output()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not find child pid: %v\n", err)
	}
	childPID := strings.TrimSpace(string(child))

	cwd, err := os.Readlink(fmt.Sprintf("/proc/%s/cwd", childPID))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not resolve link: %v\n", err)
	}
	fmt.Fprint(os.Stdout, cwd)
}

func focusedPID(node *node) int {
	if node.Focused && node.Type == "con" {
		return node.PID
	}
	for _, n := range node.Nodes {
		if pid := focusedPID(n); pid != 0 {
			return pid
		}
	}
	return 0
}

type node struct {
	Focused bool    `json:"focused"`
	Type    string  `json:"type"`
	PID     int     `json:"pid"`
	Nodes   []*node `json:"nodes"`
}
