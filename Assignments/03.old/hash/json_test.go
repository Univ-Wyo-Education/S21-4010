package hash

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/Univ-Wyo-Education/S19-4010/a/02/lib"
)

func Test_ToJSON(t *testing.T) {

	type Test1 struct {
		Sh SealType
		Bh BlockHashType
	}

	t1 := Test1{
		Sh: []byte{22, 33, 11},
		Bh: Keccak256([]byte("abcdef")),
	}

	got := lib.SVarI(t1)
	exp := `{
	"Sh": "16210b",
	"Bh": "acd0c377fe36d5b209125185bc3ac41155ed1bf7103ef9f0c2aff4320460b6df"
}`

	if got != exp {
		t.Errorf("Got ->%s<- expected ->%s<-\n", got, exp)
	}

}

func Test_FromJSON(t *testing.T) {

	type Test1 struct {
		Sh SealType
		Bh BlockHashType
	}

	var got Test1

	exp := Test1{
		Sh: []byte{22, 33, 11},
		Bh: Keccak256([]byte("abcdef")),
	}

	have := `{
	"Sh": "16210b",
	"Bh": "acd0c377fe36d5b209125185bc3ac41155ed1bf7103ef9f0c2aff4320460b6df"
}`

	err := json.Unmarshal([]byte(have), &got)
	if err != nil {
		t.Fatalf("Error unmarshal: %s\n", err)
	}

	if !reflect.DeepEqual(got, exp) {
		t.Errorf("Got ->%+v<- expected ->%+v<-\n", got, exp)
	}

}
