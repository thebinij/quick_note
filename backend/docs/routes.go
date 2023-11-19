package docs

import "quick-note/internal/user"

// swagger:route POST /login payloads PostLogin
// Post credentials.
// responses:
//   200: PostLoginRes200

// swagger:parameters PostLogin
type PostLoginReq struct {
	// in:body
	Body user.LoginUserReq
}

// swagger:response PostLoginRes200
type PostLoginRes200 struct {
	// in:body
	Body user.LoginUserRes
}

// swagger:route POST /signup payloads CreateUser
// Post credentials.
// responses:
//   200: CreateUserRes200

// swagger:parameters CreateUser
type CreateUserReq struct {
	// in:body
	Body user.CreateUserReq
}

// swagger:response CreateUserRes200
type CreateUserRes200 struct {
	// in:body
	Body user.CreateUserRes
}
