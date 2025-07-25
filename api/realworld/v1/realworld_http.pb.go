// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.8.4
// - protoc             v3.19.4
// source: realworld/v1/realworld.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationRealWorldAddComment = "/realworld.v1.RealWorld/AddComment"
const OperationRealWorldCreateArticle = "/realworld.v1.RealWorld/CreateArticle"
const OperationRealWorldDeleteArticle = "/realworld.v1.RealWorld/DeleteArticle"
const OperationRealWorldDeleteComment = "/realworld.v1.RealWorld/DeleteComment"
const OperationRealWorldFavoriteArticle = "/realworld.v1.RealWorld/FavoriteArticle"
const OperationRealWorldFeedListArticles = "/realworld.v1.RealWorld/FeedListArticles"
const OperationRealWorldFollowUser = "/realworld.v1.RealWorld/FollowUser"
const OperationRealWorldGetArticle = "/realworld.v1.RealWorld/GetArticle"
const OperationRealWorldGetComment = "/realworld.v1.RealWorld/GetComment"
const OperationRealWorldGetCurrentUser = "/realworld.v1.RealWorld/GetCurrentUser"
const OperationRealWorldGetProfile = "/realworld.v1.RealWorld/GetProfile"
const OperationRealWorldGetTags = "/realworld.v1.RealWorld/GetTags"
const OperationRealWorldListArticles = "/realworld.v1.RealWorld/ListArticles"
const OperationRealWorldLogin = "/realworld.v1.RealWorld/Login"
const OperationRealWorldRegister = "/realworld.v1.RealWorld/Register"
const OperationRealWorldUnFavoriteArticle = "/realworld.v1.RealWorld/UnFavoriteArticle"
const OperationRealWorldUnFollowUser = "/realworld.v1.RealWorld/UnFollowUser"
const OperationRealWorldUpdateArticle = "/realworld.v1.RealWorld/UpdateArticle"
const OperationRealWorldUpdateUser = "/realworld.v1.RealWorld/UpdateUser"

type RealWorldHTTPServer interface {
	AddComment(context.Context, *AddCommentRequest) (*SingleCommentReply, error)
	CreateArticle(context.Context, *CreateArticleRequest) (*SingleArticleReply, error)
	DeleteArticle(context.Context, *DeleteArticleRequest) (*SingleArticleReply, error)
	DeleteComment(context.Context, *DeleteCommentRequest) (*SingleCommentReply, error)
	FavoriteArticle(context.Context, *FavoriteArticleRequest) (*SingleArticleReply, error)
	FeedListArticles(context.Context, *FeedArticlesRequest) (*MultipleArticlesReply, error)
	FollowUser(context.Context, *FollowUserRequest) (*ProfileReply, error)
	GetArticle(context.Context, *GetArticleRequest) (*SingleArticleReply, error)
	GetComment(context.Context, *GetCommentRequest) (*SingleCommentReply, error)
	GetCurrentUser(context.Context, *GetCurrentRequest) (*UserReply, error)
	GetProfile(context.Context, *GetProfileRequest) (*GetProfileReply, error)
	GetTags(context.Context, *GetTagsRequest) (*TagListReply, error)
	ListArticles(context.Context, *ListArticlesRequest) (*MultipleArticlesReply, error)
	Login(context.Context, *LoginRequest) (*LoginReply, error)
	Register(context.Context, *RegisterRequest) (*UserReply, error)
	UnFavoriteArticle(context.Context, *UnFavoriteArticleRequest) (*SingleArticleReply, error)
	UnFollowUser(context.Context, *UnFollowUserRequest) (*ProfileReply, error)
	UpdateArticle(context.Context, *UpdateArticleRequest) (*SingleArticleReply, error)
	UpdateUser(context.Context, *UpdateUserRequest) (*UserReply, error)
}

