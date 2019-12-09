package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"io/ioutil"
	"os"

	"github.com/BurntSushi/toml"
)

// Config ... Store data
type Config struct {
	FileData []*Data
}

// Data ... Store for parameter store data.
type Data struct {
	Name        string `toml:"name"`
	Value       string `toml:"value"`
	Type        string `toml:"type"`
	Description string `toml:"description"`
	KmsAlias    string `toml:"kmsAlias"`
}

// DeleteData ... Store delete data for pamrameter store.
type DeleteData struct {
	Name string
}

// Load ... Load toml file
func (c *Config) Load(pstoreKey string) error {
	output, err := DecryptFile(pstoreKey)
	if err != nil {
		return err
	}

	if _, err := toml.Decode(string(output), c); err != nil {
		return err
	}

	return nil

}

// WriteFile ... Config struct encode to toml data and Write to file
func (c *Config) WriteFile() error {
	f, err := os.Create(FilePath)
	if err != nil {
		return err
	}
	defer f.Close()

	buf := new(bytes.Buffer)
	if err := toml.NewEncoder(buf).Encode(*c); err != nil {
		return err
	}

	if _, err := f.Write(buf.Bytes()); err != nil {
		return err
	}

	return nil
}

// DecryptFileAndWriteFileWithAddData ... Encrypt data and Write file.
func (d *Data) DecryptFileAndWriteFileWithAddData(pstoreKey string) error {
	config := &Config{}

	if IsFileNotEmpty(FilePath) {
		c := &Config{}
		if err := c.Load(pstoreKey); err != nil {
			return err
		}

		var data []*Data
		var overwrite bool
		for _, v := range c.FileData {
			if v.Name == d.Name {
				data = append(data, &Data{Name: d.Name, Value: d.Value, Type: d.Type, Description: d.Description, KmsAlias: d.KmsAlias})
				overwrite = true
				continue
			}
			data = append(data, v)
		}
		if !overwrite {
			data = append(data, &Data{Name: d.Name, Value: d.Value, Type: d.Type, Description: d.Description, KmsAlias: d.KmsAlias})
		}

		config = &Config{data}
	} else {
		config = &Config{[]*Data{d}}
	}

	if err := config.WriteFile(); err != nil {
		return err
	}

	return nil
}

// DecryptFileAndWriteFileWithDeleteData ... Encrypt data and Write file.
func (d *Data) DecryptFileAndWriteFileWithDeleteData(pstoreKey, deleteName string) error {
	config := &Config{}

	if IsFileNotEmpty(FilePath) {
		c := &Config{}
		if err := c.Load(pstoreKey); err != nil {
			return err
		}

		var data []*Data
		for _, v := range c.FileData {
			if v.Name == d.Name {
				continue
			}
			data = append(data, v)
		}

		config = &Config{data}
	} else {
		config = &Config{[]*Data{d}}
	}

	if err := config.WriteFile(); err != nil {
		return err
	}

	return nil
}


// EncryptFile ... Encrypt file
func EncryptFile(pstoreKey string) error {
	input, err := ioutil.ReadFile(FilePath)
	if err != nil {
		return err
	}

	block, err := aes.NewCipher([]byte(pstoreKey))

	output := make([]byte, aes.BlockSize+len(input))

	iv := output[:aes.BlockSize]
	res := output[aes.BlockSize:]
	_, err = io.ReadFull(rand.Reader, iv)
	if err != nil {
		return err
	}
	stream := cipher.NewCTR(block, iv)
	stream.XORKeyStream(res, input)

	if err := ioutil.WriteFile(FilePath, output, 0644); err != nil {
		return err
	}

	return nil
}

// DecryptFile ... Decrypt file
func DecryptFile(pstoreKey string) ([]byte, error) {
	input, err := ioutil.ReadFile(FilePath)
	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher([]byte(pstoreKey))

	output := make([]byte, len(input[aes.BlockSize:]))
	iv := input[:aes.BlockSize]
	stream := cipher.NewCTR(block, iv)
	stream.XORKeyStream(output, input[aes.BlockSize:])

	return output, nil
}
