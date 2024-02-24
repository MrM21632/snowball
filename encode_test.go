package snowball

import "testing"

func TestBinaryEncode(t *testing.T) {
	t.Setenv("SNOWBALL_EPOCH_MS", "1704121810000")
	t.Setenv("SNOWBALL_NODE_ID", "32")

	tests := []struct {
		name string
		arg  SnowballID
		want string
	}{
		{
			name: "Test 1",
			arg:  19638199173316608,
			want: "1000101110001001101011010001101110000010000000000000000",
		},
		{
			name: "Test 2",
			arg:  19638200368693248,
			want: "1000101110001001101011011010101000000010000000000000000",
		},
		{
			name: "Test 3",
			arg:  19638197952774144,
			want: "1000101110001001101011001000101000000010000000000000000",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			encoded := tt.arg.ToBinary()
			if encoded != tt.want {
				t.Errorf("ToBinary() decoded: %v, want: %v", encoded, tt.want)
			}
		})
	}
}

func TestBinaryDecode(t *testing.T) {
	t.Setenv("SNOWBALL_EPOCH_MS", "1704121810000")
	t.Setenv("SNOWBALL_NODE_ID", "32")

	tests := []struct {
		name    string
		arg     string
		want    SnowballID
		wantErr bool
	}{
		{
			name:    "Test 1",
			arg:     "1000101110001001101011010001101110000010000000000000000",
			want:    19638199173316608,
			wantErr: false,
		},
		{
			name:    "Test 2",
			arg:     "1000101110001001101011011010101000000010000000000000000",
			want:    19638200368693248,
			wantErr: false,
		},
		{
			name:    "Test 3",
			arg:     "1000101110001001101011001000101000000010000000000000000",
			want:    19638197952774144,
			wantErr: false,
		},
		{
			name:    "Invalid binary string",
			arg:     "1000101110001001101011010001101110000010000002000000000",
			want:    0,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			decoded, err := FromBinary(tt.arg)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromBinary() error: %v, wantErr: %v", err, tt.wantErr)
			}
			if decoded != tt.want {
				t.Errorf("FromBinary() decoded: %v, want: %v", decoded, tt.want)
			}
		})
	}
}

func TestHexEncode(t *testing.T) {
	t.Setenv("SNOWBALL_EPOCH_MS", "1704121810000")
	t.Setenv("SNOWBALL_NODE_ID", "32")

	tests := []struct {
		name string
		arg  SnowballID
		want string
	}{
		{
			name: "Test 1",
			arg:  19638199173316608,
			want: "45c4d68dc10000",
		},
		{
			name: "Test 2",
			arg:  19638200368693248,
			want: "45c4d6d5010000",
		},
		{
			name: "Test 3",
			arg:  19638197952774144,
			want: "45c4d645010000",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			encoded := tt.arg.ToHex()
			if encoded != tt.want {
				t.Errorf("ToHex() decoded: %v, want: %v", encoded, tt.want)
			}
		})
	}
}

func TestHexDecode(t *testing.T) {
	t.Setenv("SNOWBALL_EPOCH_MS", "1704121810000")
	t.Setenv("SNOWBALL_NODE_ID", "32")

	tests := []struct {
		name    string
		arg     string
		want    SnowballID
		wantErr bool
	}{
		{
			name:    "Test 1",
			arg:     "45c4d68dc10000",
			want:    19638199173316608,
			wantErr: false,
		},
		{
			name:    "Test 2",
			arg:     "45c4d6d5010000",
			want:    19638200368693248,
			wantErr: false,
		},
		{
			name:    "Test 3",
			arg:     "45c4d645010000",
			want:    19638197952774144,
			wantErr: false,
		},
		{
			name:    "Invalid hex string",
			arg:     "45c4d645010g00",
			want:    0,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			decoded, err := FromHex(tt.arg)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromHex() error: %v, wantErr: %v", err, tt.wantErr)
			}
			if decoded != tt.want {
				t.Errorf("FromHex() decoded: %v, want: %v", decoded, tt.want)
			}
		})
	}
}

