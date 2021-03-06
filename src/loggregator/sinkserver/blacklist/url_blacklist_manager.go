package blacklist

import (
	"errors"
	"loggregator/iprange"
	"net/url"
)

type URLBlacklistManager struct {
	blacklistIPs    []iprange.IPRange
	blacklistedURLs []string
}

func New(blacklistIPs []iprange.IPRange) *URLBlacklistManager {
	return &URLBlacklistManager{blacklistIPs: blacklistIPs}
}

func (blacklistManager *URLBlacklistManager) CheckUrl(rawUrl string) (outputURL *url.URL, err error) {
	outputURL, err = url.Parse(rawUrl)
	if err != nil {
		return nil, err
	}

	ipNotBlacklisted, err := iprange.IpOutsideOfRanges(*outputURL, blacklistManager.blacklistIPs)
	if err != nil {
		return nil, err
	}
	if !ipNotBlacklisted {
		return nil, errors.New("Syslog Drain URL is blacklisted")
	}
	return outputURL, nil
}
