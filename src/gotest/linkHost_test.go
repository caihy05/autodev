package gotest
import "testing"
import (
	"cLinks"
//	"fmt"
)
//测试连接
func Test_Connect(t *testing.T)  {
	_, err := cLinks.Connect("root","root1234", "172.50.10.10", 22)
	if err != nil {
		t.Error("测试结果失败")

	}
	t.Log("测试通过")
}