func TestBase32Encode(t *testing.T) {
	t.Setenv("SNOWBALL_EPOCH_MS", "1704121810000")
	t.Setenv("SNOWBALL_NODE_ID", "32")

	tests := []struct {
		name string
		arg  SnowballID
		want string
	}{
		{
			name: "Test 1",
			arg:  19638199173316608,
			want: "AAAMDDOWYRCQA===",
		},
		{
			name: "Test 2",
			arg:  19638200368693248,
			want: "AAAADVOWYRCQA===",
		},
		{
			name: "Test 3",
			arg:  19638197952774144,
			want: "AAAACROWYRCQA===",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			encoded := tt.arg.ToBase32()
			if encoded != tt.want {
				t.Errorf("ToBase32() decoded: %v, want: %v", encoded, tt.want)
			}
		})
	}
}

func TestBase32Decode(t *testing.T) {
	t.Setenv("SNOWBALL_EPOCH_MS", "1704121810000")
	t.Setenv("SNOWBALL_NODE_ID", "32")

	tests := []struct {
		name    string
		arg     string
		want    SnowballID
		wantErr bool
	}{
		{
			name:    "Test 1",
			arg:     "AAAMDDOWYRCQA===",
			want:    19638199173316608,
			wantErr: false,
		},
		{
			name:    "Test 2",
			arg:     "AAAADVOWYRCQA===",
			want:    19638200368693248,
			wantErr: false,
		},
		{
			name:    "Test 3",
			arg:     "AAAACROWYRCQA===",
			want:    19638197952774144,
			wantErr: false,
		},
		{
			name:    "Invalid characters in base32 encoding",
			arg:     "AAAACROWYR8QA===",
			want:    0,
			wantErr: true,
		},
		{
			name:    "Lowercase characters in base32 encoding",
			arg:     "aaaacrowyrcqa===",
			want:    0,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			decoded, err := FromBase32(tt.arg)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromBase32() error: %v, wantErr: %v", err, tt.wantErr)
			}
			if decoded != tt.want {
				t.Errorf("FromBase32() decoded: %v, want: %v", decoded, tt.want)
			}
		})
	}
}

func TestBase62Encode(t *testing.T) {
	t.Setenv("SNOWBALL_EPOCH_MS", "1704121810000")
	t.Setenv("SNOWBALL_NODE_ID", "32")

	tests := []struct {
		name string
		arg  SnowballID
		want string
	}{
		{
			name: "Test 1",
			arg:  19638199173316608,
			want: "1RwTVZKtLU",
		},
		{
			name: "Test 2",
			arg:  19638200368693248,
			want: "1RwTWsEZPs",
		},
		{
			name: "Test 3",
			arg:  19638197952774144,
			want: "1RwTUEjcUi",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			encoded := tt.arg.ToBase62()
			if encoded != tt.want {
				t.Errorf("ToBase62() decoded: %v, want: %v", encoded, tt.want)
			}
		})
	}
}

func TestBase62Decode(t *testing.T) {
	t.Setenv("SNOWBALL_EPOCH_MS", "1704121810000")
	t.Setenv("SNOWBALL_NODE_ID", "32")

	tests := []struct {
		name    string
		arg     string
		want    SnowballID
		wantErr bool
	}{
		{
			name:    "Test 1",
			arg:     "1RwTVZKtLU",
			want:    19638199173316608,
			wantErr: false,
		},
		{
			name:    "Test 2",
			arg:     "1RwTWsEZPs",
			want:    19638200368693248,
			wantErr: false,
		},
		{
			name:    "Test 3",
			arg:     "1RwTUEjcUi",
			want:    19638197952774144,
			wantErr: false,
		},
		{
			name:    "Invalid characters in base62 encoding",
			arg:     "1RwT+EjcUi",
			want:    0,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			decoded, err := FromBase62(tt.arg)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromBase62() error: %v, wantErr: %v", err, tt.wantErr)
			}
			if decoded != tt.want {
				t.Errorf("FromBase62() decoded: %v, want: %v", decoded, tt.want)
			}
		})
	}
}

func TestBase64Encode(t *testing.T) {
	t.Setenv("SNOWBALL_EPOCH_MS", "1704121810000")
	t.Setenv("SNOWBALL_NODE_ID", "32")

	tests := []struct {
		name string
		arg  SnowballID
		want string
	}{
		{
			name: "Test 1",
			arg:  19638199173316608,
			want: "AADBjdbERQA=",
		},
		{
			name: "Test 2",
			arg:  19638200368693248,
			want: "AAAB1dbERQA=",
		},
		{
			name: "Test 3",
			arg:  19638197952774144,
			want: "AAABRdbERQA=",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			encoded := tt.arg.ToBase64()
			if encoded != tt.want {
				t.Errorf("ToBase64() decoded: %v, want: %v", encoded, tt.want)
			}
		})
	}
}

