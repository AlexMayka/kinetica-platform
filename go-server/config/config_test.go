package config

import (
	"encoding/json"
	"os"
	"reflect"
	"testing"
)

var (
	testFileValidOne = `
		{
		  "BLE": {
			"command_char_uuid" : "0000abf3-0000-1000-8000-00805f9b34fb"
		  },
		  "TCP": {
			"server":   "localhost",
			"port":     "9090"
		  }
		}
	`

	testFileValidTwo = `
		{
		  "BLE": {
			"command_char_uuid" : "0000abf3-0000-1000-8000-00805f9b34fb"
		  },
		  "TCP": {
			"server":   "localhost",
			"port":     "9090"
		  }
		}
	`

	testFileInvalidOne = `
		{
		  "TCP": {
			"server":   "localhost",
			"port":     "9090"
		  }
		}
	`

	testFileInvalidTwo = `
		{
		  "BLE": {
			"command_char_uuid" : ""
		  },
		  "TCP": {
			"server":   "localhost",
			"port":     "9090"
		  }
		}
	`

	testFileInvalidThree = `
		{
		  "BLE": {
			"command_char_uuid" : "0000abf3-0000-1000-8000-00805f9b34fb"
		  },
		  "TCP": {
			"port":     "9090"
		  }
		}
	`

	testFileInvalidFour = `
	`
)

func createTestingFile(t *testing.T, inputData string) string {
	t.Helper()

	tmpFile, err := os.CreateTemp("", "test-config-*.json")
	if err != nil {
		t.Fatalf("error creating temp file: %v", err)
	}
	defer func() { _ = tmpFile.Close() }()

	if _, err := tmpFile.Write([]byte(inputData)); err != nil {
		t.Fatalf("error writing to temp file: %v", err)
	}

	return tmpFile.Name()
}

func TestReadData_FilePresence(t *testing.T) {
	t.Run("Test 1: File exists", func(t *testing.T) {
		path := createTestingFile(t, testFileValidOne)
		t.Cleanup(func() {
			_ = os.Remove(path)
		})

		data, err := readData(path)
		if err != nil {
			t.Fatalf("expected no error, got: %v", err)
		}
		if data == nil || len(*data) == 0 {
			t.Errorf("expected data, got nil or empty")
		}
	})

	t.Run("Test 2: File does not exist", func(t *testing.T) {
		_, err := readData("nonexistent_file_123.json")
		if err == nil {
			t.Errorf("expected error when file is missing, got nil")
		}
	})

	t.Run("Test 3: Parsed config must match file content", func(t *testing.T) {
		path := createTestingFile(t, testFileValidOne)
		t.Cleanup(func() {
			_ = os.Remove(path)
		})

		// Чтение через функцию
		value, err := readData(path)
		if err != nil {
			t.Fatalf("expected no error, got: %v", err)
		}

		var fromReadData Config
		if err := json.Unmarshal(*value, &fromReadData); err != nil {
			t.Fatalf("failed to unmarshal from readData: %v", err)
		}

		raw, err := os.ReadFile(path)
		if err != nil {
			t.Fatalf("failed to read file directly: %v", err)
		}

		var fromDisk Config
		if err := json.Unmarshal(raw, &fromDisk); err != nil {
			t.Fatalf("failed to unmarshal from disk: %v", err)
		}

		if !reflect.DeepEqual(fromReadData, fromDisk) {
			t.Errorf("configs not equal:\nfrom readData: %+v\nfrom file: %+v", fromReadData, fromDisk)
		}
	})
}

func TestInitConfig(t *testing.T) {
	t.Run("Parse Data", func(t *testing.T) {
		tests := []struct {
			name      string
			input     string
			shouldErr bool
		}{
			{"Test 1: Valid Data", testFileValidOne, false},
			{"Test 2: Valid Data", testFileValidTwo, false},
			{"Test 3: Invalid Data (missing BLE)", testFileInvalidOne, true},
			{"Test 4: Invalid Data (empty BLE)", testFileInvalidTwo, true},
			{"Test 5: Invalid Data (missing TCP.server)", testFileInvalidThree, true},
			{"Test 6: Completely empty", testFileInvalidFour, true},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				config = &Config{}
				path := createTestingFile(t, tt.input)
				t.Cleanup(func() { _ = os.Remove(path) })

				err := Load(path)

				if tt.shouldErr {
					if err == nil {
						t.Errorf("expected error, got nil")
					}
					return
				}

				if err != nil {
					t.Errorf("expected no error, got: %v", err)
					return
				}

				var expected Config
				if err := json.Unmarshal([]byte(tt.input), &expected); err != nil {
					t.Fatalf("failed to unmarshal input into Config: %v", err)
				}

				got := GetConfig()
				if !reflect.DeepEqual(expected, *got) {
					t.Errorf("config mismatch:\nexpected: %+v\ngot:      %+v", expected, *got)
				}
			})
		}
	})
}
