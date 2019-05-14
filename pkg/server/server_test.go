package server

import (
	"testing"

	"github.com/fyuan1316/proxyserver/cmd"
)

func Test_run(t *testing.T) {
	addr := ":9999"
	model := "testmodel"
	target := "http://www.baidu.com"
	prefix := "/alauda/"
	s := NewMyServer(addr,
		cmd.ModelName(model),
		cmd.TargetURL(target),
		cmd.ProxyPrefixURL(prefix))

	s.start()
}
