package aoc2023

// func TestGetCalibrationValue(t *testing.T) {
// 	tests := []struct {
// 		input    string
// 		expected int
// 	}{
// 		{"1abc2", 12},
// 		{"pqr3stu8vwx", 38},
// 		{"a1b2c3d4e5f", 15},
// 		{"treb7uchet", 7},
// 	}

// 	for _, test := range tests {
// 		result := GetCalibrationValue(test.input)
// 		if result != test.expected {
// 			t.Errorf("GetCalibrationValue(%s) = %d, want %d", test.input, result, test.expected)
// 		}
// 	}
// }

// func TestSumCalibrationValues(t *testing.T) {
// 	tests := []struct {
// 		input    string
// 		expected int
// 	}{
// 		{"1abc2\npqr3stu8vwx\na1b2c3d4e5f\ntreb7uchet", 142},
// 		{"1\n2\n3\n4\n5", 15},
// 		{"10\n20\n30\n40\n50", 150},
// 	}

// 	for _, test := range tests {
// 		r := strings.NewReader(test.input)
// 		result, err := SumCalibrationValues(r)
// 		if err != nil {
// 			t.Errorf("SumCalibrationValues(%s) returned an error: %v", test.input, err)
// 		}
// 		if result != test.expected {
// 			t.Errorf("SumCalibrationValues(%s) = %d, want %d", test.input, result, test.expected)
// 		}
// 	}
// }
