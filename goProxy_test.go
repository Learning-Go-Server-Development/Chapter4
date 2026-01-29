package goproxy_test

import (
	"fmt"
	"net/http"
	"testing"

	px "github.com/Learning-Go-Server-Development/Chapter4~"
)

func TestGoProxy_Do(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		req   *http.Request
		url   string
		obj   any
		want  bool
		want2 int
	}{
		// TODO: Add test cases.
		{
			name:  "test 1",
			url:   "http://localhost:3001/rs/orders/get/12345",
			want:  true,
			want2: 200,
		},
		{
			name:  "test 2",
			url:   "http://localhost:3001/rs/orders/get/r12345",
			want:  false,
			want2: 400,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, rErr := http.NewRequest("GET", tt.url, nil)
			if rErr != nil {
				fmt.Println("request err: ", rErr)
			}
			var uRes []px.Order
			tt.obj = &uRes
			// TODO: construct the receiver type.
			var px px.GoProxy
			p := px.New()
			got, got2 := p.Do(r, tt.obj)
			// TODO: update the condition below to compare got with tt.want.
			if got != tt.want {
				t.Errorf("Do() = %v, want %v", got, tt.want)
			}
			if got2 != tt.want2 {
				t.Errorf("Do() = %v, want %v", got2, tt.want2)
			}
		})
	}
}
