package config

import "testing"

func TestMineBlock(t *testing.T) {

	tests := []struct {
		FileName     string
		expectedPath string
	}{
		{
			FileName:     "./testdata/a.json",
			expectedPath: "./a/b/c.json",
		},
	}

	for ii, test := range tests {
		ReadConfig(test.FileName)
		cfg := GetGlobalConfig()
		if cfg.DataDir != test.expectedPath {
			t.Errorf("Test %d, expected %s got %s\n", ii, test.expectedPath, cfg.DataDir)
		}
	}

}
