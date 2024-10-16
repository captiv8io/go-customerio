package customerio_test

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/customerio/go-customerio/v3"
)

func TestAPIOptions(t *testing.T) {

	client := customerio.NewAPIClient("mykey")
	if client.URL != customerio.RegionUS.ApiURL {
		t.Errorf("wrong default url. got: %s, want: %s", client.URL, customerio.RegionUS.ApiURL)
	}

	client = customerio.NewAPIClient("mykey", customerio.WithRegion(customerio.RegionEU))
	if client.URL != customerio.RegionEU.ApiURL {
		t.Errorf("wrong url. got: %s, want: %s", client.URL, customerio.RegionEU.ApiURL)
	}

	hc := &http.Client{}
	client = customerio.NewAPIClient("mykey", customerio.WithHTTPClient(hc))
	if !reflect.DeepEqual(client.Client, hc) {
		t.Errorf("wrong http client. got: %#v, want: %#v", client.Client, hc)
	}

	customUserAgent := "Customer.io"
	client = customerio.NewAPIClient("mykey", customerio.WithUserAgent(customUserAgent))
	if client.UserAgent != customUserAgent {
		t.Errorf("wrong user-agent. got: %s, want: %s", client.UserAgent, customUserAgent)
	}
}

func TestTrackOptions(t *testing.T) {

	client := customerio.NewTrackClient("site_id", "api_key")
	if client.URL != customerio.RegionUS.TrackURL {
		t.Errorf("wrong default url. got: %s, want: %s", client.URL, customerio.RegionUS.TrackURL)
	}

	client = customerio.NewTrackClient("site_id", "api_key", customerio.WithRegion(customerio.RegionEU))
	if client.URL != customerio.RegionEU.TrackURL {
		t.Errorf("wrong url. got: %s, want: %s", client.URL, customerio.RegionEU.TrackURL)
	}

	hc := &http.Client{}
	client = customerio.NewTrackClient("site_id", "api_key", customerio.WithHTTPClient(hc))
	if !reflect.DeepEqual(client.Client, hc) {
		t.Errorf("wrong http client. got: %#v, want: %#v", client.Client, hc)
	}

	customUserAgent := "Customer.io"
	client = customerio.NewTrackClient("site_id", "api_key", customerio.WithUserAgent(customUserAgent))
	if client.UserAgent != customUserAgent {
		t.Errorf("wrong user-agent. got: %s, want: %s", client.UserAgent, customUserAgent)
	}

	emailIDType := "email"
	client = customerio.NewTrackClient("site_id", "api_key", customerio.WithIDType(emailIDType))
	if client.IDType != emailIDType {
		t.Errorf("wrong id_type. got: %s, want: %s", client.IDType, emailIDType)
	}
}
