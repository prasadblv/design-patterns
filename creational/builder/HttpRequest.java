import java.util.Map;
import java.util.HashMap;

public class HttpRequest {

    private String url;
    private String method;
    private Map<String, String> headers;
    private String body;

    public static class HttpBuilder {

        private HttpRequest req = new HttpRequest();

        public HttpBuilder url(final String url) {
            req.url = url;
            return this;
        }

        public HttpBuilder method(final String method) {
            req.method = method;
            return this;
        }

        public HttpBuilder body(final String body) {
            req.body = body;
            return this;
        }

        public HttpBuilder headers(final String key, final String value) {
            if (req.headers == null) {
                req.headers = new HashMap<>();
            }
            req.headers.put(key, value);
            return this;
        }

        public HttpRequest build() {
            if (req.url == null || req.method == null) {
                throw new IllegalStateException("URl should not be empty");
            }
            return req;
        }
    }

    public static void main(String... v) {
        HttpRequest req = new HttpBuilder().url("http://localhost:8080/test").method("GET")
                .headers("Content-type", "applictaion/json").build();

    }
}
