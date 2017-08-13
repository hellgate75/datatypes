package utils

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func ExistsFile(file string) bool {
	_, err := os.Stat(file)
	return err == nil
}

func DeleteIfExists(file string) error {
	_, err := os.Stat(file)
	if err == nil {
		return os.Remove(file)
	}
	return err
}

func MakeFolderIfNotExists(folder string) error {
	if _, err := os.Stat(folder); err != nil {
		err := os.MkdirAll(folder, 0777)
		return err
	}
	return nil
}

func EncodeBytes(decodedByteArray []byte) []byte {
	newBytes := make([]byte, 0)
	for _, byteElem := range decodedByteArray {
		newBytes = append(newBytes, byteElem-20)
	}
	return newBytes
	//return decodedByteArray
}

func DecodeBytes(encodedByteArray []byte) []byte {
	newBytes := make([]byte, 0)
	for _, byteElem := range encodedByteArray {
		newBytes = append(newBytes, byteElem+20)
	}
	return newBytes
	//return encodedByteArray
}

func CorrectInput(input string) string {
	return strings.TrimSpace(strings.ToLower(input))
}

func StringToInt(s string) (int, error) {
	return strconv.Atoi(s)
}

func IntToString(n int) string {
	return strconv.Itoa(n)
}

// go binary decoder
func GetJSONFromObj(m interface{}, prettify bool) []byte {
	if prettify {
		bytes, err := json.MarshalIndent(m, "", "  ")
		if err != nil {
			return []byte{}
		}
		return bytes
	}
	bytes, err := json.Marshal(m)
	if err != nil {
		return []byte{}
	}
	return bytes
}

func GetYAMLFromObj(m interface{}) []byte {
	bytes, err := yaml.Marshal(m)
	if err != nil {
		return []byte{}
	}
	return bytes
}

func GetXMLFromObj(m interface{}, prettify bool) []byte {
	if prettify {
		bytes, err := xml.MarshalIndent(m, "", "  ")
		if err != nil {
			return []byte{}
		}
		return bytes
	}
	bytes, err := xml.Marshal(m)
	if err != nil {
		return []byte{}
	}
	return bytes
}

func ExportStructureToFile(File string, Format string, structure interface{}) error {
	var bytesArray []byte = make([]byte, 0)
	var err error
	if CorrectInput(Format) == "json" {
		bytesArray, err = GetJSONFromElem(structure, true)
		if err != nil {
			return err
		}
	} else if CorrectInput(Format) == "xml" {
		bytesArray, err = GetXMLFromElem(structure, true)
		if err != nil {
			return err
		}
	} else if CorrectInput(Format) == "yaml" {
		bytesArray, err = GetYAMLFromElem(structure)
		if err != nil {
			return err
		}
	}  else {
		return errors.New("File Format '" + Format + "' not available")
	}
	return ioutil.WriteFile(File, bytesArray, 0777)
}

// go binary decoder
func GetJSONFromElem(m interface{}, prettify bool) ([]byte, error) {
	if prettify {
		return json.MarshalIndent(m, "", "  ")
	}
	return json.Marshal(m)
}

func GetXMLFromElem(m interface{}, prettify bool) ([]byte, error) {
	if prettify {
		return xml.MarshalIndent(m, "", "  ")
	}
	return xml.Marshal(m)
}

func GetYAMLFromElem(m interface{}) ([]byte, error) {
	return yaml.Marshal(m)
}
