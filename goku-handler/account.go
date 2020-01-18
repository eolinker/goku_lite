package goku_handler

import (
	"net/http"
)

//Account 账号处理器
type Account interface {
	CheckLogin(r *http.Request) (int, error)
	CheckPermission(pre string, isEdit bool, userID int) (bool, error)
}

//AccountHandler 账号处理器
type AccountHandler struct {
	account    Account
	handler    http.Handler
	permission string
	isEdit     bool
}

func (h *AccountHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// 检查登录操作
	userID, err := h.account.CheckLogin(r)
	if err != nil {
		WriteError(w, "100001", "user", err.Error(), err)
		return
	}

	// 检查权限操作
	if _, err := h.account.CheckPermission(h.permission, h.isEdit, userID); err != nil {
		WriteError(w, "100002", "user", err.Error(), err)
		return
	}

	r = SetUserIDToRequest(r, userID)
	h.handler.ServeHTTP(w, r)
}

//AccountHandlerFactory 账号处理工厂
type AccountHandlerFactory struct {
	account Account
}

//NewAccountHandlerFactory new AccountHandlerFactory
func NewAccountHandlerFactory(account Account) *AccountHandlerFactory {
	return &AccountHandlerFactory{account: account}
}

//NewAccountHandler new accountHandler
func (f *AccountHandlerFactory) NewAccountHandler(permission string, isEdit bool, handler http.Handler) http.Handler {
	return &AccountHandler{
		account:    f.account,
		handler:    handler,
		permission: permission,
		isEdit:     isEdit,
	}
}

//NewAccountHandleFunction new accountHandleFunction
func (f *AccountHandlerFactory) NewAccountHandleFunction(permission string, isEdit bool, handleFunc func(w http.ResponseWriter, r *http.Request)) http.Handler {
	return f.NewAccountHandler(permission, isEdit, http.HandlerFunc(handleFunc))
}
