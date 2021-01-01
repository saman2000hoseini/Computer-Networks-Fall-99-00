package config

// DefaultPacketSniffer return default configuration
func DefaultPacketSniffer() PacketSniffer {
	return PacketSniffer{
		Path: "reports",
	}
}
