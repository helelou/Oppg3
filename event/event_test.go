package event

import (
	"testing"
)

func TestPutIn(t *testing.T) {
	wanted := "[kylling korn ---\\ \\_hs+rev_/ _________________/---]"
	got := PutIn("Rev")
	if got != wanted {
		t.Errorf("Feil, fikk %q, ønsket %q", got, wanted)
	}

}

func TestCrossRiver(t *testing.T) {
	wanted := "[kylling korn ---\\  _________________\\_hs+rev_//---]"
	got := CrossRiver("Rev")
	if got != wanted {
		t.Errorf("Feil, fikk %q, ønsket %q", got, wanted)
	}
}

func TestTakeOut(t *testing.T) {
	wanted := "[kylling korn ---\\  _________________\\__//---hs+rev]"
	got := TakeOut("Rev")
	if got != wanted {
		t.Errorf("Feil, fikk %q, ønsket %q", got, wanted)
	}
}
