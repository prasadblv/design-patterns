#!/usr/bin/env python3


class HttpRequest:
    def __init__(self):
        self.url: str | None
        self.method: str | None
        self.body: str | None
        self.headers: map[str, str] = {}

    class HttpBuilder:
        def __init__(self):
            self._req = HttpRequest()

        def url(self, url: str) -> "HttpRequest.HttpBuilder":
            self._req.url = url
            return self

        def method(self, method: str) -> "HttpRequest.HttpBuilder":
            self._req.method = method
            return self

        def body(self, body: str) -> "HttpRequest.HttpBuilder":
            self._req.body = body
            return self

        def headers(self, key: str, val: str) -> "HttpRequest.HttpBuilder":
            self._req.headers[key] = val
            return self

        def build(self) -> "HttpRequest":
            if self._req.url is None:
                raise ValueError("URL Is mandatory!")
            return self._req


if __name__ == "__main__":
    req: HttpRequest = (
        HttpRequest.HttpBuilder().url("http://localhost:8080").method("GET").build()
    )
    print(req.method)
    print(req.url)
