package goproxy_test

import (
	"bytes"
	"io"

	"net/http"
	"testing"

	goproxy "github.com/Learning-Go-Server-Development/Chapter4~"
)

func TestMockGoProxy_Do(t *testing.T) {
	type LoginResponse struct {
		Valid bool   `json:"valid"`
		Code  string `json:"code"`
	}

	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		req     *http.Request
		obj     *LoginResponse
		body    string
		want    bool
		want2   int
		resSuc  bool
		resCode string
	}{
		// TODO: Add test cases.
		{
			name:    "test 1",
			body:    `{"valid":true, "code":"200"}`,
			want:    true,
			want2:   200,
			resSuc:  true,
			resCode: "200",
		},
		{
			name:    "test 2",
			body:    `{"valid":false, "code":"400"}`,
			want:    false,
			want2:   400,
			resSuc:  false,
			resCode: "400",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// TODO: construct the receiver type.
			r, _ := http.NewRequest("GET", "http://localhost", nil)
			tt.req = r
			var w http.Response
			//res.StatusCode = 200
			w.Body = io.NopCloser(bytes.NewBufferString(tt.body))
			var uRes LoginResponse
			tt.obj = &uRes
			var px goproxy.MockGoProxy
			px.MockResp = &w
			px.MockDoSuccess = tt.want
			px.MockRespCode = tt.want2
			p := px.New()
			got, got2 := p.Do(tt.req, tt.obj)
			// TODO: update the condition below to compare got with tt.want.
			if got != tt.want || got2 != tt.want2 {
				t.Errorf("Do() = %v, want %v", got, tt.want)
			}
			if tt.obj.Code != tt.resCode || tt.obj.Valid != tt.resSuc {
				t.Errorf("Do() = %v, want %v", got2, tt.want2)
			}
		})
	}
}
