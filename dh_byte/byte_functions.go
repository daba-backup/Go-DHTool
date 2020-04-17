package dh_byte

import (
	"bytes"
	"encoding/binary"
)

func Encode(v interface{}, le_flag bool) ([]byte, error) {
	buf := new(bytes.Buffer)
	var err error
	if le_flag == true {
		err = binary.Write(buf, binary.LittleEndian, v)
	} else {
		err = binary.Write(buf, binary.BigEndian, v)
	}

	if err != nil {
		return nil, err
	}

	ret := make([]byte, buf.Len())
	copy(ret, buf.Bytes())

	return ret, nil
}
func DecodeToFloat32(b []byte, le_flag bool) (float32, error) {
	buf := bytes.NewReader(b)
	var v float32

	var err error
	if le_flag == true {
		err = binary.Read(buf, binary.LittleEndian, &v)
	} else {
		err = binary.Read(buf, binary.BigEndian, &v)
	}

	if err != nil {
		return 0.0, err
	}

	return v, nil
}
func DecodeToInt16(b []byte, le_flag bool) (int16, error) {
	buf := bytes.NewReader(b)
	var v int16

	var err error
	if le_flag == true {
		err = binary.Read(buf, binary.LittleEndian, &v)
	} else {
		err = binary.Read(buf, binary.BigEndian, &v)
	}

	if err != nil {
		return 0.0, err
	}

	return v, nil
}
func DecodeToUint16(b []byte, le_flag bool) (uint16, error) {
	buf := bytes.NewReader(b)
	var v uint16

	var err error
	if le_flag == true {
		err = binary.Read(buf, binary.LittleEndian, &v)
	} else {
		err = binary.Read(buf, binary.BigEndian, &v)
	}

	if err != nil {
		return 0.0, err
	}

	return v, nil
}
