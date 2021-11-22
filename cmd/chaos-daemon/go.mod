module entropie.ai/chaos-daemon

go 1.17

replace entropie.ai/pkg/capture => ./pkg/capture

require entropie.ai/pkg/capture v0.0.0-00010101000000-000000000000

require (
	github.com/google/gopacket v1.1.19 // indirect
	golang.org/x/sys v0.0.0-20190412213103-97732733099d // indirect
)
