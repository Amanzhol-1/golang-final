package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type SubscriptionChecker interface {
	HasActiveSubscription(ctx context.Context, userID string) (bool, error)
}

type SubscriptionClient struct {
	baseURL    string
	httpClient *http.Client
}

func NewSubscriptionClient(baseURL string) SubscriptionChecker {
	return &SubscriptionClient{
		baseURL:    baseURL,
		httpClient: http.DefaultClient,
	}
}

func (c *SubscriptionClient) HasActiveSubscription(ctx context.Context, userID string) (bool, error) {
	log.Printf("✔ checking subscription for user %q", userID)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet,
		fmt.Sprintf("%s/api/v1/subscriptions?user_id=%s", c.baseURL, userID),
		nil,
	)
	if err != nil {
		return false, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		// assume body = { "active": true }
		var body struct {
			Active string `json:"Status"`
		}
		if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
			return false, err
		}
		log.Printf("✔ status code ok %q", body.Active)
		return body.Active == "active", nil
	}
	if resp.StatusCode == http.StatusNotFound {
		// no subscription record
		log.Printf("✔ status code notfound %q %q", resp.Status, req.URL)
		return false, nil
	}
	return false, fmt.Errorf("subscription service returned %d", resp.StatusCode)
}

func RequireSubscription(checker SubscriptionChecker) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// 1) extract userID however you’ve authenticated
			userID := "0"
			c.Set("userID", userID)
			// 2) ask subscription service
			ok, err := checker.HasActiveSubscription(c.Request().Context(), userID)

			if err != nil {
				log.Printf("x error %q", err)
				return echo.NewHTTPError(http.StatusInternalServerError, "subscription check failed")
			}
			if !ok {
				return echo.NewHTTPError(http.StatusForbidden, "you must have an active subscription to do that")
			}
			// 3) continue to the shipment handler
			return next(c)
		}
	}
}
