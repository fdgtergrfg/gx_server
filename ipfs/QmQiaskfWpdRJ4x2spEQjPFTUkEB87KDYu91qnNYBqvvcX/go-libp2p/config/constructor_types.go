package config

import (
	"fmt"
	"reflect"

	crypto "gx/ipfs/QmPvyPwuCgJ7pDmrKDxRtsScJgBaM5h4EpRL2qQJsmXf4n/go-libp2p-crypto"
	peer "gx/ipfs/QmQsErDt8Qgw1XrsXf2BpEzDgGWtB1YLsTAARBup5b6B9W/go-libp2p-peer"
	host "gx/ipfs/QmRRCrNRs4qxotXx7WJT6SpCvSNEhXvyBcVjXY2K71pcjE/go-libp2p-host"
	filter "gx/ipfs/QmSW4uNHbvQia8iZDXzbwjiyHQtnyo9aFqfQAMasj3TJ6Y/go-maddr-filter"
	pnet "gx/ipfs/QmW7Ump7YyBMr712Ta3iEVh3ZYcfVvJaPryfbCnyE826b4/go-libp2p-interface-pnet"
	inet "gx/ipfs/QmX5J1q63BrrDTbpcHifrFbxH3cMZsvaNajy6u3zCpzBXs/go-libp2p-net"
	mux "gx/ipfs/QmY9JXR3FupnYAYJWK9aMr9bCpqWKcToQ1tz8DVGTrHpHw/go-stream-muxer"
	transport "gx/ipfs/QmYr9RHifaqHTFZdAsUPLmiMAi2oNeEqA48AFKxXJAsLpJ/go-libp2p-transport"
	security "gx/ipfs/QmcGgFLHMFLcNMMvxsBC5LeLqubLR5djxjShVU3koVMtVq/go-conn-security"
	pstore "gx/ipfs/QmeKD8YT7887Xu6Z86iZmpYNxrLogJexqxEugSmaf14k64/go-libp2p-peerstore"
	tptu "gx/ipfs/QmfNvpHX396fhMeauERV6eFnSJg78rUjhjpFf1JvbjxaYM/go-libp2p-transport-upgrader"
)

var (
	// interfaces
	hostType      = reflect.TypeOf((*host.Host)(nil)).Elem()
	networkType   = reflect.TypeOf((*inet.Network)(nil)).Elem()
	transportType = reflect.TypeOf((*transport.Transport)(nil)).Elem()
	muxType       = reflect.TypeOf((*mux.Transport)(nil)).Elem()
	securityType  = reflect.TypeOf((*security.Transport)(nil)).Elem()
	protectorType = reflect.TypeOf((*pnet.Protector)(nil)).Elem()
	privKeyType   = reflect.TypeOf((*crypto.PrivKey)(nil)).Elem()
	pubKeyType    = reflect.TypeOf((*crypto.PubKey)(nil)).Elem()
	pstoreType    = reflect.TypeOf((*pstore.Peerstore)(nil)).Elem()

	// concrete types
	peerIDType   = reflect.TypeOf((peer.ID)(""))
	filtersType  = reflect.TypeOf((*filter.Filters)(nil))
	upgraderType = reflect.TypeOf((*tptu.Upgrader)(nil))
)

var argTypes = map[reflect.Type]constructor{
	upgraderType:  func(h host.Host, u *tptu.Upgrader) interface{} { return u },
	hostType:      func(h host.Host, u *tptu.Upgrader) interface{} { return h },
	networkType:   func(h host.Host, u *tptu.Upgrader) interface{} { return h.Network() },
	muxType:       func(h host.Host, u *tptu.Upgrader) interface{} { return u.Muxer },
	securityType:  func(h host.Host, u *tptu.Upgrader) interface{} { return u.Secure },
	protectorType: func(h host.Host, u *tptu.Upgrader) interface{} { return u.Protector },
	filtersType:   func(h host.Host, u *tptu.Upgrader) interface{} { return u.Filters },
	peerIDType:    func(h host.Host, u *tptu.Upgrader) interface{} { return h.ID() },
	privKeyType:   func(h host.Host, u *tptu.Upgrader) interface{} { return h.Peerstore().PrivKey(h.ID()) },
	pubKeyType:    func(h host.Host, u *tptu.Upgrader) interface{} { return h.Peerstore().PubKey(h.ID()) },
	pstoreType:    func(h host.Host, u *tptu.Upgrader) interface{} { return h.Peerstore() },
}

func newArgTypeSet(types ...reflect.Type) map[reflect.Type]constructor {
	result := make(map[reflect.Type]constructor, len(types))
	for _, ty := range types {
		c, ok := argTypes[ty]
		if !ok {
			panic(fmt.Sprintf("missing constructor for type %s", ty))
		}
		result[ty] = c
	}
	return result
}
