package collector

import (
	"github.com/prometheus-community/windows_exporter/internal/collector/ad"
	"github.com/prometheus-community/windows_exporter/internal/collector/adcs"
	"github.com/prometheus-community/windows_exporter/internal/collector/adfs"
	"github.com/prometheus-community/windows_exporter/internal/collector/cache"
	"github.com/prometheus-community/windows_exporter/internal/collector/container"
	"github.com/prometheus-community/windows_exporter/internal/collector/cpu"
	"github.com/prometheus-community/windows_exporter/internal/collector/cpu_info"
	"github.com/prometheus-community/windows_exporter/internal/collector/cs"
	"github.com/prometheus-community/windows_exporter/internal/collector/dfsr"
	"github.com/prometheus-community/windows_exporter/internal/collector/dhcp"
	"github.com/prometheus-community/windows_exporter/internal/collector/diskdrive"
	"github.com/prometheus-community/windows_exporter/internal/collector/dns"
	"github.com/prometheus-community/windows_exporter/internal/collector/exchange"
	"github.com/prometheus-community/windows_exporter/internal/collector/filetime"
	"github.com/prometheus-community/windows_exporter/internal/collector/fsrmquota"
	"github.com/prometheus-community/windows_exporter/internal/collector/hyperv"
	"github.com/prometheus-community/windows_exporter/internal/collector/iis"
	"github.com/prometheus-community/windows_exporter/internal/collector/license"
	"github.com/prometheus-community/windows_exporter/internal/collector/logical_disk"
	"github.com/prometheus-community/windows_exporter/internal/collector/logon"
	"github.com/prometheus-community/windows_exporter/internal/collector/memory"
	"github.com/prometheus-community/windows_exporter/internal/collector/mscluster"
	"github.com/prometheus-community/windows_exporter/internal/collector/msmq"
	"github.com/prometheus-community/windows_exporter/internal/collector/mssql"
	"github.com/prometheus-community/windows_exporter/internal/collector/net"
	"github.com/prometheus-community/windows_exporter/internal/collector/netframework"
	"github.com/prometheus-community/windows_exporter/internal/collector/nps"
	"github.com/prometheus-community/windows_exporter/internal/collector/os"
	"github.com/prometheus-community/windows_exporter/internal/collector/perfdata"
	"github.com/prometheus-community/windows_exporter/internal/collector/physical_disk"
	"github.com/prometheus-community/windows_exporter/internal/collector/printer"
	"github.com/prometheus-community/windows_exporter/internal/collector/process"
	"github.com/prometheus-community/windows_exporter/internal/collector/remote_fx"
	"github.com/prometheus-community/windows_exporter/internal/collector/scheduled_task"
	"github.com/prometheus-community/windows_exporter/internal/collector/service"
	"github.com/prometheus-community/windows_exporter/internal/collector/smb"
	"github.com/prometheus-community/windows_exporter/internal/collector/smbclient"
	"github.com/prometheus-community/windows_exporter/internal/collector/smtp"
	"github.com/prometheus-community/windows_exporter/internal/collector/system"
	"github.com/prometheus-community/windows_exporter/internal/collector/tcp"
	"github.com/prometheus-community/windows_exporter/internal/collector/terminal_services"
	"github.com/prometheus-community/windows_exporter/internal/collector/textfile"
	"github.com/prometheus-community/windows_exporter/internal/collector/thermalzone"
	"github.com/prometheus-community/windows_exporter/internal/collector/time"
	"github.com/prometheus-community/windows_exporter/internal/collector/update"
	"github.com/prometheus-community/windows_exporter/internal/collector/vmware"
)

