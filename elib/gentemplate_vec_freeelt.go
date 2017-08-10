// autogenerated: do not edit!
// generated from gentemplate [gentemplate -d Package=elib -id freeElt -d VecType=freeEltVec -d Type=freeElt vec.tmpl]

// Copyright 2016 Platina Systems, Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package elib

type freeEltVec []freeElt

func (p *freeEltVec) Resize(n uint) {
	c := uint(cap(*p))
	l := uint(len(*p)) + n
	if l > c {
		c = NextResizeCap(l)
		q := make([]freeElt, l, c)
		copy(q, *p)
		*p = q
	}
	*p = (*p)[:l]
}

func (p *freeEltVec) validate(new_len uint, zero freeElt) *freeElt {
	c := uint(cap(*p))
	lʹ := uint(len(*p))
	l := new_len
	if l <= c {
		// Need to reslice to larger length?
		if l > lʹ {
			*p = (*p)[:l]
			for i := lʹ; i < l; i++ {
				(*p)[i] = zero
			}
		}
		return &(*p)[l-1]
	}
	return p.validateSlowPath(zero, c, l, lʹ)
}

func (p *freeEltVec) validateSlowPath(zero freeElt, c, l, lʹ uint) *freeElt {
	if l > c {
		cNext := NextResizeCap(l)
		q := make([]freeElt, cNext, cNext)
		copy(q, *p)
		for i := c; i < cNext; i++ {
			q[i] = zero
		}
		*p = q[:l]
	}
	if l > lʹ {
		*p = (*p)[:l]
	}
	return &(*p)[l-1]
}

func (p *freeEltVec) Validate(i uint) *freeElt {
	var zero freeElt
	return p.validate(i+1, zero)
}

func (p *freeEltVec) ValidateInit(i uint, zero freeElt) *freeElt {
	return p.validate(i+1, zero)
}

func (p *freeEltVec) ValidateLen(l uint) (v *freeElt) {
	if l > 0 {
		var zero freeElt
		v = p.validate(l, zero)
	}
	return
}

func (p *freeEltVec) ValidateLenInit(l uint, zero freeElt) (v *freeElt) {
	if l > 0 {
		v = p.validate(l, zero)
	}
	return
}

func (p *freeEltVec) ResetLen() {
	if *p != nil {
		*p = (*p)[:0]
	}
}

func (p freeEltVec) Len() uint { return uint(len(p)) }
