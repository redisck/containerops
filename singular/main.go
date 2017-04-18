package main

import (
	"fmt"

	"github.com/Huawei/containerops_he3io/singular/nodes"
	//"containerops/singular/cmd"
	//"github.com/Huawei/containerops_he3io/singular/cmd"
)

//var nodes = [2][2]string{{"192.168.60.141", "centos-master"}, {"192.168.60.150", "centos-minion"}}
var m = make(map[string]string)

func main() {

	m["centos-master"] = "192.168.60.141"
	m["centos-minion"] = "192.168.60.150"
	for k, v := range m {
		fmt.Printf("k=%v, v=%v\n", k, v)
		if k == "centos-master" {
			nodes.Deploymaster(m, v)
		}
		if k == "centos-minion" {
			nodes.Deploynode(m, v)
		}
	}

}
