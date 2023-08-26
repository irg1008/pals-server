package auth

import (
	"irg1008/pals/ent/authrequest"
	"irg1008/pals/ent/user"
	"irg1008/pals/pkg/db"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

var typesOfRequest = []authrequest.Type{
	authrequest.TypeConfirmEmail,
	authrequest.TypeResetPassword,
}

func TestMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}

func setupService(t *testing.T) *AuthService {
	mockDB := db.GetMockedDB(t)
	return &AuthService{mockDB}
}

func TestInitService(t *testing.T) {
	s := setupService(t)

	if s.DB == nil {
		t.Error("DB is nil")
	}

	if s.DB.Ctx == nil {
		t.Error("DB context is nil")
	}
}

func Test_SignUp(t *testing.T) {
	s := setupService(t)
	defer s.DB.Close()

	email, password := "first@user.com", "pwd@1234"

	// Try inserting a simple user.
	insertedUser, err := s.CreateUser(email, password)

	// Must be unconfirmed by default.
	if insertedUser.IsConfirmed {
		t.Error("Error, user is confirmed by default")
	}

	// Password must have been hashed.
	if insertedUser.Password == password {
		t.Error("Error, password was not hashed")
	}

	// Check is actually hashed.
	if err := bcrypt.CompareHashAndPassword([]byte(insertedUser.Password), []byte(password)); err != nil {
		t.Error("Error, password was not hashed correctly")
	}

	// Try inserting the same user again. Should fail.
	_, err = s.CreateUser(email, password)
	if err == nil {
		t.Error("Error, duplicated user was inserted")
	}

	overThresholdPassword := strings.Repeat("@", 500)
	_, err = s.CreateUser(email, overThresholdPassword)
	if err == nil {
		t.Error("Error, new user with password over threshold")
	}
}

func Test_Login(t *testing.T) {
	s := setupService(t)
	defer s.DB.Close()

	email, password := "notnew@user.com", "pwd@1234"

	_, err := s.CreateUser(email, password)

	// Get user from DB with email.
	user, err := s.GetUserByEmail(email)
	if err != nil {
		t.Error("Error getting user by email", err)
	}

	// Get user with ID.
	_, err = s.GetUserById(user.ID)
	if err != nil {
		t.Error("Error getting user by ID", err)
	}
}

func Test_CreateConfirmationRequest(t *testing.T) {
	s := setupService(t)
	defer s.DB.Close()

	email, password := "notconfirmed@user.com", "pwd@1234"
	s.CreateUser(email, password)

	// Create confirmation request.
	request, err := s.CreateConfirmationRequest(email)
	if err != nil {
		t.Error("Error creating confirmation request", err)
	}

	// When creating a new request, the previous must be invalidated.
	newRequest, err := s.CreateConfirmationRequest(email)
	invalidated, err := s.DB.AuthRequest.Query().
		Where(authrequest.ID(request.ID)).
		Select(authrequest.FieldActive).
		Only(s.DB.Ctx)

	if err != nil {
		t.Error("Error getting previously invalidated request", err)
	}

	if invalidated.Active {
		t.Error("Error, previous request was not invalidated")
	}

	// User must be inactive until confirmed.
	sameUser, err := s.DB.User.Query().
		Where(user.ID(newRequest.UserID)).
		Select(user.FieldIsConfirmed).
		Only(s.DB.Ctx)

	if err != nil {
		t.Error("Error getting user", err)
	}

	if sameUser.IsConfirmed {
		t.Error("Error, user is confirmed before the confirm request is confirmed")
	}
}

func Test_ConfirmEmail(t *testing.T) {
	s := setupService(t)
	defer s.DB.Close()

	email, password := "notconfirmed@user.com", "pwd@1234"
	user, _ := s.CreateUser(email, password)

	request, err := s.CreateConfirmationRequest(email)
	if err != nil {
		t.Error("Error creating confirmation request", err)
	}

	// We send email or any other notification with token.
	token := request.Token.String()
	confirmedUser, _ := s.ConfirmEmail(token)

	// User must be confirmed.
	if !confirmedUser.IsConfirmed {
		t.Error("Error, user is not confirmed")
	}

	// Refetch the user from DB.
	user, err = s.GetUserByEmail(email)
	if err != nil {
		t.Error("Error getting user by email", err)
	}

	// User must be confirmed.
	if !user.IsConfirmed {
		t.Error("Error, user is not confirmed")
	}

	// Try to confirm with invalid token
	invalidToken := "iamhackermestealyourdatahaha"
	_, err = s.ConfirmEmail(invalidToken)
	if err == nil {
		t.Error("Error confirming user for valid request", err)
	}
}

