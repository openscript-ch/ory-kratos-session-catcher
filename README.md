# Ory Kratos Session Catcher

Unfortunately with [Ory Kratos](https://www.ory.sh/kratos/) it is currently impossible to offer authentication via one of the social login providers (Google, Apple, ..) for non-browser API apps like mobile apps. The browser API cannot be used, because the session is sent as a HttpOnly cookie to the client. It is hard or impossible to extract the session from a mobile webview.

This service can be called after a session is returned from Kratos to a browser endpoint. It will read the cookie and redirects the user back to a non-browser app via deep linking.

## Status

A solution without this service seems to be planed. As soon as there is a solution integrated into Kratos this project becomes obsolete.

 - https://github.com/ory/kratos/pull/2346