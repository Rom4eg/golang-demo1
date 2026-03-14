package target

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/Rom4eg/golang-demo1/internal/target/mock"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestTarget_Check(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		url     string
		prepare func(*testing.T, *mock.MockClient)
		expect  func(*testing.T, error, bool)
	}{
		{
			name: "FAIL: incorrect url",
			url:  "test",
			prepare: func(t *testing.T, c *mock.MockClient) {
				c.EXPECT().Head(gomock.Any()).Return(nil, context.DeadlineExceeded)
			},
			expect: func(t *testing.T, err error, b bool) {
				assert.Error(t, err)
				assert.False(t, b)
			},
		},
		{
			name: "FAIL: unexpected status code",
			url:  "http://localhost",
			prepare: func(t *testing.T, c *mock.MockClient) {
				resp := http.Response{
					StatusCode: http.StatusInternalServerError,
				}
				c.EXPECT().Head(gomock.Any()).Return(&resp, nil)
			},
			expect: func(t *testing.T, err error, b bool) {
				assert.ErrorIs(t, err, ErrUnexpectedStatusCode)
				assert.False(t, b)
			},
		},
		{
			name: "FAIL: no content-Length header",
			url:  "http://localhost",
			prepare: func(t *testing.T, mc *mock.MockClient) {
				resp := http.Response{
					StatusCode: http.StatusOK,
					Header: http.Header{
						"Accept-Ranges":  []string{"bytes"},
						"Content-Length": []string{"test"},
					},
				}
				mc.EXPECT().Head(gomock.Any()).Return(&resp, nil)
			},
			expect: func(t *testing.T, err error, b bool) {
				assert.Error(t, err)
				assert.False(t, b)
			},
		},
		{
			name: "FAIL: invalid file",
			url:  "http://localhost",
			prepare: func(t *testing.T, mc *mock.MockClient) {
				resp := http.Response{
					StatusCode: http.StatusOK,
					Header: http.Header{
						"Accept-Ranges":  []string{"bytes"},
						"Content-Length": []string{fmt.Sprint(0)},
					},
				}
				mc.EXPECT().Head(gomock.Any()).Return(&resp, nil)
			},
			expect: func(t *testing.T, err error, b bool) {
				assert.ErrorIs(t, err, ErrIncorrectFile)
				assert.False(t, b)
			},
		},
		{
			name: "PASS: OK",
			url:  "http://localhost",
			prepare: func(t *testing.T, mc *mock.MockClient) {
				resp := http.Response{
					StatusCode: http.StatusOK,
					Header: http.Header{
						"Accept-Ranges":  []string{"bytes"},
						"Content-Length": []string{fmt.Sprint(1024)},
					},
				}
				mc.EXPECT().Head(gomock.Any()).Return(&resp, nil)
			},

			expect: func(t *testing.T, err error, b bool) {
				assert.NoError(t, err)
				assert.True(t, b)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			c := mock.NewMockClient(ctrl)
			tt.prepare(t, c)

			target := New(tt.url)
			target.Client = c

			check, err := target.Check()
			tt.expect(t, err, check)
		})
	}
}