func TestBase64Decode(t *testing.T) {
	t.Setenv("SNOWBALL_EPOCH_MS", "1704121810000")
	t.Setenv("SNOWBALL_NODE_ID", "32")

	tests := []struct {
		name    string
		arg     string
		want    SnowballID
		wantErr bool
	}{
		{
			name:    "Test 1",
			arg:     "AADBjdbERQA=",
			want:    19638199173316608,
			wantErr: false,
		},
		{
			name:    "Test 2",
			arg:     "AAAB1dbERQA=",
			want:    19638200368693248,
			wantErr: false,
		},
		{
			name:    "Test 3",
			arg:     "AAABRdbERQA=",
			want:    19638197952774144,
			wantErr: false,
		},
		{
			name:    "Invalid characters in base64 encoding",
			arg:     "AAABRdb+RQA=",
			want:    0,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			decoded, err := FromBase64(tt.arg)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromBase64() error: %v, wantErr: %v", err, tt.wantErr)
			}
			if decoded != tt.want {
				t.Errorf("FromBase64() decoded: %v, want: %v", decoded, tt.want)
			}
		})
	}
}

func BenchmarkToBinary(b *testing.B) {
	b.Setenv("SNOWBALL_EPOCH_MS", "1704121810000")
	b.Setenv("SNOWBALL_NODE_ID", "32")

	node, _ := InitNode()
	id := node.GenerateID()

	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		id.ToBinary()
	}
}

func BenchmarkFromBinary(b *testing.B) {
	b.Setenv("SNOWBALL_EPOCH_MS", "1704121810000")
	b.Setenv("SNOWBALL_NODE_ID", "32")

	node, _ := InitNode()
	id := node.GenerateID()
	sid := id.ToBinary()

	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		FromBinary(sid)
	}
}

func BenchmarkToHex(b *testing.B) {
	b.Setenv("SNOWBALL_EPOCH_MS", "1704121810000")
	b.Setenv("SNOWBALL_NODE_ID", "32")

	node, _ := InitNode()
	id := node.GenerateID()

	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		id.ToHex()
	}
}

func BenchmarkFromHex(b *testing.B) {
	b.Setenv("SNOWBALL_EPOCH_MS", "1704121810000")
	b.Setenv("SNOWBALL_NODE_ID", "32")

	node, _ := InitNode()
	id := node.GenerateID()
	sid := id.ToHex()
	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		FromHex(sid)
	}
}

func BenchmarkToBase32(b *testing.B) {
	b.Setenv("SNOWBALL_EPOCH_MS", "1704121810000")
	b.Setenv("SNOWBALL_NODE_ID", "32")

	node, _ := InitNode()
	id := node.GenerateID()

	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		id.ToBase32()
	}
}

func BenchmarkFromBase32(b *testing.B) {
	b.Setenv("SNOWBALL_EPOCH_MS", "1704121810000")
	b.Setenv("SNOWBALL_NODE_ID", "32")

	node, _ := InitNode()
	id := node.GenerateID()
	sid := id.ToBase32()

	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		FromBase32(sid)
	}
}

func BenchmarkToBase62(b *testing.B) {
	b.Setenv("SNOWBALL_EPOCH_MS", "1704121810000")
	b.Setenv("SNOWBALL_NODE_ID", "32")

	node, _ := InitNode()
	id := node.GenerateID()

	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		id.ToBase62()
	}
}

func BenchmarkFromBase62(b *testing.B) {
	b.Setenv("SNOWBALL_EPOCH_MS", "1704121810000")
	b.Setenv("SNOWBALL_NODE_ID", "32")

	node, _ := InitNode()
	id := node.GenerateID()
	sid := id.ToBase62()

	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		FromBase62(sid)
	}
}

func BenchmarkToBase64(b *testing.B) {
	b.Setenv("SNOWBALL_EPOCH_MS", "1704121810000")
	b.Setenv("SNOWBALL_NODE_ID", "32")

	node, _ := InitNode()
	id := node.GenerateID()

	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		id.ToBase64()
	}
}

func BenchmarkFromBase64(b *testing.B) {
	b.Setenv("SNOWBALL_EPOCH_MS", "1704121810000")
	b.Setenv("SNOWBALL_NODE_ID", "32")

	node, _ := InitNode()
	id := node.GenerateID()
	sid := id.ToBase64()

	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		FromBase64(sid)
	}
}
