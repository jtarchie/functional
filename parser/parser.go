
//line parser/parser.rl:1
package parser

import "fmt"


//line parser/parser.go:9
var _functional_key_offsets []byte = []byte{
	0, 0, 4, 9, 10, 11, 13, 
}

var _functional_trans_keys []byte = []byte{
	65, 90, 97, 122, 32, 65, 90, 97, 
	122, 61, 32, 48, 57, 48, 57, 
}

var _functional_single_lengths []byte = []byte{
	0, 0, 1, 1, 1, 0, 0, 
}

var _functional_range_lengths []byte = []byte{
	0, 2, 2, 0, 0, 1, 1, 
}

var _functional_index_offsets []byte = []byte{
	0, 0, 3, 7, 9, 11, 13, 
}

var _functional_trans_targs []byte = []byte{
	2, 2, 0, 3, 2, 2, 0, 4, 
	0, 5, 0, 6, 0, 6, 0, 
}

const functional_start int = 1
const functional_first_final int = 6
const functional_error int = 0

const functional_en_main int = 1


//line parser/parser.rl:8


func Parse(data string) error {
  cs, p, pe := 0, 0, len(data)
  
  
//line parser/parser.go:50
	{
	cs = functional_start
	}

//line parser/parser.go:55
	{
	var _klen int
	var _trans int
	var _keys int
	if p == pe {
		goto _test_eof
	}
	if cs == 0 {
		goto _out
	}
_resume:
	_keys = int(_functional_key_offsets[cs])
	_trans = int(_functional_index_offsets[cs])

	_klen = int(_functional_single_lengths[cs])
	if _klen > 0 {
		_lower := int(_keys)
		var _mid int
		_upper := int(_keys + _klen - 1)
		for {
			if _upper < _lower {
				break
			}

			_mid = _lower + ((_upper - _lower) >> 1)
			switch {
			case data[p] < _functional_trans_keys[_mid]:
				_upper = _mid - 1
			case data[p] > _functional_trans_keys[_mid]:
				_lower = _mid + 1
			default:
				_trans += int(_mid - int(_keys))
				goto _match
			}
		}
		_keys += _klen
		_trans += _klen
	}

	_klen = int(_functional_range_lengths[cs])
	if _klen > 0 {
		_lower := int(_keys)
		var _mid int
		_upper := int(_keys + (_klen << 1) - 2)
		for {
			if _upper < _lower {
				break
			}

			_mid = _lower + (((_upper - _lower) >> 1) & ^1)
			switch {
			case data[p] < _functional_trans_keys[_mid]:
				_upper = _mid - 2
			case data[p] > _functional_trans_keys[_mid + 1]:
				_lower = _mid + 2
			default:
				_trans += int((_mid - int(_keys)) >> 1)
				goto _match
			}
		}
		_trans += _klen
	}

_match:
	cs = int(_functional_trans_targs[_trans])

	if cs == 0 {
		goto _out
	}
	p++
	if p != pe {
		goto _resume
	}
	_test_eof: {}
	_out: {}
	}

//line parser/parser.rl:23


  if cs < functional_first_final {
    return fmt.Errorf("could not parse to byte %d", p)
  }

  return nil
}