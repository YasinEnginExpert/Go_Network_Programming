// Package main implements a closed-loop network automation controller.
//
// Closed-loop automation is a network management paradigm where the system
// continuously monitors the actual state (operational state) and compares it
// to the desired state (intent). When drift is detected, the system automatically
// applies corrections to bring the network back to the desired state.
//
// Architecture:
//
//	                Closed-Loop Automation
//	┌──────────────────────────────────────────────────────┐
//	│                                                      │
//	│   ┌─────────┐    Compare    ┌─────────────────────┐  │
//	│   │ Intent  │◄─────────────►│  Operational State  │  │
//	│   │ (YAML)  │               │  (Device State)     │  │
//	│   └────┬────┘               └──────────┬──────────┘  │
//	│        │                               │             │
//	│        │         ┌───────────┐         │             │
//	│        └────────►│ Controller │◄───────┘             │
//	│                  └─────┬─────┘                       │
//	│                        │                             │
//	│                        ▼                             │
//	│          ┌──────────────────────────┐                │
//	│          │ Apply Configuration      │                │
//	│          │ (if drift detected)      │                │
//	│          └──────────────────────────┘                │
//	│                                                      │
//	└──────────────────────────────────────────────────────┘
//
// Key Concepts:
//   - Intent: The desired configuration state defined in YAML
//   - Operational State: The current actual state of the device
//   - Drift: Difference between intent and operational state
//   - Reconciliation: Process of applying intent to eliminate drift
//
// This example simulates a gRPC service configuration controller that
// ensures the service port and TLS settings match the defined intent.
//
// Usage:
//
//	go run main.go
//
// Requires input.yml file with the following structure:
//
//	device:
//	  hostname: router1
//	  platform: ios-xr
//	intent:
//	  service: gnmi
//	  port: 57400
//	  tls: true
package main

import (
	"fmt"
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

// Config represents the complete configuration loaded from YAML.
// It contains both device information and the desired intent.
type Config struct {
	Device DeviceConfig `yaml:"device"`
	Intent Intent       `yaml:"intent"`
}

// DeviceConfig holds device identification information.
type DeviceConfig struct {
	Hostname string `yaml:"hostname"` // Device hostname or identifier
	Platform string `yaml:"platform"` // Platform type (e.g., ios-xr, junos)
}

// Intent represents the desired configuration state.
// This is what we want the device to be configured as.
type Intent struct {
	Service string `yaml:"service"` // Service name (e.g., grpc, gnmi)
	Port    int    `yaml:"port"`    // Service port number
	TLS     bool   `yaml:"tls"`     // TLS enabled flag
}

// OperState represents the actual operational state of a device.
// This is retrieved from the device to compare against intent.
type OperState struct {
	Service string // Currently running service
	Port    int    // Current port configuration
	TLS     bool   // Current TLS setting
}

// Device represents a network device with its current state.
type Device struct {
	Name  string    // Device name/hostname
	State OperState // Current operational state
}

// NewDevice creates a new Device with default operational state.
// In a real implementation, this would query the actual device.
//
// Parameters:
//   - name: The device hostname or identifier
//
// Returns:
//   - Pointer to the new Device instance
func NewDevice(name string) *Device {
	return &Device{
		Name: name,
		// Simulate initial device state (different from intent)
		State: OperState{
			Service: "grpc",
			Port:    57777,
			TLS:     false,
		},
	}
}

// GetOperState retrieves the current operational state from the device.
// In a real implementation, this would use gNMI, NETCONF, or similar.
//
// Returns:
//   - The current operational state
func (d *Device) GetOperState() OperState {
	fmt.Printf("  [%s] Retrieving operational state...\n", d.Name)
	return d.State
}

// isCompliant checks if the operational state matches the intent.
//
// Parameters:
//   - intent: The desired configuration
//   - oper: The current operational state
//
// Returns:
//   - true if all fields match, false otherwise
func isCompliant(intent Intent, oper OperState) bool {
	return intent.Service == oper.Service &&
		intent.Port == oper.Port &&
		intent.TLS == oper.TLS
}

// ApplyIntent applies the desired configuration to the device.
// In a real implementation, this would push configuration via gNMI/NETCONF.
//
// Parameters:
//   - intent: The desired configuration to apply
func (d *Device) ApplyIntent(intent Intent) {
	fmt.Printf("  [%s] DRIFT DETECTED - Applying configuration...\n", d.Name)
	fmt.Printf("       Service: %s -> %s\n", d.State.Service, intent.Service)
	fmt.Printf("       Port:    %d -> %d\n", d.State.Port, intent.Port)
	fmt.Printf("       TLS:     %v -> %v\n", d.State.TLS, intent.TLS)

	// Apply the intent (simulate configuration push)
	d.State.Service = intent.Service
	d.State.Port = intent.Port
	d.State.TLS = intent.TLS

	fmt.Printf("  [%s] Configuration applied successfully\n", d.Name)
}

// enforcementLoop runs the continuous reconciliation loop.
// It periodically checks the device state and corrects any drift.
//
// Parameters:
//   - device: The device to manage
//   - intent: The desired configuration
//
// This function runs indefinitely, checking every 5 seconds.
func enforcementLoop(device *Device, intent Intent) {
	iteration := 0

	for {
		iteration++
		fmt.Printf("\n========== Enforcement Loop #%d ==========\n", iteration)

		// Get current operational state
		oper := device.GetOperState()

		// Compare against intent
		if isCompliant(intent, oper) {
			fmt.Printf("  [%s] COMPLIANT - No action needed\n", device.Name)
		} else {
			fmt.Printf("  [%s] NON-COMPLIANT - Drift detected\n", device.Name)
			device.ApplyIntent(intent)
		}

		// Display current state
		fmt.Printf("\n  Current State:\n")
		fmt.Printf("    Service: %s\n", device.State.Service)
		fmt.Printf("    Port:    %d\n", device.State.Port)
		fmt.Printf("    TLS:     %v\n", device.State.TLS)

		// Wait before next check
		fmt.Println("\n  Sleeping for 5 seconds...")
		time.Sleep(5 * time.Second)
	}
}

func main() {
	fmt.Println("===========================================")
	fmt.Println("  Closed-Loop Network Automation Demo")
	fmt.Println("===========================================")

	// Load configuration from YAML file
	data, err := os.ReadFile("input.yml")
	if err != nil {
		fmt.Printf("Error reading input.yml: %v\n", err)
		fmt.Println("\nPlease create input.yml with the following structure:")
		fmt.Println("  device:")
		fmt.Println("    hostname: router1")
		fmt.Println("    platform: ios-xr")
		fmt.Println("  intent:")
		fmt.Println("    service: gnmi")
		fmt.Println("    port: 57400")
		fmt.Println("    tls: true")
		os.Exit(1)
	}

	// Parse YAML configuration
	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		fmt.Printf("Error parsing YAML: %v\n", err)
		os.Exit(1)
	}

	// Display loaded configuration
	fmt.Println("\nConfiguration loaded successfully:")
	fmt.Printf("  Device:   %s (%s)\n", cfg.Device.Hostname, cfg.Device.Platform)
	fmt.Printf("  Intent:   Service=%s, Port=%d, TLS=%v\n",
		cfg.Intent.Service, cfg.Intent.Port, cfg.Intent.TLS)

	// Create device instance
	device := NewDevice(cfg.Device.Hostname)

	// Start the enforcement loop
	fmt.Println("\nStarting closed-loop automation...")
	enforcementLoop(device, cfg.Intent)
}
