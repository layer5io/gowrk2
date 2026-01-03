# gowrk2

Meshery supports user choice of load generator. See the [Integrating Meshery with Load Generators](https://docs.google.com/document/d/1jZMOih3Qn1ixrq1Ve6fkdZZ4TuKNoF_qfClDZTnLNcw/edit) design document for more details.

To test:
```sh
go run main.go 
```

```sh
{"AvgLatency":"98339.20","Bytes":"12606","BytesTransferPerSec":"6261.34","DurationInMicroseconds":"2013308.00","MaxLatency":"240768.00","MinLatency":"13968.00","Percentiles":[{"Count":"10","Percent":"50","Value":"83.839000"},{"Count":"15","Percent":"75","Value":"120.063000"},{"Count":"18","Percent":"90","Value":"170.751000"},{"Count":"20","Percent":"99","Value":"240.895000"},{"Count":"20","Percent":"99.9","Value":"240.895000"},{"Count":"20","Percent":"99.99","Value":"240.895000"},{"Count":"20","Percent":"99.999","Value":"240.895000"},{"Count":"20","Percent":"100","Value":"240.895000"}],"RequestsPerSec":"10.93","StdDev":"53840.26","UrlRequestCount_1":8,"UrlRequestCount_2":14,"Url_1":"https://gmail.com:443","Url_2":"https://gmail.com:443"}
```
<p style="clear:both;">
<h2><a name="contributing"></a><a name="community"></a> <a href="http://slack.layer5.io">Community</a> and <a href="https://github.com/layer5io/layer5/blob/master/CONTRIBUTING.md">Contributing</a></h2>
Our projects are community-built and welcome collaboration. 👍 Be sure to see the <a href="https://layer5.io/community/newcomers">Layer5 Community Welcome Guide</a> for a tour of resources available to you and jump into our <a href="http://slack.layer5.io">Slack</a>! Contributors are expected to adhere to the <a href="https://github.com/cncf/foundation/blob/master/code-of-conduct.md">CNCF Code of Conduct</a>.

<a href="https://slack.meshery.io">

<picture align="right">
  <source media="(prefers-color-scheme: dark)" srcset=".github/readme/images//slack-dark-128.png"  width="110px" align="right" style="margin-left:10px;margin-top:10px;">
  <source media="(prefers-color-scheme: light)" srcset=".github/readme/images//slack-128.png" width="110px" align="right" style="margin-left:10px;padding-top:5px;">
  <img alt="Shows an illustrated light mode meshery logo in light color mode and a dark mode meshery logo dark color mode." src=".github/readme/images//slack-128.png" width="110px" align="right" style="margin-left:10px;padding-top:13px;">
</picture>
</a>


<a href="https://meshery.io/community"><img alt="Layer5 Cloud Native Community" src=".github/readme/images/community.png" style="margin-right:8px;padding-top:5px;" width="140px" align="left" /></a>

<p>
✔️ <em><strong>Join</strong></em> any or all of the weekly meetings on <a href="https://calendar.google.com/calendar/b/1?cid=bGF5ZXI1LmlvX2VoMmFhOWRwZjFnNDBlbHZvYzc2MmpucGhzQGdyb3VwLmNhbGVuZGFyLmdvb2dsZS5jb20">community calendar</a>.<br />
✔️ <em><strong>Watch</strong></em> Community <a href="https://www.youtube.com/channel/UCFL1af7_wdnhHXL1InzaMvA?sub_confirmation=1">Meeting Recordings</a>.<br />
✔️ <em><strong>Access</strong></em> the <a href="https://drive.google.com/drive/u/4/folders/0ABH8aabN4WAKUk9PVA">Community Drive</a> by completing a community <a href="https://layer5.io/newcomer">Member Form</a>.<br />
✔️ <em><strong>Discuss</strong></em> in the <a href="https://discuss.layer5.io">Community Forum</a>.<br />
</p>
<p align="center">
<i>Not sure where to start?</i> Grab an open issue with the <a href="https://github.com/issues?q=is%3Aopen+is%3Aissue+archived%3Afalse+org%3Alayer5io+org%3Ameshery+org%3Aservice-mesh-performance+org%3Aservice-mesh-patterns+label%3A%22help+wanted%22+">help-wanted label</a>.
</p>

## About Layer5

[Layer5](https://layer5.io)'s cloud native application and infrastructure management software enables organizations to expect more from their infrastructure. We embrace developer-defined infrastructure. We empower engineer to change how they write applications, support operators in rethinking how they run modern infrastructure and enable product owners to regain full control over their product portfolio.


**License**

This repository and site are available as open source under the terms of the [Apache 2.0 License](https://opensource.org/licenses/Apache-2.0).
