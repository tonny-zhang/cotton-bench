// Copyright 2018 The xujiajun/gorouter Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package cottonbench

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"runtime"
	"strings"
	"testing"

	beegomux "github.com/beego/mux"
	"github.com/go-chi/chi"

	"github.com/go-zoo/bone"
	"github.com/julienschmidt/httprouter"
	triemux "github.com/teambition/trie-mux/mux"
	"github.com/tonny-zhang/cotton"
	"github.com/xujiajun/gorouter"
)

type route struct {
	method string
	path   string
}

// http://developer.github.com/v3/
var githubAPI = []route{
	// OAuth Authorizations
	{"GET", "/authorizations"},
	{"GET", "/authorizations/:id"},
	{"GET", "/authorizations/:id/test"},
	{"GET", "/authorizations/:id/tt"},
	{"POST", "/authorizations"},
	//{"PUT", "/authorizations/clients/:client_id"},
	//{"PATCH", "/authorizations/:id"},
	{"DELETE", "/authorizations/:id"},
	{"GET", "/applications/:client_id/tokens/:access_token"},
	{"DELETE", "/applications/:client_id/tokens"},
	{"DELETE", "/applications/:client_id/tokens/:access_token"},

	// Activity
	{"GET", "/events"},
	{"GET", "/repos/:owner/:repo/events"},
	{"GET", "/networks/:owner/:repo/events"},
	{"GET", "/orgs/:org/events"},
	{"GET", "/users/:user/received_events"},
	{"GET", "/users/:user/received_events/public"},
	{"GET", "/users/:user/events"},
	{"GET", "/users/:user/events/public"},
	{"GET", "/users/:user/events/orgs/:org"},
	{"GET", "/feeds"},
	{"GET", "/notifications"},
	{"GET", "/repos/:owner/:repo/notifications"},
	{"PUT", "/notifications"},
	{"PUT", "/repos/:owner/:repo/notifications"},
	{"GET", "/notifications/threads/:id"},
	//{"PATCH", "/notifications/threads/:id"},
	{"GET", "/notifications/threads/:id/subscription"},
	{"PUT", "/notifications/threads/:id/subscription"},
	{"DELETE", "/notifications/threads/:id/subscription"},
	{"GET", "/repos/:owner/:repo/stargazers"},
	{"GET", "/users/:user/starred"},
	{"GET", "/user/starred"},
	{"GET", "/user/starred/:owner/:repo"},
	{"PUT", "/user/starred/:owner/:repo"},
	{"DELETE", "/user/starred/:owner/:repo"},
	{"GET", "/repos/:owner/:repo/subscribers"},
	{"GET", "/users/:user/subscriptions"},
	{"GET", "/user/subscriptions"},
	{"GET", "/repos/:owner/:repo/subscription"},
	{"PUT", "/repos/:owner/:repo/subscription"},
	{"DELETE", "/repos/:owner/:repo/subscription"},
	{"GET", "/user/subscriptions/:owner/:repo"},
	{"PUT", "/user/subscriptions/:owner/:repo"},
	{"DELETE", "/user/subscriptions/:owner/:repo"},

	// Gists
	{"GET", "/users/:user/gists"},
	{"GET", "/gists"},
	//{"GET", "/gists/public"},
	//{"GET", "/gists/starred"},
	{"GET", "/gists/:id"},
	{"POST", "/gists"},
	//{"PATCH", "/gists/:id"},
	{"PUT", "/gists/:id/star"},
	{"DELETE", "/gists/:id/star"},
	{"GET", "/gists/:id/star"},
	{"POST", "/gists/:id/forks"},
	{"DELETE", "/gists/:id"},

	// Git Data
	{"GET", "/repos/:owner/:repo/git/blobs/:sha"},
	{"POST", "/repos/:owner/:repo/git/blobs"},
	{"GET", "/repos/:owner/:repo/git/commits/:sha"},
	{"POST", "/repos/:owner/:repo/git/commits"},
	//{"GET", "/repos/:owner/:repo/git/refs/*ref"},
	{"GET", "/repos/:owner/:repo/git/refs"},
	{"POST", "/repos/:owner/:repo/git/refs"},
	//{"PATCH", "/repos/:owner/:repo/git/refs/*ref"},
	//{"DELETE", "/repos/:owner/:repo/git/refs/*ref"},
	{"GET", "/repos/:owner/:repo/git/tags/:sha"},
	{"POST", "/repos/:owner/:repo/git/tags"},
	{"GET", "/repos/:owner/:repo/git/trees/:sha"},
	{"POST", "/repos/:owner/:repo/git/trees"},

	// Issues
	{"GET", "/issues"},
	{"GET", "/user/issues"},
	{"GET", "/orgs/:org/issues"},
	{"GET", "/repos/:owner/:repo/issues"},
	{"GET", "/repos/:owner/:repo/issues/:number"},
	{"POST", "/repos/:owner/:repo/issues"},
	//{"PATCH", "/repos/:owner/:repo/issues/:number"},
	{"GET", "/repos/:owner/:repo/assignees"},
	{"GET", "/repos/:owner/:repo/assignees/:assignee"},
	{"GET", "/repos/:owner/:repo/issues/:number/comments"},
	//{"GET", "/repos/:owner/:repo/issues/comments"},
	//{"GET", "/repos/:owner/:repo/issues/comments/:id"},
	{"POST", "/repos/:owner/:repo/issues/:number/comments"},
	//{"PATCH", "/repos/:owner/:repo/issues/comments/:id"},
	//{"DELETE", "/repos/:owner/:repo/issues/comments/:id"},
	{"GET", "/repos/:owner/:repo/issues/:number/events"},
	//{"GET", "/repos/:owner/:repo/issues/events"},
	//{"GET", "/repos/:owner/:repo/issues/events/:id"},
	{"GET", "/repos/:owner/:repo/labels"},
	{"GET", "/repos/:owner/:repo/labels/:name"},
	{"POST", "/repos/:owner/:repo/labels"},
	//{"PATCH", "/repos/:owner/:repo/labels/:name"},
	{"DELETE", "/repos/:owner/:repo/labels/:name"},
	{"GET", "/repos/:owner/:repo/issues/:number/labels"},
	{"POST", "/repos/:owner/:repo/issues/:number/labels"},
	{"DELETE", "/repos/:owner/:repo/issues/:number/labels/:name"},
	{"PUT", "/repos/:owner/:repo/issues/:number/labels"},
	{"DELETE", "/repos/:owner/:repo/issues/:number/labels"},
	{"GET", "/repos/:owner/:repo/milestones/:number/labels"},
	{"GET", "/repos/:owner/:repo/milestones"},
	{"GET", "/repos/:owner/:repo/milestones/:number"},
	{"POST", "/repos/:owner/:repo/milestones"},
	//{"PATCH", "/repos/:owner/:repo/milestones/:number"},
	{"DELETE", "/repos/:owner/:repo/milestones/:number"},

	// Miscellaneous
	{"GET", "/emojis"},
	{"GET", "/gitignore/templates"},
	{"GET", "/gitignore/templates/:name"},
	{"POST", "/markdown"},
	{"POST", "/markdown/raw"},
	{"GET", "/meta"},
	{"GET", "/rate_limit"},

	// Organizations
	{"GET", "/users/:user/orgs"},
	{"GET", "/user/orgs"},
	{"GET", "/orgs/:org"},
	//{"PATCH", "/orgs/:org"},
	{"GET", "/orgs/:org/members"},
	{"GET", "/orgs/:org/members/:user"},
	{"DELETE", "/orgs/:org/members/:user"},
	{"GET", "/orgs/:org/public_members"},
	{"GET", "/orgs/:org/public_members/:user"},
	{"PUT", "/orgs/:org/public_members/:user"},
	{"DELETE", "/orgs/:org/public_members/:user"},
	{"GET", "/orgs/:org/teams"},
	{"GET", "/teams/:id"},
	{"POST", "/orgs/:org/teams"},
	//{"PATCH", "/teams/:id"},
	{"DELETE", "/teams/:id"},
	{"GET", "/teams/:id/members"},
	{"GET", "/teams/:id/members/:user"},
	{"PUT", "/teams/:id/members/:user"},
	{"DELETE", "/teams/:id/members/:user"},
	{"GET", "/teams/:id/repos"},
	{"GET", "/teams/:id/repos/:owner/:repo"},
	{"PUT", "/teams/:id/repos/:owner/:repo"},
	{"DELETE", "/teams/:id/repos/:owner/:repo"},
	{"GET", "/user/teams"},

	// Pull Requests
	{"GET", "/repos/:owner/:repo/pulls"},
	{"GET", "/repos/:owner/:repo/pulls/:number"},
	{"POST", "/repos/:owner/:repo/pulls"},
	//{"PATCH", "/repos/:owner/:repo/pulls/:number"},
	{"GET", "/repos/:owner/:repo/pulls/:number/commits"},
	{"GET", "/repos/:owner/:repo/pulls/:number/files"},
	{"GET", "/repos/:owner/:repo/pulls/:number/merge"},
	{"PUT", "/repos/:owner/:repo/pulls/:number/merge"},
	{"GET", "/repos/:owner/:repo/pulls/:number/comments"},
	//{"GET", "/repos/:owner/:repo/pulls/comments"},
	//{"GET", "/repos/:owner/:repo/pulls/comments/:number"},
	{"PUT", "/repos/:owner/:repo/pulls/:number/comments"},
	//{"PATCH", "/repos/:owner/:repo/pulls/comments/:number"},
	//{"DELETE", "/repos/:owner/:repo/pulls/comments/:number"},

	// Repositories
	{"GET", "/user/repos"},
	{"GET", "/users/:user/repos"},
	{"GET", "/orgs/:org/repos"},
	{"GET", "/repositories"},
	{"POST", "/user/repos"},
	{"POST", "/orgs/:org/repos"},
	{"GET", "/repos/:owner/:repo"},
	//{"PATCH", "/repos/:owner/:repo"},
	{"GET", "/repos/:owner/:repo/contributors"},
	{"GET", "/repos/:owner/:repo/languages"},
	{"GET", "/repos/:owner/:repo/teams"},
	{"GET", "/repos/:owner/:repo/tags"},
	{"GET", "/repos/:owner/:repo/branches"},
	{"GET", "/repos/:owner/:repo/branches/:branch"},
	{"DELETE", "/repos/:owner/:repo"},
	{"GET", "/repos/:owner/:repo/collaborators"},
	{"GET", "/repos/:owner/:repo/collaborators/:user"},
	{"PUT", "/repos/:owner/:repo/collaborators/:user"},
	{"DELETE", "/repos/:owner/:repo/collaborators/:user"},
	{"GET", "/repos/:owner/:repo/comments"},
	{"GET", "/repos/:owner/:repo/commits/:sha/comments"},
	{"POST", "/repos/:owner/:repo/commits/:sha/comments"},
	{"GET", "/repos/:owner/:repo/comments/:id"},
	//{"PATCH", "/repos/:owner/:repo/comments/:id"},
	{"DELETE", "/repos/:owner/:repo/comments/:id"},
	{"GET", "/repos/:owner/:repo/commits"},
	{"GET", "/repos/:owner/:repo/commits/:sha"},
	{"GET", "/repos/:owner/:repo/readme"},
	//{"GET", "/repos/:owner/:repo/contents/*path"},
	//{"PUT", "/repos/:owner/:repo/contents/*path"},
	//{"DELETE", "/repos/:owner/:repo/contents/*path"},
	//{"GET", "/repos/:owner/:repo/:archive_format/:ref"},
	{"GET", "/repos/:owner/:repo/keys"},
	{"GET", "/repos/:owner/:repo/keys/:id"},
	{"POST", "/repos/:owner/:repo/keys"},
	//{"PATCH", "/repos/:owner/:repo/keys/:id"},
	{"DELETE", "/repos/:owner/:repo/keys/:id"},
	{"GET", "/repos/:owner/:repo/downloads"},
	{"GET", "/repos/:owner/:repo/downloads/:id"},
	{"DELETE", "/repos/:owner/:repo/downloads/:id"},
	{"GET", "/repos/:owner/:repo/forks"},
	{"POST", "/repos/:owner/:repo/forks"},
	{"GET", "/repos/:owner/:repo/hooks"},
	{"GET", "/repos/:owner/:repo/hooks/:id"},
	{"POST", "/repos/:owner/:repo/hooks"},
	//{"PATCH", "/repos/:owner/:repo/hooks/:id"},
	{"POST", "/repos/:owner/:repo/hooks/:id/tests"},
	{"DELETE", "/repos/:owner/:repo/hooks/:id"},
	{"POST", "/repos/:owner/:repo/merges"},
	{"GET", "/repos/:owner/:repo/releases"},
	{"GET", "/repos/:owner/:repo/releases/:id"},
	{"POST", "/repos/:owner/:repo/releases"},
	//{"PATCH", "/repos/:owner/:repo/releases/:id"},
	{"DELETE", "/repos/:owner/:repo/releases/:id"},
	{"GET", "/repos/:owner/:repo/releases/:id/assets"},
	{"GET", "/repos/:owner/:repo/stats/contributors"},
	{"GET", "/repos/:owner/:repo/stats/commit_activity"},
	{"GET", "/repos/:owner/:repo/stats/code_frequency"},
	{"GET", "/repos/:owner/:repo/stats/participation"},
	{"GET", "/repos/:owner/:repo/stats/punch_card"},
	{"GET", "/repos/:owner/:repo/statuses/:ref"},
	{"POST", "/repos/:owner/:repo/statuses/:ref"},

	// Search
	{"GET", "/search/repositories"},
	{"GET", "/search/code"},
	{"GET", "/search/issues"},
	{"GET", "/search/users"},
	{"GET", "/legacy/issues/search/:owner/:repository/:state/:keyword"},
	{"GET", "/legacy/repos/search/:keyword"},
	{"GET", "/legacy/user/search/:keyword"},
	{"GET", "/legacy/user/email/:email"},

	// Users
	{"GET", "/users/:user"},
	{"GET", "/user"},
	//{"PATCH", "/user"},
	{"GET", "/users"},
	{"GET", "/user/emails"},
	{"POST", "/user/emails"},
	{"DELETE", "/user/emails"},
	{"GET", "/users/:user/followers"},
	{"GET", "/user/followers"},
	{"GET", "/users/:user/following"},
	{"GET", "/user/following"},
	{"GET", "/user/following/:user"},
	{"GET", "/users/:user/following/:target_user"},
	{"PUT", "/user/following/:user"},
	{"DELETE", "/user/following/:user"},
	{"GET", "/users/:user/keys"},
	{"GET", "/user/keys"},
	{"GET", "/user/keys/:id"},
	{"POST", "/user/keys"},
	//{"PATCH", "/user/keys/:id"},
	{"DELETE", "/user/keys/:id"},
}

