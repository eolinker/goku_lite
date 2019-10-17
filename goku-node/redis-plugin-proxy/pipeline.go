package redis_plugin_proxy

import (
	goku_plugin "github.com/eolinker/goku-plugin"
	"github.com/go-redis/redis"
)

//PipelineProxy 转发管道
type PipelineProxy struct {
	RedisProxy
	pipeliner redis.Pipeliner
}

//Auth auth
func (p *PipelineProxy) Auth(password string) goku_plugin.StatusCmd {
	return p.pipeliner.Auth(password)
}

//Select select
func (p *PipelineProxy) Select(index int) goku_plugin.StatusCmd {
	return p.pipeliner.Select(index)
}

//SwapDB swapDB
func (p *PipelineProxy) SwapDB(index1, index2 int) goku_plugin.StatusCmd {
	return p.pipeliner.SwapDB(index1, index2)
}

//ClientSetName clientSetName
func (p *PipelineProxy) ClientSetName(name string) goku_plugin.BoolCmd {
	return p.pipeliner.ClientSetName(name)
}

//Do do
func (p *PipelineProxy) Do(args ...interface{}) goku_plugin.Cmd {
	return p.pipeliner.Do(args...)
}

//Process process
func (p *PipelineProxy) Process(cmd goku_plugin.Cmder) error {
	arg := cmd.Args()
	return p.pipeliner.Process(redis.NewCmd(arg...))
}

//Close close
func (p *PipelineProxy) Close() error {
	return p.pipeliner.Close()
}

//Discard discard
func (p *PipelineProxy) Discard() error {
	return p.pipeliner.Discard()
}

//Exec exec
func (p *PipelineProxy) Exec() ([]goku_plugin.Cmder, error) {

	cmders, err := p.pipeliner.Exec()
	if err != nil {
		return nil, err
	}

	cmds := make([]goku_plugin.Cmder, 0, len(cmders))
	for _, c := range cmders {
		cmds = append(cmds, c)
	}
	return cmds, nil
}
