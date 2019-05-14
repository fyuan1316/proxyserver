package option

import "net/url"

// Configuration Configuration for reserveproxy server
type Configuration struct {
	Addr                    string
	ModelName               string
	TargetURL               string //*url.URL
	ProxyPrefixURLCondition string
}

// SetAddr set Address
func (conf *Configuration) SetAddr(addr string) {
	conf.Addr = addr
}

//ModelName setModelName
func ModelName(name string) func(*Configuration) error {
	return func(conf *Configuration) error {
		conf.ModelName = name
		return nil
	}
}

//TargetURL setTargetURL
func TargetURL(target string) func(*Configuration) error {
	return func(conf *Configuration) error {
		_, err := url.Parse(target)
		if err != nil {
			return err
		}
		conf.TargetURL = target
		return nil
	}
}

//ProxyPrefixURL setProxyPrefixURL
func ProxyPrefixURL(prefix string) func(*Configuration) error {
	return func(conf *Configuration) error {
		conf.ProxyPrefixURLCondition = prefix
		return nil
	}
}
