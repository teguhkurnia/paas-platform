package usecase

import (
	"bytes"
	"encoding/json"
	"gofiber-boilerplate/internal/model"
	"gofiber-boilerplate/internal/util"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type ConnectGithubBody struct {
	Name           string            `json:"name"`
	URL            string            `json:"url"`
	HookAttributes map[string]string `json:"hook_attributes"`
	RedirectURL    string            `json:"redirect_url"`  // url that redirect after initiate app
	CallbackURLs   string            `json:"callback_urls"` // urls that redirect after authorize app
	SetupURL       string            `json:"setup_url"`
	Public         bool              `json:"public"`
}

type GitUseCase struct {
	Log      *logrus.Logger
	Config   *viper.Viper
	Validate *validator.Validate
	GitUtil  *util.GitUtil
}

func NewGitUseCase(
	config *viper.Viper,
	log *logrus.Logger,
	validate *validator.Validate,
	util *util.GitUtil,
) *GitUseCase {
	return &GitUseCase{
		Config:   config,
		Log:      log,
		Validate: validate,
		GitUtil:  util,
	}
}

func (u *GitUseCase) ConnectGithub() error {
	body := &ConnectGithubBody{
		Name: u.Config.GetString("app.name") + " GitHub App",
		HookAttributes: map[string]string{
			"url": u.Config.GetString("app.url") + "/api/v1/webhook/github",
		},
		RedirectURL:  u.Config.GetString("app.url") + "/api/v1/git/connect/github/",
		CallbackURLs: u.Config.GetString("app.url") + "/api/v1/git/connect/github/callback",
		SetupURL:     u.Config.GetString("app.url") + "/api/v1/git/connect/github/setup",
		Public:       false,
	}

	bodyJson, err := json.Marshal(body)
	if err != nil {
		return err
	}

	res, err := http.Post("https://github.com/settings/apps/new", "application/json", bytes.NewBuffer(bodyJson))
	if err != nil {
		return err
	}

	defer res.Body.Close()

	u.Log.Debugf("Github Response %+v", json.NewDecoder(res.Body))

	return nil
}

func (u *GitUseCase) CloneRepository(request *model.CloneGitRequest) error {
	err := u.Validate.Struct(request)
	if err != nil {
		u.Log.Warnf("Invalid request body: %v", err)
		return fiber.ErrBadRequest
	}

	uuid := uuid.New().String()
	destinationPath := "./tmp/git/" + uuid

	err = u.GitUtil.GitClone(request.URL, destinationPath)
	if err != nil {
		return fiber.ErrInternalServerError
	}

	return nil
}
