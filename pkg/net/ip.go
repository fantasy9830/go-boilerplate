package net

import (
	"net"
	"net/http"
	"strings"
)

// IsPrivate reports whether ip is a private address
func IsPrivateIP(ip string) bool {
	p := net.ParseIP(ip)

	return p.IsLoopback() || p.IsPrivate()
}

// ClientIP 取得 Client IP
func ClientIP(r *http.Request) string {
	xForwardedFor := r.Header.Get("X-Forwarded-For")
	ip := strings.TrimSpace(strings.Split(xForwardedFor, ",")[0])
	if ip != "" {
		return ip
	}

	ip = strings.TrimSpace(r.Header.Get("X-Real-Ip"))
	if ip != "" {
		return ip
	}

	if ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr)); err == nil {
		return ip
	}

	return ""
}

// ClientPublicIP 取得 Client IP
func ClientPublicIP(r *http.Request) string {
	var ip string
	for _, ip = range strings.Split(r.Header.Get("X-Forwarded-For"), ",") {
		ip = strings.TrimSpace(ip)
		if ip != "" && !IsPrivateIP(ip) {
			return ip
		}
	}

	ip = strings.TrimSpace(r.Header.Get("X-Real-Ip"))
	if ip != "" && !IsPrivateIP(ip) {
		return ip
	}

	if ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr)); err == nil {
		if !IsPrivateIP(ip) {
			return ip
		}
	}

	return ""
}

// GetIP GetIP
func GetIP(r *http.Request) string {
	if ip := ClientPublicIP(r); ip != "" {
		return ip
	}

	return ClientIP(r)
}

// MatchCIDR MatchCIDR
func MatchCIDR(addr string, cidr string) bool {
	ip := net.ParseIP(addr)
	_, ipnet, err := net.ParseCIDR(cidr)
	if err != nil {
		return false
	}

	return ipnet.Contains(ip)
}
