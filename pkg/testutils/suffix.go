package testutils

type PublicSuffixList struct {}

func (l *PublicSuffixList) PublicSuffix(domain string) string {
	return ""
}

func (l *PublicSuffixList) String() string {
	return ""
}
