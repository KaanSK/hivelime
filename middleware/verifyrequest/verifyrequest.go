package verifyrequest

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

type HMACSignature struct {
	Timestamp int64
	Signature string
}

func New(config Config) fiber.Handler {
	cfg := configDefault(config)

	return func(c *fiber.Ctx) error {
		if cfg.Next != nil && cfg.Next(c) {
			return c.Next()
		}

		reqSignature := c.Get(cfg.Header, "")
		if reqSignature == "" {
			return cfg.Unauthorized(c)
		}
		body := string(c.Body())

		kvPairs := strings.Split(reqSignature, ",")
		reqTimestamp, err := strconv.ParseInt(strings.Split(string(kvPairs[0]), "=")[1], 10, 64)
		if err != nil {
			return cfg.Unauthorized(c)
		}
		reqHMACSignature := HMACSignature{
			Timestamp: reqTimestamp,
			Signature: strings.Split(string(kvPairs[1]), "=")[1],
		}

		signatureDatetime := time.Unix(reqHMACSignature.Timestamp, 0)
		if signatureDatetime.Before(time.Now().Add(-time.Minute * time.Duration(cfg.Expiration))) {
			return cfg.Unauthorized(c)
		}

		payloadToSign := fmt.Sprintf("%d.%s", reqHMACSignature.Timestamp, body)
		h := hmac.New(sha256.New, []byte(cfg.Secret))

		h.Write([]byte(payloadToSign))
		signedPayload := hex.EncodeToString(h.Sum(nil))

		if reqHMACSignature.Signature != signedPayload {
			return cfg.Unauthorized(c)
		}

		return c.Next()
	}
}
