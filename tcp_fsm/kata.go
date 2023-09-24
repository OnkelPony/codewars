package kata


func TraverseTCPStates(events []string) string {
  transitions := map[string]map[string]string{
    "CLOSED": {"APP_PASSIVE_OPEN": "LISTEN", "APP_ACTIVE_OPEN": "SYN_SENT"},
    "LISTEN": {"RCVD_SYN": "SYN_RCVD", "APP_SEND": "SYN_SENT", "APP_CLOSE": "CLOSED"},
    "SYN_RCVD": {"APP_CLOSE": "FIN_WAIT_1", "RCV_ACK": "ESTABLISHED"},
    "SYN_SENT": {"RCV_SYN": "SYN_RCVD", "RCV_SYN_ACK": "ESTABLISHED", "APP_CLOSE": "CLOSED"},
    return "ERROR"
}
