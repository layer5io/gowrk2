# gowrk2

Meshery supports user choice of load generator. See the [Integrating Meshery with Load Generators](https://docs.google.com/document/d/1jZMOih3Qn1ixrq1Ve6fkdZZ4TuKNoF_qfClDZTnLNcw/edit) design document for more details.

To test:
```
go run main.go 
```
```
{"AvgLatency":"98339.20","Bytes":"12606","BytesTransferPerSec":"6261.34","DurationInMicroseconds":"2013308.00","MaxLatency":"240768.00","MinLatency":"13968.00","Percentiles":[{"Count":"10","Percent":"50","Value":"83.839000"},{"Count":"15","Percent":"75","Value":"120.063000"},{"Count":"18","Percent":"90","Value":"170.751000"},{"Count":"20","Percent":"99","Value":"240.895000"},{"Count":"20","Percent":"99.9","Value":"240.895000"},{"Count":"20","Percent":"99.99","Value":"240.895000"},{"Count":"20","Percent":"99.999","Value":"240.895000"},{"Count":"20","Percent":"100","Value":"240.895000"}],"RequestsPerSec":"10.93","StdDev":"53840.26","UrlRequestCount_1":8,"UrlRequestCount_2":14,"Url_1":"https://gmail.com:443","Url_2":"https://gmail.com:443"}
```
<p style="clear:both;">
<h2><a name="contributing"></a><a name="community"></a> <a href="http://slack.layer5.io">Community</a> and <a href="https://github.com/layer5io/layer5/blob/master/CONTRIBUTING.md">Contributing</a></h2>
Our projects are community-built and welcome collaboration. 👍 Be sure to see the <a href="https://docs.google.com/document/d/17OPtDE_rdnPQxmk2Kauhm3GwXF1R5dZ3Cj8qZLKdo5E/edit">Layer5 Community Welcome Guide</a> for a tour of resources available to you and jump into our <a href="http://slack.layer5.io">Slack</a>! Contributors are expected to adhere to the <a href="https://github.com/cncf/foundation/blob/master/code-of-conduct.md">CNCF Code of Conduct</a>.

<a href="https://meshery.io/community"><img alt="Layer5 Service Mesh Community" src=".github/readme/images//slack-128.png" style="margin-left:10px;padding-top:5px;" width="110px" align="right" /></a>

<a href="http://slack.layer5.io"><img alt="Layer5 Service Mesh Community" src=".github/readme/images//community.svg" style="margin-right:8px;padding-top:5px;" width="140px" align="left" /></a>

<p>
✔️ <em><strong>Join</strong></em> <a href="https://drive.google.com/open?id=1c07UO9dS7_tFD-ClCWHIrEzRnzUJoFQ10EzfJTpS7FY">weekly community meeting</a> on <a href="https://bit.ly/2SbrRhe">Fridays from 10am - 11am Central</a>.<br />
✔️ <em><strong>Watch</strong></em> community <a href="https://www.youtube.com/channel/UCFL1af7_wdnhHXL1InzaMvA?sub_confirmation=1">meeting recordings</a>.<br />
✔️ <em><strong>Access</strong></em> the <a href="https://drive.google.com/drive/u/4/folders/0ABH8aabN4WAKUk9PVA">community drive</a>.<br />
</p>
<p align="center">
<i>Not sure where to start?</i> Grab an open issue with the <a href="https://github.com/issues?utf8=✓&q=is%3Aopen+is%3Aissue+archived%3Afalse+org%3Alayer5io+label%3A%22help+wanted%22+">help-wanted label</a>.
</p>

## About Layer5

**Community First**
<p>The <a href="https://layer5.io">Layer5</a> community represents the largest collection of service mesh projects and their maintainers in the world.</p>

**Open Source First**
<p>We build projects to provide learning environments, deployment and operational best practices, performance benchmarks, create documentation, share networking opportunities, and more. Our shared commitment to the open source spirit pushes Layer5 projects forward.</p>

**License**

This repository and site are available as open source under the terms of the [Apache 2.0 License](https://opensource.org/licenses/Apache-2.0).
