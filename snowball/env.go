package snowball

import (
	"errors"
	"fmt"
	"net"
	"os"
	"strconv"
)

var ErrEmptyEnvVar = errors.New("getenv: Specified environment variable is empty or undefined")

func GetenvStr(key string) (string, error) {
	result := os.Getenv(key)
	if result == "" {
		return result, ErrEmptyEnvVar
	}

	return result, nil
}

func GetenvInteger(key string) (uint64, error) {
	str, err := GetenvStr(key)
	if err != nil {
		return 0, err
	}

	result, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		return 0, err
	}

	return result, nil
}

func GetenvBoolean(key string) (bool, error) {
	str, err := GetenvStr(key)
	if err != nil {
		return false, err
	}

	result, err := strconv.ParseBool(str)
	if err != nil {
		return false, err
	}

	return result, nil
}

// The epoch used by Snowball is entirely configurable via the SNOWBALL_EPOCH_MS system environment
// variable. This makes using Snowball in containerized deployments straightforward - just supply
// the expected environment variable and Snowball will take it from there.
//
// If the epoch is not provided, Snowball defaults to the epoch used by the original Snowflake algorithm.
func GetEpoch() uint64 {
	var epoch uint64
	var err error
	if epoch, err = GetenvInteger("SNOWBALL_EPOCH_MS"); err != nil {
		fmt.Println("get epoch failed: envvar SNOWBALL_EPOCH_MS not found, using default")
		return 1288834974657
	}

	return epoch
}

// The server ID used by Snowball is entirely configurable via the SNOWBALL_NODE_ID system environment
// variable. This makes using Snowball in containerized deployments straightforward - just supply the
// expected environment variable and Snowball will take it from there.
//
// If the ID is not provided, Snowball defaults to 0.
func GetServerId() uint64 {
	var serverId uint64
	var err error
	if serverId, err = GetenvInteger("SNOWBALL_NODE_ID"); err != nil {
		fmt.Println("get epoch failed: envvar SNOWBALL_NODE_ID not found, using default")
		return 0
	}

	return serverId
}

// Alternatively, the server ID used by Snowball can be based on the machine or pod IP address, via the
// SERVER_IP_ADDRESS system environment variable. This can be leveraged for containerized deployments that
// are expected to scale, and providing multiple unique server IDs is more difficult.
func GetServerIdFromIPAddress() uint64 {
	ipStr, err := GetenvStr("SERVER_IP_ADDRESS")
	if err != nil || len(ipStr) == 0 {
		fmt.Println("get server id failed: env var SERVER_IP_ADDRESS not found or undefined, using default")
		return 0
	}

	ip := net.ParseIP(ipStr)
	return uint64(ip[3])
}
