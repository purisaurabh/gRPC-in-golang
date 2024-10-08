package serializer

import (
	"fmt"
	"os"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func WriteProtobufToJSONFile(message proto.Message, filename string) error {
	data, err := protojson.Marshal(message)
	if err != nil {
		return fmt.Errorf("cannot marshal proto message to JSON: %w", err)
	}

	err = os.WriteFile(filename, []byte(data), 0644)
	if err != nil {
		return fmt.Errorf("cannot write JSON data to file: %w", err)
	}

	return nil
}

func WriteProtobufToBinaryFile(message proto.Message, fileName string) error {
	data, err := proto.Marshal(message)
	if err != nil {
		return fmt.Errorf("could not marshal proto message into binary : %w", err)
	}

	err = os.WriteFile(fileName, data, 0644)
	if err != nil {
		return fmt.Errorf("could not write binary data to file : %w", err)
	}
	return nil
}

func ReadProtobufFromBinaryFile(filename string, message proto.Message) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("cannot read binary data from file: %w", err)
	}

	err = proto.Unmarshal(data, message)
	if err != nil {
		return fmt.Errorf("cannot unmarshal binary to proto message: %w", err)
	}

	return nil
}
