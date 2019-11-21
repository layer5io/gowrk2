# gowrk2

Meshery supports user choice of load generator. See the [Integrating Meshery with Load Generators](https://docs.google.com/document/d/1jZMOih3Qn1ixrq1Ve6fkdZZ4TuKNoF_qfClDZTnLNcw/edit) design document for more details.

To test:
```
go run main.go 
```
```
{"AvgLatency":"98339.20","Bytes":"12606","BytesTransferPerSec":"6261.34","DurationInMicroseconds":"2013308.00","MaxLatency":"240768.00","MinLatency":"13968.00","Percentiles":[{"Count":"10","Percent":"50","Value":"83.839000"},{"Count":"15","Percent":"75","Value":"120.063000"},{"Count":"18","Percent":"90","Value":"170.751000"},{"Count":"20","Percent":"99","Value":"240.895000"},{"Count":"20","Percent":"99.9","Value":"240.895000"},{"Count":"20","Percent":"99.99","Value":"240.895000"},{"Count":"20","Percent":"99.999","Value":"240.895000"},{"Count":"20","Percent":"100","Value":"240.895000"}],"RequestsPerSec":"10.93","StdDev":"53840.26","UrlRequestCount_1":8,"UrlRequestCount_2":14,"Url_1":"https://gmail.com:443","Url_2":"https://gmail.com:443"}
```

## Contributing
Please do! Contributions, updates, [discrepancy reports](/../../issues) and [pull requests](/../../pulls) are welcome. This project is community-built and welcomes collaboration. Contributors are expected to adhere to the [Contributor Covenant](http://contributor-covenant.org) code of conduct.

Not sure where to start? See the [newcomers welcome guide](https://docs.google.com/document/d/14Fofs9BysojB5igihXBI_SsFWoSUu-QRsGnnFqUvR0M/edit) for how, where and why to contribute. Or grab an open issue with the [help-wanted label](https://github.com/issues?utf8=✓&q=is%3Aopen+is%3Aissue+archived%3Afalse+org%3Alayer5io+label%3A%22help+wanted%22+) and jump in.

#### About Layer5
[Layer5.io](https://layer5.io) is a service mesh community, serving as a repository for information pertaining to the surrounding technology ecosystem (service meshes, api gateways, edge proxies, ingress and egress controllers) of microservice management in cloud native environments.
