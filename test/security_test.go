package test

import (
	"encoding/json"
	"errors"
	"fmt"
	"go01/security"
	"go01/util"
	"testing"
)

func TestMd5(t *testing.T) {
	plain := "hello"
	encodeString := security.Md5(plain)
	fmt.Println(plain, ",md5:", encodeString, ",len:", len(encodeString))
	t.FailNow()
}

type Person struct {
	B      bool
	Int1   int
	Int2   int8
	Int3   int16
	Int4   int32
	Int5   int64
	Uint1  uint
	Uint2  uint8
	Uint3  uint16
	Uint4  uint32
	Uint5  uint64
	Float1 float32
	Float2 float64
	//Float3 complex64
	//Float4 complex128
	Array2  [16]int
	Array1  [16]int
	Map2    map[string]string
	Ptr1    *Person
	Slice1  []Person
	Slice2  []int
	String1 string
	Struct1 *Person
	Struct2 Tag
}

type Tag struct {
	Id   int
	Name string
}

func TestGenString(t *testing.T) {
	a := Person{Map2: map[string]string{"Array2": "xx", "Array1": "xx", "B": "xx"}}
	bs, _ := json.Marshal(a)
	fmt.Println(string(bs))

}

func TestFloat(t *testing.T) {
	a := 1.123456789455666555885245524524522
	fmt.Println(a)
}

func TestRsa(t *testing.T) {
	privateKey, publicKey, err := security.GenRsaKeys(2048)
	if err != nil {
		panic(err)
	}
	fmt.Println("private key:", privateKey, "public key:", publicKey)
}

func TestError(t *testing.T) {
	//var pathError *fs.PathError
	//var err error = &fs.PathError{Path: "/"}
	//if errors.As(err, &pathError) {
	//	fmt.Println("Failed at path:", pathError.Path)
	//} else {
	//	fmt.Println(err)
	//}
	var err error = &util.ServiceError{Message: "xxx"}
	var target *util.ServiceError
	if errors.As(err, &target) {
		fmt.Println("yes")
	} else {
		fmt.Println(err)
	}
}
