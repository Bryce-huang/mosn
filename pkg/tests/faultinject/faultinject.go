/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package main

import (
	"fmt"

	"gitlab.alipay-inc.com/afe/mosn/pkg/api/v2"
	"gitlab.alipay-inc.com/afe/mosn/pkg/log"
	"gitlab.alipay-inc.com/afe/mosn/pkg/server"
	"gitlab.alipay-inc.com/afe/mosn/pkg/server/config/filter/network"
	"gitlab.alipay-inc.com/afe/mosn/pkg/types"
	"gitlab.alipay-inc.com/afe/mosn/pkg/upstream/cluster"

	"net"
	"net/http"
	"runtime"
	"time"
)

const (
	TestCluster    = "tstCluster"
	TestListener   = "tstListener"
	RealServerAddr = "127.0.0.1:8080"
	MeshServerAddr = "127.0.0.1:2048"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	go func() {
		// pprof server
		http.ListenAndServe("0.0.0.0:9090", nil)
	}()

	log.InitDefaultLogger("stdout", log.DEBUG)

	stopChan := make(chan bool)

	go func() {
		nf := &network.FaultInjectFilterConfigFactory{
			FaultInject: &v2.FaultInject{
				DelayPercent:  100,
				DelayDuration: 2000,
			},
			Proxy: tcpProxyConfig(),
		}

		// mesh
		cmf := &clusterManagerFilter{}
		cm := cluster.NewClusterManager(nil, nil, nil, false, false)

		srv := server.NewServer(nil, cmf, cm)
		srv.AddListener(tcpListener(), nf, nil)
		cmf.cccb.UpdateClusterConfig(clusters())
		cmf.chcb.UpdateClusterHost(TestCluster, 0, hosts("11.162.169.38:80"))

		srv.Start()

		select {
		case <-stopChan:
			srv.Close()
		}
	}()

	select {
	case <-time.After(time.Second * 1800):
		stopChan <- true
		fmt.Println("[MAIN]closing..")
	}
}

func tcpListener() *v2.ListenerConfig {
	addr, _ := net.ResolveTCPAddr("tcp", MeshServerAddr)

	return &v2.ListenerConfig{
		Name:                    TestListener,
		Addr:                    addr,
		BindToPort:              true,
		PerConnBufferLimitBytes: 1024 * 32,
		LogPath:                 "stdout",
		LogLevel:                uint8(log.DEBUG),
	}
}

type clusterManagerFilter struct {
	cccb types.ClusterConfigFactoryCb
	chcb types.ClusterHostFactoryCb
}

func (cmf *clusterManagerFilter) OnCreated(cccb types.ClusterConfigFactoryCb, chcb types.ClusterHostFactoryCb) {
	cmf.cccb = cccb
	cmf.chcb = chcb
}

func tcpProxyConfig() *v2.TcpProxy {
	tcpProxyConfig := &v2.TcpProxy{}
	tcpProxyConfig.Routes = append(tcpProxyConfig.Routes, &v2.TcpRoute{
		Cluster: TestCluster,
	})

	return tcpProxyConfig
}

func clusters() []v2.Cluster {
	var configs []v2.Cluster
	configs = append(configs, v2.Cluster{
		Name:                 TestCluster,
		ClusterType:          v2.SIMPLE_CLUSTER,
		LbType:               v2.LB_RANDOM,
		MaxRequestPerConn:    1024,
		ConnBufferLimitBytes: 32 * 1024,
	})

	return configs
}

func hosts(host string) []v2.Host {
	var hosts []v2.Host

	if host == "" {
		host = RealServerAddr
	}

	hosts = append(hosts, v2.Host{
		Address: host,
		Weight:  100,
	})

	return hosts
}