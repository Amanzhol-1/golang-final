package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type AuthInfo struct {
	UserId string
	Role   string
}

type AuthChecker interface {
	IsLoggedIn(ctx context.Context, token string) (AuthInfo, error)
}

type AuthClient struct {
	baseURL    string
	httpClient *http.Client
}

func NewAuthClient(baseURL string) AuthChecker {
	return &AuthClient{
		baseURL:    baseURL,
		httpClient: http.DefaultClient,
	}
}

func (c *AuthClient) IsLoggedIn(ctx context.Context, token string) (AuthInfo, error) {
	log.Printf("✔ checking auth for token %q", token)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet,
		fmt.Sprintf("%s/profile", c.baseURL),
		nil,
	)

	req.Header.Set("Authorization", fmt.Sprintf("%s", token))

	if err != nil {
		return AuthInfo{UserId: ""}, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return AuthInfo{UserId: ""}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		var body struct {
			UserId int    `json:"user_id"`
			Role   string `json:"role"`
		}
		if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
			return AuthInfo{UserId: ""}, err
		}
		log.Printf("✔ status code ok %q", body.UserId)
		return AuthInfo{UserId: strconv.Itoa(body.UserId), Role: body.Role}, nil
	}
	if resp.StatusCode == http.StatusNotFound {
		// no subscription record
		log.Printf("✔ status code notfound %q %q", resp.Status, req.URL)
		return AuthInfo{UserId: ""}, nil
	}
	return AuthInfo{UserId: ""}, fmt.Errorf("subscription service returned %d", resp.StatusCode)
}

func RequireAuth(checker AuthChecker) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			token := c.Request().Header.Get("Authorization")

			ok, err := checker.IsLoggedIn(c.Request().Context(), token)
			c.Set("AuthInfo", ok)

			if err != nil {
				log.Printf("x error %q", err)
				return echo.NewHTTPError(http.StatusInternalServerError, "auth check failed")
			}
			if ok.UserId == "" {
				return echo.NewHTTPError(http.StatusForbidden, "your token is invalid")
			}
			// 3) continue to the shipment handler
			return next(c)
		}
	}
}
