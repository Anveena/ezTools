//go:build ezDebug

package log

import (
	"fmt"
	"runtime"
	"strings"
	"sync"
)

var debugLogFmtStr = "%s %s:%d%s%s\n"

func startGRPCClient() error {
	wg := sync.WaitGroup{}
	wg.Add(1)
	wg.Wait()
	return nil
}
func Log(level LogLv, msg ...any) {
	_, file, line, _ := runtime.Caller(2)
	header := lvHeaderMap[level]
	fmt.Printf(debugLogFmtStr, header, file, line, header,
		strings.ReplaceAll(fmt.Sprintln(msg...), "\n", header))
}
func LogWithTag(level LogLv, tag string, msg ...any) {
	_, file, line, _ := runtime.Caller(2)
	header := fmt.Sprintf("%s[%s]", lvHeaderMap[level], tag)
	fmt.Printf(debugLogFmtStr, header, file, line, header,
		strings.ReplaceAll(fmt.Sprintln(msg...), "\n", header))
}
func sendToDing(_ LogLv, _ string, _ string) {}