func RegisterRealWorldHTTPServer(s *http.Server, srv RealWorldHTTPServer) {
	r := s.Route("/")
	r.POST("/api/users/login", _RealWorld_Login0_HTTP_Handler(srv))
	r.POST("/api/users", _RealWorld_Register0_HTTP_Handler(srv))
	r.GET("/api/user", _RealWorld_GetCurrentUser0_HTTP_Handler(srv))
	r.PUT("/api/user", _RealWorld_UpdateUser0_HTTP_Handler(srv))
	r.GET("/api/profile/{username}", _RealWorld_GetProfile0_HTTP_Handler(srv))
	r.POST("/api/profile/{username}/follow", _RealWorld_FollowUser0_HTTP_Handler(srv))
	r.DELETE("/api/profile/{username}/follow", _RealWorld_UnFollowUser0_HTTP_Handler(srv))
	r.GET("/api/articles", _RealWorld_ListArticles0_HTTP_Handler(srv))
	r.GET("/api/articles/feed", _RealWorld_FeedListArticles0_HTTP_Handler(srv))
	r.GET("/api/article/{slug}", _RealWorld_GetArticle0_HTTP_Handler(srv))
	r.POST("/api/article", _RealWorld_CreateArticle0_HTTP_Handler(srv))
	r.PUT("/api/article/{slug}", _RealWorld_UpdateArticle0_HTTP_Handler(srv))
	r.DELETE("/api/article/{slug}", _RealWorld_DeleteArticle0_HTTP_Handler(srv))
	r.POST("/api/articles/{slug}/comments", _RealWorld_AddComment0_HTTP_Handler(srv))
	r.GET("/api/articles/{slug}/comments", _RealWorld_GetComment0_HTTP_Handler(srv))
	r.DELETE("/api/articles/{slug}/comments/{id}", _RealWorld_DeleteComment0_HTTP_Handler(srv))
	r.POST("/api/articles/{slug}/favorite", _RealWorld_FavoriteArticle0_HTTP_Handler(srv))
	r.DELETE("/api/articles/{slug}/favorite", _RealWorld_UnFavoriteArticle0_HTTP_Handler(srv))
	r.GET("/api/tags", _RealWorld_GetTags0_HTTP_Handler(srv))
}

func _RealWorld_Login0_HTTP_Handler(srv RealWorldHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in LoginRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationRealWorldLogin)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Login(ctx, req.(*LoginRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*LoginReply)
		return ctx.Result(200, reply)
	}
}

func _RealWorld_Register0_HTTP_Handler(srv RealWorldHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in RegisterRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationRealWorldRegister)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Register(ctx, req.(*RegisterRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UserReply)
		return ctx.Result(200, reply)
	}
}

func _RealWorld_GetCurrentUser0_HTTP_Handler(srv RealWorldHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetCurrentRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationRealWorldGetCurrentUser)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetCurrentUser(ctx, req.(*GetCurrentRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UserReply)
		return ctx.Result(200, reply)
	}
}

func _RealWorld_UpdateUser0_HTTP_Handler(srv RealWorldHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateUserRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationRealWorldUpdateUser)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateUser(ctx, req.(*UpdateUserRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UserReply)
		return ctx.Result(200, reply)
	}
}

func _RealWorld_GetProfile0_HTTP_Handler(srv RealWorldHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetProfileRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationRealWorldGetProfile)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetProfile(ctx, req.(*GetProfileRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetProfileReply)
		return ctx.Result(200, reply)
	}
}

func _RealWorld_FollowUser0_HTTP_Handler(srv RealWorldHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in FollowUserRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationRealWorldFollowUser)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.FollowUser(ctx, req.(*FollowUserRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ProfileReply)
		return ctx.Result(200, reply)
	}
}

func _RealWorld_UnFollowUser0_HTTP_Handler(srv RealWorldHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UnFollowUserRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationRealWorldUnFollowUser)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UnFollowUser(ctx, req.(*UnFollowUserRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ProfileReply)
		return ctx.Result(200, reply)
	}
}

func _RealWorld_ListArticles0_HTTP_Handler(srv RealWorldHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListArticlesRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationRealWorldListArticles)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListArticles(ctx, req.(*ListArticlesRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*MultipleArticlesReply)
		return ctx.Result(200, reply)
	}
}

func _RealWorld_FeedListArticles0_HTTP_Handler(srv RealWorldHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in FeedArticlesRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationRealWorldFeedListArticles)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.FeedListArticles(ctx, req.(*FeedArticlesRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*MultipleArticlesReply)
		return ctx.Result(200, reply)
	}
}

func _RealWorld_GetArticle0_HTTP_Handler(srv RealWorldHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetArticleRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationRealWorldGetArticle)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetArticle(ctx, req.(*GetArticleRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SingleArticleReply)
		return ctx.Result(200, reply)
	}
}

func _RealWorld_CreateArticle0_HTTP_Handler(srv RealWorldHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateArticleRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationRealWorldCreateArticle)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateArticle(ctx, req.(*CreateArticleRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SingleArticleReply)
		return ctx.Result(200, reply)
	}
}

func _RealWorld_UpdateArticle0_HTTP_Handler(srv RealWorldHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateArticleRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationRealWorldUpdateArticle)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateArticle(ctx, req.(*UpdateArticleRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SingleArticleReply)
		return ctx.Result(200, reply)
	}
}

func _RealWorld_DeleteArticle0_HTTP_Handler(srv RealWorldHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteArticleRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationRealWorldDeleteArticle)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteArticle(ctx, req.(*DeleteArticleRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SingleArticleReply)
		return ctx.Result(200, reply)
	}
}

func _RealWorld_AddComment0_HTTP_Handler(srv RealWorldHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in AddCommentRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationRealWorldAddComment)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.AddComment(ctx, req.(*AddCommentRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SingleCommentReply)
		return ctx.Result(200, reply)
	}
}

