# Fakedevices

Fakedevices is a test application for generating `IoT device` output to MQTT.

*!! Please note that this is a proof of concept ONLY at this point !!*


## Build

Download the source code into the $GOPATH:

```
    cd $GOPATH
    git clone https://github.com/istherepie/fakedevices
    ...
```

Run the build command (requires make):

```
    make all
```

This will generate a `build` directory with a current build (and shortcut).


## Usage

Assuming the build step was completed...

```
    cd /path/to/build/directory
    ./fakedevices --help
```

Otherwise use:

```
    cd /path/to/fakedevices
    go run cmd/main.go --help
```

The following arguments must be passed:

 * -d <Path to device file> *REQUIRED*
    For more information on the device file, please see the test sample (test/data/test_device_file.yml).

 * -h <Hostname/IP address of the MQTT broker>
    This defaults to `localhost`.

 * -p <Port of the MQTT broker>
    This defaults to `1883`.

Example:

```
    ./fakedevices -d test/data/test_device_file.yml -h 10.20.30.155 -p 18883
```


## License

MIT © Steffen Park <dev@istherepie.com>