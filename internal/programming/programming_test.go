package programming

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

const (
	UUIDWithHyphens    = "1ce44be5-fe68-46f7-a153-51c1c91a4ae4"
	UUIDWithoutHyphens = "1ce44be5fe6846f7a15351c1c91a4ae4"
)

func TestPostUUID(t *testing.T) {
	tests := []struct {
		url                    string
		expectedStatus         int
		expectedWithoutHyphens bool
		expectedUUID           string
	}{
		{
			url:                    "/programming/uuid",
			expectedStatus:         http.StatusOK,
			expectedWithoutHyphens: false,
			expectedUUID:           UUIDWithHyphens,
		},
		{
			url:                    "/programming/uuid?no-hyphens=false",
			expectedStatus:         http.StatusOK,
			expectedWithoutHyphens: false,
			expectedUUID:           UUIDWithHyphens,
		},
		{
			url:                    "/programming/uuid?no-hyphens=true",
			expectedStatus:         http.StatusOK,
			expectedWithoutHyphens: true,
			expectedUUID:           UUIDWithoutHyphens,
		},
		{
			url:                    "/programming/uuid?no-hyphens=yes",
			expectedStatus:         http.StatusOK,
			expectedWithoutHyphens: false,
			expectedUUID:           UUIDWithHyphens,
		},
		{
			url:                    "/programming/uuid?blah",
			expectedStatus:         http.StatusOK,
			expectedWithoutHyphens: false,
			expectedUUID:           UUIDWithHyphens,
		},
	}

	gin.SetMode(gin.TestMode)

	for _, tt := range tests {
		t.Run(tt.url, func(t *testing.T) {
			router := gin.Default()

			service := NewMockService(t)
			service.EXPECT().NewUUID(tt.expectedWithoutHyphens).Return(tt.expectedUUID)
			controller := NewController(service)
			controller.AddRoutes(router.Group("/"))

			req, _ := http.NewRequest(http.MethodPost, tt.url, nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			output := Response{}
			if err := json.Unmarshal(w.Body.Bytes(), &output); err != nil {
				t.Fatalf("failed to unmarshal response: %v", err)
			}

			assert.Equal(t, tt.expectedStatus, w.Code)
			assert.Equal(t, tt.expectedUUID, output.UUID)
		})
	}
}
