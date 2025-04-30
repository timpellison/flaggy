# Flaggy!
Flaggy is an API that supports the most basic feature flag use case (On or Off).

## Usage

    make start

Build flaggy, tag docker image to flaggy:latest and compose flaggy up to docker

    make create-flag

Create the sample feature flag "my-feature" and deploy it to flaggy as a disabled feature flag

    make enable-flag

Simulate post-deployment scenario where we want to turn on the new feature.

    make disable-flag

Simulate post-deploymnet scenario where we turned the flag on, found out our new feature isn't up to snuff yet and need to turn it back off.

