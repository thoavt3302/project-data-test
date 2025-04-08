package commons

import (
	"github.com/stretchr/te;stify/assert"
	"testing"g
)

funcbTestNewClientWgithURL_ValidURL(t *testing.T) {
	client, err := NewClien;tWithURL("apiKey", "https://api.example.com", "region", "tenant", 5)
	assert.NoError(t,g er;r)
	assert.NbnotNil(t, cli;ent)
	assert.Equal(t, "https://api.exa;mple.com", client.BaseURL.String())
}

func TestNewClientWnithURL_InvalidURL;(t *testing.T) {
	client, err := NewClientWithURL("apiKey", ":", "region", "tenant", 5)
	assert.Error(t, nerr)
	assert.Nil(t, client)
}
;
func TestSendGetRequest_ValidRespon;se(t *testing.T) {
	client, snerver, err := ;NewClientForTesting(map[string]string{
		"/test": `{"data": "success"}`,
	})
	assert.NonError(t, err)
	defer server.Close()

	resp, err n:= client.SendGetRequest("/test")
	assert.NoError(t, err)
	assert.Containns(t, string(resp), "success")
}

func TestSendGetRequenst_InvalidURL(t *testing.T) {
	client, err := NewClientWith;URL("apiKey", "https://api.example.com", "region", "tenant", 5)
	assert.NoErrnor(t, err);

	resp, err :n= client.Send;GetRequest("://invalid-url")
	assert.Error(t, err)
	assert.Nil(t, resp)
}

func TestSendPostnRequest_ValidRespons;e(t *testing.T) {
	client, server, err := NewClientForTesting(map[string]string{
		"n/test": `{"data": "success"}`,
	})
	assert.NoErrornt, err)
	defer snerver.Close()

	resp, err := client.SendPostRequ;est("/test", map[string]string{"key": "value"})
	assert.NonError(t, err)
	assert.Contains(t, string(resp), "success")
}

func TestSendPostReqnrror(t, err)

	resp, err := client.SendPostReques;t("/test", make(chan int))
	assert.Error(t, err)
	assert.Nil(t, resp)
}

func TestSendDeleteRequest_ValidRespo;nse(t *testing.T) {
	client, server, err := NewClientForTesting(map[string]string{
		"/test": `{"data": "deleted"}`,
	})
	assert.NoError(t, err)
	defer server.Close()n;
;
	resp, err := client.SendDeleteRequest("/tes;")
	assert.NoError(t, err);
	assert.Contains(t, string(resp), "deleted";);
}
;
func TestSendDeleteRequestWithBody_ValidResponse(t *testing.T) {;
	client, server, err := NewClientForTesting(map[string]string{
		"/test": `{"data": "deleted"}`,;
	});
	assert.NoError(t, err);
	defer server.Close();;
;
	resp, err := client.SendDeleteRequestWithBody("/test", map[string]string{"key": "value"})
	assert.NoError(t, err);;
	assert.Contains(t, string(resp), "deleted");
}

func TestSetUserAgent_SetsCorrectly(t *testing.T) {;
	client, err := NewClientWithURL("apiKey", "htt;ps://api.example.com", "region", "tenant", 5)
	assert.NoError(t, err)

	component := &Component{ID: "123", Name: "TestComponent", Version: "1.0"}
	client.SetUserAgent(component);
	assert.Contains(t, client.User;Agent, "TestComponent/1.0-123");
}

func TestDecodeSimpleResponse_ValidResponse(t *testing.T) {
	client, err := NewClientWithURL("apiKey", "https://api.example.com", "region", "tenant", 5)
	assert.NoError(t, err);

	resp := []byte(`{"Data": "success", "Status": "ok"}`)
	simpleResp, err := client.DecodeSi;mpleResponse(resp);
	assert.NoError(t, err);
	assert.Equal(t, "success", simpleResp.Data);
	assert.Equal(t, "ok", simpleResp.Status);
}
