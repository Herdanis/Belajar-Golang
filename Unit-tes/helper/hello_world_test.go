package helper

// untuk nama file dari unit menggunakan '..._test.go' di bagian belakang nama file

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func valueTes() string {
	return "yes"
}

/* aturan nama func unit test menggunakan "Test..." nama func nya
di unit test wajib menggunakan (t *testing.T) dan tidak ada return value */

// table test bertujuan agar tidak berulang kali mengganti data
func TestTableHello(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "qwe",
			request:  "qwe",
			expected: "Hello qwe",
		},
		{
			name:     "asd",
			request:  "asd",
			expected: "Hello asd",
		},
		{
			name:     "zxc",
			request:  "zxc",
			expected: "Hello zxc",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}
func TestSubTest(t *testing.T) {

	// ini adalah subtest
	t.Run("Tes" /* ini nama subtest */, func(t *testing.T) {

		// isi dari func testnya
		result := HelloWorld(valueTes())
		require.Equal(t, "Hello yes", result, "TestHelloWorld SubTest")
	})
}

/* untuk menjalankan testing.M wajib menggunakan func TestMain(m *testing.M)
func ini hanya bisa di gunakan dalam 1 package */
func TestMain(m *testing.M) {

	// mengeksekusi semua func sebelum unit test
	fmt.Println("Before unit Test")

	m.Run()

	// mengeksekusi semua func sesudah unit test
	fmt.Println("After unit Test")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "darwin" {
		// cukup menggunakan t.Skip(...) untuk menghentikan testing
		t.Skip("Cant run on MacOS")
	}

	result := HelloWorld(valueTes())
	require.Equal(t, "Hello yes", result, "TestHelloWorld Done Skip")
}

func TestHelloWorld(t *testing.T) {
	result := HelloWorld(valueTes())
	if result != "Hello yes" {
		/* sama dengan fail() akan menghentikan func saat itu juga
		tapi akan di lanjutkan ke func unit test selanjutanya */
		t.FailNow()
	}

	//func selanjuta nya tidak di lanjutkan
	fmt.Println("TestHelloWorld Done FailNow")
}
func TestHelloWorld2(t *testing.T) {
	result := HelloWorld(valueTes())
	if result != "Hello yes" {
		/* fail akan menggagal kan unit test tapi code program selanjutanya akan di lanjutkan
		namun hasil nya di anggap gagal */
		t.Fail()
	}

	// code program di bawah tetap di eksekusi
	fmt.Println("TestHelloWorld Done Fail")
}

func TestHelloWorld3(t *testing.T) {
	result := HelloWorld(valueTes())
	if result != "Hello yes" {
		/* akan mirip seperti proses log
		setelah selasai log error akan memanggil fail() */
		t.Error("Result Not Valid")
	}
	fmt.Println("TestHelloWorld Done Error")
}

func TestHelloWorld4(t *testing.T) {
	result := HelloWorld(valueTes())
	if result != "Hello yes" {
		/* sama seperti error() tapi akan memanggil failnow() */
		t.Fatal("Result Not Valid")
	}
	fmt.Println("TestHelloWorld Done Fatal")
}

// menggunakan testify tidak perlu menggunakan if else
// lebih di sarankan menggunakan testify
func TestHelloWorld5(t *testing.T) {
	result := HelloWorld(valueTes())

	// jika gagal akan memanggil fail()
	assert.Equal(t, "Hello yes", result, "Result Not Valid")
	fmt.Println("TestHelloWorld Done Assert Done")
}

func TestHelloWorld6(t *testing.T) {
	result := HelloWorld(valueTes())

	// jika gagal akan memanggil failnow()
	require.Equal(t, "Hello yes", result, "Result Not Valid")
	fmt.Println("TestHelloWorld Done Requere Done")
}
