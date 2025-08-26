package utils

import (
	"fmt"
)

// From : https://github.com/Nv7-GitHub/googlesearch/blob/master/googlesearch/user_agents.py
func GetUserAgent() string {
	lynxVersion := fmt.Sprintf("Lynx/%d.%d.%d", Random(2, 3), Random(8, 9), Random(0, 2))
	libwwwVersion := fmt.Sprintf("libwww-FM/%d.%d", Random(2, 3), Random(13, 15))
	sslmmVersion := fmt.Sprintf("SSL-MM/%d.%d", Random(1, 2), Random(3, 5))
	opensslVersion := fmt.Sprintf("OpenSSL/%d.%d.%d", Random(1, 3), Random(0, 4), Random(0, 9))
	return fmt.Sprintf("%s %s %s %s", lynxVersion, libwwwVersion, sslmmVersion, opensslVersion)
}
