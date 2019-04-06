package ikegw

import (
    "testing"
    "reflect"

    "github.com/PaloAltoNetworks/pango/testdata"
    "github.com/PaloAltoNetworks/pango/version"
)


func TestFwNormalization(t *testing.T) {
    testCases := []struct{
        desc string
        version version.Number
        conf Entry
    }{
        {"test1", version.Number{6, 1, 0, ""}, Entry{
            Name: "test1",
            Version: Ikev1,
            PeerIpType: PeerTypeIp,
            PeerIpValue: "10.1.1.1",
            Interface: "tunnel.1",
            LocalIpAddressType: LocalTypeIp,
            LocalIpAddressValue: "10.2.1.1",
            AuthType: AuthPreSharedKey,
            PreSharedKey: "secret",
            LocalIdType: "ipaddr",
            LocalIdValue: "10.3.1.1",
            PeerIdType: "fqdn",
            PeerIdValue: "example.com",
            PeerIdCheck: "exact",
        }},
        {"test2", version.Number{6, 1, 0, ""}, Entry{
            Name: "test2",
            Version: Ikev1,
            PeerIpType: PeerTypeDynamic,
            Interface: "tunnel.2",
            LocalIpAddressType: LocalTypeIp,
            LocalIpAddressValue: "10.2.1.1",
            AuthType: AuthCertificate,
            LocalCert: "myCert",
            CertProfile: "my cert profile",
            CertPermitPayloadMismatch: true,
            CertEnableStrictValidation: true,
            LocalIdType: "ipaddr",
            LocalIdValue: "10.3.1.1",
            PeerIdType: "fqdn",
            PeerIdValue: "example.com",
            PeerIdCheck: "wildcard",
            Ikev1ExchangeMode: "auto",
            Ikev1CryptoProfile: "my crypto profile",
        }},
        {"test3", version.Number{6, 1, 0, ""}, Entry{
            Name: "test3",
            Version: Ikev1,
            PeerIpType: PeerTypeDynamic,
            Interface: "tunnel.3",
            LocalIpAddressType: LocalTypeIp,
            LocalIpAddressValue: "10.2.1.1",
            AuthType: AuthPreSharedKey,
            PreSharedKey: "another secret",
            LocalIdType: "ipaddr",
            LocalIdValue: "10.3.1.1",
            PeerIdType: "fqdn",
            PeerIdValue: "example.com",
            PeerIdCheck: "wildcard",
            EnablePassiveMode: true,
            EnableNatTraversal: true,
            NatTraversalKeepAlive: 7,
            NatTraversalEnableUdpChecksum: true,
        }},
        {"test4", version.Number{6, 1, 0, ""}, Entry{
            Name: "test4",
            Version: Ikev1,
            PeerIpType: PeerTypeIp,
            PeerIpValue: "10.3.4.5",
            Interface: "tunnel.4",
            LocalIpAddressType: LocalTypeIp,
            LocalIpAddressValue: "10.2.1.1",
            AuthType: AuthPreSharedKey,
            PreSharedKey: "more secrets",
            LocalIdType: "ipaddr",
            LocalIdValue: "10.3.1.1",
            PeerIdType: "fqdn",
            PeerIdValue: "example.com",
            PeerIdCheck: "wildcard",
            EnableDeadPeerDetection: true,
            DeadPeerDetectionInterval: 7,
            DeadPeerDetectionRetry: 4,
            EnableFragmentation: true,
        }},
        {"test5", version.Number{7, 0, 0, ""}, Entry{
            Name: "test5",
            Version: Ikev1,
            PeerIpType: PeerTypeIp,
            PeerIpValue: "10.1.1.1",
            Interface: "tunnel.1",
            LocalIpAddressType: LocalTypeIp,
            LocalIpAddressValue: "10.2.1.1",
            AuthType: AuthPreSharedKey,
            PreSharedKey: "secret",
            LocalIdType: "ipaddr",
            LocalIdValue: "10.3.1.1",
            PeerIdType: "fqdn",
            PeerIdValue: "example.com",
            PeerIdCheck: "exact",
        }},
        {"test6", version.Number{7, 0, 0, ""}, Entry{
            Name: "test6",
            Version: Ikev1,
            PeerIpType: PeerTypeDynamic,
            Interface: "tunnel.2",
            LocalIpAddressType: LocalTypeIp,
            LocalIpAddressValue: "10.2.1.1",
            AuthType: AuthCertificate,
            LocalCert: "myCert",
            CertProfile: "my cert profile",
            CertPermitPayloadMismatch: true,
            CertEnableStrictValidation: true,
            LocalIdType: "ipaddr",
            LocalIdValue: "10.3.1.1",
            PeerIdType: "fqdn",
            PeerIdValue: "example.com",
            PeerIdCheck: "wildcard",
            Ikev1ExchangeMode: "auto",
            Ikev1CryptoProfile: "my crypto profile",
        }},
        {"test7", version.Number{7, 0, 0, ""}, Entry{
            Name: "test7",
            Version: Ikev1,
            PeerIpType: PeerTypeDynamic,
            Interface: "tunnel.3",
            LocalIpAddressType: LocalTypeIp,
            LocalIpAddressValue: "10.2.1.1",
            AuthType: AuthPreSharedKey,
            PreSharedKey: "another secret",
            LocalIdType: "ipaddr",
            LocalIdValue: "10.3.1.1",
            PeerIdType: "fqdn",
            PeerIdValue: "example.com",
            PeerIdCheck: "wildcard",
            EnablePassiveMode: true,
            EnableNatTraversal: true,
            NatTraversalKeepAlive: 7,
            NatTraversalEnableUdpChecksum: true,
        }},
        {"test8", version.Number{7, 0, 0, ""}, Entry{
            Name: "test8",
            Version: Ikev1,
            PeerIpType: PeerTypeIp,
            PeerIpValue: "10.3.4.5",
            Interface: "tunnel.4",
            LocalIpAddressType: LocalTypeIp,
            LocalIpAddressValue: "10.2.1.1",
            AuthType: AuthPreSharedKey,
            PreSharedKey: "more secrets",
            LocalIdType: "ipaddr",
            LocalIdValue: "10.3.1.1",
            PeerIdType: "fqdn",
            PeerIdValue: "example.com",
            PeerIdCheck: "wildcard",
            EnableDeadPeerDetection: true,
            DeadPeerDetectionInterval: 7,
            DeadPeerDetectionRetry: 4,
            EnableFragmentation: true,
        }},
        {"test9", version.Number{7, 0, 0, ""}, Entry{
            Name: "test9",
            Version: Ikev2,
            PeerIpType: PeerTypeIp,
            PeerIpValue: "10.3.4.5",
            Interface: "tunnel.4",
            LocalIpAddressType: LocalTypeIp,
            LocalIpAddressValue: "10.2.1.1",
            AuthType: AuthPreSharedKey,
            PreSharedKey: "more secrets",
            LocalIdType: "ipaddr",
            LocalIdValue: "10.3.1.1",
            PeerIdType: "fqdn",
            PeerIdValue: "example.com",
            PeerIdCheck: "wildcard",
            Ikev2CryptoProfile: "v2 crypto profile",
            Ikev2CookieValidation: true,
        }},
        {"test10", version.Number{7, 0, 0, ""}, Entry{
            Name: "test10",
            Version: Ikev2,
            PeerIpType: PeerTypeIp,
            PeerIpValue: "10.3.4.5",
            Interface: "tunnel.4",
            LocalIpAddressType: LocalTypeIp,
            LocalIpAddressValue: "10.2.1.1",
            AuthType: AuthPreSharedKey,
            PreSharedKey: "more secrets",
            LocalIdType: "ipaddr",
            LocalIdValue: "10.3.1.1",
            PeerIdType: "fqdn",
            PeerIdValue: "example.com",
            PeerIdCheck: "wildcard",
            EnableLivenessCheck: true,
            LivenessCheckInterval: 7,
        }},
        {"test11", version.Number{7, 0, 0, ""}, Entry{
            Name: "test11",
            Version: Ikev2,
            PeerIpType: PeerTypeIp,
            PeerIpValue: "10.3.4.5",
            Interface: "tunnel.4",
            LocalIpAddressType: LocalTypeIp,
            LocalIpAddressValue: "10.2.1.1",
            AuthType: AuthPreSharedKey,
            PreSharedKey: "more secrets",
            LocalIdType: "ipaddr",
            LocalIdValue: "10.3.1.1",
            PeerIdType: "fqdn",
            PeerIdValue: "example.com",
            PeerIdCheck: "wildcard",
            Ikev2CryptoProfile: "v2 crypto profile",
            Ikev2CookieValidation: true,
            EnableLivenessCheck: true,
            LivenessCheckInterval: 7,
            EnableFragmentation: true,
        }},
        {"test12", version.Number{7, 0, 0, ""}, Entry{
            Name: "test12",
            Version: Ikev2Preferred,
            PeerIpType: PeerTypeIp,
            PeerIpValue: "10.3.4.5",
            Interface: "tunnel.4",
            LocalIpAddressType: LocalTypeIp,
            LocalIpAddressValue: "10.2.1.1",
            AuthType: AuthPreSharedKey,
            PreSharedKey: "more secrets",
            LocalIdType: "ipaddr",
            LocalIdValue: "10.3.1.1",
            PeerIdType: "fqdn",
            PeerIdValue: "example.com",
            PeerIdCheck: "wildcard",
            Ikev1ExchangeMode: "auto",
            Ikev1CryptoProfile: "my crypto profile",
            EnableDeadPeerDetection: true,
            DeadPeerDetectionInterval: 7,
            DeadPeerDetectionRetry: 4,
            Ikev2CryptoProfile: "v2 crypto profile",
            Ikev2CookieValidation: true,
            EnableLivenessCheck: true,
            LivenessCheckInterval: 11,
            EnableFragmentation: true,
        }},
        {"test13", version.Number{7, 1, 0, ""}, Entry{
            Name: "test13",
            Version: Ikev1,
            PeerIpType: PeerTypeIp,
            PeerIpValue: "10.1.1.1",
            Interface: "tunnel.1",
            LocalIpAddressType: LocalTypeIp,
            LocalIpAddressValue: "10.2.1.1",
            AuthType: AuthPreSharedKey,
            PreSharedKey: "secret",
            LocalIdType: "ipaddr",
            LocalIdValue: "10.3.1.1",
            PeerIdType: "fqdn",
            PeerIdValue: "example.com",
            PeerIdCheck: "exact",
        }},
        {"test14", version.Number{7, 1, 0, ""}, Entry{
            Name: "test14",
            Version: Ikev1,
            PeerIpType: PeerTypeDynamic,
            Interface: "tunnel.2",
            LocalIpAddressType: LocalTypeIp,
            LocalIpAddressValue: "10.2.1.1",
            AuthType: AuthCertificate,
            LocalCert: "myCert",
            CertProfile: "my cert profile",
            CertPermitPayloadMismatch: true,
            CertEnableStrictValidation: true,
            LocalIdType: "ipaddr",
            LocalIdValue: "10.3.1.1",
            PeerIdType: "fqdn",
            PeerIdValue: "example.com",
            PeerIdCheck: "wildcard",
            Ikev1ExchangeMode: "auto",
            Ikev1CryptoProfile: "my crypto profile",
        }},
        {"test15", version.Number{7, 1, 0, ""}, Entry{
            Name: "test15",
            Version: Ikev1,
            PeerIpType: PeerTypeDynamic,
            Interface: "tunnel.3",
            LocalIpAddressType: LocalTypeIp,
            LocalIpAddressValue: "10.2.1.1",
            AuthType: AuthPreSharedKey,
            PreSharedKey: "another secret",
            LocalIdType: "ipaddr",
            LocalIdValue: "10.3.1.1",
            PeerIdType: "fqdn",
            PeerIdValue: "example.com",
            PeerIdCheck: "wildcard",
            EnablePassiveMode: true,
            EnableNatTraversal: true,
            NatTraversalKeepAlive: 7,
            NatTraversalEnableUdpChecksum: true,
        }},
        {"test16", version.Number{7, 1, 0, ""}, Entry{
            Name: "test16",
            Version: Ikev1,
            PeerIpType: PeerTypeIp,
            PeerIpValue: "10.3.4.5",
            Interface: "tunnel.4",
            LocalIpAddressType: LocalTypeIp,
            LocalIpAddressValue: "10.2.1.1",
            AuthType: AuthPreSharedKey,
            PreSharedKey: "more secrets",
            LocalIdType: "ipaddr",
            LocalIdValue: "10.3.1.1",
            PeerIdType: "fqdn",
            PeerIdValue: "example.com",
            PeerIdCheck: "wildcard",
            EnableDeadPeerDetection: true,
            DeadPeerDetectionInterval: 7,
            DeadPeerDetectionRetry: 4,
            EnableFragmentation: true,
        }},
        {"test17", version.Number{7, 1, 0, ""}, Entry{
            Name: "test17",
            Version: Ikev2,
            PeerIpType: PeerTypeIp,
            PeerIpValue: "10.3.4.5",
            Interface: "tunnel.4",
            LocalIpAddressType: LocalTypeIp,
            LocalIpAddressValue: "10.2.1.1",
            AuthType: AuthPreSharedKey,
            PreSharedKey: "more secrets",
            LocalIdType: "ipaddr",
            LocalIdValue: "10.3.1.1",
            PeerIdType: "fqdn",
            PeerIdValue: "example.com",
            PeerIdCheck: "wildcard",
            Ikev2CryptoProfile: "v2 crypto profile",
            Ikev2CookieValidation: true,
        }},
        {"test18", version.Number{7, 1, 0, ""}, Entry{
            Name: "test18",
            Version: Ikev2,
            PeerIpType: PeerTypeIp,
            PeerIpValue: "10.3.4.5",
            Interface: "tunnel.4",
            LocalIpAddressType: LocalTypeIp,
            LocalIpAddressValue: "10.2.1.1",
            AuthType: AuthPreSharedKey,
            PreSharedKey: "more secrets",
            LocalIdType: "ipaddr",
            LocalIdValue: "10.3.1.1",
            PeerIdType: "fqdn",
            PeerIdValue: "example.com",
            PeerIdCheck: "wildcard",
            EnableLivenessCheck: true,
            LivenessCheckInterval: 7,
        }},
        {"test19", version.Number{7, 1, 0, ""}, Entry{
            Name: "test19",
            Version: Ikev2,
            PeerIpType: PeerTypeIp,
            PeerIpValue: "10.3.4.5",
            Interface: "tunnel.4",
            LocalIpAddressType: LocalTypeIp,
            LocalIpAddressValue: "10.2.1.1",
            AuthType: AuthPreSharedKey,
            PreSharedKey: "more secrets",
            LocalIdType: "ipaddr",
            LocalIdValue: "10.3.1.1",
            PeerIdType: "fqdn",
            PeerIdValue: "example.com",
            PeerIdCheck: "wildcard",
            Ikev2CryptoProfile: "v2 crypto profile",
            Ikev2CookieValidation: true,
            EnableLivenessCheck: true,
            LivenessCheckInterval: 7,
            EnableFragmentation: true,
        }},
        {"test20", version.Number{7, 1, 0, ""}, Entry{
            Name: "test20",
            Version: Ikev2Preferred,
            PeerIpType: PeerTypeIp,
            PeerIpValue: "10.3.4.5",
            Interface: "tunnel.4",
            LocalIpAddressType: LocalTypeIp,
            LocalIpAddressValue: "10.2.1.1",
            AuthType: AuthPreSharedKey,
            PreSharedKey: "more secrets",
            LocalIdType: "ipaddr",
            LocalIdValue: "10.3.1.1",
            PeerIdType: "fqdn",
            PeerIdValue: "example.com",
            PeerIdCheck: "wildcard",
            Ikev1ExchangeMode: "auto",
            Ikev1CryptoProfile: "my crypto profile",
            EnableDeadPeerDetection: true,
            DeadPeerDetectionInterval: 7,
            DeadPeerDetectionRetry: 4,
            Ikev2CryptoProfile: "v2 crypto profile",
            Ikev2CookieValidation: true,
            EnableLivenessCheck: true,
            LivenessCheckInterval: 11,
            EnableFragmentation: true,
        }},
        {"test21", version.Number{7, 1, 0, ""}, Entry{
            Name: "test21",
            Version: Ikev2Preferred,
            PeerIpType: PeerTypeIp,
            PeerIpValue: "10.3.4.5",
            Interface: "tunnel.4",
            LocalIpAddressType: LocalTypeFloatingIp,
            LocalIpAddressValue: "10.2.1.1",
            AuthType: AuthPreSharedKey,
            PreSharedKey: "more secrets",
            LocalIdType: "ipaddr",
            LocalIdValue: "10.3.1.1",
            PeerIdType: "fqdn",
            PeerIdValue: "example.com",
            PeerIdCheck: "wildcard",
            Ikev1ExchangeMode: "auto",
            Ikev1CryptoProfile: "my crypto profile",
            EnableDeadPeerDetection: true,
            DeadPeerDetectionInterval: 7,
            DeadPeerDetectionRetry: 4,
            Ikev2CryptoProfile: "v2 crypto profile",
            Ikev2CookieValidation: true,
            EnableLivenessCheck: true,
            LivenessCheckInterval: 11,
            EnableFragmentation: true,
        }},
        {"test22", version.Number{8, 1, 0, ""}, Entry{
            Name: "test22",
            Version: Ikev1,
            PeerIpType: PeerTypeIp,
            PeerIpValue: "10.1.1.1",
            Interface: "tunnel.1",
            LocalIpAddressType: LocalTypeIp,
            LocalIpAddressValue: "10.2.1.1",
            AuthType: AuthPreSharedKey,
            PreSharedKey: "secret",
            LocalIdType: "ipaddr",
            LocalIdValue: "10.3.1.1",
            PeerIdType: "fqdn",
            PeerIdValue: "example.com",
            PeerIdCheck: "exact",
        }},
        {"test23", version.Number{8, 1, 0, ""}, Entry{
            Name: "test23",
            Version: Ikev1,
            PeerIpType: PeerTypeDynamic,
            Interface: "tunnel.2",
            LocalIpAddressType: LocalTypeIp,
            LocalIpAddressValue: "10.2.1.1",
            AuthType: AuthCertificate,
            LocalCert: "myCert",
            CertProfile: "my cert profile",
            CertPermitPayloadMismatch: true,
            CertEnableStrictValidation: true,
            LocalIdType: "ipaddr",
            LocalIdValue: "10.3.1.1",
            PeerIdType: "fqdn",
            PeerIdValue: "example.com",
            PeerIdCheck: "wildcard",
            Ikev1ExchangeMode: "auto",
            Ikev1CryptoProfile: "my crypto profile",
        }},
        {"test24", version.Number{8, 1, 0, ""}, Entry{
            Name: "test24",
            Version: Ikev1,
            PeerIpType: PeerTypeDynamic,
            Interface: "tunnel.3",
            LocalIpAddressType: LocalTypeIp,
            LocalIpAddressValue: "10.2.1.1",
            AuthType: AuthPreSharedKey,
            PreSharedKey: "another secret",
            LocalIdType: "ipaddr",
            LocalIdValue: "10.3.1.1",
            PeerIdType: "fqdn",
            PeerIdValue: "example.com",
            PeerIdCheck: "wildcard",
            EnablePassiveMode: true,
            EnableNatTraversal: true,
            NatTraversalKeepAlive: 7,
            NatTraversalEnableUdpChecksum: true,
        }},
        {"test25", version.Number{8, 1, 0, ""}, Entry{
            Name: "test25",
            Version: Ikev1,
            PeerIpType: PeerTypeIp,
            PeerIpValue: "10.3.4.5",
            Interface: "tunnel.4",
            LocalIpAddressType: LocalTypeIp,
            LocalIpAddressValue: "10.2.1.1",
            AuthType: AuthPreSharedKey,
            PreSharedKey: "more secrets",
            LocalIdType: "ipaddr",
            LocalIdValue: "10.3.1.1",
            PeerIdType: "fqdn",
            PeerIdValue: "example.com",
            PeerIdCheck: "wildcard",
            EnableDeadPeerDetection: true,
            DeadPeerDetectionInterval: 7,
            DeadPeerDetectionRetry: 4,
            EnableFragmentation: true,
        }},
        {"test26", version.Number{8, 1, 0, ""}, Entry{
            Name: "test26",
            Version: Ikev2,
            PeerIpType: PeerTypeIp,
            PeerIpValue: "10.3.4.5",
            Interface: "tunnel.4",
            LocalIpAddressType: LocalTypeIp,
            LocalIpAddressValue: "10.2.1.1",
            AuthType: AuthPreSharedKey,
            PreSharedKey: "more secrets",
            LocalIdType: "ipaddr",
            LocalIdValue: "10.3.1.1",
            PeerIdType: "fqdn",
            PeerIdValue: "example.com",
            PeerIdCheck: "wildcard",
            Ikev2CryptoProfile: "v2 crypto profile",
            Ikev2CookieValidation: true,
        }},
        {"test27", version.Number{8, 1, 0, ""}, Entry{
            Name: "test27",
            Version: Ikev2,
            PeerIpType: PeerTypeIp,
            PeerIpValue: "10.3.4.5",
            Interface: "tunnel.4",
            LocalIpAddressType: LocalTypeIp,
            LocalIpAddressValue: "10.2.1.1",
            AuthType: AuthPreSharedKey,
            PreSharedKey: "more secrets",
            LocalIdType: "ipaddr",
            LocalIdValue: "10.3.1.1",
            PeerIdType: "fqdn",
            PeerIdValue: "example.com",
            PeerIdCheck: "wildcard",
            EnableLivenessCheck: true,
            LivenessCheckInterval: 7,
        }},
        {"test28", version.Number{8, 1, 0, ""}, Entry{
            Name: "test28",
            Version: Ikev2,
            PeerIpType: PeerTypeIp,
            PeerIpValue: "10.3.4.5",
            Interface: "tunnel.4",
            LocalIpAddressType: LocalTypeIp,
            LocalIpAddressValue: "10.2.1.1",
            AuthType: AuthPreSharedKey,
            PreSharedKey: "more secrets",
            LocalIdType: "ipaddr",
            LocalIdValue: "10.3.1.1",
            PeerIdType: "fqdn",
            PeerIdValue: "example.com",
            PeerIdCheck: "wildcard",
            Ikev2CryptoProfile: "v2 crypto profile",
            Ikev2CookieValidation: true,
            EnableLivenessCheck: true,
            LivenessCheckInterval: 7,
            EnableFragmentation: true,
        }},
        {"test29", version.Number{8, 1, 0, ""}, Entry{
            Name: "test29",
            Version: Ikev2Preferred,
            PeerIpType: PeerTypeIp,
            PeerIpValue: "10.3.4.5",
            Interface: "tunnel.4",
            LocalIpAddressType: LocalTypeIp,
            LocalIpAddressValue: "10.2.1.1",
            AuthType: AuthPreSharedKey,
            PreSharedKey: "more secrets",
            LocalIdType: "ipaddr",
            LocalIdValue: "10.3.1.1",
            PeerIdType: "fqdn",
            PeerIdValue: "example.com",
            PeerIdCheck: "wildcard",
            Ikev1ExchangeMode: "auto",
            Ikev1CryptoProfile: "my crypto profile",
            EnableDeadPeerDetection: true,
            DeadPeerDetectionInterval: 7,
            DeadPeerDetectionRetry: 4,
            Ikev2CryptoProfile: "v2 crypto profile",
            Ikev2CookieValidation: true,
            EnableLivenessCheck: true,
            LivenessCheckInterval: 11,
            EnableFragmentation: true,
        }},
        {"test30", version.Number{8, 1, 0, ""}, Entry{
            Name: "test30",
            Version: Ikev2Preferred,
            PeerIpType: PeerTypeFqdn,
            PeerIpValue: "panw.com",
            Interface: "tunnel.4",
            LocalIpAddressType: LocalTypeIp,
            LocalIpAddressValue: "10.2.1.1",
            AuthType: AuthPreSharedKey,
            PreSharedKey: "more secrets",
            LocalIdType: "ipaddr",
            LocalIdValue: "10.3.1.1",
            PeerIdType: "fqdn",
            PeerIdValue: "example.com",
            PeerIdCheck: "wildcard",
            Ikev1ExchangeMode: "auto",
            Ikev1CryptoProfile: "my crypto profile",
            EnableDeadPeerDetection: true,
            DeadPeerDetectionInterval: 7,
            DeadPeerDetectionRetry: 4,
            Ikev2CryptoProfile: "v2 crypto profile",
            Ikev2CookieValidation: true,
            EnableLivenessCheck: true,
            LivenessCheckInterval: 11,
            EnableFragmentation: true,
        }},
        {"test31", version.Number{8, 1, 0, ""}, Entry{
            Name: "test31",
            Version: Ikev2Preferred,
            PeerIpType: PeerTypeFqdn,
            PeerIpValue: "panw.com",
            Interface: "tunnel.4",
            LocalIpAddressType: LocalTypeFloatingIp,
            LocalIpAddressValue: "10.2.1.1",
            AuthType: AuthPreSharedKey,
            PreSharedKey: "more secrets",
            LocalIdType: "ipaddr",
            LocalIdValue: "10.3.1.1",
            PeerIdType: "fqdn",
            PeerIdValue: "example.com",
            PeerIdCheck: "wildcard",
            Ikev1ExchangeMode: "auto",
            Ikev1CryptoProfile: "my crypto profile",
            EnableDeadPeerDetection: true,
            DeadPeerDetectionInterval: 7,
            DeadPeerDetectionRetry: 4,
            Ikev2CryptoProfile: "v2 crypto profile",
            Ikev2CookieValidation: true,
            EnableLivenessCheck: true,
            LivenessCheckInterval: 11,
            EnableFragmentation: true,
        }},
    }

    mc := &testdata.MockClient{}
    ns := &FwIkeGw{}
    ns.Initialize(mc)

    for _, tc := range testCases {
        t.Run(tc.desc, func(t *testing.T) {
            mc.AddResp("")
            mc.Version = tc.version
            err := ns.Set(tc.conf)
            if err != nil {
                t.Errorf("Error in set: %s", err)
            } else {
                mc.AddResp(mc.Elm)
                r, err := ns.Get(tc.conf.Name)
                if err != nil {
                    t.Errorf("Error in get: %s", err)
                } else if !reflect.DeepEqual(tc.conf, r) {
                    t.Errorf("%#v != %#v", tc.conf, r)
                }
            }
        })
    }
}
