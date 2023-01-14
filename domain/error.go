package domain

type ErrorFormat struct {
	Code    int
	Message string
}

var (
	ErrorUnauthorized     = ErrorFormat{Code: 401, Message: "Unauthorized"}
	ErrorForbidden        = ErrorFormat{Code: 403, Message: "Forbidden"}
	ErrorAuthTokenExpired = ErrorFormat{Code: 4011, Message: "auth token expired"}
	ErrorInvalidToken     = ErrorFormat{Code: 4012, Message: "invalid token"}
	ErrorBadRequest       = ErrorFormat{Code: 400, Message: "bad request"}
	ErrorServer           = ErrorFormat{Code: 500, Message: "Server Error"}

	// admin
	ErrorPasswordRules = ErrorFormat{Code: 3001, Message: "password rules do not match"}

	// 前端
	ErrorPreOfficialMintIsMintAtSameTime      = ErrorFormat{Code: 6001, Message: "Error preOfficialMint is other user mint at the same time"}
	ErrorPreOfficialMintIsMinted              = ErrorFormat{Code: 6002, Message: "Error preOfficialMint is minted"}
	ErrorPreOfficialMintNFTNotFound           = ErrorFormat{Code: 6003, Message: "Error preOfficialMint NFT not found"}
	ErrorPreOfficialMintMintOnlyOneAtSameTime = ErrorFormat{Code: 6004, Message: "Error preOfficialMint mint only one at the same time"}
	ErrorPreOfficialMintIsOver                = ErrorFormat{Code: 6005, Message: "Error preOfficialMint is sale over"}
	ErrorFrontSignature                       = ErrorFormat{Code: 6006, Message: "Error signature"}
	ErrorFrontSignatureTimeout                = ErrorFormat{Code: 6007, Message: "Error signature timeout"}
	ErrorFrontUserBanned                      = ErrorFormat{Code: 6008, Message: "User have been banned"}
	ErrorOfficialMaterialNotFound             = ErrorFormat{Code: 6009, Message: "Error official material not found"}
	ErrorNFTNotFound                          = ErrorFormat{Code: 6010, Message: "Error NFT not found"}
	ErrorIsNotInWhiteList                     = ErrorFormat{Code: 6011, Message: "Error user is not in white list"}
	ErrorNftTodaySoldOut                      = ErrorFormat{Code: 6012, Message: "Error nft today sold out"}
	ErrorSuperRewordOtherUserReword           = ErrorFormat{Code: 6013, Message: "Error get reword other user reword"}
	ErrorUserSynchronizing                    = ErrorFormat{Code: 6014, Message: "Error user synchronizing"}
	ErrorNFTOwnerNotYou                       = ErrorFormat{Code: 6015, Message: "Error NFT owner not you"}
)
