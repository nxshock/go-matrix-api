package matrixapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/google/go-querystring/query"
	"github.com/pkg/errors"
)

type Client struct {
	AccessToken string
	HomeServer  string
	UserID      string

	JoinedRooms map[string]*JoinedRoomC
	Invites     map[string]*InviteC

	InviteHandler  InviteHandler
	MessageHandler MessageHandler

	serverURL  string
	HTTPClient *http.Client

	nextBatch string
	Logger    io.Writer
}

func NewClientWithPassword(host string, user, password string) (*Client, error) {
	u, err := url.Parse(host)
	if err != nil {
		return nil, err
	}

	u2 := url.URL{
		Scheme: u.Scheme,
		Host:   u.Hostname() + ":" + u.Port()}

	client := &Client{
		serverURL:   u2.String(),
		JoinedRooms: make(map[string]*JoinedRoomC),
		Invites:     make(map[string]*InviteC),
		HTTPClient:  new(http.Client)}

	var loginReply LoginReply
	loginRequest := &LoginRequest{
		Type: AuthenticationTypePassword,
		Identifier: UserIdentifier{
			Type: IdentifierTypeUser,
			User: user},
		DeviceID: "go-matrix-api",
		Password: password}

	err = client.do("POST", "/_matrix/client/r0/login", loginRequest, &loginReply)
	if err != nil {
		return nil, err
	}

	client.AccessToken = loginReply.AccessToken
	client.HomeServer = loginReply.HomeServer
	client.UserID = loginReply.UserID

	return client, nil
}

func (client *Client) do(method, path string, request interface{}, reply interface{}) error {
	u, err := url.Parse(client.serverURL)
	if err != nil {
		return err
	}
	u.Path = path

	var req *http.Request
	val := make(url.Values)
	switch method {
	case "GET":
		val, err := query.Values(request)
		if err != nil {
			return err
		}
		if client.AccessToken != "" {
			val.Set("access_token", client.AccessToken)
		}
		u.RawQuery = val.Encode()

		req, err = http.NewRequest(method, u.String(), nil)
		if err != nil {
			return err
		}
	case "POST", "PUT":
		if client.AccessToken != "" {
			val.Set("access_token", client.AccessToken)
		}
		u.RawQuery = val.Encode()

		b, err := json.Marshal(request)
		if err != nil {
			return err
		}

		req, err = http.NewRequest(method, u.String(), bytes.NewReader(b))
		if err != nil {
			return err
		}
		req.Header.Set("Content-Type", "application/json")
	default:
		return errors.Errorf("unknown method: %s", method)
	}

	resp, err := client.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		var apiError APIError
		json.Unmarshal(b, &apiError)
		if err != nil {
			return err
		}

		return &apiError
	}

	return json.Unmarshal(b, &reply)
}

func (client *Client) getOrCreateRoom(roomID string) *JoinedRoomC {
	if client.JoinedRooms == nil {
		client.JoinedRooms = make(map[string]*JoinedRoomC)
	}

	if client.JoinedRooms[roomID] == nil {
		client.JoinedRooms[roomID] = new(JoinedRoomC)
	}

	return client.JoinedRooms[roomID]
}

func (client *Client) logln(a ...interface{}) {
	if client.Logger == nil {
		return
	}

	fmt.Fprintln(client.Logger, a...)
}

func (client *Client) log(a ...interface{}) {
	if client.Logger == nil {
		return
	}

	fmt.Fprint(client.Logger, a...)
}

func (client *Client) logf(format string, a ...interface{}) {
	if client.Logger == nil {
		return
	}

	fmt.Fprintf(client.Logger, format, a...)
}