func Test_CreatePasswordResetRequest(t *testing.T) {
	s := setupService(t)
	defer s.DB.Close()

	email, password := "user@pwdreset.com", "pwd@1234"
	s.CreateUser(email, password)

	// Create password reset request.
	request, err := s.CreatePasswordResetRequest(email)
	if err != nil {
		t.Error("Error creating password reset request")
	}

	_, err = s.CreatePasswordResetRequest(email)
	if err != nil {
		t.Error("Error creating password reset request", err)
	}

	// When creating a new request, the previous requests must keep the active state.
	invalidated, _ := s.DB.AuthRequest.Query().
		Where(authrequest.ID(request.ID)).
		Select(authrequest.FieldActive).
		Only(s.DB.Ctx)

	if !invalidated.Active {
		t.Error("Error, previous request was invalidated. Password reset requests should not invalidate previous requests")
	}
}

func Test_ResetPassword(t *testing.T) {
	s := setupService(t)
	defer s.DB.Close()

	email, password := "user@confirmpwsreset.com", "pwd@1234"
	s.CreateUser(email, password)

	// We send email or any other notification with token.
	resetRequest, _ := s.CreatePasswordResetRequest(email)
	token := resetRequest.Token.String()

	// Must give an error wheen hashing very long password.
	overThresholdPassword := strings.Repeat("#", 500)
	_, err := s.ResetPassword(token, overThresholdPassword)
	if err == nil {
		t.Error("Error, reset password with password over threshold")
	}

	// Must correctly set new password.
	newPassword := "newpwd@1234"
	userWithNewPassword, err := s.ResetPassword(token, newPassword)

	if err != nil {
		t.Error("Error resetting password for valid request", err)
	}

	// New hashed password must be set.
	if err := bcrypt.CompareHashAndPassword([]byte(userWithNewPassword.Password), []byte(newPassword)); err != nil {
		t.Error("Error, password was not hashed correctly")
	}

	// Try to reset password with invalid token.
	invalidToken := "iamgoingtohackyouwiththishehe"
	_, err = s.ResetPassword(invalidToken, newPassword)
	if err == nil {
		t.Error("Error, reset password for invalid token")
	}
}

func Test_CreateRequest(t *testing.T) {
	s := setupService(t)
	defer s.DB.Close()

	email, password := "mewanna@request.com", "pwd@1234"
	user, _ := s.CreateUser(email, password)

	for _, requestType := range typesOfRequest {
		_, err := s.createRequest("idont@exist.com", requestType, CreateRequestOptions{})
		if err == nil {
			t.Error("Error, created request for non-existing user")
		}

		request, err := s.createRequest(email, requestType, CreateRequestOptions{})
		if err != nil {
			t.Error("Error creating confirmation request", err)
		}

		if request.Type != requestType {
			t.Error("Error, created request has wrong type")
		}

		if request.UserID != user.ID {
			t.Error("Error, created request has wrong user")
		}

		// Request must be active.
		if !request.Active {
			t.Error("Error, created request is not active")
		}

		// Expiry time must be set in the future.
		if request.ExpiresAt.Before(time.Now()) {
			t.Error("Error, created request has wrong expiry time")
		}

		// Create another X requests for the same user.
		newRequestCount := 10
		newRequestsIds := make([]int, newRequestCount)
		for i := 0; i < newRequestCount; i++ {
			req, err := s.createRequest(email, requestType, CreateRequestOptions{})
			if err != nil {
				t.Error("Error creating additional requests", err)
			}
			newRequestsIds[i] = req.ID
		}

		// Create a new request with invalidate rest option.
		newValidReq, err := s.createRequest(email, requestType, CreateRequestOptions{invalidateRest: true})
		invalidRequests, err := s.DB.AuthRequest.Query().Where(
			authrequest.IDIn(newRequestsIds...),
			authrequest.TypeEQ(requestType),
		).All(s.DB.Ctx)

		if err != nil {
			t.Error("Error getting previously invalidated request", err)
		}

		for _, req := range invalidRequests {
			if req.Active {
				t.Error("Error, previous request was not invalidated", req, newValidReq)
			}
		}
	}
}