func calcMem(name string, load func()) {
	m := new(runtime.MemStats)

	// before
	runtime.GC()
	runtime.ReadMemStats(m)
	before := m.HeapAlloc

	load()

	// after
	runtime.GC()
	runtime.ReadMemStats(m)
	after := m.HeapAlloc
	fmt.Printf("%15s:%10d bytes\n", name, after-before)
}

var (
	beegoMuxRouter http.Handler
	boneRouter     http.Handler
	chiRouter      http.Handler
	httpRouter     http.Handler
	goRouter1      http.Handler
	muxRouter      http.Handler
	trieMuxRouter  http.Handler
	cottonRouter   http.Handler
)

func init() {
	println("GithubAPI Routes:", len(githubAPI))

	calcMem("cottonRouter", func() {
		router := cotton.NewRouter()
		cotton.SetMode(cotton.ModeTest)

		handler := func(c *cotton.Context) {
			c.Response.WriteHeader(204)
		}
		for _, route := range githubAPI {
			if route.method == http.MethodGet {
				router.Get(route.path, handler)
			}
			if route.method == http.MethodPost {
				router.Post(route.path, handler)
			}
			if route.method == http.MethodPut {
				router.Put(route.path, handler)
			}
			if route.method == http.MethodDelete {
				router.Delete(route.path, handler)
			}
		}
		cottonRouter = router
	})

	calcMem("BeegoMuxRouter", func() {
		router := beegomux.New()
		handler := func(w http.ResponseWriter, _ *http.Request) {
			w.WriteHeader(204)
		}
		for _, route := range githubAPI {
			router.Handle(route.method, route.path, handler)
		}
		beegoMuxRouter = router
	})

	calcMem("BoneRouter", func() {
		router := bone.New()
		handler := func(w http.ResponseWriter, _ *http.Request) {
			w.WriteHeader(204)
		}
		for _, route := range githubAPI {
			if route.method == http.MethodGet {
				router.Get(route.path, http.HandlerFunc(handler))
			}
			if route.method == http.MethodPost {
				router.Post(route.path, http.HandlerFunc(handler))
			}
			if route.method == http.MethodPut {
				router.Put(route.path, http.HandlerFunc(handler))
			}
			if route.method == http.MethodDelete {
				router.Delete(route.path, http.HandlerFunc(handler))
			}
		}
		boneRouter = router
	})

	calcMem("ChiRouter", func() {
		router := chi.NewRouter()
		handler := func(w http.ResponseWriter, _ *http.Request) {
			w.WriteHeader(204)
		}
		for _, route := range githubAPI {
			router.MethodFunc(route.method, route.path, handler)
		}
		chiRouter = router
	})

	calcMem("HttpRouter", func() {
		router := httprouter.New()
		handler := func(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
			w.WriteHeader(204)
		}
		for _, route := range githubAPI {
			router.Handle(route.method, route.path, handler)
		}
		httpRouter = router

		// router.PrintTree(http.MethodGet)
	})

	calcMem("trie-mux", func() {
		router := triemux.New()
		handler := func(w http.ResponseWriter, _ *http.Request, _ triemux.Params) {
			w.WriteHeader(204)
		}
		for _, route := range githubAPI {
			router.Handle(route.method, route.path, handler)
		}
		trieMuxRouter = router
	})

	calcMem("GoRouter1", func() {
		router := gorouter.New()
		handler := func(w http.ResponseWriter, _ *http.Request) {
			w.WriteHeader(204)
		}
		for _, route := range githubAPI {
			router.Handle(route.method, route.path, handler)
		}
		goRouter1 = router
	})

	for i := 0; i < len(githubAPI); i++ {
		githubAPI[i].path = strings.ReplaceAll(githubAPI[i].path, ":", "val")
	}
}

