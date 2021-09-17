package errors

import "errors"

var(
	ErrUnableToSave  = errors.New("unable to save data")
	ErrForgotEmail = errors.New("email is required")
 	ErrUnableToDelete  = errors.New("unable to delete data")
	ErrUnableToFetch  = errors.New("unable to fetch data")
 	ErrIDNotFound  = errors.New("id not found ")
	ErrInvalidRequest                 = errors.New("invalid_request")
	ErrUnauthorizedClient             = errors.New("unauthorized_client")
	ErrAccessDenied                   = errors.New("access_denied")
	ErrUnsupportedResponseType        = errors.New("unsupported_response_type")
	ErrInvalidScope                   = errors.New("invalid_scope")
	ErrServerError                    = errors.New("server_error")
	ErrTemporarilyUnavailable         = errors.New("temporarily_unavailable")
	ErrInvalidClient                  = errors.New("invalid_client")
	ErrInvalidGrant                   = errors.New("invalid_grant")
	ErrUnsupportedGrantType           = errors.New("unsupported_grant_type")
	ErrCodeChallengeRquired           = errors.New("invalid_request")
	ErrUnsupportedCodeChallengeMethod = errors.New("invalid_request")
	ErrInvalidCodeChallengeLen        = errors.New("invalid_request")

	ErrInvalidRedirectURI   = errors.New("invalid redirect uri")
	ErrInvalidAuthorizeCode = errors.New("invalid authorize code")
	ErrInvalidAccessToken   = errors.New("invalid access token")
	ErrInvalidRefreshToken  = errors.New("invalid refresh token")
	ErrExpiredAccessToken   = errors.New("expired access token")
	ErrExpiredRefreshToken  = errors.New("expired refresh token")
	ErrMissingCodeVerifier  = errors.New("missing code verifier")
	ErrMissingCodeChallenge = errors.New("missing code challenge")
	ErrInvalidCodeChallenge = errors.New("invalid code challenge")
)
// Descriptions error description
var Descriptions = map[error]string{
	ErrForgotEmail:"Email is forgotten",
	ErrIDNotFound:"Id not found",
	ErrUnableToSave:"unable to save",
	ErrUnableToDelete:"unable to delete",
	ErrUnableToFetch:"unanble to fetch",
	ErrInvalidRequest:                 "The request is missing a required parameter, includes an invalid parameter value, includes a parameter more than once, or is otherwise malformed",
	ErrUnauthorizedClient:             "The client is not authorized to request an authorization code using this method",
	ErrAccessDenied:                   "The resource owner or authorization server denied the request",
	ErrUnsupportedResponseType:        "The authorization server does not support obtaining an authorization code using this method",
	ErrInvalidScope:                   "The requested scope is invalid, unknown, or malformed",
	ErrServerError:                    "The authorization server encountered an unexpected condition that prevented it from fulfilling the request",
	ErrTemporarilyUnavailable:         "The authorization server is currently unable to handle the request due to a temporary overloading or maintenance of the server",
	ErrInvalidClient:                  "Client authentication failed",
	ErrInvalidGrant:                   "The provided authorization grant (e.g., authorization code, resource owner credentials) or refresh token is invalid, expired, revoked, does not match the redirection URI used in the authorization request, or was issued to another client",
	ErrUnsupportedGrantType:           "The authorization grant type is not supported by the authorization server",
	ErrCodeChallengeRquired:           "PKCE is required. code_challenge is missing",
	ErrUnsupportedCodeChallengeMethod: "Selected code_challenge_method not supported",
	ErrInvalidCodeChallengeLen:        "Code challenge length must be between 43 and 128 charachters long",
}

// StatusCodes response error HTTP status code
var StatusCodes = map[error]int{

	ErrInvalidRequest:                 400,
	ErrUnauthorizedClient:             401,
	ErrAccessDenied:                   403,
	ErrUnsupportedResponseType:        401,
	ErrInvalidScope:                   400,
	ErrServerError:                    500,
	ErrTemporarilyUnavailable:         503,
	ErrInvalidClient:                  401,
	ErrInvalidGrant:                   401,
	ErrUnsupportedGrantType:           401,
	ErrCodeChallengeRquired:           400,
	ErrUnsupportedCodeChallengeMethod: 400,
	ErrInvalidCodeChallengeLen:        400,
	ErrIDNotFound:					   404,
	ErrForgotEmail:                    422,
	ErrUnableToSave:				   422,
	ErrUnableToDelete:				   422,
	ErrUnableToFetch:				   422,
}

// StatusCodes response error HTTP status code
var ErrCodes = map[error]int{
	ErrInvalidRequest:                 4000,
	ErrUnauthorizedClient:             4001,
	ErrAccessDenied:                   4002,
	ErrUnsupportedResponseType:        4003,
	ErrInvalidScope:                   4004,
	ErrServerError:                    4005,
	ErrTemporarilyUnavailable:         4006,
	ErrInvalidClient:                  4007,
	ErrInvalidGrant:                   4008,
	ErrUnsupportedGrantType:           4009,
	ErrCodeChallengeRquired:           4010,
	ErrUnsupportedCodeChallengeMethod: 4011,
	ErrInvalidCodeChallengeLen:        4012,
	ErrIDNotFound:					   4013,
	ErrUnableToSave:				   4014,
	ErrUnableToDelete:				   4015,
	ErrUnableToFetch:				   4016,

}
