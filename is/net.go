package is

import (
	"net/netip"
	"net/url"
)

func IP(s string) bool {
	_, err := netip.ParseAddr(s)
	return err == nil
}

func IP4(s string) bool {
	addr, err := netip.ParseAddr(s)
	return err == nil && addr.Is4()
}

func IP6(s string) bool {
	addr, err := netip.ParseAddr(s)
	return err == nil && addr.Is6()
}

func URL(s string) bool {
	_, err := url.Parse(s)
	return err == nil
}

func URLWithScheme(s string, scheme string) bool {
	u, err := url.Parse(s)
	return err == nil && u.Scheme == scheme
}

func RequestURI(s string) bool {
	_, err := url.ParseRequestURI(s)
	return err == nil
}

func URLQuery(s string) bool {
	_, err := url.ParseQuery(s)
	return err == nil
}