// reference：https://github.com/julienschmidt/go-http-routing-benchmark/blob/2b136956a56bc65dddfa4bdaf7d1728ae2c90d50/bench_test.go#L76
func benchRoutes(b *testing.B, router http.Handler, routes []route) {
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/", nil)
	u := r.URL
	rq := u.RawQuery

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, route := range routes {
			r.Method = route.method
			r.RequestURI = route.path
			u.Path = route.path
			u.RawQuery = rq
			router.ServeHTTP(w, r)
		}
	}
}

// With GithubAPI (goRouter vs beegoMuxRouter vs BoneRouter vs httpRouter vs trieMuxRouter)

func BenchmarkHttpRouterWithGithubAPI(b *testing.B) {
	benchRoutes(b, httpRouter, githubAPI)
}
func BenchmarkCottonRouterWithGithubAPI(b *testing.B) {
	benchRoutes(b, cottonRouter, githubAPI)
}

func BenchmarkBeegoMuxRouterWithGithubAPI(b *testing.B) {
	benchRoutes(b, beegoMuxRouter, githubAPI)
}

func BenchmarkBoneRouterWithGithubAPI(b *testing.B) {
	benchRoutes(b, boneRouter, githubAPI)
}

func BenchmarkTrieMuxRouterWithGithubAPI(b *testing.B) {
	benchRoutes(b, trieMuxRouter, githubAPI)
}

func BenchmarkGoRouter1WithGithubAPI(b *testing.B) {
	benchRoutes(b, goRouter1, githubAPI)
}
