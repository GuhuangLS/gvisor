// automatically generated by stateify.

package uevent

import (
	"gvisor.dev/gvisor/pkg/state"
)

func (x *Protocol) beforeSave() {}
func (x *Protocol) save(m state.Map) {
	x.beforeSave()
}

func (x *Protocol) afterLoad() {}
func (x *Protocol) load(m state.Map) {
}

func init() {
	state.Register("uevent.Protocol", (*Protocol)(nil), state.Fns{Save: (*Protocol).save, Load: (*Protocol).load})
}
