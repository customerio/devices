# @customerio/devices

Utility to map device models to marketing friendly device name. The android list is pretty huge, and we've decided to cater to the most popular brands based on [market share](https://www.appbrain.com/stats/top-manufacturers). If you need a device that isn't in the list, please open an issue or a PR.

## Installation

```bash
npm i --save @customerio/devices
```

## Usage

```ts
import { getMarketingName, OsType } from '@customerio/devices';

console.log(getMarketingName('iPhone14,2', OsType.IOS));
```
Output:

```bash
iPhone 13 Pro
``` 

## Development

#### How to generate android devices json?

Simple golang script fetches data from https://storage.googleapis.com/play_public/supported_devices.html and generates json file.

```bash
cd scripts/android
go run main.go
```

TODO

#### How to generate ios devices json?

- Manually maintained 🤓 from https://www.theiphonewiki.com/wiki/Models for Apple TV, Watch, iPads and iPhones.

## Size details

[Size Limit](https://github.com/ai/size-limit) controls the size.
Small Libraries

```bash
$ npm run build && npm run size

✔ Adding to empty webpack project
✔ Running JS in headless Chrome
  
  Time limit:   3 s
  Size:         90.16 kB with all dependencies, minified and gzipped
  Loading time: 1.8 s    on slow 3G
  Running time: 266 ms   on Snapdragon 410
  Total time:   2.1 s

```

## Todo

- [ ] Protect main branch from direct commits
- [ ] Add a PR template
- [ ] Add development instructions
- [ ] Discuss the need of all 37,708 android devices
- [ ] Automate iOS data collection (far fetched)
