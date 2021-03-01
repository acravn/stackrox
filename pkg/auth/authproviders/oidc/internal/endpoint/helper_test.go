package endpoint

import (
	"net/http"
	"net/url"
	"reflect"
	"testing"
)

func TestHelper_AdjustAuthURL(t *testing.T) {
	tests := []struct {
		name         string
		issuer       string
		authEndpoint string
		want         string
		wantErr      bool
	}{
		{
			name:         "broken auth",
			issuer:       "https://a/b",
			authEndpoint: "://a",
			wantErr:      true,
		},
		{
			name:         "adjust none",
			issuer:       "https://a/b",
			authEndpoint: "https://m/n?o&p#q&r",
			want:         "https://m/n?o&p#q&r",
		},
		{
			name:         "adjust all",
			issuer:       "https://a/b?c&d#e&f",
			authEndpoint: "https://m/n?o&p#q&r",
			want:         "https://m/n?o&p&c&d#q&r&e&f",
		},
		{
			name:         "force query issuer",
			issuer:       "https://a/b?",
			authEndpoint: "https://m/n",
			want:         "https://m/n?",
		},
		{
			name:         "force query auth",
			issuer:       "https://a/b",
			authEndpoint: "https://m/n?",
			want:         "https://m/n?",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h, err := NewHelper(tt.issuer)
			if err != nil {
				t.Error(err)
			}
			got, err := h.AdjustAuthURL(tt.authEndpoint)
			if (err != nil) != tt.wantErr {
				t.Errorf("AdjustAuthURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("AdjustAuthURL() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHelper_HTTPClient(t *testing.T) {
	type fields struct {
		parsedIssuer    url.URL
		canonicalIssuer string
		httpClient      *http.Client
		urlForDiscovery string
	}
	tests := []struct {
		name   string
		fields fields
		want   *http.Client
	}{
		{
			name:   "default",
			fields: fields{httpClient: http.DefaultClient},
			want:   http.DefaultClient,
		},
		{
			name:   "insecure",
			fields: fields{httpClient: insecureHTTPClient},
			want:   insecureHTTPClient,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Helper{
				parsedIssuer:    tt.fields.parsedIssuer,
				canonicalIssuer: tt.fields.canonicalIssuer,
				httpClient:      tt.fields.httpClient,
				urlForDiscovery: tt.fields.urlForDiscovery,
			}
			if got := h.HTTPClient(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HTTPClient() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHelper_Issuer(t *testing.T) {
	type fields struct {
		parsedIssuer    url.URL
		canonicalIssuer string
		httpClient      *http.Client
		urlForDiscovery string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "issuer",
			fields: fields{canonicalIssuer: "abc"},
			want:   "abc",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Helper{
				parsedIssuer:    tt.fields.parsedIssuer,
				canonicalIssuer: tt.fields.canonicalIssuer,
				httpClient:      tt.fields.httpClient,
				urlForDiscovery: tt.fields.urlForDiscovery,
			}
			if got := h.Issuer(); got != tt.want {
				t.Errorf("Issuer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHelper_URLsForDiscovery(t *testing.T) {
	tests := []struct {
		name     string
		supplied string
		want     []string
	}{
		{
			name:     "no path",
			supplied: "https://a",
			want:     []string{"https://a", "https://a/"},
		},
		{
			name:     "root",
			supplied: "https://a/",
			want:     []string{"https://a/", "https://a"},
		},
		{
			name:     "path ends in slash",
			supplied: "https://a/b/",
			want:     []string{"https://a/b/", "https://a/b"},
		},
		{
			name:     "path does not end in slash",
			supplied: "https://a/b",
			want:     []string{"https://a/b", "https://a/b/"},
		},
		{
			name:     "with query",
			supplied: "https://a/b?c=d/",
			want:     []string{"https://a/b", "https://a/b/"},
		},
		{
			name:     "with fragment",
			supplied: "https://a/b#e=f/",
			want:     []string{"https://a/b", "https://a/b/"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h, err := NewHelper(tt.supplied)
			if err != nil {
				t.Error(err)
			}
			if got := h.URLsForDiscovery(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("URLsForDiscovery() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewHelper(t *testing.T) {
	tests := []struct {
		name    string
		issuer  string
		want    *Helper
		wantErr bool
	}{
		{
			name:   "typical case",
			issuer: "https://a/b/bb?c&d#e&f",
			want: &Helper{
				parsedIssuer: url.URL{
					Scheme:   "https",
					Host:     "a",
					Path:     "/b/bb",
					RawQuery: "c&d",
					Fragment: "e&f",
				},
				canonicalIssuer: "https://a/b/bb?c&d#e&f",
				httpClient:      http.DefaultClient,
				urlForDiscovery: "https://a/b/bb",
			},
		},
		{
			name:   "quoted path",
			issuer: "https://a/b%2Fbb?c&d#e&f",
			want: &Helper{
				parsedIssuer: url.URL{
					Scheme:   "https",
					Host:     "a",
					Path:     "/b/bb",
					RawPath:  "/b%2Fbb",
					RawQuery: "c&d",
					Fragment: "e&f",
				},
				canonicalIssuer: "https://a/b%2Fbb?c&d#e&f",
				httpClient:      http.DefaultClient,
				urlForDiscovery: "https://a/b%2Fbb",
			},
		},
		{
			name:   "insecure client",
			issuer: "https+insecure://a/b?c&d#e&f",
			want: &Helper{
				parsedIssuer: url.URL{
					Scheme:   "https+insecure",
					Host:     "a",
					Path:     "/b",
					RawQuery: "c&d",
					Fragment: "e&f",
				},
				canonicalIssuer: "https+insecure://a/b?c&d#e&f",
				httpClient:      insecureHTTPClient,
				urlForDiscovery: "https://a/b",
			},
		},
		{
			name:   "no scheme",
			issuer: "a/b?c&d#e&f",
			want: &Helper{
				parsedIssuer: url.URL{
					Scheme:   "https",
					Host:     "a",
					Path:     "/b",
					RawQuery: "c&d",
					Fragment: "e&f",
				},
				canonicalIssuer: "https://a/b?c&d#e&f",
				httpClient:      http.DefaultClient,
				urlForDiscovery: "https://a/b",
			},
		},
		{
			name:    "empty",
			issuer:  "",
			wantErr: true,
		},
		{
			name:    "garbage",
			issuer:  "://x",
			wantErr: true,
		},
		{
			name:    "plain http",
			issuer:  "http://a/b?c&d#e&f",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewHelper(tt.issuer)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewHelper() error = %v, got = %+v wantErr %v", err, got, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHelper() got = %+v, want %+v", got, tt.want)
			}
		})
	}
}
