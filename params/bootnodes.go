// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package params

import "github.com/ethereum/go-ethereum/common"

// MainnetBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the main Tango network.
var MainnetBootnodes = []string{
	// "enode://d16cf20214ae987c0b200252df331633a680866e75fec511ae25573c06585a8ce2d7379526c9ca186501fbcf1af92443a178f7a800049947a178d5a7ba130566@13.230.226.55:30303",  // kcc-mainnet-node-boot-01
}

// TestnetBootnodes are the enode URLs of the P2P bootstrap nodes running on the test network.
var TestnetBootnodes = []string{
}

var V5Bootnodes []string

// KnownDNSNetwork returns the address of a public DNS-based node list for the given
// genesis hash and protocol. See https://github.com/ethereum/discv4-dns-lists for more
// information.
func KnownDNSNetwork(genesis common.Hash, protocol string) string {
	return ""
}
