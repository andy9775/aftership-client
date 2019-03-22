# aftership-go

Unofficial go client library for the aftership api service

**WARNING** This library does not support all queries at this time. Feel free to update and issue a pull request
Currently we only support

- Create a tracking using `NewTracking`
- Get tracking info using `GetTracking`

---

The go structs are created using [json-to-go](https://mholt.github.io/json-to-go/) based on the aftership api [documentation](https://docs.aftership.com/api/4/overview) examples
