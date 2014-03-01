package main

import (
	"bufio"
	"github.com/voxelbrain/goptions"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

const HOSTS_FILE = "/etc/hosts"

type CmdOptions struct {
	Dnsmasq       bool `goptions:"--dnsmasq, description='Send SIGHUP to dnsmasq'"`
	goptions.Help `goptions:"-h, --help, description='Show this help'"`
}

type Event struct {
	name    string
	address string
	role    string
}

func getEventParams() []Event {
	var result []Event
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		result = append(result, parseEvent(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		panic("Event reading error")
	}
	return result
}

func parseEvent(line string) Event {
	items := strings.Fields(line)
	switch len(items) {
	case 2:
		return Event{items[0], items[1], ""}
	case 3:
		return Event{items[0], items[1], items[2]}
	default:
		panic("Unsupported event format")
	}
}

func getEntries(filename string) []string {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic("File reading error")
	}
	entries := strings.Split(string(data), "\n")
	if n := len(entries) - 1; entries[n] == "" {
		return entries[:n]
	} else {
		return entries
	}
}

func removeEntry(entries []string, host string) []string {
	var result []string
	for _, v := range entries {
		if !strings.Contains(v, host) {
			result = append(result, v)
		}
	}
	return result
}

func sendSIGHUP(name string) {
	exec.Command("pkill", "-HUP", name).Run()
}

func main() {
	options := CmdOptions{}
	goptions.ParseAndFail(&options)

	event := os.Getenv("SERF_EVENT")
	if !(event == "member-join" || event == "member-leave" || event == "member-failed") {
		os.Exit(1)
	}

	events := getEventParams()
	entries := getEntries(HOSTS_FILE)

	for _, event := range events {
		entries = removeEntry(entries, event.address)
	}

	if event == "member-join" {
		for _, event := range events {
			new_entry := strings.Join([]string{event.address, event.name}, "\t")
			entries = append(entries, new_entry)
		}
	}

	data := strings.Join(entries, "\n")
	ioutil.WriteFile(HOSTS_FILE, []byte(data+"\n"), 0644)

	if options.Dnsmasq {
		sendSIGHUP("dnsmasq")
	}
}
