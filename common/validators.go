package common

import (
	"reflect"
	"regexp"

	"github.com/ethereum/go-ethereum/common"
)

func IsValidAddress(iaddress interface{}) bool {
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
	switch v := iaddress.(type) {
	case string:
		return re.MatchString(v)
	case common.Address:
		return re.MatchString(v.Hex())
	default:
		return false
	}
}

func IsZeroAddress(iaddress interface{}) bool {
	var address common.Address
	switch v := iaddress.(type) {
	case string:
		address = common.HexToAddress(v)
	case common.Address:
		address = v
	default:
		return false
	}

	zeroAddressBytes := common.FromHex("0x0000000000000000000000000000000000000000")
	addressBytes := address.Bytes()
	return reflect.DeepEqual(addressBytes, zeroAddressBytes)
}

func ValidateAddress(address interface{}) bool {
	return IsValidAddress(address) && !IsZeroAddress(address)
}

func ValidateAddresses(addresses []interface{}) (
	[]interface{},
	[]interface{},
	bool,
) {
	if len(addresses) == 0 {
		return nil, nil, false
	}

	var validAddresses []interface{}
	var invalidAddresses []interface{}
	for _, address := range addresses {

		if ValidateAddress(address) {
			validAddresses = append(validAddresses, address)
		} else {
			invalidAddresses = append(invalidAddresses, address)
		}
	}

	success := len(invalidAddresses) == 0
	return validAddresses, invalidAddresses, success
}
