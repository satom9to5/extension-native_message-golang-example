package main

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"io"
)

var (
	byteOrder binary.ByteOrder = binary.LittleEndian
)

// メッセージ受け取り
func ReceiveMessage(reader io.Reader) ([]byte, error) {
	// 4byte分
	var length uint32

	if err := binary.Read(reader, byteOrder, &length); err != nil {
		return nil, err
	}

	if length == 0 {
		return nil, nil
	}

	message := make([]byte, length)

	// メッセージ読み込み
	n, err := reader.Read(message)
	if err != nil {
		return nil, err
	}
	if n != len(message) {
		return nil, errors.New("message length is different.")
	}

	return message, nil
}

// メッセージ送信
func SendMessage(message []byte, writer io.Writer) error {
	header := make([]byte, 4)

	byteOrder.PutUint32(header, (uint32)(len(message)))

	n, err := writer.Write(header)
	if err != nil {
		return err
	}

	if n != len(header) {
		return errors.New("header length is different.")
	}

	// メッセージ書き込み
	n, err = writer.Write(message)
	if err != nil {
		return err
	}

	if n != len(message) {
		return errors.New("header length is different.")
	}

	return nil
}

// JSONデータ受信
func Receive(data interface{}, reader io.Reader) error {
	message, err := ReceiveMessage(reader)

	if err != nil {
		return err
	}

	return json.Unmarshal(message, data)
}

// JSONデータ送信
func Send(data interface{}, writer io.Writer) error {
	jsonData, err := json.Marshal(data)

	if err != nil {
		return err
	}

	return SendMessage(jsonData, writer)
}
