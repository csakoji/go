// Copyright © 2015-2018 Platina Systems, Inc. All rights reserved.
// Use of this source code is governed by the GPL-2 license described in the
// LICENSE file.

package serverd

import (
	"encoding/json"
	"fmt"
	"time"
)

func register(u *[]string) (s string, err error) {
	if len(*u) < 2 {
		return "", err
	}

	err = json.Unmarshal([]byte((*u)[1]), &regReq)
	if err != nil {
		fmt.Println("There was an error:", err)
	}
	mac := regReq.Mac
	ip := regReq.IP

	ClientCfg[mac].MacAddr = mac
	ClientCfg[mac].IpAddr = ip
	ClientCfg[mac].BootState = BootStateRegistered
	ClientCfg[mac].InstallState = InstallStateInProgress
	t := time.Now()
	ClientCfg[mac].TimeRegistered = fmt.Sprintf("%10s",
		t.Format("2006-01-02 15:04:05"))
	ClientCfg[mac].TimeInstalled = fmt.Sprintf("%10s",
		t.Format("2006-01-02 15:04:05"))
	ClientCfg[mac].InstallCounter++

	regReply.Reply = RegReplyRegistered
	regReply.TorName = ClientCfg[mac].Name
	regReply.Error = nil
	jsonInfo, err := json.Marshal(regReply)
	if err != nil {
		return "404", err
	}

	return string(jsonInfo), nil
}

func dumpvars() (s string, err error) {
	s = ""
	return s, nil
}

func numclients() (s string, err error) {
	var numReply NumClntReply

	numReply.Clients = len(ClientCfg)
	jsonInfo, err := json.Marshal(numReply)
	if err != nil {
		return "404", err
	}
	return string(jsonInfo), nil
}

func clientdata(j int) (s string, err error) {
	for i, _ := range ClientCfg {
		if ClientCfg[i].Unit == j {
			jsonInfo, err := json.Marshal(ClientCfg[i])
			if err != nil {
				return "", err
			}
			return string(jsonInfo), nil
		}
	}
	err = fmt.Errorf("client number not found: %v", err)
	return "", nil
}

func clientbootdata(j int) (s string, err error) {
	for i, _ := range ClientBootCfg {
		if ClientCfg[i].Unit == j {
			jsonInfo, err := json.Marshal(ClientBootCfg[i])
			if err != nil {
				return "", err
			}
			return string(jsonInfo), nil
		}
	}
	err = fmt.Errorf("client number not found: %v", err)
	return "", nil
}

// TODO
func bootStateMachine() {
}

// TODO
func installStateMachine() {
}

