package user_controller

import (
	"fmt"
	"net/http"
	"server_crm/internal/auxiliary"
)

// Auth Handler for login
// @Summary Login
// @Description Login with email and password
// @Tags User
// @Accept  json
// @Produce  json
// @Param dto body rest_models.LoginReqDto true "Login data"
// @Success 200  "OK"
// @Failure 500  "Error"
// @Router /user/login [post]
func (h UserController) Update(w http.ResponseWriter, r *http.Request) {

	// var dto rest_models.UserUpdateReqDto

	fmt.Println(auxiliary.GetUserInfo(r.Context()))
}
