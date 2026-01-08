package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"text/template"
	"time"

	"golang.org/x/crypto/ssh"
	"gopkg.in/yaml.v2"
)

//
// -------- DATA MODEL (intent) --------
//

type Model struct {
	ASN      int    `yaml:"asn"`
	Loopback Addr   `yaml:"loopback"`
	Uplinks  []Link `yaml:"uplinks"`
	Peers    []Peer `yaml:"peers"`
}

type Addr struct {
	IP string `yaml:"ip"`
}

type Link struct {
	Name   string `yaml:"name"`   // e.g. ethernet-1/1
	Prefix string `yaml:"prefix"` // e.g. 192.0.2.0/31
}

type Peer struct {
	IP  string `yaml:"ip"`
	ASN int    `yaml:"asn"`
}

//
// -------- SR LINUX CLI TEMPLATE --------
//

const srlTemplate = `
enter candidate
{{- range .Uplinks }}
set / interface {{ .Name }} subinterface 0 ipv4 address {{ .Prefix }}
set / network-instance default interface {{ .Name }}.0
{{- end }}

{{- if .Loopback.IP }}
set / interface lo0 subinterface 0 ipv4 address {{ .Loopback.IP }}/32
set / network-instance default interface lo0.0
{{- end }}

{{- if .ASN }}
set /network-instance default protocols bgp autonomous-system {{ .ASN }}
{{- if .Loopback.IP }}
set /network-instance default protocols bgp router-id {{ .Loopback.IP }}
{{- end }}

set /network-instance default protocols bgp group EBGP peer-as {{ (index .Peers 0).ASN }}
set /network-instance default protocols bgp group EBGP ipv4-unicast admin-state enable

{{- range .Peers }}
set /network-instance default protocols bgp neighbor {{ .IP }} peer-group EBGP
{{- end }}

set /network-instance default protocols bgp ipv4-unicast admin-state enable
{{- end }}

commit now
quit
`

func loadInput(path string) (Model, error) {
	f, err := os.Open(path)
	if err != nil {
		return Model{}, err
	}
	defer f.Close()

	var m Model
	if err := yaml.NewDecoder(f).Decode(&m); err != nil {
		return Model{}, err
	}
	return m, nil
}

func generateConfig(m Model) (string, error) {
	var buf bytes.Buffer

	t, err := template.New("srl").Parse(srlTemplate)
	if err != nil {
		return "", err
	}
	if err := t.Execute(&buf, m); err != nil {
		return "", err
	}

	// CLI için sonuna newline iyi olur
	if buf.Len() == 0 || buf.Bytes()[buf.Len()-1] != '\n' {
		buf.WriteByte('\n')
	}

	return buf.String(), nil
}

func pushConfigSSH(host, user, pass, config string) error {
	sshCfg := &ssh.ClientConfig{
		User:            user,
		Auth:            []ssh.AuthMethod{ssh.Password(pass)},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         8 * time.Second,
	}

	client, err := ssh.Dial("tcp", host+":22", sshCfg)
	if err != nil {
		return fmt.Errorf("ssh dial failed: %w", err)
	}
	defer client.Close()

	session, err := client.NewSession()
	if err != nil {
		return fmt.Errorf("new session failed: %w", err)
	}
	defer session.Close()

	// PTY (SRL CLI interaktif)
	modes := ssh.TerminalModes{
		ssh.ECHO:          1,
		ssh.TTY_OP_ISPEED: 115200,
		ssh.TTY_OP_OSPEED: 115200,
	}
	if err := session.RequestPty("xterm", 40, 120, modes); err != nil {
		return fmt.Errorf("request pty failed: %w", err)
	}

	stdin, err := session.StdinPipe()
	if err != nil {
		return fmt.Errorf("stdin pipe failed: %w", err)
	}
	stdout, err := session.StdoutPipe()
	if err != nil {
		return fmt.Errorf("stdout pipe failed: %w", err)
	}
	stderr, err := session.StderrPipe()
	if err != nil {
		return fmt.Errorf("stderr pipe failed: %w", err)
	}

	if err := session.Shell(); err != nil {
		return fmt.Errorf("shell start failed: %w", err)
	}

	// Output oku
	var out bytes.Buffer
	done := make(chan struct{})
	go func() {
		_, _ = io.Copy(&out, stdout)
		close(done)
	}()
	go func() {
		_, _ = io.Copy(&out, stderr)
	}()

	log.Println("Connected. Sending SR Linux configuration...")

	// Komutları gönder
	if _, err := io.WriteString(stdin, config); err != nil {
		return fmt.Errorf("write config failed: %w", err)
	}

	// stdin kapatmak session'ın bitmesine yardım eder
	_ = stdin.Close()

	// Session kapanmasını bekle (quit sonrası)
	waitCh := make(chan error, 1)
	go func() { waitCh <- session.Wait() }()

	select {
	case err := <-waitCh:
		<-done
		fmt.Println("----- DEVICE OUTPUT -----")
		fmt.Println(out.String())
		if err != nil {
			return fmt.Errorf("session ended with error: %w", err)
		}
	case <-time.After(10 * time.Second):
		fmt.Println("----- DEVICE OUTPUT (timeout) -----")
		fmt.Println(out.String())
		return fmt.Errorf("timeout waiting for remote session to close")
	}

	return nil
}

func main() {
	model, err := loadInput("input.yml")
	if err != nil {
		log.Fatal(err)
	}

	config, err := generateConfig(model)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Generated SR Linux configuration:")
	fmt.Println(config)

	err = pushConfigSSH(
		"172.20.20.3",
		"admin",
		"admin",
		config,
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Configuration pushed successfully.")
}
