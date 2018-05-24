package ckies

// Merger uses Config and []Services to build a list of cookies
type Merger struct {
	Config   Config
	Services map[string]Service
}

func (m *Merger) filterByType(filterFor CookieType) []Cookie {
	data := m.Config.filterCookiesByType(filterFor)

	for _, service := range m.Services {
		data = append(data, service.filterCookiesByType(filterFor)...)
	}

	return data
}
