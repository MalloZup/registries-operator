/*
 * Copyright 2018 SUSE LINUX GmbH, Nuernberg, Germany..
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package assets

var Certs = map[string][]byte{"foo.crt":
[]byte(`-----BEGIN CERTIFICATE-----
MIIGKjCCBBKgAwIBAgIJANENwKQx6q0aMA0GCSqGSIb3DQEBCwUAMIGhMQswCQYD
VQQGEwJERTEQMA4GA1UECAwHQmF2YXJpYTESMBAGA1UEBwwJTnVyZW1iZXJnMRsw
GQYDVQQKDBJTVVNFIEF1dG9nZW5lcmF0ZWQxMzAxBgNVBAsMKkNhYVNQIGJiNGU5
OWViLWJiOTEtNDc1Mi04ZGQ2LWQ4YWViNDZhZjY5YzEaMBgGA1UEAwwRQ2FhU1Ag
SW50ZXJuYWwgQ0EwHhcNMTgxMjA5MTUzMTMxWhcNMjgxMjA2MTUzMTMxWjCBoTEL
MAkGA1UEBhMCREUxEDAOBgNVBAgMB0JhdmFyaWExEjAQBgNVBAcMCU51cmVtYmVy
ZzEbMBkGA1UECgwSU1VTRSBBdXRvZ2VuZXJhdGVkMTMwMQYDVQQLDCpDYWFTUCBi
YjRlOTllYi1iYjkxLTQ3NTItOGRkNi1kOGFlYjQ2YWY2OWMxGjAYBgNVBAMMEUNh
YVNQIEludGVybmFsIENBMIICIjANBgkqhkiG9w0BAQEFAAOCAg8AMIICCgKCAgEA
u8uuLLu0LOI0Ssdo1obztfXoD064PMdREyuo4D30fH80JAKtE+0Dw2WANgGTpf5s
1z/CpDJs5W2u/XXavLqp6KvJjrDS/c2AYLAoanXnsPOicN8YvdN8nSZJ8SvbRnOW
1nuAKufBBI2Ubym8qH+2h59JixeC4AfXPdZRfUsm7UONbcqBP+8bMD0JeEcUhMKz
fLEjrCV51O44WHvSllqz8j15KaejVectucLT1p3rkyzbdhiEim52ypYVNOKTqM54
mhNjURqLagnt/zdURBsXASxWMiaNyumHk22xtLAHORHna6BpO6Xuu+vUari5IM0K
yO0HNEJAQbmHL0Wbi8DsygyW1uZvU2tfkkahHypv/4CurJbYdWwB7+ItOQ8wKhtJ
ZeVDhSF1fWL6pIelvrRFhiUaBqdvX/7zE4bobIJ4sDczeX5glbBpYwBpidUIDYmp
LmybczZU5v4B7d1XWHSDqA5CBwOvpuHaD9wkVKZZZVmjdYQuCJDPR6j8eYlCtiqB
spaxN+b97UOyz0q1E41Ox1en9JcI6RxCp4UXeZfgIgv3Pvb9L+AduXaLEU+TYsvC
/2LKrglLtdlQu9ZtGhZ8gvXCy/Y1B8f2CJbW0r6mP8YkKNJc5h0wwi+Pk0TIoVYg
ai+fEF6mUwteXw8oNdn4RwNpmBngzap1XPvkc3JLcpMCAwEAAaNjMGEwHQYDVR0O
BBYEFExq8DtMx8rVaUIF9BsZgONm2imtMB8GA1UdIwQYMBaAFExq8DtMx8rVaUIF
9BsZgONm2imtMA8GA1UdEwEB/wQFMAMBAf8wDgYDVR0PAQH/BAQDAgLkMA0GCSqG
SIb3DQEBCwUAA4ICAQCCKEGme6rMzxnYYW+avZHpl4vRQ3WCd5FVRD0oxhvUGn8p
BbSCfok+zeoDkVJW0FpGolBp7gZ97QQzQbiaP35U3nw3VOpqhPIN0HlRyvuA8fVk
yAcJR/IUkFXChgNtPKfJj250dXwt3bAXQo5rYbq08nr4IJ0jUMU6mcU32LxkRH4R
w3EvZbwW868O+rbzZPRXQLhJj23LGeet7CU5UIMBAXKCMHpLdWXnQHIkEnNBdi60
0KdPbNev91zuKEA4LMohapc32enUVhcY0LCcute4+jVBCEotYnSP1VvoyvWJFb03
4F+L5wynE3X4ZT4hqAn22Fx8aCmyZ1XLBMHtBafy6gbh1qyOXmq2uu9+v88+Y1gl
susXXo9ThTTJ7L2mO4HuopsMTShr8WlS/FDbH2eGNrL29arVOBPWc+cd8ujNDiFg
OCbGDTMpBAcMXJPNLYo247fuCvfojkFCIrkPLwSQKnWwgtUb5b1wkEYwahEysz5u
foPqziQ1v9kvu6pMS7YmTjmNRACjZZYOSCYAZz2PouWL0vpt7zDr455XAYUR9Pmw
cR4ugRFQyr3rcjJojSd7WlFu4lVno/+UybY2IRKHEiVK+gdU1KgGJBgjQcqAyKVB
/x4Lq78rdVYuOz4vdBq/iTwS5FONqe+eepTMv3AOAwGduLjk4QLGGyQ8fFOhRQ==
-----END CERTIFICATE-----`)}

