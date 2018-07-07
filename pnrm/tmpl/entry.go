package tmpl

import (
    "encoding/xml"

    "github.com/PaloAltoNetworks/pango/util"
)

// Entry is a normalized, version independent representation of a template.
//
// Devices is a map where the key is the serial number of the target device and
// the value is a list of specific vsys on that device.  The list of vsys is
// nil if all vsys on that device should be included or if the device is a
// virtual firewall (and thus only has vsys1).
type Entry struct {
    Name string
    Description string
    DefaultVsys string
    MultiVsys bool
    Mode string
    VpnDisableMode bool
    Devices map[string] []string
}

// Copy copies the information from source's Entry `s` to this object.  As the
// Name field relates to the XPATH of this object, this field is not copied.
func (o *Entry) Copy(s Entry) {
    o.Description = s.Description
    o.DefaultVsys = s.DefaultVsys
    o.MultiVsys = s.MultiVsys
    o.Mode = s.Mode
    o.VpnDisableMode = s.VpnDisableMode
    o.Devices = s.Devices
}

/** Structs / functions for normalization. **/

type normalizer interface {
    Normalize() Entry
}

type container_v1 struct {
    Answer entry_v1 `xml:"result>entry"`
}

func (o *container_v1) Normalize() Entry {
    ans := Entry{
        Name: o.Answer.Name,
        Description: o.Answer.Description,
        Devices: util.VsysEntToMap(o.Answer.Devices),
    }

    if o.Answer.Settings != nil {
        ans.MultiVsys = util.AsBool(o.Answer.Settings.MultiVsys)
        ans.Mode = o.Answer.Settings.Mode
        ans.VpnDisableMode = util.AsBool(o.Answer.Settings.VpnDisableMode)
    }

    return ans
}

type container_v2 struct {
    Answer entry_v2 `xml:"result>entry"`
}

func (o *container_v2) Normalize() Entry {
    ans := Entry{
        Name: o.Answer.Name,
        Description: o.Answer.Description,
        Devices: util.VsysEntToMap(o.Answer.Devices),
    }

    if o.Answer.Settings != nil {
        ans.DefaultVsys = o.Answer.Settings.DefaultVsys
    }

    return ans
}

type entry_v1 struct {
    XMLName xml.Name `xml:"entry"`
    Name string `xml:"name,attr"`
    Description string `xml:"description,omitempty"`
    Devices *util.VsysEntryType `xml:"devices"`
    Settings *settings_v1 `xml:"settings"`
}

type settings_v1 struct {
    MultiVsys string `xml:"multi-vsys"`
    Mode string `xml:"operational-mode,omitempty"`
    VpnDisableMode string `xml:"vpn-disable-mode"`
}

func specify_v1(e Entry) interface{} {
    ans := entry_v1{
        Name: e.Name,
        Description: e.Description,
        Devices: util.MapToVsysEnt(e.Devices),
    }

    if e.MultiVsys || e.VpnDisableMode || e.Mode != "" {
        ans.Settings = &settings_v1{
            MultiVsys: util.YesNo(e.MultiVsys),
            Mode: e.Mode,
            VpnDisableMode: util.YesNo(e.VpnDisableMode),
        }
    }

    return ans
}

type entry_v2 struct {
    XMLName xml.Name `xml:"entry"`
    Name string `xml:"name,attr"`
    Description string `xml:"description,omitempty"`
    Devices *util.VsysEntryType `xml:"devices"`
    Settings *settings_v2 `xml:"settings"`
}

type settings_v2 struct {
    DefaultVsys string `xml:"default-vsys"`
}

func specify_v2(e Entry) interface{} {
    ans := entry_v2{
        Name: e.Name,
        Description: e.Description,
        Devices: util.MapToVsysEnt(e.Devices),
    }

    if e.DefaultVsys != "" {
        ans.Settings = &settings_v2{e.DefaultVsys}
    }

    return ans
}