type Config struct {
	AD               ad.Config                `yaml:"ad"`
	ADCS             adcs.Config              `yaml:"adcs"`
	ADFS             adfs.Config              `yaml:"adfs"`
	Cache            cache.Config             `yaml:"cache"`
	Container        container.Config         `yaml:"container"`
	CPU              cpu.Config               `yaml:"cpu"`
	CPUInfo          cpu_info.Config          `yaml:"cpu_info"`
	Cs               cs.Config                `yaml:"cs"`
	DFSR             dfsr.Config              `yaml:"dfsr"`
	Dhcp             dhcp.Config              `yaml:"dhcp"`
	DiskDrive        diskdrive.Config         `yaml:"disk_drive"`
	DNS              dns.Config               `yaml:"dns"`
	Exchange         exchange.Config          `yaml:"exchange"`
	Filetime         filetime.Config          `yaml:"filetime"`
	Fsrmquota        fsrmquota.Config         `yaml:"fsrmquota"`
	HyperV           hyperv.Config            `yaml:"hyper_v"`
	IIS              iis.Config               `yaml:"iis"`
	License          license.Config           `yaml:"license"`
	LogicalDisk      logical_disk.Config      `yaml:"logical_disk"`
	Logon            logon.Config             `yaml:"logon"`
	Memory           memory.Config            `yaml:"memory"`
	MSCluster        mscluster.Config         `yaml:"ms_cluster"`
	Msmq             msmq.Config              `yaml:"msmq"`
	Mssql            mssql.Config             `yaml:"mssql"`
	Net              net.Config               `yaml:"net"`
	NetFramework     netframework.Config      `yaml:"net_framework"`
	Nps              nps.Config               `yaml:"nps"`
	OS               os.Config                `yaml:"os"`
	PerfData         perfdata.Config          `yaml:"perf_data"`
	PhysicalDisk     physical_disk.Config     `yaml:"physical_disk"`
	Printer          printer.Config           `yaml:"printer"`
	Process          process.Config           `yaml:"process"`
	RemoteFx         remote_fx.Config         `yaml:"remote_fx"`
	ScheduledTask    scheduled_task.Config    `yaml:"scheduled_task"`
	Service          service.Config           `yaml:"service"`
	SMB              smb.Config               `yaml:"smb"`
	SMBClient        smbclient.Config         `yaml:"smb_client"`
	SMTP             smtp.Config              `yaml:"smtp"`
	System           system.Config            `yaml:"system"`
	TCP              tcp.Config               `yaml:"tcp"`
	TerminalServices terminal_services.Config `yaml:"terminal_services"`
	Textfile         textfile.Config          `yaml:"textfile"`
	ThermalZone      thermalzone.Config       `yaml:"thermal_zone"`
	Time             time.Config              `yaml:"time"`
	Update           update.Config            `yaml:"update"`
	Vmware           vmware.Config            `yaml:"vmware"`
}

// ConfigDefaults Is an interface to be used by the external libraries. It holds all ConfigDefaults form all collectors
//
//goland:noinspection GoUnusedGlobalVariable
var ConfigDefaults = Config{
	AD:               ad.ConfigDefaults,
	ADCS:             adcs.ConfigDefaults,
	ADFS:             adfs.ConfigDefaults,
	Cache:            cache.ConfigDefaults,
	Container:        container.ConfigDefaults,
	CPU:              cpu.ConfigDefaults,
	CPUInfo:          cpu_info.ConfigDefaults,
	Cs:               cs.ConfigDefaults,
	DFSR:             dfsr.ConfigDefaults,
	Dhcp:             dhcp.ConfigDefaults,
	DiskDrive:        diskdrive.ConfigDefaults,
	DNS:              dns.ConfigDefaults,
	Exchange:         exchange.ConfigDefaults,
	Filetime:         filetime.ConfigDefaults,
	Fsrmquota:        fsrmquota.ConfigDefaults,
	HyperV:           hyperv.ConfigDefaults,
	IIS:              iis.ConfigDefaults,
	License:          license.ConfigDefaults,
	LogicalDisk:      logical_disk.ConfigDefaults,
	Logon:            logon.ConfigDefaults,
	Memory:           memory.ConfigDefaults,
	MSCluster:        mscluster.ConfigDefaults,
	Msmq:             msmq.ConfigDefaults,
	Mssql:            mssql.ConfigDefaults,
	Net:              net.ConfigDefaults,
	NetFramework:     netframework.ConfigDefaults,
	Nps:              nps.ConfigDefaults,
	OS:               os.ConfigDefaults,
	PerfData:         perfdata.ConfigDefaults,
	PhysicalDisk:     physical_disk.ConfigDefaults,
	Printer:          printer.ConfigDefaults,
	Process:          process.ConfigDefaults,
	RemoteFx:         remote_fx.ConfigDefaults,
	ScheduledTask:    scheduled_task.ConfigDefaults,
	Service:          service.ConfigDefaults,
	SMB:              smb.ConfigDefaults,
	SMBClient:        smbclient.ConfigDefaults,
	SMTP:             smtp.ConfigDefaults,
	System:           system.ConfigDefaults,
	TCP:              tcp.ConfigDefaults,
	TerminalServices: terminal_services.ConfigDefaults,
	Textfile:         textfile.ConfigDefaults,
	ThermalZone:      thermalzone.ConfigDefaults,
	Time:             time.ConfigDefaults,
	Update:           update.ConfigDefaults,
	Vmware:           vmware.ConfigDefaults,
}
