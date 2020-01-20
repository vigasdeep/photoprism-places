PhotoPrism Places: Geocoding API
================================

[![License: AGPL](https://img.shields.io/badge/license-AGPL-blue.svg)][license]
[![Code Quality](https://goreportcard.com/badge/github.com/photoprism/photoprism-places)][goreport]
[![Build Status](https://travis-ci.org/photoprism/photoprism-places.png?branch=develop)][ci]
[![Documentation](https://readthedocs.org/projects/photoprism-docs/badge/?version=latest&style=flat)][docs]
[![GitHub contributors](https://img.shields.io/github/contributors/photoprism/photoprism-places.svg)](https://github.com/photoprism/photoprism-places/graphs/contributors/)
[![Community Chat](https://img.shields.io/badge/chat-on%20gitter-4aa087.svg)][chat]
[![Twitter](https://img.shields.io/badge/follow-@browseyourlife-00acee.svg)][twitter]

Note that this is work in progress. We are happy to assist other OSS projects that don't have the time or expertise to run their own infrastructure.

## Example Request ## 

https://places.photoprism.org/v1/location/149ce78563

```json
{
  "id":"149ce78563",
  "name":"Pink Beach",
  "category":"nature",
  "timezone":"Europe/Athens",
  "lat":35.26963621850717,
  "lng":23.53695076231683,
  "place": {
    "id":"149ce78563",
    "label":"Chrisoskalitissa, Crete, Greece",
    "city":"Chrisoskalitissa",
    "state":"Crete",
    "country":"gr"
  },
  "events":[],
  "licence":"Data © OpenStreetMap contributors"
}
```

## Privacy ##

Geocoding requests are NOT logged, but developers can of course see cached items in MariaDB without personal information. That's the point of a cache. Those will be randomly distributed with hot spots around tourist attractions and big cities.

Because of HTTPS, your internet provider can't see the exact request, just that you contacted a server.

The API approximates coordinates, encodes them with S2 and doesn't care about street or house number:

![](https://pbs.twimg.com/media/EN9AoYdWkAIqVDD?format=jpg&name=medium)

## Performance ##

First [benchmarks](https://github.com/tsliwowicz/go-wrk) show that up to 2500 req/s can be handled. Compare this with the pricing of commercial providers and you'll see the value.

If you prefer running this on-site: We use a 6-core Intel Xeon processor, 320 GB of SSD and 16 GB of memory. 
In addition you'll have to download ~100 GB of data.
Due to the properties of S2 cell IDs, scaling and sharding should be easy if needed.

## How to contribute ##

We welcome contributions of any kind. If you have a bug or an idea, read our 
[guide](https://docs.photoprism.org/en/latest/contribute/) before opening an issue.
Issues labeled [help wanted](https://github.com/photoprism/photoprism-places/labels/help%20wanted) / 
[easy](https://github.com/photoprism/photoprism-places/issues?q=is%3Aissue+is%3Aopen+label%3Aeasy) can be
good (first) contributions. Don't be afraid to ask stupid questions.

Feature requests backed by sponsors are marked with a golden [sponsor][sponsored issues] label.
Let us know if we mistakenly label an idea as [unfunded][unfunded issues].

Please follow us on [Twitter][twitter] and leave a star if you like this project, it provides additional motivation to keep going.

Thank you very much! <3

## Disclaimer ##

We'd like to remind everyone that we are no full-time marketing specialists but developers who work a lot and 
enjoy a bit of sarcasm from time to time. Please let us know when there is an issue with our "nuance and tone" 
and we'll find a solution.

[help]: https://groups.google.com/a/photoprism.org/forum/#!forum/help
[license]: https://github.com/photoprism/photoprism-places/blob/develop/LICENSE
[patreon]: https://www.patreon.com/photoprism
[paypal]: https://www.paypal.me/photoprism
[goreport]: https://goreportcard.com/report/github.com/photoprism/photoprism-places
[coverage]: https://codecov.io/gh/photoprism/photoprism-places
[ci]: https://travis-ci.org/photoprism/photoprism-places
[docs]: https://docs.photoprism.org/en/latest/
[issuehunt]: https://issuehunt.io/repos/119160553
[chat]: https://gitter.im/browseyourlife/community
[twitter]: https://twitter.com/browseyourlife
[unfunded issues]: https://github.com/photoprism/photoprism-places/issues?q=is%3Aissue+is%3Aopen+label%3Aunfunded
[sponsored issues]: https://github.com/photoprism/photoprism-places/issues?q=is%3Aissue+is%3Aopen+label%3Asponsor
