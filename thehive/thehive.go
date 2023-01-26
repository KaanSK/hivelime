package thehive

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
	"go.uber.org/zap"

	"github.com/goccy/go-json"
)

type TheHiveClient struct {
	apiKey string
	url    string
	HTTP   *fiber.Agent
	Logger *zap.Logger
}

func SetHeaders(req *fasthttp.Request, method string, apiKey string) {
	req.Header.SetMethod(method)
	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")
}

func GetHiveClient(url string, key string, logger *zap.Logger) (*TheHiveClient, error) {
	if key == "" {
		return nil, errors.New("empty Hive API key")
	}

	return &TheHiveClient{
		apiKey: key,
		url:    url,
		Logger: logger,
	}, nil
}

func NewAlert() Alert {
	return Alert{}
}

func (a *Alert) AddObservable(obsType string, obs string, tags []string) {
	obsInstance := Observable{
		Data:     obs,
		DataType: obsType,
		Tags:     tags,
	}
	a.Observables = append(a.Observables, obsInstance)
}

func (s *TheHiveClient) CreateAlert(alert Alert) (id string, err error) {
	if s == nil {
		return id, errors.New("not initialized hive client")
	}
	payload, err := json.Marshal(alert)
	if err != nil {
		return id, err
	}

	agent := fiber.AcquireAgent()
	res := fiber.AcquireResponse()
	defer fiber.ReleaseAgent(agent)
	defer fiber.ReleaseResponse(res)

	req := agent.Request()
	SetHeaders(req, fiber.MethodPost, s.apiKey)

	req.SetRequestURI(s.url + "/api/v1/alert")
	req.SetBodyString(string(payload))
	if err := agent.Parse(); err != nil {
		return id, err
	}
	code, body, errs := agent.Bytes()
	if errs != nil {
		return id, errs[0]
	}

	if code != fiber.StatusCreated {
		if err != nil {
			return id, err
		}
		return id, errors.New(string(body))
	}
	createdAlert := NewAlert()
	err = json.Unmarshal(body, &createdAlert)
	if err != nil {
		return id, errors.New(err.Error())
	}
	return createdAlert.Id, nil
}
