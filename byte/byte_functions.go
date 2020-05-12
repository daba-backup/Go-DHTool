package byte

import (
	"bytes"
	"encoding/binary"
	"fmt"
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
	if len(b) != 4 {
		return 0.0, fmt.Errorf("Invalid slice length: %v", len(b))
	}

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
	if len(b) != 2 {
		return 0, fmt.Errorf("Invalid slice length: %v", len(b))
	}

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
	if len(b) != 2 {
		return 0, fmt.Errorf("Invalid slice length: %v", len(b))
	}

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

func GetFloat32ValueFromBin_LE(bin []byte, pos int) (float32, error) {
	var buffer [4]byte
	buffer[0] = bin[pos]
	buffer[1] = bin[pos+1]
	buffer[2] = bin[pos+2]
	buffer[3] = bin[pos+3]

	ret, err := DecodeToFloat32(buffer[:], true)
	if err != nil {
		return 0.0, err
	}

	return ret, nil
}
func GetInt16ValueFromBin_LE(bin []byte, pos int) (int16, error) {
	var buffer [2]byte
	buffer[0] = bin[pos]
	buffer[1] = bin[pos+1]

	ret, err := DecodeToInt16(buffer[:], true)
	if err != nil {
		return 0, err
	}

	return ret, nil
}
func GetUint16ValueFromBin_LE(bin []byte, pos int) (uint16, error) {
	var buffer [2]byte
	buffer[0] = bin[pos]
	buffer[1] = bin[pos+1]

	ret, err := DecodeToUint16(buffer[:], true)
	if err != nil {
		return 0, err
	}

	return ret, nil
}
