// autogenerated: do not edit!
// generated from gentemplate [gentemplate -d Package=mctree -id level -d VecType=level_vec -d Type=level github.com/platinasystems/go/elib/vec.tmpl]

package mctree

import (
	"github.com/platinasystems/go/elib"
)

type level_vec []level

func (p *level_vec) Resize(n uint) {
	c := elib.Index(cap(*p))
	l := elib.Index(len(*p)) + elib.Index(n)
	if l > c {
		c = elib.NextResizeCap(l)
		q := make([]level, l, c)
		copy(q, *p)
		*p = q
	}
	*p = (*p)[:l]
}

func (p *level_vec) validate(new_len uint, zero *level) *level {
	c := elib.Index(cap(*p))
	lʹ := elib.Index(len(*p))
	l := elib.Index(new_len)
	if l <= c {
		// Need to reslice to larger length?
		if l >= lʹ {
			*p = (*p)[:l]
		}
		return &(*p)[l-1]
	}
	return p.validateSlowPath(zero, c, l, lʹ)
}

func (p *level_vec) validateSlowPath(zero *level,
	c, l, lʹ elib.Index) *level {
	if l > c {
		cNext := elib.NextResizeCap(l)
		q := make([]level, cNext, cNext)
		copy(q, *p)
		if zero != nil {
			for i := c; i < cNext; i++ {
				q[i] = *zero
			}
		}
		*p = q[:l]
	}
	if l > lʹ {
		*p = (*p)[:l]
	}
	return &(*p)[l-1]
}

func (p *level_vec) Validate(i uint) *level {
	return p.validate(i+1, (*level)(nil))
}

func (p *level_vec) ValidateInit(i uint, zero level) *level {
	return p.validate(i+1, &zero)
}

func (p *level_vec) ValidateLen(l uint) (v *level) {
	if l > 0 {
		v = p.validate(l, (*level)(nil))
	}
	return
}

func (p *level_vec) ValidateLenInit(l uint, zero level) (v *level) {
	if l > 0 {
		v = p.validate(l, &zero)
	}
	return
}

func (p level_vec) Len() uint { return uint(len(p)) }