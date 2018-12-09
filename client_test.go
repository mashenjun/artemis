package artemis

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"log"
	"testing"
)

func TestNew(t *testing.T) {
	_, err := New("https://open8200.hikvision.com", "24341259", "M5llsRpDovRZcB3WkhTk")
	if err != nil {
		t.Fatalf("could not new client: %v\n", err)
	}
}

func TestClient_ControlUnits(t *testing.T) {
	cli, err := New("https://open8200.hikvision.com", "24341259", "M5llsRpDovRZcB3WkhTk")
	if err != nil {
		t.Fatalf("could not new client: %v\n", err)
	}
	resp, err := cli.ControlUnits(context.TODO(), 10,0)
	if err != nil {
		t.Fatalf("could not get control units: %v\n", err)
	}
	log.Printf("%+v\n", resp)
}

func TestClient_ChildrenControlUnits(t *testing.T) {
	cli, err := New("https://open8200.hikvision.com", "24341259", "M5llsRpDovRZcB3WkhTk")
	if err != nil {
		t.Fatalf("could not new client: %v\n", err)
	}
	resp, err := cli.ChildrenControlUnits(context.TODO(), "0")
	if err != nil {
		t.Fatalf("could not get children control units: %v\n", err)
	}
	log.Printf("%+v\n", resp)
}

func TestClient_SecurityInfo(t *testing.T) {
	cli, err := New("https://open8200.hikvision.com", "24341259", "M5llsRpDovRZcB3WkhTk")
	if err != nil {
		t.Fatalf("could not new client: %v\n", err)
	}
	resp, err := cli.SecurityInfo(context.TODO(), "24341259")
	if err != nil {
		t.Fatalf("could not get security info: %v\n", err)
	}
	log.Printf("%+v\n", resp)
}

func TestClient_Cameras(t *testing.T) {
	cli, err := New("https://open8200.hikvision.com", "24341259", "M5llsRpDovRZcB3WkhTk")
	if err != nil {
		t.Fatalf("could not new client: %v\n", err)
	}
	resp, err := cli.Cameras(context.TODO(), 10, 0)
	if err != nil {
		t.Fatalf("could not get cameras: %v\n", err)
	}
	log.Printf("%+v\n", resp)
}

func TestClient_ChildrenCameras(t *testing.T) {
	cli, err := New("https://open8200.hikvision.com", "24341259", "M5llsRpDovRZcB3WkhTk")
	if err != nil {
		t.Fatalf("could not new client: %v\n", err)
	}
	resp, err := cli.ChildrenCameras(context.TODO(), 10, 0, "11")
	if err != nil {
		t.Fatalf("could not get cameras: %v\n", err)
	}
	log.Printf("%+v\n", resp)
}

func TestBase64(t *testing.T) {
	base := `GET
*/*
application/text;charset=UTF-8
x-ca-key:24341259
x-ca-nonce:70064c35-ddff-41f0-9769-291b4764137c
x-ca-timestamp:1544367183633
/artemis/api/common/v1/remoteControlUnitRestService/findControlUnitPage?size=10&start=0`
	h := hmac.New(sha256.New, []byte("M5llsRpDovRZcB3WkhTk"))
	h.Write([]byte(base))
	signed := base64.StdEncoding.EncodeToString(h.Sum(nil))
	log.Printf("%s\n", signed)
}