func _RealWorld_GetComment0_HTTP_Handler(srv RealWorldHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetCommentRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationRealWorldGetComment)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetComment(ctx, req.(*GetCommentRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SingleCommentReply)
		return ctx.Result(200, reply)
	}
}

func _RealWorld_DeleteComment0_HTTP_Handler(srv RealWorldHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteCommentRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationRealWorldDeleteComment)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteComment(ctx, req.(*DeleteCommentRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SingleCommentReply)
		return ctx.Result(200, reply)
	}
}

func _RealWorld_FavoriteArticle0_HTTP_Handler(srv RealWorldHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in FavoriteArticleRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationRealWorldFavoriteArticle)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.FavoriteArticle(ctx, req.(*FavoriteArticleRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SingleArticleReply)
		return ctx.Result(200, reply)
	}
}

func _RealWorld_UnFavoriteArticle0_HTTP_Handler(srv RealWorldHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UnFavoriteArticleRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationRealWorldUnFavoriteArticle)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UnFavoriteArticle(ctx, req.(*UnFavoriteArticleRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SingleArticleReply)
		return ctx.Result(200, reply)
	}
}

func _RealWorld_GetTags0_HTTP_Handler(srv RealWorldHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetTagsRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationRealWorldGetTags)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetTags(ctx, req.(*GetTagsRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*TagListReply)
		return ctx.Result(200, reply)
	}
}

type RealWorldHTTPClient interface {
	AddComment(ctx context.Context, req *AddCommentRequest, opts ...http.CallOption) (rsp *SingleCommentReply, err error)
	CreateArticle(ctx context.Context, req *CreateArticleRequest, opts ...http.CallOption) (rsp *SingleArticleReply, err error)
	DeleteArticle(ctx context.Context, req *DeleteArticleRequest, opts ...http.CallOption) (rsp *SingleArticleReply, err error)
	DeleteComment(ctx context.Context, req *DeleteCommentRequest, opts ...http.CallOption) (rsp *SingleCommentReply, err error)
	FavoriteArticle(ctx context.Context, req *FavoriteArticleRequest, opts ...http.CallOption) (rsp *SingleArticleReply, err error)
	FeedListArticles(ctx context.Context, req *FeedArticlesRequest, opts ...http.CallOption) (rsp *MultipleArticlesReply, err error)
	FollowUser(ctx context.Context, req *FollowUserRequest, opts ...http.CallOption) (rsp *ProfileReply, err error)
	GetArticle(ctx context.Context, req *GetArticleRequest, opts ...http.CallOption) (rsp *SingleArticleReply, err error)
	GetComment(ctx context.Context, req *GetCommentRequest, opts ...http.CallOption) (rsp *SingleCommentReply, err error)
	GetCurrentUser(ctx context.Context, req *GetCurrentRequest, opts ...http.CallOption) (rsp *UserReply, err error)
	GetProfile(ctx context.Context, req *GetProfileRequest, opts ...http.CallOption) (rsp *GetProfileReply, err error)
	GetTags(ctx context.Context, req *GetTagsRequest, opts ...http.CallOption) (rsp *TagListReply, err error)
	ListArticles(ctx context.Context, req *ListArticlesRequest, opts ...http.CallOption) (rsp *MultipleArticlesReply, err error)
	Login(ctx context.Context, req *LoginRequest, opts ...http.CallOption) (rsp *LoginReply, err error)
	Register(ctx context.Context, req *RegisterRequest, opts ...http.CallOption) (rsp *UserReply, err error)
	UnFavoriteArticle(ctx context.Context, req *UnFavoriteArticleRequest, opts ...http.CallOption) (rsp *SingleArticleReply, err error)
	UnFollowUser(ctx context.Context, req *UnFollowUserRequest, opts ...http.CallOption) (rsp *ProfileReply, err error)
	UpdateArticle(ctx context.Context, req *UpdateArticleRequest, opts ...http.CallOption) (rsp *SingleArticleReply, err error)
	UpdateUser(ctx context.Context, req *UpdateUserRequest, opts ...http.CallOption) (rsp *UserReply, err error)
}

type RealWorldHTTPClientImpl struct {
	cc *http.Client
}

func NewRealWorldHTTPClient(client *http.Client) RealWorldHTTPClient {
	return &RealWorldHTTPClientImpl{client}
}

