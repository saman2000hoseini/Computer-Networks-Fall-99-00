package sniffer

import (
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/saman2000hoseini/Computer-Networks-Fall-99-00/PacketSniffer/model"
	"github.com/saman2000hoseini/Computer-Networks-Fall-99-00/PacketSniffer/utils"
	"github.com/sirupsen/logrus"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/google/gopacket/pcap"
)

const (
	linkName        = "Link Layer"
	networkName     = "Network Layer"
	transportName   = "Transport Layer"
	applicationName = "Application Layer"
	endpointsName   = "Endpoints"
	fragmentsName   = "Fragments"
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
	fragments := make(map[string]uint64)

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
				if ipv4 := packet.Layer(layers.LayerTypeIPv4); ipv4 != nil {
					ip := ipv4.(*layers.IPv4)
					if ip.Flags.String() == "DF" {
						fragments["Dont Fragment"]++
					} else {
						fragments["Fragment"]++
					}
				}

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
	writeStatistics(sFile, fragmentsName, fragments)

	if len(linkLayer) > 0 {
		utils.DrawPieChart(linkLayer, generateChartPath(fPath, linkName))
	}

	utils.DrawPieChart(networkLayer, generateChartPath(fPath, networkName))
	utils.DrawPieChart(transportLayer, generateChartPath(fPath, transportName))
	utils.DrawPieChart(applicationLayer, generateChartPath(fPath, applicationName))
	utils.DrawPieChart(endpoints, generateChartPath(fPath, endpointsName))
	utils.DrawPieChart(fragments, generateChartPath(fPath, fragmentsName))

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

func generateChartPath(path, name string) string {
	return fmt.Sprintf("%s/%s", path, strings.Replace(name, " ", "", -1))
}

func sortEndpoints(endpoints map[string]uint64) model.PairList {
	el := make(model.PairList, len(endpoints))
	i := 0
	for k, v := range endpoints {
		el[i] = model.Pair{Key: k, Value: v}
		i++
	}

	sort.Sort(sort.Reverse(el))
	return el
}

func endpointsString(src, dest gopacket.Endpoint) string {
	return fmt.Sprintf("%s ==> %s", src.String(), dest.String())
}
