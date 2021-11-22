package capture

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/google/gopacket"
	"github.com/google/gopacket/tcpassembly"
	"github.com/google/gopacket/tcpassembly/tcpreader"
)

type HttpStreamFactory struct{}

type HttpStream struct {
	net, transport gopacket.Flow
	r              tcpreader.ReaderStream
}

func (h *HttpStreamFactory) New(net, transport gopacket.Flow) tcpassembly.Stream {
	hstream := &HttpStream{
		net:       net,
		transport: transport,
		r:         tcpreader.NewReaderStream(),
	}
	go hstream.run()

	return &hstream.r
}

func (h *HttpStream) run() {
	buf := bufio.NewReader(&h.r)
	for {
		req, err := http.ReadRequest(buf)
		if err == io.EOF {
			return
		} else if err != nil {
			fmt.Println("Error reading stream", h.net, h.transport, ":", err)
		} else if req.Method == "POST" {
			var data interface{}
			err = json.NewDecoder(req.Body).Decode(&data)
			if err != nil {
				return
			}
			payload, err := json.Marshal(data)
			if err != nil {
				return
			}

			headers := make(map[string]string)
			for k, v := range req.Header {
				headers[k] = strings.Join(v[:], ",")
			}
			h, err := json.Marshal(headers)
			if err != nil {
				return
			}

			log.Println("=================")
			log.Println("header:", string(h))
			log.Println(string(payload))
			defer req.Body.Close()

		}
	}
}
