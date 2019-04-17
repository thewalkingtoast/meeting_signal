## Meeting Signal

#### Pre-requisites

Uses the `ykush` command to control USB power via a Yepkit board.
See https://github.com/Yepkit/ykush and https://www.yepkit.com/product/300115/YKUSHXS.

For the signal lamp itself, I'm using Target's DIY Lightbox:
https://www.target.com/p/diy-lightbox-novelty-led-table-lamp-black-room-essentials-153/-/A-53216203

Setup:

```
D/C Power -> Raspberry PI ZeroW -> Ykushxs -> Signal Lamp
```

#### Compile

For Raspberry PI Zero:

`GOOS=linux GOARCH=arm GOARM=5 go build`
