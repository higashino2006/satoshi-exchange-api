package controllers_test

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"se-api/src/internal/lib/common"
	"se-api/src/internal/services"
	test_lib "se-api/src/internal/tests/lib"
	"strings"
	"testing"
)

func TestBuyCrypto(t *testing.T) {
	// truncate database
	test_lib.TruncateAllTables()

	// test user id and jpy balance
	testUserID := "test_user1"
	initialJPYBalance := 10000

	// prepare service
	userService := services.NewUserService()

	// prepare test user
	err := prepareTestUser(userService, testUserID, initialJPYBalance)
	if err != nil {
		t.Fatalf("Failed to prepare test user: %v", err)
	}

	// request and expected response
	requestBody := `{"satoshi": 10000}`
	expectedResponseBody := `{"satoshi_balance": 10000, "jpy_balance": 7000}`

	// send post request and get response
	resp, err := sendBuyCryptoRequest(testUserID, requestBody)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	// check status code
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Expected status 200 OK, got %v", resp.Status)
	}

	// read response
	actualResponseBody, err := readResponseBody(resp.Body)
	if err != nil {
		t.Fatalf("Failed to read response body: %v", err)
	}

	// check response
	if strings.TrimSpace(actualResponseBody) != strings.TrimSpace(expectedResponseBody) {
		t.Fatalf("Expected body %q, got %q", expectedResponseBody, actualResponseBody)
	}
}

func prepareTestUser(userService *services.UserService, userID string, initialBalance int) error {
	// create user
	_, err := userService.CreateUserFromID(userID)
	if err != nil {
		return fmt.Errorf("Error creating user: %v", err)
	}

	// set user balance
	err = userService.UpdateUser(userID, &common.KeyValue{
		"jpy_balance": initialBalance,
	})
	if err != nil {
		return fmt.Errorf("Error updating user balance: %v", err)
	}

	return nil
}

func sendBuyCryptoRequest(userID, requestBody string) (*http.Response, error) {
	// create a request
	req, err := http.NewRequest(
		"POST",
		common.JoinWithBackendUrl("/buy_crypto"),
		bytes.NewBufferString(requestBody),
	)
	if err != nil {
		return nil, err
	}

	// custom header
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Test-User-ID", userID)

	// send request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func readResponseBody(body io.Reader) (string, error) {
	respBody, err := io.ReadAll(body)
	if err != nil {
		return "", fmt.Errorf("Failed to read response body: %v", err)
	}
	return string(respBody), nil
}