func (c *RealWorldHTTPClientImpl) AddComment(ctx context.Context, in *AddCommentRequest, opts ...http.CallOption) (*SingleCommentReply, error) {
	var out SingleCommentReply
	pattern := "/api/articles/{slug}/comments"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationRealWorldAddComment))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *RealWorldHTTPClientImpl) CreateArticle(ctx context.Context, in *CreateArticleRequest, opts ...http.CallOption) (*SingleArticleReply, error) {
	var out SingleArticleReply
	pattern := "/api/article"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationRealWorldCreateArticle))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *RealWorldHTTPClientImpl) DeleteArticle(ctx context.Context, in *DeleteArticleRequest, opts ...http.CallOption) (*SingleArticleReply, error) {
	var out SingleArticleReply
	pattern := "/api/article/{slug}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationRealWorldDeleteArticle))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *RealWorldHTTPClientImpl) DeleteComment(ctx context.Context, in *DeleteCommentRequest, opts ...http.CallOption) (*SingleCommentReply, error) {
	var out SingleCommentReply
	pattern := "/api/articles/{slug}/comments/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationRealWorldDeleteComment))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *RealWorldHTTPClientImpl) FavoriteArticle(ctx context.Context, in *FavoriteArticleRequest, opts ...http.CallOption) (*SingleArticleReply, error) {
	var out SingleArticleReply
	pattern := "/api/articles/{slug}/favorite"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationRealWorldFavoriteArticle))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *RealWorldHTTPClientImpl) FeedListArticles(ctx context.Context, in *FeedArticlesRequest, opts ...http.CallOption) (*MultipleArticlesReply, error) {
	var out MultipleArticlesReply
	pattern := "/api/articles/feed"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationRealWorldFeedListArticles))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *RealWorldHTTPClientImpl) FollowUser(ctx context.Context, in *FollowUserRequest, opts ...http.CallOption) (*ProfileReply, error) {
	var out ProfileReply
	pattern := "/api/profile/{username}/follow"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationRealWorldFollowUser))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *RealWorldHTTPClientImpl) GetArticle(ctx context.Context, in *GetArticleRequest, opts ...http.CallOption) (*SingleArticleReply, error) {
	var out SingleArticleReply
	pattern := "/api/article/{slug}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationRealWorldGetArticle))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *RealWorldHTTPClientImpl) GetComment(ctx context.Context, in *GetCommentRequest, opts ...http.CallOption) (*SingleCommentReply, error) {
	var out SingleCommentReply
	pattern := "/api/articles/{slug}/comments"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationRealWorldGetComment))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *RealWorldHTTPClientImpl) GetCurrentUser(ctx context.Context, in *GetCurrentRequest, opts ...http.CallOption) (*UserReply, error) {
	var out UserReply
	pattern := "/api/user"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationRealWorldGetCurrentUser))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *RealWorldHTTPClientImpl) GetProfile(ctx context.Context, in *GetProfileRequest, opts ...http.CallOption) (*GetProfileReply, error) {
	var out GetProfileReply
	pattern := "/api/profile/{username}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationRealWorldGetProfile))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *RealWorldHTTPClientImpl) GetTags(ctx context.Context, in *GetTagsRequest, opts ...http.CallOption) (*TagListReply, error) {
	var out TagListReply
	pattern := "/api/tags"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationRealWorldGetTags))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *RealWorldHTTPClientImpl) ListArticles(ctx context.Context, in *ListArticlesRequest, opts ...http.CallOption) (*MultipleArticlesReply, error) {
	var out MultipleArticlesReply
	pattern := "/api/articles"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationRealWorldListArticles))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *RealWorldHTTPClientImpl) Login(ctx context.Context, in *LoginRequest, opts ...http.CallOption) (*LoginReply, error) {
	var out LoginReply
	pattern := "/api/users/login"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationRealWorldLogin))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *RealWorldHTTPClientImpl) Register(ctx context.Context, in *RegisterRequest, opts ...http.CallOption) (*UserReply, error) {
	var out UserReply
	pattern := "/api/users"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationRealWorldRegister))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *RealWorldHTTPClientImpl) UnFavoriteArticle(ctx context.Context, in *UnFavoriteArticleRequest, opts ...http.CallOption) (*SingleArticleReply, error) {
	var out SingleArticleReply
	pattern := "/api/articles/{slug}/favorite"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationRealWorldUnFavoriteArticle))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *RealWorldHTTPClientImpl) UnFollowUser(ctx context.Context, in *UnFollowUserRequest, opts ...http.CallOption) (*ProfileReply, error) {
	var out ProfileReply
	pattern := "/api/profile/{username}/follow"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationRealWorldUnFollowUser))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *RealWorldHTTPClientImpl) UpdateArticle(ctx context.Context, in *UpdateArticleRequest, opts ...http.CallOption) (*SingleArticleReply, error) {
	var out SingleArticleReply
	pattern := "/api/article/{slug}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationRealWorldUpdateArticle))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *RealWorldHTTPClientImpl) UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...http.CallOption) (*UserReply, error) {
	var out UserReply
	pattern := "/api/user"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationRealWorldUpdateUser))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
