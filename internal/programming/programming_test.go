package programming

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	programmingLib "github.com/brentellingson/learning-golang-lib/programming"
)

func TestPostUUID(t *testing.T) {
	tests := []struct {
		name               string
		queryParam         string
		expectedStatus     int
		expectedLength     int
		expectedPredicates []func(string) bool
	}{
		{
			name:           "UUID no-hyphens=null",
			queryParam:     "",
			expectedStatus: http.StatusOK,
			expectedLength: 36,
			expectedPredicates: []func(string) bool{
				func(s string) bool { return strings.Contains(s, "-") },
			},
		},
		{
			name:           "UUID no-hyphens=false",
			queryParam:     "?no-hyphens=false",
			expectedStatus: http.StatusOK,
			expectedLength: 36,
			expectedPredicates: []func(string) bool{
				func(s string) bool { return strings.Contains(s, "-") },
			},
		},
		{
			name:           "UUID no-hyphens=true",
			queryParam:     "?no-hyphens=true",
			expectedStatus: http.StatusOK,
			expectedLength: 32,
			expectedPredicates: []func(string) bool{
				func(s string) bool { return !strings.Contains(s, "-") },
			},
		},
	}

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	controller := NewController(programmingLib.NewService())
	controller.AddRoutes(router.Group("/"))

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, _ := http.NewRequest(http.MethodPost, "/programming/uuid"+tt.queryParam, nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			rsp := Response{}
			if err := json.Unmarshal(w.Body.Bytes(), &rsp); err != nil {
				t.Fatalf("failed to unmarshal response: %v", err)
			}

			assert.Equal(t, tt.expectedStatus, w.Code)
			assert.Contains(t, w.Body.String(), "uuid")
			assert.Len(t, rsp.UUID, tt.expectedLength)
			for _, pred := range tt.expectedPredicates {
				assert.True(t, pred(rsp.UUID))
			}
		})
	}
}
