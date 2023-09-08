package kata

import (
	"strconv"
	"strings"
)

func IpsBetween(start, end string) int {
    startIPs := makeOctets(start)
    endIPs := makeOctets(end)
    return (endIPs[0] - startIPs[0]) << 24 + (endIPs[1] - startIPs[1]) << 16 + (endIPs[2] - startIPs[2]) << 8 + (endIPs[3] - startIPs[3])

}

func makeOctets(start string) [4]int {
    var result [4]int
    for i, octet := range strings.Split(start, ".") {
        result[i], _ = strconv.Atoi(octet)
    }
    return result
}