package aggregate

import (
	"github.com/PaloAltoNetworks/pango/version"
)

type tc struct {
	desc    string
	version version.Number
	conf    Entry
}

func getTests() []tc {
	return []tc{
		{"v1 ha", version.Number{7, 1, 0, ""}, Entry{
			Name:    "ae1",
			Mode:    ModeHa,
			Comment: "my description",
		}},
		{"v1 decrypt mirror", version.Number{7, 1, 0, ""}, Entry{
			Name: "ae1",
			Mode: ModeDecryptMirror,
		}},
		{"v1 virtual wire basic", version.Number{7, 1, 0, ""}, Entry{
			Name: "ae1",
			Mode: ModeVirtualWire,
		}},
		{"v1 virtual wire full", version.Number{7, 1, 0, ""}, Entry{
			Name:           "ae1",
			Mode:           ModeVirtualWire,
			NetflowProfile: "my netflow profile",
			raw: map[string]string{
				"vwsi": "subinterfaces",
			},
		}},
		{"v1 l2 basic", version.Number{7, 1, 0, ""}, Entry{
			Name: "ae1",
			Mode: ModeLayer2,
		}},
		{"v1 l2 full", version.Number{7, 1, 0, ""}, Entry{
			Name:           "ae1",
			Mode:           ModeLayer2,
			NetflowProfile: "my netflow profile",
			raw: map[string]string{
				"l2si": "subinterfaces",
			},
		}},
		{"v1 l3 static ips", version.Number{7, 1, 0, ""}, Entry{
			Name:              "ae1",
			Mode:              ModeLayer3,
			StaticIps:         []string{"10.2.3.1/24", "10.2.4.1/24"},
			ManagementProfile: "allow pings",
			NetflowProfile:    "some netflow profile",
		}},
		{"v1 l3 dhcp", version.Number{7, 1, 0, ""}, Entry{
			Name:                   "ae1",
			Mode:                   ModeLayer3,
			EnableDhcp:             true,
			CreateDhcpDefaultRoute: true,
			DhcpDefaultRouteMetric: 42,
		}},
		{"v1 l3 other stuff", version.Number{7, 1, 0, ""}, Entry{
			Name:                       "ae1",
			Mode:                       ModeLayer3,
			EnableUntaggedSubinterface: true,
			AdjustTcpMss:               true,
			Ipv4MssAdjust:              11,
			Ipv6MssAdjust:              12,
			Mtu:                        13,
			raw: map[string]string{
				"ndp":    "ndp config",
				"l3si":   "layer3 subinterfaces",
				"arp":    "arp config",
				"v6addr": "ipv6 address info",
				"v6nd":   "ipv6 neighbor discovery config",
			},
		}},
		{"v1 l3 ipv6", version.Number{7, 1, 0, ""}, Entry{
			Name:            "ae1",
			Mode:            ModeLayer3,
			Ipv6Enabled:     true,
			Ipv6InterfaceId: "ipv6 interface id",
		}},
		{"v2 ha", version.Number{8, 1, 0, ""}, Entry{
			Name:    "ae1",
			Mode:    ModeHa,
			Comment: "my description",
		}},
		{"v2 decrypt mirror", version.Number{8, 1, 0, ""}, Entry{
			Name: "ae1",
			Mode: ModeDecryptMirror,
		}},
		{"v2 virtual wire basic", version.Number{8, 1, 0, ""}, Entry{
			Name: "ae1",
			Mode: ModeVirtualWire,
		}},
		{"v2 virtual wire full", version.Number{8, 1, 0, ""}, Entry{
			Name:           "ae1",
			Mode:           ModeVirtualWire,
			NetflowProfile: "my netflow profile",
			raw: map[string]string{
				"vwsi": "subinterfaces",
			},
		}},
		{"v2 l2 basic", version.Number{8, 1, 0, ""}, Entry{
			Name: "ae1",
			Mode: ModeLayer2,
		}},
		{"v2 l2 full", version.Number{8, 1, 0, ""}, Entry{
			Name:           "ae1",
			Mode:           ModeLayer2,
			NetflowProfile: "my netflow profile",
			raw: map[string]string{
				"l2si": "subinterfaces",
			},
		}},
		{"v2 l3 static ips", version.Number{8, 1, 0, ""}, Entry{
			Name:              "ae1",
			Mode:              ModeLayer3,
			StaticIps:         []string{"10.2.3.1/24", "10.2.4.1/24"},
			ManagementProfile: "allow pings",
			NetflowProfile:    "some netflow profile",
		}},
		{"v2 l3 dhcp", version.Number{8, 1, 0, ""}, Entry{
			Name:                   "ae1",
			Mode:                   ModeLayer3,
			EnableDhcp:             true,
			CreateDhcpDefaultRoute: true,
			DhcpDefaultRouteMetric: 42,
		}},
		{"v2 l3 dhcp", version.Number{8, 1, 0, ""}, Entry{
			Name:                   "ae1",
			Mode:                   ModeLayer3,
			EnableDhcp:             true,
			CreateDhcpDefaultRoute: true,
			DhcpDefaultRouteMetric: 42,
		}},
		{"v2 l3 other stuff", version.Number{8, 1, 0, ""}, Entry{
			Name:                       "ae1",
			Mode:                       ModeLayer3,
			EnableUntaggedSubinterface: true,
			DecryptForward:             true,
			AdjustTcpMss:               true,
			Ipv4MssAdjust:              11,
			Ipv6MssAdjust:              12,
			Mtu:                        13,
			raw: map[string]string{
				"ndp":    "ndp config",
				"l3si":   "layer3 subinterfaces",
				"arp":    "arp config",
				"v6addr": "ipv6 address info",
				"v6nd":   "ipv6 neighbor discovery config",
			},
		}},
		{"v2 l3 ipv6", version.Number{8, 1, 0, ""}, Entry{
			Name:            "ae1",
			Mode:            ModeLayer3,
			Ipv6Enabled:     true,
			Ipv6InterfaceId: "ipv6 interface id",
		}},
		{"v3 ha", version.Number{9, 0, 0, ""}, Entry{
			Name:    "ae1",
			Mode:    ModeHa,
			Comment: "my description",
		}},
		{"v3 decrypt mirror", version.Number{9, 0, 0, ""}, Entry{
			Name: "ae1",
			Mode: ModeDecryptMirror,
		}},
		{"v3 virtual wire basic", version.Number{9, 0, 0, ""}, Entry{
			Name: "ae1",
			Mode: ModeVirtualWire,
		}},
		{"v3 virtual wire full", version.Number{9, 0, 0, ""}, Entry{
			Name:           "ae1",
			Mode:           ModeVirtualWire,
			NetflowProfile: "my netflow profile",
			raw: map[string]string{
				"vwsi": "subinterfaces",
			},
		}},
		{"v3 l2 basic", version.Number{9, 0, 0, ""}, Entry{
			Name: "ae1",
			Mode: ModeLayer2,
		}},
		{"v3 l2 full", version.Number{9, 0, 0, ""}, Entry{
			Name:           "ae1",
			Mode:           ModeLayer2,
			NetflowProfile: "my netflow profile",
			raw: map[string]string{
				"l2si": "subinterfaces",
			},
		}},
		{"v3 l3 static ips", version.Number{9, 0, 0, ""}, Entry{
			Name:              "ae1",
			Mode:              ModeLayer3,
			StaticIps:         []string{"10.2.3.1/24", "10.2.4.1/24"},
			ManagementProfile: "allow pings",
			NetflowProfile:    "some netflow profile",
		}},
		{"v3 l3 dhcp", version.Number{9, 0, 0, ""}, Entry{
			Name:                   "ae1",
			Mode:                   ModeLayer3,
			EnableDhcp:             true,
			CreateDhcpDefaultRoute: true,
			DhcpDefaultRouteMetric: 42,
		}},
		{"v3 l3 dhcp", version.Number{9, 0, 0, ""}, Entry{
			Name:                   "ae1",
			Mode:                   ModeLayer3,
			EnableDhcp:             true,
			CreateDhcpDefaultRoute: true,
			DhcpDefaultRouteMetric: 42,
		}},
		{"v3 l3 dhcp with send hostname", version.Number{9, 0, 0, ""}, Entry{
			Name:                   "ae1",
			Mode:                   ModeLayer3,
			EnableDhcp:             true,
			CreateDhcpDefaultRoute: true,
			DhcpDefaultRouteMetric: 42,
			DhcpSendHostnameEnable: true,
			DhcpSendHostnameValue:  "example.com",
		}},
		{"v3 l3 other stuff", version.Number{9, 0, 0, ""}, Entry{
			Name:                       "ae1",
			Mode:                       ModeLayer3,
			EnableUntaggedSubinterface: true,
			DecryptForward:             true,
			AdjustTcpMss:               true,
			Ipv4MssAdjust:              11,
			Ipv6MssAdjust:              12,
			Mtu:                        13,
			raw: map[string]string{
				"ndp":    "ndp config",
				"l3si":   "layer3 subinterfaces",
				"arp":    "arp config",
				"v6addr": "ipv6 address info",
				"v6nd":   "ipv6 neighbor discovery config",
			},
		}},
		{"v3 l3 ipv6", version.Number{9, 0, 0, ""}, Entry{
			Name:            "ae1",
			Mode:            ModeLayer3,
			Ipv6Enabled:     true,
			Ipv6InterfaceId: "ipv6 interface id",
		}},
	}
}
