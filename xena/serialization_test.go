package xena

import (
	"github.com/xenaex/client-go/xena/fixjson"
	"github.com/xenaex/client-go/xena/messages"
	"testing"
)

func TestUnmarshalLogon(t *testing.T) {
	s := `{"35":"MsgTypeValue","108":123,"1328":"RejectTextBalue","1":[1,2,3],"52":1234567890,"553":"UsernameValue","554":"PasswordValue"}`
	o2 := new(messages.Logon)
	err := fixjson.Unmarshal([]byte(s), o2)
	if err != nil {
		t.Errorf("error: %s on Unmarshal(%s)", err, s)
	}
}

func TestLogon(t *testing.T) {
	o := messages.Logon{
		MsgType:     "MsgTypeValue",
		HeartBtInt:  123,
		RejectText:  "RejectTextBalue",
		Account:     []uint64{1, 2, 3},
		SendingTime: 1234567890,
		Username:    "UsernameValue",
		Password:    "PasswordValue",
	}

	j, err := fixjson.Marshal(o)
	if err != nil {
		t.Errorf("error: %s on marshal(%v)", err, o)
	}
	s := string(j)

	expected := `{"35":"MsgTypeValue","108":123,"1328":"RejectTextBalue","1":[1,2,3],"52":1234567890,"553":"UsernameValue","554":"PasswordValue"}`
	if s != expected {
		t.Errorf("got: %s, but expected: %s", s, expected)
	}

	o2 := new(messages.Logon)
	err = fixjson.Unmarshal(j, o2)
	if err != nil {
		t.Errorf("error: %s on Unmarshal(%s)", err, s)
	}

	j2, err := fixjson.Marshal(o2)
	if err != nil {
		t.Errorf("error: %s on marshal(%v)", err, o2)
	}
	s2 := string(j2)
	if s2 != s {
		t.Errorf("json marshal results are not equals %s != %s", s2, s)
	}

	if o.MsgType != o2.MsgType || o.HeartBtInt != o2.HeartBtInt || o.RejectText != o2.RejectText ||
		o.SendingTime != o2.SendingTime || o.Username != o2.Username || o.Password != o2.Password ||
		o.Account[0] != o2.Account[0] || o.Account[1] != o2.Account[1] || o.Account[2] != o2.Account[2] {
		t.Errorf("objects are not equals %v != %v", o, o2)
	}
}
