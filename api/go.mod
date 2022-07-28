module github.com/winnix/proreporting/api

go 1.16

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gin-contrib/sessions v0.0.3
	github.com/gin-gonic/gin v1.7.1
	github.com/go-playground/validator/v10 v10.6.1
	github.com/google/uuid v1.2.0
	github.com/kr/text v0.2.0 // indirect
	github.com/markbates/goth v1.67.1
	github.com/niemeyer/pretty v0.0.0-20200227124842-a10e7caefd8e // indirect
	github.com/sirupsen/logrus v1.8.1
	github.com/stretchr/testify v1.7.0 // indirect
	github.com/valyala/quicktemplate v1.6.3
	github.com/winnix/proreporting/api/clients/bookingclient v0.0.0-00010101000000-000000000000
	github.com/winnix/proreporting/api/clients/integrationclient v0.0.0-00010101000000-000000000000
	go.mongodb.org/mongo-driver v1.5.3
	golang.org/x/oauth2 v0.0.0-20210514164344-f6687ab2804c // indirect
	gopkg.in/check.v1 v1.0.0-20200227125254-8fa46927fb4f // indirect
	gopkg.in/yaml.v2 v2.4.0
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
)

replace github.com/winnix/proreporting/api/clients/bookingclient => ./clients/bookingclient

replace github.com/winnix/proreporting/api/clients/integrationclient => ./clients/integrationclient
