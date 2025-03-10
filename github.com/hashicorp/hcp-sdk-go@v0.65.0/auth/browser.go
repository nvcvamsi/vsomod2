// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package auth

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/mitchellh/colorstring"
	"github.com/skratchdot/open-golang/open"
	"golang.org/x/oauth2"
)

// Browser implements the browser handler for the OAuth2 login flow.
type Browser interface {
	GetTokenFromBrowser(context.Context, *oauth2.Config) (*oauth2.Token, error)
}

// oauthBrowser implements the Browser interface using the real OAuth2 login flow.
type oauthBrowser struct{}

// GetTokenFromBrowser opens a browser window for the user to log in and handles the OAuth2 flow to obtain a token.
func (b *oauthBrowser) GetTokenFromBrowser(ctx context.Context, conf *oauth2.Config) (*oauth2.Token, error) {
	// Launch a request to Auth0's authorization endpoint.
	colorstring.Printf("[bold][yellow]The default web browser has been opened at %s. Please continue the login in the web browser.\n", conf.Endpoint.AuthURL)

	// Prepare the /authorize request with randomly generated state, offline access option, and audience
	aud := "https://api.hashicorp.cloud"
	opt := oauth2.SetAuthURLParam("audience", aud)
	authzURL := conf.AuthCodeURL(generateRandomString(32), oauth2.AccessTypeOffline, opt)

	// Handle ctrl-c while waiting for the callback
	sigintCh := make(chan os.Signal, 1)
	signal.Notify(sigintCh, os.Interrupt)
	defer signal.Stop(sigintCh)

	if err := open.Start(authzURL); err != nil {
		return nil, fmt.Errorf("failed to open browser at URL %q: %w", authzURL, err)
	}

	// Start callback server
	callbackEndpoint := &callbackEndpoint{}
	callbackEndpoint.shutdownSignal = make(chan error)
	server := &http.Server{
		Addr:           ":8443",
		Handler:        nil,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	callbackEndpoint.server = server
	http.Handle("/oidc/callback", callbackEndpoint)

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			callbackEndpoint.shutdownSignal <- fmt.Errorf("failed to start callback server: %w", err)
		}
	}()

	// Wait for either the callback to finish, SIGINT to be received or up to 2 minutes
	select {
	case err := <-callbackEndpoint.shutdownSignal:

		if err != nil {
			return nil, err
		}

		err = callbackEndpoint.server.Shutdown(context.Background())
		if err != nil {
			return nil, fmt.Errorf("failed to shutdown callback server: %w", err)
		}

		// Exchange the code returned in the callback for a token.
		tok, err := conf.Exchange(ctx, callbackEndpoint.code)
		if err != nil {
			return nil, fmt.Errorf("failed to exchange code for token: %w", err)
		}

		// Save the token to config file.
		cache := Cache{
			AccessToken:       tok.AccessToken,
			RefreshToken:      tok.RefreshToken,
			AccessTokenExpiry: tok.Expiry,
			SessionExpiry:     time.Now().Add(SessionMaxAge),
		}
		err = Write(cache)
		if err != nil {
			return nil, fmt.Errorf("failed to write token to file: %w", err)
		}

		return tok, nil
	case <-sigintCh:
		return nil, errors.New("interrupted")
	case <-time.After(2 * time.Minute):
		return nil, errors.New("timed out waiting for response from provider")
	}
}

// callbackEndpoint exposes the confiugration for the callback server.
type callbackEndpoint struct {
	server         *http.Server
	code           string
	shutdownSignal chan error
}

// callbackEndpoint endpoint ServeHTTP confirms if an Authorization code was received from Auth0.
func (h *callbackEndpoint) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	code := r.URL.Query().Get("code")
	if code != "" {
		h.code = code
		fmt.Fprintln(w, "Login is successful. You may close the browser and return to the command line.")
		colorstring.Println("[bold][green]Success!")
	} else {
		fmt.Fprintln(w, "Login is not successful. You may close the browser and try again.")
	}
	h.shutdownSignal <- nil
}

// generateRandomString returns a URL-safe, base64 encoded
// securely generated random string.
func generateRandomString(n int) string {
	b := make([]byte, n)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		panic(err)
	}

	return base64.RawURLEncoding.EncodeToString(b)
}
