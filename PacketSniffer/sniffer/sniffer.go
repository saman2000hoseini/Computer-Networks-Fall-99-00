package sniffer

import (
	"fmt"
	"github.com/google/gopacket"
	"github.com/sirupsen/logrus"
	"math"
	"os"
	"sort"
	"strconv"
	"time"

	"github.com/google/gopacket/pcap"
)

const (
	linkName        = "Link Layer"
	networkName     = "Network Layer"
	transportName   = "Transport Layer"
	applicationName = "Application Layer"
	endpointsName   = "Endpoints"
	separator       = "------------------------------------------------------------------------------------"
)

type Sniffer struct {
	Finished *bool
}

func (s *Sniffer) Capture(rPath, device string) {
	fPath := generateName(rPath)

	os.Mkdir(fPath, 0755)

	report := fPath + "/packets.txt"

	file, err := os.Create(report)
	if err != nil {
		logrus.Errorf("error creating file: %s", err.Error())
		return
	}
	defer file.Close()

	linkLayer := make(map[string]uint64)
	networkLayer := make(map[string]uint64)
	transportLayer := make(map[string]uint64)
	applicationLayer := make(map[string]uint64)
	endpoints := make(map[string]uint64)

	minLength := math.MaxInt64
	maxLength := math.MinInt64
	count := uint64(0)
	avg := uint64(0)

	if handle, err := pcap.OpenLive(device, 1600, true, pcap.BlockForever); err != nil {
		logrus.Errorf("error openning source: %s", err.Error())
		return
	} else {
		packetSource := gopacket.NewPacketSource(handle, handle.LinkType())

		for packet := range packetSource.Packets() {

			if packet.LinkLayer() != nil {
				linkLayer[packet.LinkLayer().LayerType().String()]++
			}

			if packet.NetworkLayer() != nil {
				networkLayer[packet.NetworkLayer().LayerType().String()]++

				src, dest := packet.NetworkLayer().NetworkFlow().Endpoints()
				endpoints[endpointsString(src, dest)]++
			}

			if packet.TransportLayer() != nil {
				transportLayer[packet.TransportLayer().LayerType().String()]++
			}

			if packet.ApplicationLayer() != nil {
				applicationLayer[packet.ApplicationLayer().LayerType().String()]++
			}

			length := packet.Metadata().Length
			if minLength > length {
				minLength = length
			}

			if maxLength < length {
				maxLength = length
			}

			avg *= count
			count++
			avg += uint64(length)
			avg /= count

			file.WriteString(packet.String())

			if *s.Finished {
				break
			}
		}
	}

	sFile, err := os.Create(fPath + "/stats.txt")
	if err != nil {
		logrus.Errorf("error creating stats file: %s", err.Error())
		return
	}
	defer sFile.Close()

	sFile.WriteString("Average length: " + strconv.FormatUint(avg, 10) + "\n")
	sFile.WriteString("Min length: " + strconv.FormatInt(int64(minLength), 10) + "\n")
	sFile.WriteString("Max length: " + strconv.FormatInt(int64(maxLength), 10) + "\n")

	writeStatistics(sFile, linkName, linkLayer)
	writeStatistics(sFile, networkName, networkLayer)
	writeStatistics(sFile, transportName, transportLayer)
	writeStatistics(sFile, applicationName, applicationLayer)

	eFile, err := os.Create(fPath + "/endpoints.txt")
	if err != nil {
		logrus.Errorf("error creating endpoints file: %s", err.Error())
		return
	}
	defer eFile.Close()

	endpointsList := sortEndpoints(endpoints)
	for _, endpoint := range endpointsList {
		eFile.WriteString(endpoint.Key + ": " + strconv.FormatUint(endpoint.Value, 10) + "\n")
	}

	fmt.Println("Reports saved to: " + fPath)
}

func (s *Sniffer) GetDevices() []pcap.Interface {
	devs, err := pcap.FindAllDevs()
	if err != nil {
		return nil
	}

	return devs
}

func writeStatistics(file *os.File, title string, stats map[string]uint64) {
	file.WriteString(separator + "\n")
	file.WriteString(title + ":\n\n")

	for key, value := range stats {
		file.WriteString(key + ": " + strconv.FormatUint(value, 10) + "\n")
	}
}

func generateName(path string) string {
	t := time.Now()

	return fmt.Sprintf("%s/%d%02d%02d_%02d%02d%02d",
		path, t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
}

type pair struct {
	Key   string
	Value uint64
}

type pairList []pair

func (p pairList) Len() int           { return len(p) }
func (p pairList) Less(i, j int) bool { return p[i].Value < p[j].Value }
func (p pairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func sortEndpoints(endpoints map[string]uint64) pairList {
	el := make(pairList, len(endpoints))
	i := 0
	for k, v := range endpoints {
		el[i] = pair{k, v}
		i++
	}

	sort.Sort(sort.Reverse(el))
	return el
}

func endpointsString(src, dest gopacket.Endpoint) string {
	return fmt.Sprintf("%s ==> %s", src.String(), dest.String())
}
