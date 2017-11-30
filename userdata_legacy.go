// +build !go1.9

package tox

import (
	"fmt"
	"runtime"

	"github.com/streamrail/concurrent-map"
)

/*
#include <tox/tox.h>
#include <tox/toxav.h>
*/
import "C"

type userData struct {
	ud0 map[*C.Tox]*Tox
	ud1 cmap.ConcurrentMap
	cc  bool // concurrent?
}

func newUserData() *userData {
	cc := true
	var ud0 map[*C.Tox]*Tox
	var ud1 cmap.ConcurrentMap

	if runtime.GOMAXPROCS(0) == 1 {
		cc = false
		ud0 = make(map[*C.Tox]*Tox, 0)
	} else {
		ud1 = cmap.New()
	}

	return &userData{ud0: ud0, ud1: ud1, cc: cc}
}

func (this *userData) set(ctox *C.Tox, gtox *Tox) {
	if this.cc {
		key := this.obj2Str(ctox)
		this.ud1.Set(key, gtox)
	} else {
		this.ud0[ctox] = gtox
	}
}

func (this *userData) get(ctox *C.Tox) *Tox {
	if this.cc {
		key := this.obj2Str(ctox)
		ival, ok := this.ud1.Get(key)
		if !ok {
			return nil
		}
		return ival.(*Tox)
	} else {
		if _, ok := this.ud0[ctox]; ok {
			return this.ud0[ctox]
		} else {
			return nil
		}
	}
}

func (this *userData) del(ctox *C.Tox) {
	if this.cc {
		key := this.obj2Str(ctox)
		this.ud1.Remove(key)
	} else {
		if _, ok := this.ud0[ctox]; ok {
			delete(this.ud0, ctox)
		}
	}
}

func (this *userData) obj2Str(ctox *C.Tox) string {
	return fmt.Sprintf("%p", ctox)
}

type userDataAV struct {
	ud0 map[*C.ToxAV]*ToxAV
	ud1 cmap.ConcurrentMap
	cc  bool // concurrent?
}

func newUserDataAV() *userDataAV {
	cc := true
	var ud0 map[*C.ToxAV]*ToxAV
	var ud1 cmap.ConcurrentMap

	if runtime.GOMAXPROCS(0) == 1 {
		cc = false
		ud0 = make(map[*C.ToxAV]*ToxAV, 0)
	} else {
		ud1 = cmap.New()
	}

	return &userDataAV{ud0: ud0, ud1: ud1, cc: cc}
}

func (this *userDataAV) set(ctox *C.ToxAV, gtox *ToxAV) {
	if this.cc {
		key := this.obj2Str(ctox)
		this.ud1.Set(key, gtox)
	} else {
		this.ud0[ctox] = gtox
	}
}

func (this *userDataAV) get(ctox *C.ToxAV) *ToxAV {
	if this.cc {
		key := this.obj2Str(ctox)
		ival, ok := this.ud1.Get(key)
		if !ok {
			return nil
		}
		return ival.(*ToxAV)
	} else {
		if _, ok := this.ud0[ctox]; ok {
			return this.ud0[ctox]
		} else {
			return nil
		}
	}
}

func (this *userDataAV) del(ctox *C.ToxAV) {
	if this.cc {
		key := this.obj2Str(ctox)
		this.ud1.Remove(key)
	} else {
		if _, ok := this.ud0[ctox]; ok {
			delete(this.ud0, ctox)
		}
	}
}

func (this *userDataAV) obj2Str(ctox *C.ToxAV) string {
	return fmt.Sprintf("%p", ctox)
}
