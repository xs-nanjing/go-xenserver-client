package client

import (
	"github.com/nilshell/xmlrpc"
)

type Host XenAPIObject

func (self *Host) CallPlugin(plugin, method string, params map[string]string) (response string, err error) {
	result := APIResult{}
	params_rec := make(xmlrpc.Struct)
	for key, value := range params {
		params_rec[key] = value
	}
	err = self.Client.APICall(&result, "host.call_plugin", self.Ref, plugin, method, params_rec)
	if err != nil {
		return "", err
	}
	response = result.Value.(string)
	return
}
