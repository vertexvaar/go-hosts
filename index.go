package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type IPv4 struct {
	blk1 int
	blk2 int
	blk3 int
	blk4 int
}

type HostsEntry struct  {
	ip IPv4
	domains []string
}


func main() {
	const hostsFile string = "/etc/hosts"

	file, err := os.Open(hostsFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	r := regexp.MustCompile("^(\\d{1,3})\\.(\\d{1,3})\\.(\\d{1,3})\\.(\\d{1,3})(?P<host>(?:\\s[\\w\\d\\.\\-]+)+)")

	// There is one caveat: Scanner does not deal well with lines longer than 65536 characters.
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		bytes := scanner.Text()
		submatch := r.FindAllStringSubmatch(bytes, -1)

		if (submatch == nil) {
			continue
		}

		fmt.Printf("\nEntry:  %#v\n", bytes)
		fmt.Printf(" - Matched: %#v\n", submatch)

		blk1, err := strconv.Atoi(submatch[0][1])
		check(err)
		blk2, err := strconv.Atoi(submatch[0][2])
		check(err)
		blk3, err := strconv.Atoi(submatch[0][3])
		check(err)
		blk4, err := strconv.Atoi(submatch[0][4])
		check(err)

		ip := IPv4{blk1, blk2, blk3, blk4}
		fmt.Printf("%#v\n", ip)
		fmt.Printf("%#v\n", formatIpv4(ip))

		domainList := submatch[0][5]
		domains := strings.Fields(domainList)
		entry := HostsEntry{ip, domains}
		fmt.Printf("%#v\n", entry)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func formatIpv4(ip IPv4) (formatted string) {
	return strconv.Itoa(ip.blk1) + "." + strconv.Itoa(ip.blk2) + "." + strconv.Itoa(ip.blk3) + "." + strconv.Itoa(ip.blk4)
}