package conf

// AgentConf is agent conf
type AgentConf struct {
	PDHost       string
	TimeInterval int
	AdapterListen string
}

// RunTimeInfo is run time info
var RunTimeInfo AgentConf

// RunTimeMap is run time info map
var RunTimeMap map[string]AgentConf
