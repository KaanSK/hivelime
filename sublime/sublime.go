package sublime

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type SublimeClient struct {
	apiKey string
	url    string
	HTTP   *fiber.Agent
	Logger *zap.Logger
}

func GetSublimeClient(url string, key string, logger *zap.Logger) (*SublimeClient, error) {
	if key == "" {
		return nil, errors.New("empty Sublime API key")
	}

	return &SublimeClient{
		apiKey: key,
		url:    url,
		Logger: logger,
	}, nil
}

func (s *SublimeClient) GetMessageGroup(id string) (mg MessageGroup, err error) {
	agent := fiber.AcquireAgent()
	res := fiber.AcquireResponse()
	defer fiber.ReleaseAgent(agent)
	defer fiber.ReleaseResponse(res)

	req := agent.Request()
	req.Header.SetMethod(fiber.MethodGet)
	req.Header.Set("Authorization", "Bearer "+s.apiKey)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("accept", "application/json")

	req.SetRequestURI(fmt.Sprintf("%s/v1/messages/groups/%s", s.url, id))
	if err := agent.Parse(); err != nil {
		return mg, err
	}
	code, body, errs := agent.Bytes()
	if errs != nil {
		return mg, errs[0]
	}

	if code != fiber.StatusOK {
		if err != nil {
			return mg, err
		}
		return mg, errors.New(string(body))
	}

	if err := json.Unmarshal(body, &mg); err != nil {
		return mg, err
	}

	return mg, nil
}
