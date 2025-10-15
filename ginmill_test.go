package ginmill

import (
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
)

func routes() (r gin.RoutesInfo) {
	e := gin.New()

	e.GET("/get", func(c *gin.Context) {})

	return r
}

func TestFeatures_RoutesField(t *testing.T) {
	r := routes()
	f := NewFeatures(r)
	if !reflect.DeepEqual(f.routes, r) {
		t.Errorf("Features.routes = %v, want %v", f.routes, r)
	}
}

func TestServer_Initialization(t *testing.T) {
	engine := gin.New()
	s := &Server{Engine: engine}
	if s.Engine != engine {
		t.Errorf("Server.Engine = %v, want %v", s.Engine, engine)
	}
}

func TestNewFeatures_EmptyRoutes(t *testing.T) {
	f := NewFeatures(gin.RoutesInfo{})
	if f == nil || len(f.routes) != 0 {
		t.Errorf("NewFeatures with empty routes should have empty Features.routes, got %v", f.routes)
	}
}

func TestServer_With_MultipleCalls(t *testing.T) {
	engine := gin.New()
	s := &Server{Engine: engine}
	r := routes()
	f := NewFeatures(r)
	s.With(f)
	// Call With again to check idempotency or side effects
	s.With(f)
	if !reflect.DeepEqual(r, engine.Routes()) {
		t.Errorf("Server.With() after multiple calls, routes = %v, want %v", engine.Routes(), r)
	}
}

func TestNewFeatures(t *testing.T) {
	type args struct {
		routes gin.RoutesInfo
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test NewFeatures",
			args: args{
				routes: routes(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewFeatures(tt.args.routes)
			r := reflect.TypeOf(got)
			if r.Kind() != reflect.Ptr || r.Elem().Name() != "Features" {
				t.Errorf("NewFeatures() = %v, want *Features", got)
			}
		})
	}
}

func TestFeatures_GetRoutes(t *testing.T) {
	r := routes()
	f := NewFeatures(r)
	got := f.GetRoutes()

	if len(got) != len(f.routes) {
		t.Errorf("GetRoutes() length = %d, want %d", len(got), len(f.routes))
	}

	for i := range got {
		if got[i].Method != f.routes[i].Method || got[i].Path != f.routes[i].Path {
			t.Errorf("GetRoutes() element %d not equal", i)
		}
	}

	if len(got) > 0 {
		got[0].Path = "/changed"
		if got[0].Path == f.routes[0].Path {
			t.Errorf("GetRoutes should return a copy, not a reference")
		}
	}
}

func TestServer_With(t *testing.T) {
	type args struct {
		r gin.RoutesInfo
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test With",
			args: args{
				r: routes(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			engine := gin.New()
			s := &Server{
				Engine: engine,
			}
			f := NewFeatures(tt.args.r)

			if s.With(f); !reflect.DeepEqual(tt.args.r, engine.Routes()) {
				t.Errorf("Server.With() r = %v, routes %v", tt.args.r, engine.Routes())
			}
		})
	}
}
