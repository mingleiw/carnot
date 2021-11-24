package capture

import (
	"fmt"
	"log"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/examples/util"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"github.com/google/gopacket/tcpassembly"
)

type Capture struct {
	iface string
	port  string
}

func (c *Capture) WithPort(port string) *Capture {
	c.port = port

	return c
}

func (c *Capture) WithIface(iface string) *Capture {
	c.iface = iface

	return c
}

func (c *Capture) Start() {
	defer util.Run()()
	var handle *pcap.Handle
	var err error

	var snaplen int32 = 65535
	bpfFilter := fmt.Sprintf("%s%s", "tcp and dst port ", c.port)

	handle, err = pcap.OpenLive(c.iface, snaplen, true, pcap.BlockForever)
	defer handle.Close()
	if err != nil {
		log.Fatal(err)
		return
	}

	if err := handle.SetBPFFilter(bpfFilter); err != nil {
		log.Fatal(err)
		return
	}

	streamFactory := &HttpStreamFactory{}
	streamPool := tcpassembly.NewStreamPool(streamFactory)
	assembler := tcpassembly.NewAssembler(streamPool)

	log.Println("reading in packets ...")

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	packets := packetSource.Packets()

	ticker := time.Tick(time.Minute)
	for {
		select {
		case packet := <-packets:
			if packet == nil {
				return
			}
			if packet.NetworkLayer() == nil || packet.TransportLayer() == nil || packet.TransportLayer().LayerType() != layers.LayerTypeTCP {
				log.Println("Unusable packet")
				continue
			}
			tcp := packet.TransportLayer().(*layers.TCP)
			assembler.AssembleWithTimestamp(packet.NetworkLayer().NetworkFlow(), tcp, packet.Metadata().Timestamp)

		case <-ticker:
			assembler.FlushOlderThan(time.Now().Add(time.Minute * -2))
		}
	}
}
