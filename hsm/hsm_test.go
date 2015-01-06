package hsm

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"testing"
)

type Test_Struct struct {
	Key string
}

type TestStruct2 struct {
	StringField         string
	DecimalNumberField1 uint
	DecimalNumberField2 uint
	HexNumberField      uint
	RawData             []byte
}




func Test_read_key(t *testing.T) {

	var key_holder Test_Struct
	buf := bytes.NewBuffer([]byte("e045fe3ea2a2f47d007a3030"))
	read_key(buf, &key_holder.Key)

	if key_holder.Key != "e045fe3ea2a2f47d" {
		t.Fail()
	}

}

func Test_read_key2(t *testing.T) {

	var key_holder Test_Struct
	buf := bytes.NewBuffer([]byte("Ue045fe3ea2a2f47d007afe3ea2a2f47d007a3030"))
	read_key(buf, &key_holder.Key)
	if key_holder.Key != "Ue045fe3ea2a2f47d007afe3ea2a2f47d" {
		t.Fail()
	}

}

func Test_read_key3(t *testing.T) {

	var key_holder Test_Struct
	buf := bytes.NewBuffer([]byte("Te045fe3ea2a2f47d007afe3ea2a2f47d007a30307ee5eae300"))
	read_key(buf, &key_holder.Key)

	if key_holder.Key != "Te045fe3ea2a2f47d007afe3ea2a2f47d007a30307ee5eae3" {
		t.Fail()
	}

}

func Test_read_key4(t *testing.T) {

	var key_holder Test_Struct
	buf := bytes.NewBuffer([]byte("Ke045fe3ea2a2f47d007afe3ea2a2f47d007a30307ee5eae300"))

	result := read_key(buf, &key_holder.Key)

	if !result {
		t.Fail()
	}

}

func TestReadFieldsFromStruct(t *testing.T) {

	var fields_struct TestStruct2
	data,_:=hex.DecodeString("303030303030303030303032323334303041310001020304")
	buf := bytes.NewBuffer(data)

	result := read_fixed_field(buf, &fields_struct.StringField, 12, String)
	if !result {
		t.Fail()
	}
	result = read_fixed_field(buf, &fields_struct.DecimalNumberField1, 1, DecimalInt)
	if !result {
		t.Fail()
	}
	result = read_fixed_field(buf, &fields_struct.DecimalNumberField2, 2, DecimalInt)
	if !result {
		t.Fail()
	}
	result = read_fixed_field(buf, &fields_struct.HexNumberField, 4, HexadecimalInt)
	if !result {
		t.Fail()
	}

	result = read_fixed_field(buf, &fields_struct.RawData, 5, Binary)
	if !result {
		t.Fail()
	}

	fmt.Println(Dump(fields_struct))

	if (fields_struct.StringField == "000000000002" && 
		fields_struct.DecimalNumberField1 == 2 && 
		fields_struct.DecimalNumberField2 == 34 && 
		fields_struct.HexNumberField == 161) {
		
		if hex.EncodeToString(fields_struct.RawData) != "0001020304" {
			t.Fail()
			
		}
	} else {
		t.Fail()
	}

}