func Test_ConfirmRequest(t *testing.T) {
	s := setupService(t)
	defer s.DB.Close()

	email, password := "user@confirmrequest.com", "pwd@1234"
	s.CreateUser(email, password)

	for _, requestType := range typesOfRequest {

		// Should not confirm on invalid token.
		invalidToken := "comeonareyouseriousthisisnotavalidtoken"
		_, err := s.confirmRequest(invalidToken, requestType, ConfirmRequestOptions{})
		if err == nil {
			t.Error("Error, confirmed request with invalid token")
		}

		// Valid token but no request found.
		_, err = s.confirmRequest(uuid.New().String(), requestType, ConfirmRequestOptions{})
		if err == nil {
			t.Error("Error, found request for non existing token")
		}

		request, _ := s.createRequest(email, requestType, CreateRequestOptions{})
		requestUser, err := s.confirmRequest(request.Token.String(), requestType, ConfirmRequestOptions{})
		if err != nil {
			t.Error("Request not found for provided token", err)
		}

		// Request should be invalidated.
		sameRequest, _ := s.DB.AuthRequest.Query().Where(
			authrequest.ID(request.ID),
			authrequest.TypeEQ(requestType),
		).Only(s.DB.Ctx)

		if sameRequest.Active {
			t.Error("Error, request was not invalidated on confirmation")
		}

		// Request should return valid user
		if request.UserID == 0 {
			t.Error("Error, request has no user")
		}

		if requestUser.Email != email {
			t.Error("Error, request has wrong user")
		}

		// Now invalid requests confirmed
		inactiveReq, _ := s.createRequest(email, requestType, CreateRequestOptions{})
		inactiveReq.Update().SetActive(false).Exec(s.DB.Ctx)

		_, err = s.confirmRequest(inactiveReq.Token.String(), requestType, ConfirmRequestOptions{})
		if err == nil {
			t.Error("Error, confirmed request that was inactive")
		}

		// No expired requests confirmed
		lastYear := time.Now().Add(-365 * 24 * time.Hour)
		expiredReq, _ := s.createRequest(email, requestType, CreateRequestOptions{})
		expiredReq.Update().SetExpiresAt(lastYear).Exec(s.DB.Ctx)

		_, err = s.confirmRequest(expiredReq.Token.String(), requestType, ConfirmRequestOptions{})
		if err == nil {
			t.Error("Error, confirmed request that was expired")
		}

		// If option to delete all is set, the created request must be deleted after confirmation
		toBeDeleted, _ := s.createRequest(email, requestType, CreateRequestOptions{})
		_, err = s.confirmRequest(toBeDeleted.Token.String(), requestType, ConfirmRequestOptions{deleteAll: true})

		if err != nil {
			t.Error("Error, could not confirm to-be-deleted request", err)
		}

		deletedReq, err := s.DB.AuthRequest.Query().Where(authrequest.ID(toBeDeleted.ID)).Only(s.DB.Ctx)
		if err == nil {
			t.Error("Error, request was not deleted", deletedReq)
		}
	}

	// Try to create request for invalid type
	_, err := s.createRequest(email, "invalid-type", CreateRequestOptions{})
	if err == nil {
		t.Error("Error, created request for invalid type")
	}
}

func Test_DeleteRequests(t *testing.T) {
	s := setupService(t)
	defer s.DB.Close()

	email, password := "user@deleterequests.com", "pwd@1234"
	user, _ := s.CreateUser(email, password)

	for _, requestType := range typesOfRequest {

		// Create X requests for the same user.
		for i := 0; i < 10; i++ {
			_, err := s.createRequest(email, requestType, CreateRequestOptions{})
			if err != nil {
				t.Error("Error creating additional requests", err)
			}
		}

		// Delete all requests for the same user.
		s.deleteRequests(user.ID, requestType)

		// Check that all requests are deleted.
		deletedRequests, err := s.DB.AuthRequest.Query().Where(
			authrequest.UserID(user.ID),
			authrequest.TypeEQ(requestType),
		).Count(s.DB.Ctx)

		if err != nil {
			t.Error("Error getting deleted requests", err)
		}

		if deletedRequests != 0 {
			t.Error("Error, requests were not deleted", deletedRequests)
		}
	}
}
