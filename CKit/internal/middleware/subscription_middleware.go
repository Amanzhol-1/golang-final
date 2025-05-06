package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type SubscriptionInfo struct {
	IsActive  bool
	StartDate time.Time
	EndDate   time.Time
}

type SubscriptionChecker interface {
	HasActiveSubscription(ctx context.Context, userID string) (SubscriptionInfo, error)
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

func (c *SubscriptionClient) HasActiveSubscription(ctx context.Context, userID string) (SubscriptionInfo, error) {
	log.Printf("✔ checking subscription for user %q", userID)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet,
		fmt.Sprintf("%s/api/v1/subscriptions?user_id=%s", c.baseURL, userID),
		nil,
	)
	if err != nil {
		return SubscriptionInfo{IsActive: false}, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return SubscriptionInfo{IsActive: false}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		var body struct {
			Active    string    `json:"Status"`
			StartDate time.Time `json:"StartDate"`
			EndDate   time.Time `json:"EndDate"`
		}
		if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
			return SubscriptionInfo{IsActive: false}, err
		}
		log.Printf("✔ status code ok %q", body.Active)
		return SubscriptionInfo{IsActive: body.Active == "active", StartDate: body.StartDate, EndDate: body.EndDate}, nil
	}
	if resp.StatusCode == http.StatusNotFound {
		// no subscription record
		log.Printf("✔ status code notfound %q %q", resp.Status, req.URL)
		return SubscriptionInfo{IsActive: false}, nil
	}
	return SubscriptionInfo{IsActive: false}, fmt.Errorf("subscription service returned %d", resp.StatusCode)
}

func RequireSubscription(checker SubscriptionChecker) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// 1) extract userID however you’ve authenticated
			userID := c.Get("AuthInfo").(AuthInfo).UserId
			fmt.Printf("userID: %s", userID)
			// 2) ask subscription service
			ok, err := checker.HasActiveSubscription(c.Request().Context(), userID)
			fmt.Printf("okStatus: %s", ok)
			c.Set("SubscriptionInfo", ok)

			if err != nil {
				log.Printf("x error %q", err)
				return echo.NewHTTPError(http.StatusInternalServerError, "subscription check failed")
			}
			if !ok.IsActive {
				return echo.NewHTTPError(http.StatusForbidden, "you must have an active subscription to do that")
			}
			// 3) continue to the shipment handler
			return next(c)
		}
	}
}
