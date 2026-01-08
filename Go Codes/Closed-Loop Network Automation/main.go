package main

import (
	"fmt"
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Device DeviceConfig `yaml:"device"`
	Intent Intent       `yaml:"intent"`
}

type DeviceConfig struct {
	Hostname string `yaml:"hostname"`
	Platform string `yaml:"platform"`
}

type Intent struct {
	Service string `yaml:"service"`
	Port    int    `yaml:"port"`
	TLS     bool   `yaml:"tls"`
}

type OperState struct {
	Service string
	Port    int
	TLS     bool
}

type Device struct {
	Name  string
	State OperState
}

func NewDevice(name string) *Device {
	return &Device{
		Name: name,
		State: OperState{
			Service: "grpc",
			Port:    57777,
			TLS:     false,
		},
	}
}

func (d *Device) GetOperState() OperState {
	fmt.Println("Operasyonel durum alındı")
	return d.State
}

func isCompliant(intent Intent, oper OperState) bool {
	return intent.Service == oper.Service &&
		intent.Port == oper.Port &&
		intent.TLS == oper.TLS
}

func (d *Device) ApplyIntent(intent Intent) {
	fmt.Println(" Drift bulundu → Konfigürasyon uygulanıyor")

	d.State.Service = intent.Service
	d.State.Port = intent.Port
	d.State.TLS = intent.TLS
}

func enforcementLoop(device *Device, intent Intent) {
	for {
		fmt.Println("\n Kontrol döngüsü")

		oper := device.GetOperState()

		if isCompliant(intent, oper) {
			fmt.Println("Sistem istenen durumda")
		} else {
			fmt.Println("Drift tespit edildi")
			device.ApplyIntent(intent)
		}

		fmt.Printf(
			"State → Service:%s Port:%d TLS:%v\n",
			device.State.Service,
			device.State.Port,
			device.State.TLS,
		)

		time.Sleep(5 * time.Second)
	}
}

func main() {

	data, err := os.ReadFile("input.yml")
	if err != nil {
		panic(err)
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		panic(err)
	}

	fmt.Println("YAML yüklendi")
	fmt.Printf("Intent → %+v\n", cfg.Intent)

	device := NewDevice(cfg.Device.Hostname)

	fmt.Println("Closed-loop automation başladı")
	enforcementLoop(device, cfg.Intent)
}
