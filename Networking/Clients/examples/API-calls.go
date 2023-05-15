r1, err := http.Get("http://www.google.com/robots.txt")
// Read response body. Not shown.
defer r1.Body.Close()

r2, err := http.Head("http://www.google.com/robots.txt")
// Read response body. Not shown.
defer r2.Body.Close()

form := url.Values{}
form.Add("foo", "bar")
r3, err = http.Post(
	"https://www.google.com/robots.txt",
	"application/x-www-form-urlencoded",
	strings.NewReader(form.Encode()),
)
// Read response body. Not Shown.
defer r3.Body.Close()

// PostForm() removes tediousness of setting values and manually encoding every request
form := url.Values{}
form.Add("foo", "bar")
r3, err := http.PostForm("https://www.google.com/robots.txt",form)
// Read response body and close.

// Generate request - DELETE example
req, err := http.NewRequest("DELETE", "https://www.google.com/robots.txt", nil)
var client http.Client
resp, err := client.Do(req)
// Read response body and close.

// Generate request - PUT example
form := url.Values()
form.Add("foo", "bar")
var client http.Client
req, err := http.NewRequest(
	"PUT",
	"https://www.google.com/robots.txt",
	strings.NewReader(form.Encode())
)
resp, err := client.Do(req)
//Read response body and close.
