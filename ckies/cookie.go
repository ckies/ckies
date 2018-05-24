package ckies

// CookieType is one of the four types
type CookieType string

const (
	// CookieTypeNecessary is
	CookieTypeNecessary CookieType = "necessary"
	// CookieTypeFunctional is
	CookieTypeFunctional CookieType = "functional"
	// CookieTypePerformance is
	CookieTypePerformance CookieType = "performance"
	// CookieTypeMarketing is
	CookieTypeMarketing CookieType = "marketing"
)

// Cookie contains information about a single cookie
type Cookie struct {
	Name    string
	Info    string
	Type    CookieType
	Expires string
}

// ContainsCookies adds simple filtering for array of Cookies
type ContainsCookies struct {
	Cookies []Cookie
}

func (s *ContainsCookies) filterCookiesByType(filterFor CookieType) []Cookie {
	var data []Cookie

	for _, c := range s.Cookies {
		if c.Type == filterFor {
			data = append(data, c)
		}
	}

	return data
}

// Necessary filters list of cookies for necessary cookies
func (s *ContainsCookies) Necessary() []Cookie {
	return s.filterCookiesByType(CookieTypeNecessary)
}

// Functional filters list of cookies for necessary cookies
func (s *ContainsCookies) Functional() []Cookie {
	return s.filterCookiesByType(CookieTypeFunctional)
}

// Performance filters list of cookies for necessary cookies
func (s *ContainsCookies) Performance() []Cookie {
	return s.filterCookiesByType(CookieTypePerformance)
}

// Marketing filters list of cookies for necessary cookies
func (s *ContainsCookies) Marketing() []Cookie {
	return s.filterCookiesByType(CookieTypeMarketing)
}