func readClientCfgDB() (err error) {
	// TODO try reading from cloud DB

	// TODO try reading from local DB

	// default to literal for testing
	ClientCfg["01:02:03:04:05:06"] = &Client{
		Unit:           1,
		Name:           "Invader10",
		Machine:        "ToR MK1",
		MacAddr:        "01:02:03:04:05:06",
		IpAddr:         "0.0.0.0",
		BootState:      BootStateNotRegistered,
		InstallState:   InstallStateFactory,
		AutoInstall:    true,
		CertPresent:    false,
		DistroType:     Debian,
		TimeRegistered: "0000-00-00:00:00:00",
		TimeInstalled:  "0000-00-00:00:00:00",
		InstallCounter: 0,
	}
	ClientCfg["01:02:03:04:05:07"] = &Client{
		Unit:           2,
		Name:           "Invader11",
		Machine:        "ToR MK1",
		MacAddr:        "01:02:03:04:05:07",
		IpAddr:         "0.0.0.0",
		BootState:      BootStateNotRegistered,
		InstallState:   InstallStateFactory,
		AutoInstall:    true,
		CertPresent:    false,
		DistroType:     Debian,
		TimeRegistered: "0000-00-00:00:00:00",
		TimeInstalled:  "0000-00-00:00:00:00",
		InstallCounter: 0,
	}
	ClientCfg["01:02:03:04:05:08"] = &Client{
		Unit:           3,
		Name:           "Invader12",
		Machine:        "ToR MK1",
		MacAddr:        "01:02:03:04:05:08",
		IpAddr:         "0.0.0.0",
		BootState:      BootStateNotRegistered,
		InstallState:   InstallStateFactory,
		AutoInstall:    true,
		CertPresent:    false,
		DistroType:     Debian,
		TimeRegistered: "0000-00-00:00:00:00",
		TimeInstalled:  "0000-00-00:00:00:00",
		InstallCounter: 0,
	}
	ClientBootCfg["01:02:03:04:05:06"] = &BootcConfig{
		Install:         false,
		BootSda1:        false,
		BootSda6Cnt:     3,
		EraseSda6:       false,
		IAmMaster:       false,
		MyIpAddr:        "192.168.101.129",
		MyGateway:       "192.168.101.1",
		MyNetmask:       "255.255.255.0",
		MasterAddresses: []string{"198.168.101.142"},
		ReInstallK:      "/newroot/sda1/boot/vmlinuz",
		ReInstallI:      "/newroot/sda1/boot/initrd.gz",
		ReInstallC:      `netcfg/get_hostname=platina netcfg/get_domain=platinasystems.com interface=auto auto locale=en_US preseed/file=/hd-media/preseed.cfg`,
		Sda1K:           "/newroot/sda1/boot/vmlinuz-3.16.0-4-amd64",
		Sda1I:           "/newroot/sda1/boot/initrd.img-3.16.0-4-amd64",
		Sda1C:           "::eth0:none",
		Sda6K:           "/newroot/sda6/boot/vmlinuz-3.16.0-4-amd64",
		Sda6I:           "/newroot/sda6/boot/initrd.img-3.16.0-4-amd64",
		Sda6C:           "::eth0:none",
		ISO1Name:        "debian-8.10.0-amd64-DVD-1.iso",
		ISO1Desc:        "Jessie debian-8.10.0",
		ISO2Name:        " ",
		ISO2Desc:        " ",
		ISOlastUsed:     1,
	}
	ClientBootCfg["01:02:03:04:05:07"] = &BootcConfig{
		Install:         false,
		BootSda1:        false,
		BootSda6Cnt:     3,
		EraseSda6:       false,
		IAmMaster:       false,
		MyIpAddr:        "192.168.101.130",
		MyGateway:       "192.168.101.1",
		MyNetmask:       "255.255.255.0",
		MasterAddresses: []string{"198.168.101.142"},
		ReInstallK:      "/newroot/sda1/boot/vmlinuz",
		ReInstallI:      "/newroot/sda1/boot/initrd.gz",
		ReInstallC:      `netcfg/get_hostname=platina netcfg/get_domain=platinasystems.com interface=auto auto locale=en_US preseed/file=/hd-media/preseed.cfg`,
		Sda1K:           "/newroot/sda1/boot/vmlinuz-3.16.0-4-amd64",
		Sda1I:           "/newroot/sda1/boot/initrd.img-3.16.0-4-amd64",
		Sda1C:           "::eth0:none",
		Sda6K:           "/newroot/sda6/boot/vmlinuz-3.16.0-4-amd64",
		Sda6I:           "/newroot/sda6/boot/initrd.img-3.16.0-4-amd64",
		Sda6C:           "::eth0:none",
		ISO1Name:        "debian-8.10.0-amd64-DVD-1.iso",
		ISO1Desc:        "Jessie debian-8.10.0",
		ISO2Name:        "",
		ISO2Desc:        "",
		ISOlastUsed:     1,
	}
	ClientBootCfg["01:02:03:04:05:07"] = &BootcConfig{
		Install:         false,
		BootSda1:        false,
		BootSda6Cnt:     3,
		EraseSda6:       false,
		IAmMaster:       false,
		MyIpAddr:        "192.168.101.131",
		MyGateway:       "192.168.101.1",
		MyNetmask:       "255.255.255.0",
		MasterAddresses: []string{"198.168.101.142"},
		ReInstallK:      "/newroot/sda1/boot/vmlinuz",
		ReInstallI:      "/newroot/sda1/boot/initrd.gz",
		ReInstallC:      `netcfg/get_hostname=platina netcfg/get_domain=platinasystems.com interface=auto auto locale=en_US preseed/file=/hd-media/preseed.cfg`,
		Sda1K:           "/newroot/sda1/boot/vmlinuz-3.16.0-4-amd64",
		Sda1I:           "/newroot/sda1/boot/initrd.img-3.16.0-4-amd64",
		Sda1C:           "::eth0:none",
		Sda6K:           "/newroot/sda6/boot/vmlinuz-3.16.0-4-amd64",
		Sda6I:           "/newroot/sda6/boot/initrd.img-3.16.0-4-amd64",
		Sda6C:           "::eth0:none",
		ISO1Name:        "debian-8.10.0-amd64-DVD-1.iso",
		ISO1Desc:        "Jessie debian-8.10.0",
		ISO2Name:        "",
		ISO2Desc:        "",
		ISOlastUsed:     1,
	}
	return nil
}
