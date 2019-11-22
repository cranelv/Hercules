package runningSpace
type RunningSpaceHandler interface {
	GetRunningSpace(string) RunningSpaceInterface
}
