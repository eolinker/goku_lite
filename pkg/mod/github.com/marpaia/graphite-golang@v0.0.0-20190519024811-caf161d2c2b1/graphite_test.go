package graphite

import (
	"fmt"
	"net"
	"strings"
	"testing"
)

const TCP = "tcp"
const UDP = "udp"
const NOP = "nop"

// Change these to be your own graphite server if you so please
var graphiteHost = "carbon.hostedgraphite.com"
var graphitePort = 2003

func TestNewGraphite(t *testing.T) {
	gh, err := NewGraphite(graphiteHost, graphitePort)
	if err != nil {
		t.Error(err)
	}

	if _, ok := gh.conn.(*net.TCPConn); !ok {
		t.Error("GraphiteHost.conn is not a TCP connection")
	}
}

func TestNewGraphiteWithMetricPrefix(t *testing.T) {
	prefix := "test"
	gh, err := NewGraphiteWithMetricPrefix(graphiteHost, graphitePort, prefix)
	if err != nil {
		t.Error(err)
	}

	if _, ok := gh.conn.(*net.TCPConn); !ok {
		t.Error("GraphiteHost.conn is not a TCP connection")
	}
}

func TestNewGraphiteUDP(t *testing.T) {
	gh, err := NewGraphiteUDP(graphiteHost, graphitePort)
	if err != nil {
		t.Error(err)
	}

	if _, ok := gh.conn.(*net.UDPConn); !ok {
		t.Error("GraphiteHost.conn is not a UDP connection")
	}
}

func TestGraphiteFactoryTCP(t *testing.T) {
	gr, err := GraphiteFactory(TCP, graphiteHost, graphitePort, "")

	if err != nil {
		t.Error(err)
	}

	if _, ok := gr.conn.(*net.TCPConn); !ok {
		t.Error("GraphiteHost.conn is not a TCP connection")
	}
}

func TestGraphiteFactoryTCPWithPrefix(t *testing.T) {
	prefix := "test"
	gr, err := GraphiteFactory(TCP, graphiteHost, graphitePort, prefix)

	if err != nil {
		t.Error(err)
	}

	if _, ok := gr.conn.(*net.TCPConn); !ok {
		t.Error("GraphiteHost.conn is not a TCP connection")
	}

	if !strings.EqualFold(prefix, gr.Prefix) {
		t.Error(fmt.Sprintf("Wrong prefix is set expected %s actual %s",
			prefix,
			gr.Prefix))
	}
}

func TestGraphiteFactoryUDP(t *testing.T) {
	gr, err := GraphiteFactory(UDP, graphiteHost, graphitePort, "")

	if err != nil {
		t.Error(err)
	}

	if _, ok := gr.conn.(*net.UDPConn); !ok {
		t.Error("GraphiteHost.conn is not a UDP connection")
	}
}

func TestGraphiteFactoryNop(t *testing.T) {
	gr, err := GraphiteFactory(NOP, graphiteHost, graphitePort, "")

	if err != nil {
		t.Error(err)
	}

	if !gr.IsNop() {
		t.Error("GraphiteHost is not NOP")
	}
}

// Uncomment the following method to test sending an actual metric to graphite
//
//func TestSendMetric(t *testing.T) {
//	gh, err := NewGraphite(graphiteHost, graphitePort)
//	if err != nil {
//		t.Error(err)
//	}
//	err = gh.SimpleSend("stats.test.metric11", "1")
//	if err != nil {
//		t.Error(err)
//	}
//}
