# Android GUI applications written entirely in Go

See [main.go](main.go) for example application.

## State of the project

Right now project is not ready-to-use library, but it's intended to.

* Can be done from Go:
  * subscribe to any sensor values (using normal refresh interval);
  * list View widgets from any layout;
  * call any method with simple arguments for View widgets (accepting strings,
    numbers and boolean);
  * calling simple View methods, like SetText; bindings to all methods for View
    widgets are autogenerated from the official documentation; all calls are
    syncronous, call will return described in documentation type;
  * subscribing to onClick, onTouch events;
  * *WIP* subscribing on other View widget events;

## What a feature is that?

Everybody loves Go. Everybody wants to Go mobile.
But Go can not into android GUI, right? Not exactly.

It is time to build android applications written in pure Go
that are able to use android GUI (all that buttons, text fields
and so forth).

## But... how?

It's just slightly modified example from https://github.com/golang/mobile.

Beware muggles: magic inside. Android documentation parsing,
RPC to cross-communicate between Go and java, autogenerated code
via gobind, a couldron of java reflections concealed deep inside
the project.

## I've got the power?

As a developer, you can get all the views (resources) described in layout files
inside android project. Having this knowledge, you can call view methods
(for example, TextView.SetText) to change UI. Also, you can set up handlers
for view events (button clicks etc).

As much as for view events, listeners can be set for sensors. For example,
you can get actual data from accelerometer.

See "State of project" and issues list to get the picture on current state.

## Install and run

Prepare your system for android development. Android NDK and Android SDK (v19)
are required. You also need Apache Ant to build your project and
ADB to upload the project to android device.

There's a chance that no simple way to install all the needed software exists,
but you can start from here: https://github.com/golang/mobile.

Happy Arch Linux users can follow detailed instruction:
https://github.com/seletskiy/kb/blob/master/howto/go-for-android-setup.md.

After preparation, clone project and start cute little all.bash:
```
go get -d github.com/seletskiy/go-android-rpc
cd $GOPATH/src/github.com/seletskiy/go-android-rpc
setup-android-env-go ./generate.bash
setup-android-env-go ./all.bash
```

Details about setup-android-env-go can be found here:
https://github.com/seletskiy/kb/blob/master/howto/go-for-android-setup.md.
