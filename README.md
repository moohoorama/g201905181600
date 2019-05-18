# Go port of "Inovation 2007" by Omega

## Original Work

http://o-mega.sakura.ne.jp/product/ino.html

## Releases

### Web Browsers

http://hajimehoshi.github.io/go-inovation/

### Android

<a href='https://play.google.com/store/apps/details?id=com.hajimehoshi.goinovation&utm_source=global_co&utm_medium=prtnr&utm_content=Mar2515&utm_campaign=PartBadge&pcampaignid=MKT-Other-global-all-co-prtnr-py-PartBadge-Mar2515-1'><img alt='Get it on Google Play' src='https://play.google.com/intl/en_us/badges/images/generic/en_badge_web_generic.png' width="210px" height="80px"/></a>

### iOS

<a href="https://itunes.apple.com/us/app/%E3%81%84%E3%81%AE-%E3%81%B9%E3%83%BC%E3%81%97%E3%82%87%E3%82%93-2007/id1132624266?mt=8"><img src="https://linkmaker.itunes.apple.com/assets/shared/badges/en-us/appstore-lrg.svg" alt="Download on the App Store" width="135" height="40"></a>

## How to install and run on desktops

```
go get github.com/hajimehoshi/go-inovation
cd $GOPATH/src/github.com/hajimehoshi/go-inovation
go run main.go
```

## How to build for Android

At this directory, run

```
gomobile bind -target android -javapkg com.hajimehoshi.goinovation -o ./mobile/android/inovation/inovation.aar github.com/hajimehoshi/go-inovation/mobile
```

and run the Android Studio project in `./mobile/android`.

`GO111MODULE=off` might be required.

gomobile bind -target android -javapkg com.github.moohoorama.g201905181600 -o ./mobile/android/inovation/inovation.aar github.com/moohoorama/g201905181600/mobile


## How to build for iOS

At this directory, run

```
gomobile bind -target ios -o ./mobile/ios/Mobile.framework github.com/hajimehoshi/go-inovation/mobile
```

and run the Xcode project in `./mobile/ios`.

`GO111MODULE=off` might be required.
